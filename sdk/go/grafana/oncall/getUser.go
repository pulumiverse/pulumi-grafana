// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/onCall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := onCall.GetUser(ctx, &oncall.GetUserArgs{
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
func GetUser(ctx *pulumi.Context, args *GetUserArgs, opts ...pulumi.InvokeOption) (*GetUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserResult
	err := ctx.Invoke("grafana:onCall/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The username of the user.
	Username string `pulumi:"username"`
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// The email of the user.
	Email string `pulumi:"email"`
	// The ID of the user.
	Id string `pulumi:"id"`
	// The role of the user.
	Role string `pulumi:"role"`
	// The username of the user.
	Username string `pulumi:"username"`
}

func GetUserOutput(ctx *pulumi.Context, args GetUserOutputArgs, opts ...pulumi.InvokeOption) GetUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserResultOutput, error) {
			args := v.(GetUserArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetUserResult
			secret, err := ctx.InvokePackageRaw("grafana:onCall/getUser:getUser", args, &rv, "", opts...)
			if err != nil {
				return GetUserResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetUserResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetUserResultOutput), nil
			}
			return output, nil
		}).(GetUserResultOutput)
}

// A collection of arguments for invoking getUser.
type GetUserOutputArgs struct {
	// The username of the user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (GetUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type GetUserResultOutput struct{ *pulumi.OutputState }

func (GetUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserResult)(nil)).Elem()
}

func (o GetUserResultOutput) ToGetUserResultOutput() GetUserResultOutput {
	return o
}

func (o GetUserResultOutput) ToGetUserResultOutputWithContext(ctx context.Context) GetUserResultOutput {
	return o
}

// The email of the user.
func (o GetUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// The ID of the user.
func (o GetUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The role of the user.
func (o GetUserResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Role }).(pulumi.StringOutput)
}

// The username of the user.
func (o GetUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserResultOutput{})
}
