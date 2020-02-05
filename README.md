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

You can see a few [sample examples](https://github.com/rockset/rockset-go-client/tree/master/examples) of how to create a collection, how to put documents in a collection and how to use SQL to query your collections.

## Testing

Tests are available in the [Test](https://github.com/rockset/rockset-go-client/tree/master/test) folder.

Set ROCKSET_APIKEY and ROCKSET_APISERVER endpoint in the environment variables. To run tests:
```
go test ./test
```

## Support

Feel free to log issues against this client through GitHub.

## License

The Rockset Go Client is licensed under the [Apache 2.0 License](https://github.com/rockset/rockset-go-client/blob/master/LICENSE)
