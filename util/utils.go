package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

type InvalidUrlError struct {
	msg string
}

func (i InvalidUrlError) Error() string {
	return i.msg
}

func GetUrlParts(url string) (string, string, string, string) {
	re := regexp.MustCompile(`(.*)://(.+?)(:([0-9]+))?(/.*)`)
	submatch := re.FindStringSubmatch(url)
	protocol, host, port, path := submatch[1], submatch[2], submatch[4], submatch[5]
	if protocol != "http" {
		log.WithField("protocol", protocol).WithError(InvalidUrlError{msg: "protocol is invalid"}).Fatalln()
	}
	if len(port) == 0 {
		port = "80"
	}

	return protocol, host, port, path
}

func ConvArrToMap(arr []string) map[string]string {
	if len(arr) == 0 {
		return nil
	}
	m := make(map[string]string)
	for _, element := range arr {
		temp := strings.Split(element, ": ")
		fmt.Println(temp)
		key, value := temp[0], temp[1]
		m[key] = value
	}
	return m
}
