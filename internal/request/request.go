package request

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/amirsadraabdollahi/curl/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Response interface{}

type Requester interface {
	Get(ctx context.Context, url string) Response
	Post(ctx context.Context, url string) Response
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

func (h *HttpRequester) Get(ctx context.Context, url string) Response {
	protocol, host, port, path := util.GetUrlParts(url)
	method := "GET"

	ctx = context.WithValue(ctx, "stream", "out")
	util.PrintPreRequest(ctx, h.printer, method, host, port, path)

	finalUrl := protocol + "://" + host + ":" + port + path
	resp, err := http.Get(finalUrl)
	if err != nil {
		log.WithField("url", url).WithError(err).Fatalln()
	}

	util.PrintResponse(ctx, h.printer, *resp)

	return resp
}

func (h *HttpRequester) Post(ctx context.Context, url string) Response {
	protocol, host, port, path := util.GetUrlParts(url)
	method := "POST"

	ctx = context.WithValue(ctx, "stream", "out")
	util.PrintPreRequest(ctx, h.printer, method, host, port, path)

	finalUrl := protocol + "://" + host + ":" + port + path
	body := ctx.Value("body")
	out, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(method, finalUrl, bytes.NewBuffer(out))
	if err != nil {
		log.WithError(err).Fatalln()
	}

	headers, _ := ctx.Value("headers").(map[string]string)
	util.SetHeaders(req, headers)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Fatalln()
	}

	ctx = context.WithValue(ctx, "stream", "in")
	util.PrintResponse(ctx, h.printer, *resp)

	return resp
}

func (h *HttpRequester) Put() Response {
	//TODO implement me
	panic("implement me")
}
