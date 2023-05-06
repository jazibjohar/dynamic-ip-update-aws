package main

import (
	"os"

	externalip "github.com/glendc/go-external-ip"
)

func main() {

	// Create the default consensus,
	// using the default configuration and no logger.
	consensus := externalip.DefaultConsensus(nil, nil)
	// By default Ipv4 or Ipv6 is returned,
	// use the function below to limit yourself to IPv4,
	// or pass in `6` instead to limit yourself to IPv6.
	consensus.UseIPProtocol(4)

	// Get your IP,
	// which is never <nil> when err is <nil>.
	ip, err := consensus.ExternalIP()
	if err != nil {
		panic(err)
	}

	UpsertIp(UpdateIP{
		HostPublicIP: ip.String(), // print IPv4/IPv6 in string format
		HostedZoneId: os.Getenv("HOSTED_ZONE_ID"),
		RecordName:   os.Getenv("DOMAIN_TO_UPDATE"),
	})
}
