package mlib

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

const HTTP_GET = 0
const HTTP_POST = 1

func Write200(w *http.ResponseWriter) {
	(*w).WriteHeader(200)
}

func Write501(w *http.ResponseWriter) {
	(*w).WriteHeader(501)
}

func Write503(w *http.ResponseWriter) {
	(*w).WriteHeader(501)
}

func Write400(w *http.ResponseWriter, msg string) {
	(*w).WriteHeader(400)
	fmt.Fprintf((*w), "%s", msg)
}

func Write401(w *http.ResponseWriter, msg string) {
	(*w).WriteHeader(401)
	fmt.Fprintf((*w), "%s", msg)
}

func Write404(w *http.ResponseWriter, msg string) {
	(*w).WriteHeader(404)
	fmt.Fprintf((*w), "%s", msg)
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func GetHttpClient(url *string, mezod int, insecSsl bool) (*http.Client, *http.Request) {
	var _mezod string
	if mezod == HTTP_GET {
		_mezod = "GET"
	} else if mezod == HTTP_POST {
		_mezod = "POST"
	}

	tr := &http.Transport{
		ResponseHeaderTimeout: time.Second * 30,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: insecSsl},
		Dial:                  dialTimeout,
	}

	client := &http.Client{
		Transport: tr}

	req, err := http.NewRequest(_mezod, *url, nil)
	if err != nil {
		fmt.Printf("error in get req url: %s error %s\n", url, err)
		return nil, nil
	}

	return client, req
}
