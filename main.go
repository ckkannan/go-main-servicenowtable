package main

import (
	"encoding/json"
	"fmt"
	"gt-ss-dp/capam"
	"os"
	"strconv"
)

const HostURL string = "https://p1ehowld170.prudential.com"

var h string = HostURL
var user string
var pass string
var host *string = &h
var puser *string = &user
var ppass *string = &pass

type TLSconfig struct {
	InsecureSkipVerify bool
}

func main() {

	h = os.Getenv("CAPAM_HOST")
	user = os.Getenv("CAPAM_APIUSER")
	pass = os.Getenv("CAPAM_APIPASSWORD")
	sslignore, err := strconv.ParseBool(os.Getenv("CAPAM_SSLIGNORE"))
	tlsconfig := &capam.TLSconfig{InsecureSkipVerify: sslignore}
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
