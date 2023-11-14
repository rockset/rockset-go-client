package fake

import "github.com/rockset/rockset-go-client/openapi"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// TODO generate the below list automatically
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.APIKeysApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.AliasesApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.CollectionsApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.CustomRolesApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.DocumentsApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.IntegrationsApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.OrganizationsApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.QueriesApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.QueryLambdasApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.SharedLambdasApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.SourcesApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.UsersApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.ViewsApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.VirtualInstancesApi
//counterfeiter:generate -o . github.com/rockset/rockset-go-client/openapi.WorkspacesApi

// NewAPIClient returns a fake API Client that can be passed to rockset.NewRockClient()
//
//	f := fake.NewAPIClient()
//	rc, err := rockset.NewClient(rockset.WithAPIClient(f))
func NewAPIClient() *openapi.APIClient {
	// TODO generate this dynamically
	return &openapi.APIClient{
		APIKeysApi:          &FakeAPIKeysApi{},
		AliasesApi:          &FakeAliasesApi{},
		CollectionsApi:      &FakeCollectionsApi{},
		CustomRolesApi:      &FakeCustomRolesApi{},
		DocumentsApi:        &FakeDocumentsApi{},
		IntegrationsApi:     &FakeIntegrationsApi{},
		OrganizationsApi:    &FakeOrganizationsApi{},
		QueriesApi:          &FakeQueriesApi{},
		QueryLambdasApi:     &FakeQueryLambdasApi{},
		SharedLambdasApi:    &FakeSharedLambdasApi{},
		SourcesApi:          &FakeSourcesApi{},
		UsersApi:            &FakeUsersApi{},
		ViewsApi:            &FakeViewsApi{},
		VirtualInstancesApi: &FakeVirtualInstancesApi{},
		WorkspacesApi:       &FakeWorkspacesApi{},
	}
}
