// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Manages Grafana public dashboards.
//
// **Note:** This resource is available only with Grafana 10.2+.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/share-dashboards-panels/shared-dashboards/)
// * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/dashboard_public/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Optional (On-premise, not supported in Grafana Cloud): Create an organization
//			myOrg, err := oss.NewOrganization(ctx, "my_org", &oss.OrganizationArgs{
//				Name: pulumi.String("test 1"),
//			})
//			if err != nil {
//				return err
//			}
//			// Create resources (optional: within the organization)
//			myFolder, err := oss.NewFolder(ctx, "my_folder", &oss.FolderArgs{
//				OrgId: myOrg.OrgId,
//				Title: pulumi.String("test Folder"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"title": "My Terraform Dashboard",
//				"uid":   "my-dashboard-uid",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			testDash, err := oss.NewDashboard(ctx, "test_dash", &oss.DashboardArgs{
//				OrgId:      myOrg.OrgId,
//				Folder:     myFolder.ID(),
//				ConfigJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewDashboardPublic(ctx, "my_public_dashboard", &oss.DashboardPublicArgs{
//				OrgId:                myOrg.OrgId,
//				DashboardUid:         testDash.Uid,
//				Uid:                  pulumi.String("my-custom-public-uid"),
//				AccessToken:          pulumi.String("e99e4275da6f410d83760eefa934d8d2"),
//				TimeSelectionEnabled: pulumi.Bool(true),
//				IsEnabled:            pulumi.Bool(true),
//				AnnotationsEnabled:   pulumi.Bool(true),
//				Share:                pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			// Optional (On-premise, not supported in Grafana Cloud): Create an organization
//			myOrg2, err := oss.NewOrganization(ctx, "my_org2", &oss.OrganizationArgs{
//				Name: pulumi.String("test 2"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"title": "My Terraform Dashboard2",
//				"uid":   "my-dashboard-uid2",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			testDash2, err := oss.NewDashboard(ctx, "test_dash2", &oss.DashboardArgs{
//				OrgId:      myOrg2.OrgId,
//				ConfigJson: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewDashboardPublic(ctx, "my_public_dashboard2", &oss.DashboardPublicArgs{
//				OrgId:        myOrg2.OrgId,
//				DashboardUid: testDash2.Uid,
//				Share:        pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import grafana:oss/dashboardPublic:DashboardPublic name "{{ dashboardUID }}:{{ publicDashboardUID }}"
// ```
//
// ```sh
// $ pulumi import grafana:oss/dashboardPublic:DashboardPublic name "{{ orgID }}:{{ dashboardUID }}:{{ publicDashboardUID }}"
// ```
type DashboardPublic struct {
	pulumi.CustomResourceState

	// A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
	AccessToken pulumi.StringOutput `pulumi:"accessToken"`
	// Set to `true` to show annotations. The default value is `false`.
	AnnotationsEnabled pulumi.BoolPtrOutput `pulumi:"annotationsEnabled"`
	// The unique identifier of the original dashboard.
	DashboardUid pulumi.StringOutput `pulumi:"dashboardUid"`
	// Set to `true` to enable the public dashboard. The default value is `false`.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Set the share mode. The default value is `public`.
	Share pulumi.StringPtrOutput `pulumi:"share"`
	// Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
	TimeSelectionEnabled pulumi.BoolPtrOutput `pulumi:"timeSelectionEnabled"`
	// The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewDashboardPublic registers a new resource with the given unique name, arguments, and options.
func NewDashboardPublic(ctx *pulumi.Context,
	name string, args *DashboardPublicArgs, opts ...pulumi.ResourceOption) (*DashboardPublic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DashboardUid == nil {
		return nil, errors.New("invalid value for required argument 'DashboardUid'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/dashboardPublic:DashboardPublic"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DashboardPublic
	err := ctx.RegisterResource("grafana:oss/dashboardPublic:DashboardPublic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardPublic gets an existing DashboardPublic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardPublic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardPublicState, opts ...pulumi.ResourceOption) (*DashboardPublic, error) {
	var resource DashboardPublic
	err := ctx.ReadResource("grafana:oss/dashboardPublic:DashboardPublic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardPublic resources.
type dashboardPublicState struct {
	// A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
	AccessToken *string `pulumi:"accessToken"`
	// Set to `true` to show annotations. The default value is `false`.
	AnnotationsEnabled *bool `pulumi:"annotationsEnabled"`
	// The unique identifier of the original dashboard.
	DashboardUid *string `pulumi:"dashboardUid"`
	// Set to `true` to enable the public dashboard. The default value is `false`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Set the share mode. The default value is `public`.
	Share *string `pulumi:"share"`
	// Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
	TimeSelectionEnabled *bool `pulumi:"timeSelectionEnabled"`
	// The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
	Uid *string `pulumi:"uid"`
}

type DashboardPublicState struct {
	// A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
	AccessToken pulumi.StringPtrInput
	// Set to `true` to show annotations. The default value is `false`.
	AnnotationsEnabled pulumi.BoolPtrInput
	// The unique identifier of the original dashboard.
	DashboardUid pulumi.StringPtrInput
	// Set to `true` to enable the public dashboard. The default value is `false`.
	IsEnabled pulumi.BoolPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Set the share mode. The default value is `public`.
	Share pulumi.StringPtrInput
	// Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
	TimeSelectionEnabled pulumi.BoolPtrInput
	// The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
	Uid pulumi.StringPtrInput
}

func (DashboardPublicState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardPublicState)(nil)).Elem()
}

type dashboardPublicArgs struct {
	// A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
	AccessToken *string `pulumi:"accessToken"`
	// Set to `true` to show annotations. The default value is `false`.
	AnnotationsEnabled *bool `pulumi:"annotationsEnabled"`
	// The unique identifier of the original dashboard.
	DashboardUid string `pulumi:"dashboardUid"`
	// Set to `true` to enable the public dashboard. The default value is `false`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Set the share mode. The default value is `public`.
	Share *string `pulumi:"share"`
	// Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
	TimeSelectionEnabled *bool `pulumi:"timeSelectionEnabled"`
	// The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
	Uid *string `pulumi:"uid"`
}

// The set of arguments for constructing a DashboardPublic resource.
type DashboardPublicArgs struct {
	// A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
	AccessToken pulumi.StringPtrInput
	// Set to `true` to show annotations. The default value is `false`.
	AnnotationsEnabled pulumi.BoolPtrInput
	// The unique identifier of the original dashboard.
	DashboardUid pulumi.StringInput
	// Set to `true` to enable the public dashboard. The default value is `false`.
	IsEnabled pulumi.BoolPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Set the share mode. The default value is `public`.
	Share pulumi.StringPtrInput
	// Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
	TimeSelectionEnabled pulumi.BoolPtrInput
	// The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
	Uid pulumi.StringPtrInput
}

func (DashboardPublicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardPublicArgs)(nil)).Elem()
}

type DashboardPublicInput interface {
	pulumi.Input

	ToDashboardPublicOutput() DashboardPublicOutput
	ToDashboardPublicOutputWithContext(ctx context.Context) DashboardPublicOutput
}

func (*DashboardPublic) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardPublic)(nil)).Elem()
}

