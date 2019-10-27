# rust-grpc-web-wasm-test

## Currently supported

A go gRPC-web server with the built-in Improbable proxy. Run with:

```bash
$ go run main.go
```

There's one flag, `--html-root`, which can be used to specify
which directory to serve the frontend from. The Go gRPC backend
has been sprinkled with helpful log messages throughout, which should
hopefully aid in debugging any problems.

## Remaining work

1. Manually implement the Rust WASM client to call the gRPC-Web API.
   The following API calls are supported (here presented as JS Fetch calls):
   GetBook:
   ```js
   await fetch("https://localhost:8080/library.BookService/GetBook", {
    "credentials": "include",
    "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:69.0) Gecko/20100101 Firefox/69.0",
        "Accept": "*/*",
        "Accept-Language": "en-GB,en;q=0.5",
        "content-type": "application/grpc-web+proto", // Content-Type
        "x-grpc-web": "1",
    },
    "referrer": "https://localhost:8080/",
    "body": "\u0000\u0000\u0000\u0000\u0005\b½·áB", // binary protobuf body
    //       ^----^                                 // Compression flag, 1 byte
    //             ^----------------------^         // Size of body, 4 bytes
    //                                     ^----^   // Body
    "method": "POST",
    "mode": "cors"
   });
   ```
   QueryBooks (stream):
   ```js
   await fetch("https://localhost:8080/library.BookService/QueryBooks", {
    "credentials": "include",
    "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:69.0) Gecko/20100101 Firefox/69.0",
        "Accept": "*/*",
        "Accept-Language": "en-GB,en;q=0.5",
        "content-type": "application/grpc-web+proto",
        "x-grpc-web": "1",
    },
    "referrer": "https://localhost:8080/",
    "body": "\u0000\u0000\u0000\u0000\b\n\u0006George", // binary protobuf body
    //       ^----^                                     // Compression flag, 1 byte
    //             ^--------------------------^         // Size of body, 4 bytes
    //                                         ^----^   // Body
    "method": "POST",
    "mode": "cors"
   });
   ```
1. After successful manual unary and server streaming requests are handled,
move to implement Rust WASM protobuf generator.
1. Iterate on designs until ready to share.

## Further reading:

gRPC-Web protocol spec: https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md
Official JS gRPC-Web client (XHR based): https://github.com/grpc/grpc-web/tree/master/javascript/net/grpc/web
Improbable TypeScript gRPC-Web client (Fetch and XHR based): https://github.com/improbable-eng/grpc-web/tree/master/client/grpc-web/src
Dart gRPC-Web client (XHR based): https://github.com/grpc/grpc-dart/commit/91564ff7aa33214b55d050bed44be1b875a6713b

