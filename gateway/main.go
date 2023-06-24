package main

import (
	"fmt"
	"github.com/zcubbs/zlogger/pkg/logger"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func run() {
	scannerServerURL, err := url.Parse("http://localhost:8001")
	if err != nil {
		log.Fatal("invalid scanner server URL")
	}
	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		logger.L().Infof("[reverse proxy server] received request at: %s\n", time.Now())

		if strings.HasPrefix(req.RequestURI, "/v1/scan") {
			req.Host = scannerServerURL.Host
			req.URL.Host = scannerServerURL.Host
			req.URL.Scheme = scannerServerURL.Scheme
			req.RequestURI = ""

			// send a request to the origin server
			scannerServerResponse, err := http.DefaultClient.Do(req)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprint(rw, "Internal Server Error")
				logger.L().Error(err)
				return
			}

			// return response to the client
			rw.WriteHeader(http.StatusOK)
			_, err = io.Copy(rw, scannerServerResponse.Body)
			if err != nil {
				logger.L().Error(err)
				return
			}

			return
		}

		rw.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(rw, "Not Found")
		return
	})

	logger.L().Infof("Starting HTTP server on port 8000")
	logger.L().Fatal(http.ListenAndServe(":8000", reverseProxy))
}

func main() {
	logger.SetupLogger(logger.ZapLogger)

	run()
}
