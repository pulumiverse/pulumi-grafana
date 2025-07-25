// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
//
// This data source uses Grafana's admin APIs for reading users which
// does not currently work with API Tokens. You must use basic auth.
// This data source is also not compatible with Grafana Cloud, as it does not allow basic auth.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewUser(ctx, "test_all_users", &oss.UserArgs{
//				Email:    pulumi.String("all_users@example.com"),
//				Name:     pulumi.String("Testing grafana_users"),
//				Login:    pulumi.String("test-grafana-users"),
//				Password: pulumi.String("my-password"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.GetUsers(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetUsers(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("grafana:oss/getUsers:getUsers", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Grafana instance's users.
	Users []GetUsersUser `pulumi:"users"`
}

func GetUsersOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetUsersResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("grafana:oss/getUsers:getUsers", nil, GetUsersResultOutput{}, options).(GetUsersResultOutput), nil
	}).(GetUsersResultOutput)
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Grafana instance's users.
func (o GetUsersResultOutput) Users() GetUsersUserArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUser { return v.Users }).(GetUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
