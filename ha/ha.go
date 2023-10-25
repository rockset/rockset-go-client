package ha

import (
	"context"
	"errors"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// Querier is an interface implemented by the rockset.RockClient
type Querier interface {
	Query(context.Context, string, ...option.QueryOption) (openapi.QueryResponse, error)
}

type Client struct {
	clients []Querier
}

// New creates a new HA query client, and takes a list of rockset.RockClient which will be used for running
// queries in parallel.
func New(client ...Querier) *Client {
	return &Client{client}
}

// Query executes a SQL query in parallel, and returns the first response.
func (ha *Client) Query(ctx context.Context, query string,
	options ...option.QueryOption) (openapi.QueryResponse, []error) {
	log := zerolog.Ctx(ctx)

	var wg sync.WaitGroup
	resultCh := make(chan openapi.QueryResponse, len(ha.clients))
	errorCh := make(chan error, len(ha.clients))

	wg.Add(len(ha.clients))
	for idx, c := range ha.clients {
		go func(i int, cl Querier) {
			res, err := cl.Query(ctx, query, options...)
			if err != nil {
				log.Error().Err(err).Int("idx", i).Msg("failed to query")
				errorCh <- err
			} else {
				log.Trace().Int("idx", i).Msg("got query response")
				resultCh <- res
			}
			wg.Done()
		}(idx, c) // avoid using the loop variables as they mutate each iteration
	}

	// This go routine waits for all parallel query go routines to complete,
	// which happens once the first request is returned and the rest get cancelled,
	// or when all requests have failed. It then closes all channels.
	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()

	// create a sub-context, so we can cancel all HTTP requests after getting the first answer
	subCtx, cancel := context.WithCancel(ctx)

	// make sure to cancel any pending query result as we don't need them when we return
	defer cancel()

	return returnFirst(subCtx, resultCh, errorCh)
}

func returnFirst(ctx context.Context, resultCh chan openapi.QueryResponse,
	errorCh chan error) (openapi.QueryResponse, []error) {
	var errs []error
	for {
		select {
		case res, ok := <-resultCh:
			if !ok {
				log.Warn().Msg("receive from closed results channel")
				// avoid getting selected again once the channel is closed, which can happen when we just had errors
				resultCh = nil
				continue
			}

			// log any error we got BEFORE the successful response
			for _, err := range errs {
				log.Error().Err(err).Msg("before the response")
			}
			// log any error we got AFTER the successful response
			for err := range errorCh {
				if errors.Is(err, context.Canceled) {
					log.Error().Err(err).Msg("after the response")
				}
			}

			return res, nil
		case <-ctx.Done():
			log.Error().Err(ctx.Err()).Msg("context cancelled")
			errs = append(errs, ctx.Err())

			return openapi.QueryResponse{}, errs
		case err, ok := <-errorCh:
			// if the errorCh is closed, ok will be false, and it is time to call it a day
			if !ok {
				return openapi.QueryResponse{}, errs
			}
			errs = append(errs, err)
		}
	}
}
