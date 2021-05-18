package rockset

import (
	"context"
	"log"
	"time"
)

const (
	collectionStatusREADY = "READY"
)

// WaitUntilAliasAvailable waits until the alias is available.
func (rc *RockClient) WaitUntilAliasAvailable(ctx context.Context, workspace, alias string) error {
	return rc.WaitUntil(ctx, func() (bool, error) {
		_, _, err := rc.AliasesApi.GetAlias(ctx, workspace, alias).Execute()
		if err == nil {
			return true, nil
		}

		if _, ok := IsNotFoundError(err); ok {
			return false, nil
		}

		return false, err
	})
}

// WaitUntilCollectionReady waits until the collection is ready.
func (rc *RockClient) WaitUntilCollectionReady(ctx context.Context, workspace, name string) error {
	return rc.WaitUntil(ctx, rc.collectionHasState(ctx, workspace, name, collectionStatusREADY))
}

// WaitUntilCollectionGone waits until the a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (rc *RockClient) WaitUntilCollectionGone(ctx context.Context, workspace, name string) error {
	return rc.WaitUntil(ctx, rc.collectionIsGone(ctx, workspace, name))
}

// WaitUntilCollectionDocuments waits until the collection has at least count new documents
func (rc *RockClient) WaitUntilCollectionDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.WaitUntil(ctx, waiter.collectionHasNewDocs(ctx, workspace, name, count))
}

// WaitFunc is the function WaitUntil will continually call until it either returns true or an error.
type WaitFunc func() (bool, error)

// WaitUntil waits until the waitFn returns true, or an error. Uses exponential backoff.
func (rc *RockClient) WaitUntil(ctx context.Context, wait WaitFunc) error {
	const maxBackoff = 8 * time.Second
	waitInterval := time.Second
	t0 := time.Now()
	defer func() {
		log.Printf("wait time: %s", time.Since(t0).String())
	}()

	for {
		done, err := wait()
		if err != nil {
			return err
		}
		if done {
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

func (rc *RockClient) collectionIsGone(ctx context.Context, workspace, name string) WaitFunc {
	return func() (bool, error) {
		_, err := rc.GetCollection(ctx, workspace, name)
		if err == nil {
			return false, nil
		}
		// the state "GONE" is a special case, which is when the collection is deleted and returns a 404
		if t, ok := IsNotFoundError(err); ok {
			if t.Error() == "404 Not Found" {
				log.Printf("GetCollection() returned %s which means it is gone", t.Error())

				return true, nil
			}
		}

		return false, err
	}
}

func (rc *RockClient) collectionHasState(ctx context.Context, workspace, name, state string) WaitFunc {
	return func() (bool, error) {
		c, err := rc.GetCollection(ctx, workspace, name)
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
	rc        *RockClient
	prevCount *int64
}

func (d *docWaiter) collectionHasNewDocs(ctx context.Context, workspace, name string, count int64) WaitFunc {
	return func() (bool, error) {
		c, err := d.rc.GetCollection(ctx, workspace, name)
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