func (i *DashboardPublic) ToDashboardPublicOutput() DashboardPublicOutput {
	return i.ToDashboardPublicOutputWithContext(context.Background())
}

func (i *DashboardPublic) ToDashboardPublicOutputWithContext(ctx context.Context) DashboardPublicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPublicOutput)
}

// DashboardPublicArrayInput is an input type that accepts DashboardPublicArray and DashboardPublicArrayOutput values.
// You can construct a concrete instance of `DashboardPublicArrayInput` via:
//
//	DashboardPublicArray{ DashboardPublicArgs{...} }
type DashboardPublicArrayInput interface {
	pulumi.Input

	ToDashboardPublicArrayOutput() DashboardPublicArrayOutput
	ToDashboardPublicArrayOutputWithContext(context.Context) DashboardPublicArrayOutput
}

type DashboardPublicArray []DashboardPublicInput

func (DashboardPublicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardPublic)(nil)).Elem()
}

func (i DashboardPublicArray) ToDashboardPublicArrayOutput() DashboardPublicArrayOutput {
	return i.ToDashboardPublicArrayOutputWithContext(context.Background())
}

func (i DashboardPublicArray) ToDashboardPublicArrayOutputWithContext(ctx context.Context) DashboardPublicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPublicArrayOutput)
}

