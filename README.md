# Rockset Go Client
[![CircleCI](https://circleci.com/gh/rockset/rockset-go-client/tree/go-client-v0.8.svg?style=shield)](https://circleci.com/gh/rockset/rockset-go-client/tree/go-client-v0.8)
[![Documentation](https://godoc.org/github.com/rockset/rockset-go-client?status.svg)](http://godoc.org/github.com/rockset/rockset-go-client)
[![License](https://img.shields.io/github/license/rockset/rockset-go-client.svg?maxAge=2592000)](https://github.com/rockset/rockset-go-client/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/rockset/rockset-go-client.svg)](https://github.com/rockset/rockset-go-client/issues)
[![Release](https://img.shields.io/github/release/rockset/rockset-go-client.svg?label=Release)](https://github.com/rockset/rockset-go-client/releases)

Official Go client library for Rockset

## Installation

Install the Rockset Go client from Github:

```
go get github.com/rockset/rockset-go-client
```

or install it from a source code checkout:

```
cd $GOPATH/src/github.com
mkdir rockset
cd rockset
git clone git@github.com:rockset/rockset-go-client.git
go install rockset-go-client/rockclient.go
```

## Usage

You can see a few [examples](https://godoc.org/github.com/rockset/rockset-go-client/#pkg-overview) in the godoc
of how to create a collection, how to put documents in a collection and how to use SQL to query your collections.

```
client, err := rockset.NewClient(rockset.WithAPIKey("..."))
if err != nil {
    log.Fatal(err)
}

q := models.QueryRequest{
    Sql: &models.QueryRequestSql{
        Query: `SELECT * FROM "_events" LIMIT 1`,
    },
}

res, _, err := client.Query(q)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", res)
```

## Testing

Tests are available in the [test](https://github.com/rockset/rockset-go-client/tree/master/test) folder.

Set ROCKSET_APIKEY and ROCKSET_APISERVER endpoint in the environment variables. To run tests:
```
go test ./...
```

### Code Coverage

```
go test ./... -coverprofile cover.out
go tool cover -func=cover.out
go tool cover -html=cover.out -o coverage.html
```

## Support

Feel free to log issues against this client through GitHub.

## License

The Rockset Go Client is licensed under the [Apache 2.0 License](https://github.com/rockset/rockset-go-client/blob/master/LICENSE)
