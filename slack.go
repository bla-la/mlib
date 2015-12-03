package mlib

import (
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
)

func PostToSlackChat(slackUrl string,user string,msg string)(error){
        msg = strings.Replace(msg,"\n"," ",-1)
        slackText := "{\"username\": \""+user+"\",\"text\":\""+msg+"\"}"

	Info("Try send slack")
        resp, cliErr := http.PostForm(slackUrl,
                url.Values{"payload": { slackText }})

        if cliErr != nil {
                Error(" get req text %s error %s",slackText,cliErr)
                return cliErr
        }


        body,readBodyErr := ioutil.ReadAll(resp.Body)
        if readBodyErr != nil {
                Error("read body %s error %s",body,readBodyErr)
                return readBodyErr
        }
        Success("text:%s code %s",slackText,resp.Status)
        return nil
}