// DashboardPublicMapInput is an input type that accepts DashboardPublicMap and DashboardPublicMapOutput values.
// You can construct a concrete instance of `DashboardPublicMapInput` via:
//
//	DashboardPublicMap{ "key": DashboardPublicArgs{...} }
type DashboardPublicMapInput interface {
	pulumi.Input

	ToDashboardPublicMapOutput() DashboardPublicMapOutput
	ToDashboardPublicMapOutputWithContext(context.Context) DashboardPublicMapOutput
}

type DashboardPublicMap map[string]DashboardPublicInput

func (DashboardPublicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardPublic)(nil)).Elem()
}

func (i DashboardPublicMap) ToDashboardPublicMapOutput() DashboardPublicMapOutput {
	return i.ToDashboardPublicMapOutputWithContext(context.Background())
}

func (i DashboardPublicMap) ToDashboardPublicMapOutputWithContext(ctx context.Context) DashboardPublicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPublicMapOutput)
}

type DashboardPublicOutput struct{ *pulumi.OutputState }

func (DashboardPublicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardPublic)(nil)).Elem()
}

func (o DashboardPublicOutput) ToDashboardPublicOutput() DashboardPublicOutput {
	return o
}

func (o DashboardPublicOutput) ToDashboardPublicOutputWithContext(ctx context.Context) DashboardPublicOutput {
	return o
}

// A public unique identifier of a public dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a public dashboard.
func (o DashboardPublicOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.StringOutput { return v.AccessToken }).(pulumi.StringOutput)
}

// Set to `true` to show annotations. The default value is `false`.
func (o DashboardPublicOutput) AnnotationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.BoolPtrOutput { return v.AnnotationsEnabled }).(pulumi.BoolPtrOutput)
}

// The unique identifier of the original dashboard.
func (o DashboardPublicOutput) DashboardUid() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.StringOutput { return v.DashboardUid }).(pulumi.StringOutput)
}

// Set to `true` to enable the public dashboard. The default value is `false`.
func (o DashboardPublicOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o DashboardPublicOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Set the share mode. The default value is `public`.
func (o DashboardPublicOutput) Share() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.StringPtrOutput { return v.Share }).(pulumi.StringPtrOutput)
}

// Set to `true` to enable the time picker in the public dashboard. The default value is `false`.
func (o DashboardPublicOutput) TimeSelectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.BoolPtrOutput { return v.TimeSelectionEnabled }).(pulumi.BoolPtrOutput)
}

// The unique identifier of a public dashboard. It's automatically generated if not provided when creating a public dashboard.
func (o DashboardPublicOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardPublic) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type DashboardPublicArrayOutput struct{ *pulumi.OutputState }

func (DashboardPublicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardPublic)(nil)).Elem()
}

func (o DashboardPublicArrayOutput) ToDashboardPublicArrayOutput() DashboardPublicArrayOutput {
	return o
}

func (o DashboardPublicArrayOutput) ToDashboardPublicArrayOutputWithContext(ctx context.Context) DashboardPublicArrayOutput {
	return o
}

func (o DashboardPublicArrayOutput) Index(i pulumi.IntInput) DashboardPublicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardPublic {
		return vs[0].([]*DashboardPublic)[vs[1].(int)]
	}).(DashboardPublicOutput)
}

type DashboardPublicMapOutput struct{ *pulumi.OutputState }

func (DashboardPublicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardPublic)(nil)).Elem()
}

func (o DashboardPublicMapOutput) ToDashboardPublicMapOutput() DashboardPublicMapOutput {
	return o
}

func (o DashboardPublicMapOutput) ToDashboardPublicMapOutputWithContext(ctx context.Context) DashboardPublicMapOutput {
	return o
}

func (o DashboardPublicMapOutput) MapIndex(k pulumi.StringInput) DashboardPublicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardPublic {
		return vs[0].(map[string]*DashboardPublic)[vs[1].(string)]
	}).(DashboardPublicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPublicInput)(nil)).Elem(), &DashboardPublic{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPublicArrayInput)(nil)).Elem(), DashboardPublicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPublicMapInput)(nil)).Elem(), DashboardPublicMap{})
	pulumi.RegisterOutputType(DashboardPublicOutput{})
	pulumi.RegisterOutputType(DashboardPublicArrayOutput{})
	pulumi.RegisterOutputType(DashboardPublicMapOutput{})
}
