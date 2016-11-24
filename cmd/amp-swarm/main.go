package main

import ()

// build vars
var (
	Version string
	Build   string
	client  *swarmClient  = &swarmClient{}
	conf    *ClientConfig = &ClientConfig{}
)

func main() {
	conf.init(Version, Build)
	client.cli()
}
