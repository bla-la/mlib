package mlib

import (
	"fmt"
	"time"
	"net"
	"crypto/tls"
	"net/http"
)

func Write501(w *http.ResponseWriter){
        (*w).WriteHeader(501)
}


var timeout = time.Duration(30 * time.Second)
func dialTimeout(network, addr string) (net.Conn, error) {
        fmt.Printf("Timeout\n")
        return net.DialTimeout(network, addr, timeout)
}

func GetHttpClient(url string,mezod string) (*http.Client,*http.Request){
        tr := &http.Transport{
		ResponseHeaderTimeout: time.Second * 30,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: dialTimeout,
        }

        client := &http.Client{
		Transport: tr}

        req, err := http.NewRequest(mezod, url, nil)
        if err != nil {
                fmt.Printf("error in get req url: %s error %s\n",url,err)
                return nil,nil
        }

        return client, req
}
