package provider

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = SemVerVersionFunction{}
)

func NewSemVerVersionFunction() function.Function {
	return SemVerVersionFunction{}
}

type SemVerVersionFunction struct{}

func (r SemVerVersionFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "semver_version"
}

func (r SemVerVersionFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check whether a given string is a valid semantic version",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Name:               "version",
				Description:        "The semantic version to extract the version part from",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r SemVerVersionFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var v string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &v))
	if resp.Error != nil {
		return
	}

	_, err := version.NewVersion(v)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, err == nil))
}
