package mlib

import (
	"github.com/miekg/dns"
	"errors"
	"fmt"
)

func GettYpeA(domain *string)([]string,error){
	var ret []string
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(*domain+".", dns.TypeA)
	r, _, err := c.Exchange(m, config.Servers[0]+":"+config.Port)
	if err != nil {
		errTxt := fmt.Sprintf("Lookup %s",err)
		return nil,errors.New(errTxt)
	}
	if r.Rcode != dns.RcodeSuccess {
		errTxt := fmt.Sprintf("DNS error %d",r.Rcode)
		return nil,errors.New(errTxt)
	}
	for _,a := range r.Answer {
		if addr, ok := a.(*dns.A); ok {
			ret = append(ret,addr.A.String())
		}
	}

	return ret,nil
}
