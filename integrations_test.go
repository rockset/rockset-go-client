package rockset_test

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (s *IntegrationsSuite) TestCreateAzureBlob() {
	name := "azureBlobTest"
	connectionString := ""
	ctx := testCtx()

	integrationRequest := openapi.ApiCreateIntegrationRequest{
		ApiService: s.IntegrationsApi,
	}
	err := InjectCtx(ctx, &integrationRequest)
	s.Require().NoError(err)
	integrationRequest.ApiService = s.IntegrationsApi

	integrationResponse := openapi.CreateIntegrationResponse{}

	integration := openapi.Integration{
		Name: name,
	}

	execute := s.resp.On("Execute")
	execute.Return(integration)

	createIntegration := s.IntegrationsApi.On("CreateIntegration", mock.Anything)
	createIntegration.Return(integrationRequest)

	rc, err := rockset.NewClient()
	s.Require().NoError(err)
	rc.IntegrationsApi = s.IntegrationsApi

	createExecute := s.IntegrationsApi.On("CreateIntegrationExecute", mock.Anything)
	createExecute.Return(&integrationResponse, new(http.Response), nil)

	resp, err := rc.CreateAzureBlobStorageIntegration(ctx, name, connectionString)
	s.Require().NoError(err)
	s.Require().True(resp.HasAzureBlobStorage())
	actual := resp.GetAzureBlobStorage()

	s.Equal(name, resp.GetName())
	s.Equal(connectionString, actual.GetConnectionString())
}

func TestRockClient_S3Integration(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	name := "s3test"

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	// get the integration
	getReq := rc.IntegrationsApi.GetIntegration(ctx, name)
	_, _, err = getReq.Execute()
	if err != nil {
		// check if it is missing
		var re rockset.Error
		if errors.As(err, &re) {
			if !re.IsNotFoundError() {
				require.NoError(t, err)
			}
		}
	} else {
		// the integration exists, delete it
		deleteReq := rc.IntegrationsApi.DeleteIntegration(ctx, name)
		_, _, err = deleteReq.Execute()
		require.NoError(t, err)
	}

	// create a new integration
	_, err = rc.CreateS3Integration(ctx, name,
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription("test"))
	require.NoError(t, err)

	// list integrations and look for the newly created integration
	listReq := rc.IntegrationsApi.ListIntegrations(ctx)
	listResp, _, err := listReq.Execute()
	require.NoError(t, err)

	var found bool
	for _, i := range listResp.GetData() {
		if i.Name == name {
			found = true
		}
	}
	if !found {
		t.Errorf("could not find %s", name)
	}

	// delete
	deleteReq := rc.IntegrationsApi.DeleteIntegration(ctx, name)
	_, _, err = deleteReq.Execute()
	require.NoError(t, err)
}

func TestRockClient_CreateGCSIntegration(t *testing.T) {
	skipUnlessIntegrationTest(t)

	saKeyFile := os.Getenv("GCP_SA_KEY_JSON")
	if saKeyFile == "" {
		t.Skip("environment variable GCP_SA_KEY_JSON set")
	}

	ctx := testCtx()
	rc, err := rockset.NewClient()
	require.NoError(t, err)

	iName := "testGCSIntegration"
	gcs, err := rc.GetIntegration(ctx, iName)
	if err == nil {
		err = rc.DeleteIntegration(ctx, iName)
		require.NoError(t, err)
		log.Printf("collection deleted: %s", iName)
	} else {
		var re rockset.Error
		if !errors.As(err, &re) {
			t.Fatalf("GetIntegration failed for %s: %v", iName, err)
		}
		if !re.IsNotFoundError() {
			t.Fatalf("GetIntegration failed for %s: %v", iName, re)
		}
	}

	gcs, err = rc.CreateGCSIntegration(ctx, iName, saKeyFile,
		option.WithGCSIntegrationDescription("created by rockset integration tests"))
	require.NoError(t, err)

	log.Printf("created gcs integration: %s", gcs.GetName())
}

type IntegrationsSuite struct {
	suite.Suite
	IntegrationsApi MockIntegrationsAPI

	resp MockIntegrationResponse
}

func (s *IntegrationsSuite) SetupSuite() {
	s.IntegrationsApi = MockIntegrationsAPI{}
	s.resp = MockIntegrationResponse{}

	os.Setenv(APIKeyEnvironmentVariableName, "api key")
	os.Setenv(APIServerEnvironmentVariableName, "https://null.nothing")
}

func (s *IntegrationsSuite) TearDownSuite() {
	os.Setenv(APIKeyEnvironmentVariableName, "")
	os.Setenv(APIServerEnvironmentVariableName, "")
}

func TestIntegrations(t *testing.T) {
	s := new(IntegrationsSuite)
	suite.Run(t, s)
}

type MockIntegrationsAPI struct {
	mock.Mock
	openapi.IntegrationsApi
}

func (m MockIntegrationsAPI) CreateIntegration(ctx context.Context) openapi.ApiCreateIntegrationRequest {
	args := m.Called(ctx)
	return args.Get(0).(openapi.ApiCreateIntegrationRequest)
}

func (m MockIntegrationsAPI) CreateIntegrationExecute(r openapi.ApiCreateIntegrationRequest) (*openapi.CreateIntegrationResponse, *http.Response, error) {
	args := m.Called()
	return args.Get(0).(*openapi.CreateIntegrationResponse), args.Get(1).(*http.Response), args.Error(2)
}

type MockIntegrationRequest struct {
	mock.Mock
	openapi.ApiGetIntegrationRequest
}

func (m MockIntegrationRequest) Execute() (*openapi.GetIntegrationResponse, *http.Response, error) {
	args := m.Called()
	return args.Get(0).(*openapi.GetIntegrationResponse), args.Get(1).(*http.Response), args.Error(2)
}

type MockIntegrationResponse struct {
	mock.Mock
	openapi.GetIntegrationResponse
}

//InjectCtx is necessary because the openapi.ApiCreateIntegrationRequest
//keeps a copy of a context in an unexported field `ctx`. Since we are
//mocking the call to create it, we cannot put the context in (our
//mocks are not in the same package).
//This allows us to set that context.
//nolint
func InjectCtx(ctx context.Context, subject *openapi.ApiCreateIntegrationRequest) error {
	sv := reflect.ValueOf(subject).Elem()
	cf := sv.FieldByName("ctx")
	cfp := cf.UnsafeAddr()
	nf := reflect.NewAt(cf.Type(), unsafe.Pointer(cfp)).Elem()
	nf.Set(reflect.ValueOf(ctx))

	return nil
}
