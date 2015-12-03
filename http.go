package mlib

import (
	"fmt"
	"time"
	"net"
	"crypto/tls"
	"net/http"
)
const HTTP_GET = 0
const HTTP_POST = 1

func Write501(w *http.ResponseWriter){
        (*w).WriteHeader(501)
}


var timeout = time.Duration(30 * time.Second)
func dialTimeout(network, addr string) (net.Conn, error) {
        return net.DialTimeout(network, addr, timeout)
}

func GetHttpClient(url *string,mezod int,insecSsl bool) (*http.Client,*http.Request){
	var _mezod string
	if mezod == HTTP_GET {
		_mezod = "GET"
	}else if mezod == HTTP_POST {
		_mezod = "POST"
	}

        tr := &http.Transport{
		ResponseHeaderTimeout: time.Second * 30,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecSsl},
		Dial: dialTimeout,
        }

        client := &http.Client{
		Transport: tr}

        req, err := http.NewRequest(_mezod, *url, nil)
        if err != nil {
                fmt.Printf("error in get req url: %s error %s\n",url,err)
                return nil,nil
        }

        return client, req
}
