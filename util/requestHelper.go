package util

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func SetHeaders(r *http.Request, headers map[string]string) {
	for key, value := range headers {
		r.Header.Set(key, value)
	}
}

func PrintPreRequest(ctx context.Context, printer Printer, method, host, port, path string) {
	printer.Print(ctx, fmt.Sprintf("connecting to %s:%s", host, port))
	printer.Print(ctx, fmt.Sprintf("Sending request %s %s HTTP/1.1", method, path))
	printer.Print(ctx, fmt.Sprintf("Host: %s", host))
}

func PrintResponse(ctx context.Context, printer Printer, response http.Response) {
	printer.Print(ctx, fmt.Sprintf("%q %q", response.Proto, response.Status))
	for k, v := range response.Header {
		printer.Print(ctx, fmt.Sprintf("%q: %q", k, v))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.WithError(err).Fatalln()
		}
	}(response.Body)
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.WithError(err).Fatalln()
	}
	bodyString := string(bodyBytes)
	printer.Print(ctx, bodyString)
}
