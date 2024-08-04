package main

import (
	"context"
	"log"

	"github.com/ascopes/terraform-imagecopy-provider/internal"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	ctx := context.Background()
	opts := providerserver.ServeOpts{
		// TODO: use the published name of the Terraform provider on the Terraform/OpenTofu registry.
		Address: "github.com/ascopes/terraform-imagecopy-provider",
	}

	if err := providerserver.Serve(ctx, internal.NewProvider, opts); err != nil {
		log.Fatal(err.Error())
	}
}