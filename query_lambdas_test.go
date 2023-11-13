package rockset_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
)

type QueryLambdaSuite struct {
	suite.Suite
	rc      *rockset.RockClient
	tag     string
	name    string
	ws      string
	version string
}

func TestQueryLambdaSuite(t *testing.T) {
	rc, nameGenerator := vcrTestClient(t, t.Name())

	s := QueryLambdaSuite{
		rc:   rc,
		tag:  nameGenerator("tag"),
		name: nameGenerator("ql"),
		ws:   "commons",
	}
	suite.Run(t, &s)
}

func (s *QueryLambdaSuite) TearDownSuite() {
	ctx := test.Context()

	//make sure we clean up
	_ = s.rc.DeleteQueryLambda(ctx, s.ws, s.name)
}

func (s *QueryLambdaSuite) Test_0_Create() {
	ctx := test.Context()

	ql, err := s.rc.CreateQueryLambda(ctx, s.ws, s.name, "SELECT 1",
		option.WithQueryLambdaDescription("create"))
	s.Require().NoError(err)
	s.Equal(s.name, ql.GetName())
	s.version = ql.GetVersion()

	w := wait.New(s.rc)
	err = w.UntilQueryLambdaVersionActive(ctx, s.ws, s.name, s.version)
	s.Require().NoError(err)
}

func (s *QueryLambdaSuite) Test_1_CreateTag() {
	ctx := test.Context()

	tag, err := s.rc.CreateQueryLambdaTag(ctx, s.ws, s.name, s.version, s.tag)
	s.Require().NoError(err)
	s.Equal(s.tag, tag.GetTagName())
	v := tag.GetVersion()
	s.Equal(s.version, v.GetVersion())

	w := wait.New(s.rc)
	w.QueryLambdaTagPropagation = time.Microsecond
	err = w.UntilQueryLambdaTagPropagated(ctx, s.ws, s.name, s.tag)
	s.Require().NoError(err)
}

func (s *QueryLambdaSuite) Test_2_GetVersion() {
	ctx := test.Context()

	ql, err := s.rc.GetQueryLambdaVersion(ctx, s.ws, s.name, s.version)
	s.Require().NoError(err)
	s.Equal(s.name, ql.GetName())
}

func (s *QueryLambdaSuite) Test_3_GetTag() {
	ctx := test.Context()

	ql, err := s.rc.GetQueryLambdaVersionByTag(ctx, s.ws, s.name, s.tag)
	s.Require().NoError(err)
	v := ql.GetVersion()
	s.Equal(s.version, v.GetVersion())
}

func (s *QueryLambdaSuite) Test_4_List() {
	ctx := test.Context()

	qls, err := s.rc.ListQueryLambdas(ctx)
	s.Require().NoError(err)
	s.NotEmpty(qls)

	qls, err = s.rc.ListQueryLambdas(ctx, option.WithQueryLambdaWorkspace(s.ws))
	s.Require().NoError(err)
	s.NotEmpty(qls)

	tags, err := s.rc.ListQueryLambdaTags(ctx, s.ws, s.name)
	s.Require().NoError(err)
	s.NotEmpty(tags)

	versions, err := s.rc.ListQueryLambdaVersions(ctx, s.ws, s.name)
	s.Require().NoError(err)
	s.NotEmpty(versions)
}

func (s *QueryLambdaSuite) Test_5_Update() {
	ctx := test.Context()

	ql, err := s.rc.UpdateQueryLambda(ctx, s.ws, s.name, "SELECT 2", option.WithQueryLambdaDescription("updated"))
	s.Require().NoError(err)
	s.Equal("updated", ql.GetDescription())
	s.version = ql.GetVersion()
}

func (s *QueryLambdaSuite) Test_6_Execute() {
	ctx := test.Context()

	_, err := s.rc.ExecuteQueryLambda(ctx, s.ws, s.name, option.WithQueryLambdaRowLimit(100))
	s.NoError(err)

	_, err = s.rc.ExecuteQueryLambda(ctx, s.ws, s.name, option.WithTag(s.tag))
	s.NoError(err)

	_, err = s.rc.ExecuteQueryLambda(ctx, s.ws, s.name, option.WithVersion(s.version))
	s.NoError(err)
}

func (s *QueryLambdaSuite) Test_9_Delete() {
	ctx := test.Context()

	err := s.rc.DeleteQueryLambdaTag(ctx, s.ws, s.name, s.tag)
	s.NoError(err)

	err = s.rc.DeleteQueryLambdaVersion(ctx, s.ws, s.name, s.version)
	s.NoError(err)

	err = s.rc.DeleteQueryLambda(ctx, s.ws, s.name)
	s.NoError(err)
}
