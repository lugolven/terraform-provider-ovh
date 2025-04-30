package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// NewProvider creates a new provider instance
func NewProvider() func() provider.Provider {
	return func() provider.Provider {
		return &internalProvider{}
	}
}

type internalProvider struct {
}

func (p *internalProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ovh"
}

func (p *internalProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Sets up OVH resources.",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (p *internalProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {

}

func (p *internalProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *internalProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
