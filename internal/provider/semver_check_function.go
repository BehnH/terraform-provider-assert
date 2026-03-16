package provider

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = SemVerCheckFunction{}
)

func NewSemVerCheckFunction() function.Function {
	return SemVerCheckFunction{}
}

type SemVerCheckFunction struct{}

func (r SemVerCheckFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "semver_check"
}

func (r SemVerCheckFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a given semver string matches a constraint",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The constraint to validate against",
				Name:               "constraint",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The version  string to validate",
				Name:               "version",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r SemVerCheckFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var c, v string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &c, &v))
	if resp.Error != nil {
		return
	}

	semver, err := version.NewVersion(v)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	constraints, err := version.NewConstraint(c)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, constraints.Check(semver)))
}
