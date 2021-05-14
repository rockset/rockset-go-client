package rockset

import (
	"context"
	"log"
	"time"

	"github.com/rockset/rockset-go-client/openapi"
)

const (
	collectionStatusREADY = "READY"
)

// WaitForCollectionReady waits until the collection is ready.
func (rc *RockClient) WaitForCollectionReady(ctx context.Context, workspace, name string) error {
	return rc.waitForCollection(ctx, workspace, name, hasState(collectionStatusREADY))
}

// WaitForCollectionGone waits until the a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (rc *RockClient) WaitForCollectionGone(ctx context.Context, workspace, name string) error {
	return rc.waitForCollection(ctx, workspace, name, isGone())
}

// WaitForCollectionDocuments waits until the collection has at least count new documents
func (rc *RockClient) WaitForCollectionDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{}
	return rc.waitForCollection(ctx, workspace, name, waiter.hasNewDocs(count))
}

type waitFunc func(c openapi.Collection, err error) (bool, error)

// WaitForCollection waits until the collection reaches the desired state, determined by the waitFunc.
// Uses exponential backoff.
func (rc *RockClient) waitForCollection(ctx context.Context, workspace, name string, waitFn waitFunc) error {
	const maxBackoff = 8 * time.Second
	waitInterval := time.Second
	t0 := time.Now()
	defer func() {
		log.Printf("wait time: %s", time.Since(t0).String())
	}()

	for {
		if done, err := waitFn(rc.GetCollection(ctx, workspace, name)); err != nil {
			return err
		} else if done {
			return nil
		}

		t := time.NewTimer(waitInterval)
		select {
		case <-ctx.Done():
			log.Printf("wait cancelled")
			t.Stop()

			return ctx.Err()
		case t := <-t.C:
			log.Printf("wait time %s", t)
			if waitInterval < maxBackoff {
				waitInterval *= 2
			}
		}
	}
}

func isGone() waitFunc {
	return func(c openapi.Collection, err error) (bool, error) {
		if err != nil {
			// the state "GONE" is a special case, which is when the collection is deleted and returns a 404
			if t, ok := err.(openapi.GenericOpenAPIError); ok {
				if t.Error() == "404 Not Found" {
					log.Printf("GetCollection() returned %s which means it is gone", t.Error())

					return true, nil
				}
			}

			return false, err
		}

		return false, nil
	}
}

func hasState(state string) waitFunc {
	return func(c openapi.Collection, err error) (bool, error) {
		if err != nil {
			return false, err
		}
		if c.GetStatus() == state {
			return true, nil
		}
		return false, nil
	}
}

type docWaiter struct {
	prevCount *int64
}

func (d *docWaiter) hasNewDocs(count int64) waitFunc {
	return func(c openapi.Collection, err error) (bool, error) {
		if err != nil {
			return false, err
		}
		current := c.Stats.DocCount
		if d.prevCount == nil {
			d.prevCount = current
		}

		// log.Printf("%d - %d > %d", *current, *d.prevCount, count)
		if *current-*d.prevCount >= count {
			return true, nil
		}

		return false, nil
	}
}
