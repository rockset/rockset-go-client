# rockset-go-client
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

Usage
=====
You can see a few [sample examples](https://github.com/rockset/rockset-go-client/tree/master/examples) of how to create a collection, how to put documents in a collection and how to use SQL to query your collections.

Testing
=======

Unit tests are available in the [Test](https://github.com/rockset/rockset-go-client/tree/master/test) folder.

Run unit tests:
```
go test ./test
```

## Support

Feel free to log issues against this client through GitHub.

## License

## License

The Rockset Go Client is licensed under the [MIT License](https://github.com/rockset/rockset-go-client/blob/master/LICENSE)
