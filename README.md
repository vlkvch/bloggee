# bloggee

bloggee: a blog engine.

## Usage

If you would like to write a post, simply create a directory under `./blog` (the directory's name will be your posts's ID) and place a file named `index.md` there.

## Building from source

### Generating a TLS certificate

First of all, you have to generate a self-signed TLS certificate under `./tls`. To do so, you can use `generate_cert.go` from the `crypto/tls` package (depending on your system, the location may differ):

```
$ mkdir tls
$ cd tls
$ go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

### Running the app

To run the app straight away, you can use Task:

```
$ task run
```

If you'd like to change the directory for your posts:

```
$ task run -- -dir ./posts
```

To simply build a binary file:

```
$ task
```

## License

MIT
