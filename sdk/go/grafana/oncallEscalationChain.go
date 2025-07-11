// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myTeam, err := oss.LookupTeam(ctx, &oss.LookupTeamArgs{
//				Name: "my team",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			myTeamGetTeam, err := oncall.GetTeam(ctx, &oncall.GetTeamArgs{
//				Name: myTeam.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewEscalationChain(ctx, "default", &oncall.EscalationChainArgs{
//				Name:   pulumi.String("default"),
//				TeamId: pulumi.String(myTeamGetTeam.Id),
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
// $ pulumi import grafana:index/oncallEscalationChain:OncallEscalationChain name "{{ id }}"
// ```
//
// Deprecated: grafana.index/oncallescalationchain.OncallEscalationChain has been deprecated in favor of grafana.oncall/escalationchain.EscalationChain
type OncallEscalationChain struct {
	pulumi.CustomResourceState

	// The name of the escalation chain.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
}

// NewOncallEscalationChain registers a new resource with the given unique name, arguments, and options.
func NewOncallEscalationChain(ctx *pulumi.Context,
	name string, args *OncallEscalationChainArgs, opts ...pulumi.ResourceOption) (*OncallEscalationChain, error) {
	if args == nil {
		args = &OncallEscalationChainArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OncallEscalationChain
	err := ctx.RegisterResource("grafana:index/oncallEscalationChain:OncallEscalationChain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOncallEscalationChain gets an existing OncallEscalationChain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOncallEscalationChain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OncallEscalationChainState, opts ...pulumi.ResourceOption) (*OncallEscalationChain, error) {
	var resource OncallEscalationChain
	err := ctx.ReadResource("grafana:index/oncallEscalationChain:OncallEscalationChain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OncallEscalationChain resources.
type oncallEscalationChainState struct {
	// The name of the escalation chain.
	Name *string `pulumi:"name"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId *string `pulumi:"teamId"`
}

type OncallEscalationChainState struct {
	// The name of the escalation chain.
	Name pulumi.StringPtrInput
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrInput
}

func (OncallEscalationChainState) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallEscalationChainState)(nil)).Elem()
}

type oncallEscalationChainArgs struct {
	// The name of the escalation chain.
	Name *string `pulumi:"name"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a OncallEscalationChain resource.
type OncallEscalationChainArgs struct {
	// The name of the escalation chain.
	Name pulumi.StringPtrInput
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrInput
}

func (OncallEscalationChainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallEscalationChainArgs)(nil)).Elem()
}

type OncallEscalationChainInput interface {
	pulumi.Input

	ToOncallEscalationChainOutput() OncallEscalationChainOutput
	ToOncallEscalationChainOutputWithContext(ctx context.Context) OncallEscalationChainOutput
}

func (*OncallEscalationChain) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallEscalationChain)(nil)).Elem()
}

func (i *OncallEscalationChain) ToOncallEscalationChainOutput() OncallEscalationChainOutput {
	return i.ToOncallEscalationChainOutputWithContext(context.Background())
}

func (i *OncallEscalationChain) ToOncallEscalationChainOutputWithContext(ctx context.Context) OncallEscalationChainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallEscalationChainOutput)
}

// OncallEscalationChainArrayInput is an input type that accepts OncallEscalationChainArray and OncallEscalationChainArrayOutput values.
// You can construct a concrete instance of `OncallEscalationChainArrayInput` via:
//
//	OncallEscalationChainArray{ OncallEscalationChainArgs{...} }
type OncallEscalationChainArrayInput interface {
	pulumi.Input

	ToOncallEscalationChainArrayOutput() OncallEscalationChainArrayOutput
	ToOncallEscalationChainArrayOutputWithContext(context.Context) OncallEscalationChainArrayOutput
}

type OncallEscalationChainArray []OncallEscalationChainInput

func (OncallEscalationChainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallEscalationChain)(nil)).Elem()
}

func (i OncallEscalationChainArray) ToOncallEscalationChainArrayOutput() OncallEscalationChainArrayOutput {
	return i.ToOncallEscalationChainArrayOutputWithContext(context.Background())
}

