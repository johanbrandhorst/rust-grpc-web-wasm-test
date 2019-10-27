package main

import (
	"crypto/tls"
	"flag"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/lpar/gzipped"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/johanbrandhorst/rust-grpc-web-wasm-test/proto/gen/go/library"
	"github.com/johanbrandhorst/rust-grpc-web-wasm-test/server"
)

var logger *logrus.Logger
var htmlRoot = flag.String("html-root", "./client/html", "Root of the HTML files directory.")

func init() {
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	// Should only be done from init functions
	grpclog.SetLogger(logger)
}

func main() {
	flag.Parse()

	gs := grpc.NewServer()
	library.RegisterBookServiceServer(gs, &server.BookService{
		Logger: logger,
	})
	wrappedServer := grpcweb.WrapServer(gs)

	httpsSrv := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              ":https",
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
		Handler: hstsHandler(
			grpcTrafficSplitter(
				folderReader(
					gzipped.FileServer(http.Dir(*htmlRoot)),
				),
				wrappedServer,
			),
		),
	}

	httpsSrv.Addr = "localhost:8080"
	logger.Info("Serving on https://localhost:8080")
	logger.Fatal(httpsSrv.ListenAndServeTLS("./insecure/cert.pem", "./insecure/key.pem"))
}

// hstsHandler wraps an http.HandlerFunc such that it sets the HSTS header.
func hstsHandler(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		fn.ServeHTTP(w, r)
	})
}

func folderReader(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			// Use contents of index.html for directory, if present.
			r.URL.Path = path.Join(r.URL.Path, "index.html")
		}
		fn.ServeHTTP(w, r)
	})
}

func grpcTrafficSplitter(fallback http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Redirect gRPC and gRPC-Web requests to the gRPC Server
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			logger.Info("Serving gRPC-Web request")
			grpcHandler.ServeHTTP(w, r)
			return
		}

		logger.Info("Serving HTML request")
		fallback.ServeHTTP(w, r)
	})
}
