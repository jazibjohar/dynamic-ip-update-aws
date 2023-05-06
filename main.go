package main

import (
	"os"
	"strings"
)

func main() {
	externalIP, err := GetExternalIP(4)
	if err != nil {
		panic(err)
	}
	updateIPConfig := UpdateIP{
		HostPublicIP: externalIP,
		HostedZoneId: os.Getenv("HOSTED_ZONE_ID"),
		RecordName:   os.Getenv("DOMAIN_TO_UPDATE"),
	}

	if len(strings.TrimSpace(updateIPConfig.HostedZoneId)) == 0 &&
		len(strings.TrimSpace(updateIPConfig.RecordName)) == 0 {
		panic("Please configure HOSTED_ZONE_ID and DOMAIN_TO_UPDATE")
	}

	if err := UpsertIp(updateIPConfig); err != nil {
		panic(err)
	}
}
