// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := grafana.NewTeam(ctx, "test", &grafana.TeamArgs{
//				Email: pulumi.String("test-team-email@test.com"),
//				Preferences: &grafana.TeamPreferencesArgs{
//					Theme:    pulumi.String("dark"),
//					Timezone: pulumi.String("utc"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = grafana.LookupTeamOutput(ctx, grafana.GetTeamOutputArgs{
//				Name: test.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupTeam(ctx *pulumi.Context, args *LookupTeamArgs, opts ...pulumi.InvokeOption) (*LookupTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTeamResult
	err := ctx.Invoke("grafana:index/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type LookupTeamArgs struct {
	Name         string  `pulumi:"name"`
	OrgId        *string `pulumi:"orgId"`
	ReadTeamSync *bool   `pulumi:"readTeamSync"`
}

// A collection of values returned by getTeam.
type LookupTeamResult struct {
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id           string              `pulumi:"id"`
	Members      []string            `pulumi:"members"`
	Name         string              `pulumi:"name"`
	OrgId        *string             `pulumi:"orgId"`
	Preferences  []GetTeamPreference `pulumi:"preferences"`
	ReadTeamSync *bool               `pulumi:"readTeamSync"`
	TeamId       int                 `pulumi:"teamId"`
	TeamSyncs    []GetTeamTeamSync   `pulumi:"teamSyncs"`
}

func LookupTeamOutput(ctx *pulumi.Context, args LookupTeamOutputArgs, opts ...pulumi.InvokeOption) LookupTeamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTeamResult, error) {
			args := v.(LookupTeamArgs)
			r, err := LookupTeam(ctx, &args, opts...)
			var s LookupTeamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTeamResultOutput)
}

// A collection of arguments for invoking getTeam.
type LookupTeamOutputArgs struct {
	Name         pulumi.StringInput    `pulumi:"name"`
	OrgId        pulumi.StringPtrInput `pulumi:"orgId"`
	ReadTeamSync pulumi.BoolPtrInput   `pulumi:"readTeamSync"`
}

func (LookupTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamArgs)(nil)).Elem()
}

// A collection of values returned by getTeam.
type LookupTeamResultOutput struct{ *pulumi.OutputState }

func (LookupTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamResult)(nil)).Elem()
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutput() LookupTeamResultOutput {
	return o
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutputWithContext(ctx context.Context) LookupTeamResultOutput {
	return o
}

func (o LookupTeamResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTeamResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

func (o LookupTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTeamResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTeamResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

func (o LookupTeamResultOutput) Preferences() GetTeamPreferenceArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []GetTeamPreference { return v.Preferences }).(GetTeamPreferenceArrayOutput)
}

func (o LookupTeamResultOutput) ReadTeamSync() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTeamResult) *bool { return v.ReadTeamSync }).(pulumi.BoolPtrOutput)
}

func (o LookupTeamResultOutput) TeamId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTeamResult) int { return v.TeamId }).(pulumi.IntOutput)
}

func (o LookupTeamResultOutput) TeamSyncs() GetTeamTeamSyncArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []GetTeamTeamSync { return v.TeamSyncs }).(GetTeamTeamSyncArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTeamResultOutput{})
}
