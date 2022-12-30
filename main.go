package main

import (
	"encoding/json"
	"fmt"
	"gt-ss-dp/capam"
	"os"
)

const HostURL string = "https://p1ehowld170.prudential.com"

var h string = HostURL
var user string = "super-api-1"
var pass string = "Remove Passwd"
var host *string = &h
var puser *string = &user
var ppass *string = &pass

type TLSconfig struct {
	InsecureSkipVerify bool
}

func main() {
	tlsconfig := &capam.TLSconfig{InsecureSkipVerify: false}
	tlsconfig.InsecureSkipVerify = true
	h, err := capam.NewClient(host, puser, ppass, tlsconfig)
	if err != nil {
		fmt.Println("PAM New Client Failed")
		os.Exit(1)
	}
	var pvps []capam.PVP
	pvps, err = h.GetPVPs()
	if err != nil {
		fmt.Println("GetPVPs Failed")
		os.Exit(1)
	}
	s, _ := json.Marshal(pvps)
	fmt.Println(string(s))
}
