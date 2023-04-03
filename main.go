package main

import (
	"fmt"

	"os"

	"ckkannan/servicenowtable_client"
)

const CSN_URL string = "https://dev161016.service-now.com/"

var h string = CSN_URL
var sn_user string
var sn_pass string

// var host *string = &h
// var puser *string = &user
// var ppass *string = &pass

type TLSconfig struct {
	InsecureSkipVerify bool
}

// var version string = "dev"

func main() {

	h = os.Getenv("sn_url")
	sn_user = os.Getenv("sn_user")
	sn_pass = os.Getenv("sn_pass")
	//sslignore, err := strconv.ParseBool(os.Getenv("SN_SSLIGNORE"))
	// if err != nil {
	// 	fmt.Println("Pare error for SN_SSLIGNORE")
	// }
	sndev := servicenowtable_client.ServicenowtableProviderInput{Sn_url: h, Sn_user: sn_user, Sn_pass: sn_pass, SSLIgnore: true, Version: "1.0"}

	c, err := servicenowtable_client.NewClient(sndev)
	if err != nil {
		fmt.Println("ServiceNow New Client Failed")
		os.Exit(1)
	}

	c.Table = "x_22541_terraform_organization"
	c.Query = "to_adgroupSTARTSWITHSSDP-TFE"
	c.Fields = "sys_id,to_adgroup,to_org_name,to_org_type,id"

	OrgRows, err := c.GetRows()
	if err != nil {
		fmt.Println("Query Failed")
		os.Exit(1)
	}

	for i := range OrgRows {
		for k1, v1 := range OrgRows[i] {
			fmt.Print(i, " ", k1, " ")
			fmt.Println(v1)
		}

	}

}
