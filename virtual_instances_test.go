package rockset_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/dataset"
	"github.com/rockset/rockset-go-client/option"
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
	rc, _ := vcrTestClient(t, t.Name())

	s := VirtualInstanceSuite{
		rc:  rc,
		vID: rocksetCircleCIMainVI,
	}
	suite.Run(t, &s)
}

func (s *VirtualInstanceSuite) TestListQueries() {
	ctx := testCtx()

	queries, err := s.rc.ListVirtualInstanceQueries(ctx, s.vID)
	s.Require().NoError(err)
	for _, q := range queries {
		s.NotEmpty(q.GetQueryId())
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
		"SELECT * from commons._events LIMIT 10")
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
	path       string
}

// This test suite requires a Rockset org with "premium" support, as it needs to create a virtual instance.
func TestVirtualInstanceIntegration(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())

	name := randomName("vi")
	s := VirtualInstanceIntegrationSuite{
		rc:   rc,
		name: name,
		path: name + "." + name,
	}
	suite.Run(t, &s)
}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance_0_Create() {
	ctx := testCtx()
	t0 := time.Now()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	vi, err := rc.CreateVirtualInstance(ctx, s.name,
		option.WithVirtualInstanceSize(option.SizeSmall),
		option.WithNoMountRefresh(),
		option.WithAutoSuspend(15*time.Minute),
	)
	s.Require().NoError(err)
	s.vID = vi.GetId()
	s.T().Logf("vi %s / %s is created (%s)", vi.GetId(), vi.GetName(), time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilVirtualInstanceActive(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("vi %s is active (%s)", s.vID, time.Since(t0))
	t0 = time.Now()
}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance_1_Collection() {
	ctx := testCtx()
	t0 := time.Now()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	ws, err := rc.CreateWorkspace(ctx, s.name)
	s.Require().NoError(err)
	s.workspace = ws.GetName()
	s.T().Logf("workspace %s is created (%s)", s.workspace, time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilWorkspaceAvailable(ctx, s.workspace)
	s.Require().NoError(err)
	s.T().Logf("workspace %s is available (%s)", s.workspace, time.Since(t0))
	t0 = time.Now()

	c, err := rc.CreateCollection(ctx, s.workspace, s.name, option.WithSampleDataset(dataset.Movies))
	s.Require().NoError(err)
	s.collection = c.GetName()
	s.T().Logf("collection %s.%s is created (%s)", s.workspace, s.collection, time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilCollectionHasDocuments(ctx, s.workspace, s.collection, 2_830)
	s.Require().NoError(err)
	s.T().Log("collection has documents", time.Since(t0))
	t0 = time.Now()

}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance_2_Mount() {
	ctx := testCtx()
	t0 := time.Now()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	_, err := rc.MountCollections(ctx, s.vID, []string{s.path})
	s.Require().NoError(err)
	s.T().Logf("collection %s.%s is mounted on %s (%s)", s.workspace, s.collection, s.vID, time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilCollectionReady(ctx, s.workspace, s.collection)
	s.Require().NoError(err)
	s.T().Logf("collection %s.%s is ready (%s)", s.workspace, s.collection, time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilMountActive(ctx, s.vID, s.workspace, s.collection)
	s.Require().NoError(err)
	s.T().Logf("mount is active: %s", time.Since(t0))
	t0 = time.Now()
}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance_3_Query() {
	ctx := testCtx()
	t0 := time.Now()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	query := fmt.Sprintf("SELECT * FROM %s.%s", s.workspace, s.collection)
	_, err := rc.ExecuteQueryOnVirtualInstance(ctx, s.vID, query)
	s.Require().NoError(err)
	s.T().Logf("query ran (%s)", time.Since(t0))
	t0 = time.Now()

	queries, err := rc.ListVirtualInstanceQueries(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("listed queries (%s)", time.Since(t0))
	s.NotEmpty(queries)
	t0 = time.Now()
}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance_4_Unmount() {
	ctx := testCtx()
	t0 := time.Now()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	mounts, err := rc.ListCollectionMounts(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("listed mounts (%s)", time.Since(t0))
	s.NotEmpty(mounts)
	t0 = time.Now()

	_, err = rc.UnmountCollection(ctx, s.vID, s.path)
	s.Require().NoError(err)
	s.T().Logf("unmounted (%s)", time.Since(t0))
	t0 = time.Now()
}

func (s *VirtualInstanceIntegrationSuite) TestVirtualInstance_5_Suspend() {
	ctx := testCtx()
	t0 := time.Now()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	_, err := rc.SuspendVirtualInstance(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("vi is suspending (%s)", time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilVirtualInstanceSuspended(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("vi is suspended (%s)", time.Since(t0))
	t0 = time.Now()

	_, err = rc.ResumeVirtualInstance(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("vi is resuming (%s)", time.Since(t0))
	t0 = time.Now()

	err = rc.WaitUntilVirtualInstanceActive(ctx, s.vID)
	s.Require().NoError(err)
	s.T().Logf("vi is active (%s)", time.Since(t0))
	t0 = time.Now()

	_, err = rc.UpdateVirtualInstance(ctx, s.vID,
		option.WithVirtualInstanceName("new_"+s.name),
		option.WithVirtualInstanceDescription("new"),
	)
	s.Require().NoError(err)
	s.T().Logf("vi updated (%s)", time.Since(t0))
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
