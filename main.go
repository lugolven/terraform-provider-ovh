package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-ovh/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "hashicorp.com/edu/hashicups",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.NewProvider(), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