func (i OncallEscalationChainArray) ToOncallEscalationChainArrayOutputWithContext(ctx context.Context) OncallEscalationChainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallEscalationChainArrayOutput)
}

// OncallEscalationChainMapInput is an input type that accepts OncallEscalationChainMap and OncallEscalationChainMapOutput values.
// You can construct a concrete instance of `OncallEscalationChainMapInput` via:
//
//	OncallEscalationChainMap{ "key": OncallEscalationChainArgs{...} }
type OncallEscalationChainMapInput interface {
	pulumi.Input

	ToOncallEscalationChainMapOutput() OncallEscalationChainMapOutput
	ToOncallEscalationChainMapOutputWithContext(context.Context) OncallEscalationChainMapOutput
}

type OncallEscalationChainMap map[string]OncallEscalationChainInput

func (OncallEscalationChainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallEscalationChain)(nil)).Elem()
}

func (i OncallEscalationChainMap) ToOncallEscalationChainMapOutput() OncallEscalationChainMapOutput {
	return i.ToOncallEscalationChainMapOutputWithContext(context.Background())
}

func (i OncallEscalationChainMap) ToOncallEscalationChainMapOutputWithContext(ctx context.Context) OncallEscalationChainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallEscalationChainMapOutput)
}

type OncallEscalationChainOutput struct{ *pulumi.OutputState }

func (OncallEscalationChainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallEscalationChain)(nil)).Elem()
}

func (o OncallEscalationChainOutput) ToOncallEscalationChainOutput() OncallEscalationChainOutput {
	return o
}

func (o OncallEscalationChainOutput) ToOncallEscalationChainOutputWithContext(ctx context.Context) OncallEscalationChainOutput {
	return o
}

// The name of the escalation chain.
func (o OncallEscalationChainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OncallEscalationChain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the OnCall team (using the `onCall.getTeam` datasource).
func (o OncallEscalationChainOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalationChain) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

type OncallEscalationChainArrayOutput struct{ *pulumi.OutputState }

func (OncallEscalationChainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallEscalationChain)(nil)).Elem()
}

func (o OncallEscalationChainArrayOutput) ToOncallEscalationChainArrayOutput() OncallEscalationChainArrayOutput {
	return o
}

func (o OncallEscalationChainArrayOutput) ToOncallEscalationChainArrayOutputWithContext(ctx context.Context) OncallEscalationChainArrayOutput {
	return o
}

func (o OncallEscalationChainArrayOutput) Index(i pulumi.IntInput) OncallEscalationChainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OncallEscalationChain {
		return vs[0].([]*OncallEscalationChain)[vs[1].(int)]
	}).(OncallEscalationChainOutput)
}

type OncallEscalationChainMapOutput struct{ *pulumi.OutputState }

func (OncallEscalationChainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallEscalationChain)(nil)).Elem()
}

func (o OncallEscalationChainMapOutput) ToOncallEscalationChainMapOutput() OncallEscalationChainMapOutput {
	return o
}

func (o OncallEscalationChainMapOutput) ToOncallEscalationChainMapOutputWithContext(ctx context.Context) OncallEscalationChainMapOutput {
	return o
}

func (o OncallEscalationChainMapOutput) MapIndex(k pulumi.StringInput) OncallEscalationChainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OncallEscalationChain {
		return vs[0].(map[string]*OncallEscalationChain)[vs[1].(string)]
	}).(OncallEscalationChainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OncallEscalationChainInput)(nil)).Elem(), &OncallEscalationChain{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallEscalationChainArrayInput)(nil)).Elem(), OncallEscalationChainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallEscalationChainMapInput)(nil)).Elem(), OncallEscalationChainMap{})
	pulumi.RegisterOutputType(OncallEscalationChainOutput{})
	pulumi.RegisterOutputType(OncallEscalationChainArrayOutput{})
	pulumi.RegisterOutputType(OncallEscalationChainMapOutput{})
}
