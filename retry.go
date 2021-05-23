package rockset

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

const (
	collectionStatusREADY = "READY"
)

// WaitUntilAliasAvailable waits until the alias is available.
func (rc *RockClient) WaitUntilAliasAvailable(ctx context.Context, workspace, alias string) error {
	return rc.Retry(ctx, func() (bool, error) {
		_, _, err := rc.AliasesApi.GetAlias(ctx, workspace, alias).Execute()
		if err == nil {
			return true, nil
		}

		var re Error
		if AsError(err, &re) {
			if re.IsNotFoundError() {
				return false, nil
			}
		}

		return false, err
	})
}

// WaitUntilCollectionReady waits until the collection is ready.
func (rc *RockClient) WaitUntilCollectionReady(ctx context.Context, workspace, name string) error {
	return rc.Retry(ctx, rc.collectionHasState(ctx, workspace, name, collectionStatusREADY))
}

// WaitUntilCollectionGone waits until the a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (rc *RockClient) WaitUntilCollectionGone(ctx context.Context, workspace, name string) error {
	return rc.Retry(ctx, rc.collectionIsGone(ctx, workspace, name))
}

// WaitUntilCollectionDocuments waits until the collection has at least count new documents
func (rc *RockClient) WaitUntilCollectionDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.Retry(ctx, waiter.collectionHasNewDocs(ctx, workspace, name, count))
}

// RetryFunc is the function Retrier will continually call until it either returns true or an error.
type RetryFunc func() (retry bool, err error)

// Retrier is the interface used by the RockClient convenience methods to retry a noperation
// which returned a rockset.Error which is Retryable()
type Retrier interface {
	Retry(ctx context.Context, retryFn RetryFunc) error
}

// ExponentialRetry is used to perform retries with exponential backoff
type ExponentialRetry struct {
	// MaxBackoff is the max time the exponential backoff can wait
	MaxBackoff time.Duration
	// WaitInterval is the initial interval wait between consecutive calls
	WaitInterval time.Duration
}

// Retry retries retryFn until it returns false, or an error. Uses exponential backoff.
func (r ExponentialRetry) Retry(ctx context.Context, retryFn RetryFunc) error {
	t0 := time.Now()
	log := zerolog.Ctx(ctx)

	var maxBackoff = 8 * time.Second
	if r.MaxBackoff != 0 {
		maxBackoff = r.MaxBackoff
	}
	waitInterval := time.Second
	if r.WaitInterval != 0 {
		waitInterval = r.WaitInterval
	}

	defer func() {
		log.Debug().Str("d", time.Since(t0).String()).Msg("retry duration")
	}()

	for {
		t1 := time.Now()
		retry, err := retryFn()
		log.Debug().Str("d", time.Since(t1).String()).Msg("call curation")
		if err != nil {
			return err
		}
		if !retry {
			return nil
		}

		t := time.NewTimer(waitInterval)
		select {
		case <-ctx.Done():
			log.Debug().Msg("wait cancelled")
			t.Stop()

			return ctx.Err()
		case t := <-t.C:
			log.Trace().Str("t", t.String()).Msg("wait time")
			if waitInterval < maxBackoff {
				waitInterval *= 2
			}
		}
	}
}

func (rc *RockClient) collectionIsGone(ctx context.Context, workspace, name string) RetryFunc {
	return func() (bool, error) {
		_, err := rc.GetCollection(ctx, workspace, name)

		if err == nil {
			return false, nil
		}

		var re Error
		if AsError(err, &re) {
			// the state "GONE" is a special case, which is when the collection is deleted and returns a 404
			if re.IsNotFoundError() {
				return false, nil
			}
		}

		return true, err
	}
}

func (rc *RockClient) collectionHasState(ctx context.Context, workspace, name, state string) RetryFunc {
	return func() (bool, error) {
		log := zerolog.Ctx(ctx)

		c, err := rc.GetCollection(ctx, workspace, name)
		if err != nil {
			return false, err
		}

		log.Debug().Str("status", c.GetStatus()).Str("workspace", workspace).
			Str("collection", name).Msg("collectionHasState()")
		if c.GetStatus() == state {
			return false, nil
		}

		return true, nil
	}
}

type docWaiter struct {
	rc        *RockClient
	prevCount *int64
}

func (d *docWaiter) collectionHasNewDocs(ctx context.Context, workspace, name string, count int64) RetryFunc {
	return func() (bool, error) {
		log := zerolog.Ctx(ctx)

		c, err := d.rc.GetCollection(ctx, workspace, name)
		if err != nil {
			return false, err
		}

		current := c.Stats.DocCount
		log.Debug().Str("workspace", workspace).Int64("current", *current).
			Int64("previous", *d.prevCount).Str("collection", name).Msg("collectionHasNewDocs()")

		if d.prevCount == nil {
			d.prevCount = current
		}

		// log.Printf("%d - %d > %d", *current, *d.prevCount, count)
		if *current-*d.prevCount >= count {
			return false, nil
		}

		return true, nil
	}
}
