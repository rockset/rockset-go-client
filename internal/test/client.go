package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func Client(t *testing.T) *rockset.RockClient {
	rc, err := rockset.NewClient(rockset.WithUserAgent("rockset-go-integration-tests"))
	require.NoError(t, err)
	return rc
}

func DebugClient(t *testing.T) *rockset.RockClient {
	rc, err := rockset.NewClient(rockset.WithUserAgent("rockset-go-integration-tests"), rockset.WithHTTPDebug())
	require.NoError(t, err)
	return rc
}
