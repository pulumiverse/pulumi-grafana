// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oncall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oncall.GetTeam(ctx, &oncall.GetTeamArgs{
//				Name: "example_team",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getoncallteam.getOncallTeam has been deprecated in favor of grafana.oncall/getteam.getTeam
func GetOncallTeam(ctx *pulumi.Context, args *GetOncallTeamArgs, opts ...pulumi.InvokeOption) (*GetOncallTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOncallTeamResult
	err := ctx.Invoke("grafana:index/getOncallTeam:getOncallTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOncallTeam.
type GetOncallTeamArgs struct {
	// The team name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOncallTeam.
type GetOncallTeamResult struct {
	AvatarUrl string `pulumi:"avatarUrl"`
	Email     string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The team name.
	Name string `pulumi:"name"`
}

func GetOncallTeamOutput(ctx *pulumi.Context, args GetOncallTeamOutputArgs, opts ...pulumi.InvokeOption) GetOncallTeamResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOncallTeamResultOutput, error) {
			args := v.(GetOncallTeamArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getOncallTeam:getOncallTeam", args, GetOncallTeamResultOutput{}, options).(GetOncallTeamResultOutput), nil
		}).(GetOncallTeamResultOutput)
}

// A collection of arguments for invoking getOncallTeam.
type GetOncallTeamOutputArgs struct {
	// The team name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetOncallTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallTeamArgs)(nil)).Elem()
}

// A collection of values returned by getOncallTeam.
type GetOncallTeamResultOutput struct{ *pulumi.OutputState }

func (GetOncallTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallTeamResult)(nil)).Elem()
}

func (o GetOncallTeamResultOutput) ToGetOncallTeamResultOutput() GetOncallTeamResultOutput {
	return o
}

func (o GetOncallTeamResultOutput) ToGetOncallTeamResultOutputWithContext(ctx context.Context) GetOncallTeamResultOutput {
	return o
}

func (o GetOncallTeamResultOutput) AvatarUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallTeamResult) string { return v.AvatarUrl }).(pulumi.StringOutput)
}

func (o GetOncallTeamResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallTeamResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOncallTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// The team name.
func (o GetOncallTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOncallTeamResultOutput{})
}
