package rockset_test

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/dataset"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
	"time"
)

const (
	rocksetCircleCIMainVI = "rrn:vi:usw2a1:29e4a43c-fff4-4fe6-80e3-1ee57bc22e82"
	rocksetTestMainVI     = "rrn:vi:usw2a1:b0f23396-a8c9-45c5-9452-00c1267094a4"
)

type VirtualInstanceSuite struct {
	suite.Suite
	rc  *rockset.RockClient
	vID string
}

func TestVirtualInstance(t *testing.T) {
	skipUnlessIntegrationTest(t)

	s := VirtualInstanceSuite{
		rc:  testClient(t),
		vID: rocksetCircleCIMainVI,
	}
	suite.Run(t, &s)
}

func (s *VirtualInstanceSuite) TestListQueries() {
	ctx := testCtx()

	queries, err := s.rc.GetVirtualInstanceQueries(ctx, s.vID)
	s.Require().NoError(err)
	for _, q := range queries {
		s.T().Logf("%s: %s", q.GetQueryId(), q.GetStatus())
	}
}

func (s *VirtualInstanceSuite) TestGetCollectionMount() {
	ctx := testCtx()

	mount, err := s.rc.GetCollectionMount(ctx, s.vID, "commons.movie_releases")
	s.Require().NoError(err)
	s.Assert().Equal(strings.TrimPrefix(s.vID, "rrn:vi:usw2a1:"), mount.GetVirtualInstanceId())
}

func (s *VirtualInstanceSuite) TestListCollectionMounts() {
	ctx := testCtx()

	mounts, err := s.rc.ListCollectionMounts(ctx, s.vID)
	s.Require().NoError(err)
	s.Assert().NotEmpty(mounts)
}

func (s *VirtualInstanceSuite) TestExecuteQuery() {
	ctx := testCtx()

	result, err := s.rc.ExecuteQueryOnVirtualInstance(ctx, s.vID,
		"SELECT * from commons._events LIMIT 10",
		option.WithWarnings())
	s.Require().NoError(err)
	s.Assert().Contains(result.Collections, "commons._events")
}

func (s *VirtualInstanceSuite) TestGetVirtualInstance() {
	ctx := testCtx()

	// make sure we can access the VI using both the plain id and the RRN
	rocksetTestVIs := []string{
		strings.TrimPrefix(s.vID, "rrn:vi:usw2a1:"),
		s.vID,
	}

	for _, vi := range rocksetTestVIs {
		s.Run(vi, func() {
			main, err := s.rc.GetVirtualInstance(ctx, vi)
			s.Assert().NoError(err)
			s.Assert().Equal("main", main.GetName())
		})
	}
}

func (s *VirtualInstanceSuite) TestListVirtualInstances() {
	ctx := testCtx()

	vis, err := s.rc.ListVirtualInstances(ctx)
	s.Require().NoError(err)
	s.Assert().NotEmpty(vis)
}

type VirtualInstanceIntegrationSuite struct {
	suite.Suite
	rc         *rockset.RockClient
	vID        string
	name       string
	workspace  string
	collection string
}

// This test suite requires a Rockset org with "premium" support, as it needs to create a virtual instance.
func TestVirtualInstanceIntegration(t *testing.T) {
	skipUnlessIntegrationTest(t)
	t.Log("This test can take more than 10 minutes to run, so make sure to use `go test -timeout 15m`")

	s := VirtualInstanceIntegrationSuite{
		rc:   testClient(t),
		name: randomName("vi"),
	}
	suite.Run(t, &s)
}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance() {
	ctx := testCtx()

	vi, err := s.rc.CreateVirtualInstance(ctx, s.name, option.WithVirtualInstanceSize(option.SizeSmall),
		option.WithContinuousMountRefresh(), option.WithAutoSuspend(time.Hour))
	s.Require().NoError(err)
	s.vID = vi.GetId()
	s.T().Logf("vi %s is created", s.vID)

	err = s.rc.WaitUntilVirtualInstanceActive(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("vi %s is active", s.vID)

	ws, err := s.rc.CreateWorkspace(ctx, s.name)
	s.Require().NoError(err)
	s.workspace = ws.GetName()
	s.T().Logf("workspace %s is created", s.workspace)

	err = s.rc.WaitUntilWorkspaceAvailable(ctx, s.workspace)
	s.Require().NoError(err)
	s.T().Logf("workspace %s is available", s.workspace)

	c, err := s.rc.CreateCollection(ctx, s.workspace, s.name, option.WithSampleDataset(dataset.Movies))
	s.Require().NoError(err)
	s.collection = c.GetName()
	s.T().Logf("collection %s.%s is created", s.workspace, s.collection)

	err = s.rc.WaitUntilCollectionHasDocuments(ctx, s.workspace, s.collection, 2_830)
	s.Require().NoError(err)
	s.T().Log("collection has documents")

	path := s.workspace + "." + s.collection
	_, err = s.rc.MountCollection(ctx, s.vID, []string{path})
	s.Require().NoError(err)
	s.T().Logf("collection %s.%s is mounted on %s", s.workspace, s.collection, s.vID)

	err = s.rc.WaitUntilCollectionReady(ctx, s.workspace, s.collection)
	s.Require().NoError(err)
	s.T().Logf("collection %s.%s is ready", s.workspace, s.collection)

	query := fmt.Sprintf("SELECT * FROM %s", path)
	_, err = s.rc.Query(ctx, query)
	s.Require().NoError(err)
	s.T().Log("query ran")

	_, err = s.rc.SuspendVirtualInstance(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Log("vi is suspending")

	err = s.rc.WaitUntilVirtualInstanceSuspended(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Log("vi is suspended")

	_, err = s.rc.ResumeVirtualInstance(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Log("vi is resuming")

	err = s.rc.WaitUntilVirtualInstanceActive(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Log("vi is active")

	_, err = s.rc.Query(ctx, query)
	s.Require().NoError(err)
	s.T().Log("query after resume ran")
}

func (s *VirtualInstanceIntegrationSuite) TearDownSuite() {
	ctx := testCtx()

	err := s.rc.DeleteCollection(ctx, s.workspace, s.collection)
	s.Assert().NoError(err)

	err = s.rc.WaitUntilCollectionGone(ctx, s.workspace, s.collection)
	s.Assert().NoError(err)

	err = s.rc.DeleteWorkspace(ctx, s.workspace)
	s.Assert().NoError(err)

	err = s.rc.WaitUntilWorkspaceGone(ctx, s.workspace)
	s.Assert().NoError(err)

	_, err = s.rc.DeleteVirtualInstance(ctx, s.vID)
	s.Assert().NoError(err)
}
