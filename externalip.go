package main

import externalip "github.com/glendc/go-external-ip"

func GetExternalIP(version uint) (string, error) {
	// Create the default consensus,
	// using the default configuration and no logger.
	consensus := externalip.DefaultConsensus(nil, nil)
	// By default Ipv4 or Ipv6 is returned,
	// use the function below to limit yourself to IPv4,
	// or pass in `6` instead to limit yourself to IPv6.
	consensus.UseIPProtocol(version)

	// Get your IP,
	// which is never <nil> when err is <nil>.
	ip, err := consensus.ExternalIP()
	if err != nil {
		return "", err
	}
	return ip.String(), nil
}
