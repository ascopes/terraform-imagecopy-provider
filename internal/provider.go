package internal

import (
	"context"

	"github.com/ascopes/terraform-imagecopy-provider/internal/auth"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewProvider() provider.Provider {
	return &providerImpl{}
}

type providerImpl struct {
	authenticatorProviders map[string]auth.AuthenticatorFactory
}

func (provider *providerImpl) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
	panic("unimplemented")
}

func (provider *providerImpl) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (provider *providerImpl) Metadata(_ context.Context, _ provider.MetadataRequest, res *provider.MetadataResponse) {
	res.TypeName = "imagecopy"
}

func (provider *providerImpl) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (provider *providerImpl) Schema(_ context.Context, _ provider.SchemaRequest, res *provider.SchemaResponse) {
	res.Schema = schema.Schema{
		Description: "Provider to enable Terraforming the transfer of container images between locations.",
		Blocks: map[string]schema.Block{
			"basic_auth": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Description: "The URL of the registry.",
							Required:    true,
						},
						"username": schema.StringAttribute{
							Description: "The basic authentication username to use.",
							Required:    true,
						},
						"password": schema.StringAttribute{
							Description: "The basic authentication password to use.",
							Required:    true,
							Sensitive:   true,
						},
					},
				},
				Description: "A basic authentication configuration block.",
			},
			"bearer_auth": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Description: "The URL of the registry.",
							Required:    true,
						},
						"token": schema.StringAttribute{
							Description: "The bearer token to use.",
							Required:    true,
							Sensitive:   true,
						},
					},
				},
				Description: "A bearer authentication configuration block.",
			},
		},
	}
}
