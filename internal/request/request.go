package request

import (
	"context"
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
}

type HttpRequester struct {
}

func NewHttpRequester() HttpRequester {
	return HttpRequester{}
}

func (h HttpRequester) Get(ctx *context.Context, url string) Response {
	re := regexp.MustCompile(`(.*)://(.+)(:([0-9]+))?(/.*?)`)
	submatch := re.FindStringSubmatch(url)
	protocol, host, port, path := submatch[1], submatch[2], submatch[4], submatch[5]

	if protocol != "http" {
		log.WithField("protocol", protocol).WithError(InvalidUrlError{msg: "protocol is invalid"}).Errorln()
	}

	if len(port) == 0 {
		port = "80"
	}

	log.WithField("protocol", protocol).WithField("host", host).WithField("port", port).WithField("path", path).Info()
	resp, err := http.Get(protocol + "://" + host + ":" + port + path)
	if err != nil {
		log.WithField("url", url).WithError(err).Errorln()
	}
	return resp
}

func (h HttpRequester) Post() Response {
	//TODO implement me
	panic("implement me")
}

func (h HttpRequester) Put() Response {
	//TODO implement me
	panic("implement me")
}
