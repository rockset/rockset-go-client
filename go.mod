module github.com/rockset/rockset-go-client

go 1.16

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/rockset/rockset-go-client/openapi v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20210505024714-0287a6fb4125 // indirect
	golang.org/x/oauth2 v0.0.0-20210427180440-81ed05c6b58c // indirect
	google.golang.org/appengine v1.6.7 // indirect
)

replace (
	github.com/rockset/rockset-go-client/openapi => ./openapi
	github.com/rockset/rockset-go-client/option => ./option
)
