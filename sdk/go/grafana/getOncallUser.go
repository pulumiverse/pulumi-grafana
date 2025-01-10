// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/users/)
//
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
//			_, err := oncall.GetUser(ctx, &oncall.GetUserArgs{
//				Username: "alex",
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
// Deprecated: grafana.index/getoncalluser.getOncallUser has been deprecated in favor of grafana.oncall/getuser.getUser
func GetOncallUser(ctx *pulumi.Context, args *GetOncallUserArgs, opts ...pulumi.InvokeOption) (*GetOncallUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOncallUserResult
	err := ctx.Invoke("grafana:index/getOncallUser:getOncallUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOncallUser.
type GetOncallUserArgs struct {
	// The username of the user.
	Username string `pulumi:"username"`
}

// A collection of values returned by getOncallUser.
type GetOncallUserResult struct {
	// The email of the user.
	Email string `pulumi:"email"`
	// The ID of the user.
	Id string `pulumi:"id"`
	// The role of the user.
	Role string `pulumi:"role"`
	// The username of the user.
	Username string `pulumi:"username"`
}

func GetOncallUserOutput(ctx *pulumi.Context, args GetOncallUserOutputArgs, opts ...pulumi.InvokeOption) GetOncallUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOncallUserResultOutput, error) {
			args := v.(GetOncallUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("grafana:index/getOncallUser:getOncallUser", args, GetOncallUserResultOutput{}, options).(GetOncallUserResultOutput), nil
		}).(GetOncallUserResultOutput)
}

// A collection of arguments for invoking getOncallUser.
type GetOncallUserOutputArgs struct {
	// The username of the user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (GetOncallUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallUserArgs)(nil)).Elem()
}

// A collection of values returned by getOncallUser.
type GetOncallUserResultOutput struct{ *pulumi.OutputState }

func (GetOncallUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallUserResult)(nil)).Elem()
}

func (o GetOncallUserResultOutput) ToGetOncallUserResultOutput() GetOncallUserResultOutput {
	return o
}

func (o GetOncallUserResultOutput) ToGetOncallUserResultOutputWithContext(ctx context.Context) GetOncallUserResultOutput {
	return o
}

// The email of the user.
func (o GetOncallUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// The ID of the user.
func (o GetOncallUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The role of the user.
func (o GetOncallUserResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserResult) string { return v.Role }).(pulumi.StringOutput)
}

// The username of the user.
func (o GetOncallUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOncallUserResultOutput{})
}
