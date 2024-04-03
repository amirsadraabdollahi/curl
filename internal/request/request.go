package request

import (
	"context"
	"fmt"
	"github.com/amirsadraabdollahi/curl/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

type InvalidUrlError struct {
	msg string
}

func (i InvalidUrlError) Error() string {
	return i.msg
}

type Response interface{}

type Requester interface {
	Get(ctx *context.Context, url string) Response
	Post() Response
	Put() Response
	SetPrinter(printer util.Printer)
}

type HttpRequester struct {
	printer util.Printer
}

func (h *HttpRequester) SetPrinter(printer util.Printer) {
	h.printer = printer
}

func NewHttpRequester(printer util.Printer) *HttpRequester {
	return &HttpRequester{printer: printer}
}

func (h *HttpRequester) Get(ctx *context.Context, url string) Response {
	re := regexp.MustCompile(`(.*)://(.+?)(:([0-9]+))?(/.*)`)
	submatch := re.FindStringSubmatch(url)
	protocol, host, port, path := submatch[1], submatch[2], submatch[4], submatch[5]
	fmt.Println(submatch)
	if protocol != "http" {
		log.WithField("protocol", protocol).WithError(InvalidUrlError{msg: "protocol is invalid"}).Errorln()
	}

	if len(port) == 0 {
		port = "80"
	}
	*ctx = context.WithValue(*ctx, "stream", "out")
	h.printer.Print(*ctx, fmt.Sprintf("connecting to %s:%s", host, port))
	h.printer.Print(*ctx, fmt.Sprintf("Sending request GET %s HTTP/1.1", path))
	h.printer.Print(*ctx, fmt.Sprintf("Host: %s", host))
	resp, err := http.Get(protocol + "://" + host + ":" + port + path)
	if err != nil {
		log.WithField("url", url).WithError(err).Errorln()
	}
	*ctx = context.WithValue(*ctx, "stream", "in")
	h.printer.Print(*ctx, fmt.Sprintf("%q %q", resp.Proto, resp.Status))
	for k, v := range resp.Header {
		h.printer.Print(*ctx, fmt.Sprintf("%q: %q", k, v))
	}
	return resp
}

func (h *HttpRequester) Post() Response {
	//TODO implement me
	panic("implement me")
}

func (h *HttpRequester) Put() Response {
	//TODO implement me
	panic("implement me")
}
