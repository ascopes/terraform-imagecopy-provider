package internal

import (
	"context"

	"github.com/ascopes/terraform-provider-imagecopy/internal/auth"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NewProvider() provider.Provider {
	return &providerImpl{}
}

type providerImpl struct {
	authFactories map[string]auth.AuthFactory
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
			"registry": schema.SetNestedBlock{
				Description: "Configuration for a registry.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Description: "The base URL of the registry.",
							Required:    true,
						},
					},
					Blocks: map[string]schema.Block{
						"basic_auth": schema.SingleNestedBlock{
							Description: "Configuration for basic authentication.",
							Attributes: map[string]schema.Attribute{
								"username": schema.StringAttribute{
									Description: "The username to use.",
									Required:    true,
									Sensitive:   true,
								},
								"password": schema.StringAttribute{
									Description: "The password to use.",
									Required:    true,
									Sensitive:   true,
								},
							},
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("bearer_auth")),
							},
						},
						"bearer_auth": schema.SingleNestedBlock{
							Description: "Configuration for bearer authentication.",
							Attributes: map[string]schema.Attribute{
								"token": schema.StringAttribute{
									Description: "The token to use.",
									Required:    true,
									Sensitive:   true,
								},
							},
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("basic_auth")),
							},
						},
					},
				},
			},
		},
	}
}
