// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

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
// $ pulumi import grafana:onCall/escalationChain:EscalationChain name "{{ id }}"
// ```
type EscalationChain struct {
	pulumi.CustomResourceState

	// The name of the escalation chain.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
}

// NewEscalationChain registers a new resource with the given unique name, arguments, and options.
func NewEscalationChain(ctx *pulumi.Context,
	name string, args *EscalationChainArgs, opts ...pulumi.ResourceOption) (*EscalationChain, error) {
	if args == nil {
		args = &EscalationChainArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/oncallEscalationChain:OncallEscalationChain"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EscalationChain
	err := ctx.RegisterResource("grafana:onCall/escalationChain:EscalationChain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEscalationChain gets an existing EscalationChain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEscalationChain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EscalationChainState, opts ...pulumi.ResourceOption) (*EscalationChain, error) {
	var resource EscalationChain
	err := ctx.ReadResource("grafana:onCall/escalationChain:EscalationChain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EscalationChain resources.
type escalationChainState struct {
	// The name of the escalation chain.
	Name *string `pulumi:"name"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId *string `pulumi:"teamId"`
}

type EscalationChainState struct {
	// The name of the escalation chain.
	Name pulumi.StringPtrInput
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrInput
}

func (EscalationChainState) ElementType() reflect.Type {
	return reflect.TypeOf((*escalationChainState)(nil)).Elem()
}

type escalationChainArgs struct {
	// The name of the escalation chain.
	Name *string `pulumi:"name"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId *string `pulumi:"teamId"`
}

// The set of arguments for constructing a EscalationChain resource.
type EscalationChainArgs struct {
	// The name of the escalation chain.
	Name pulumi.StringPtrInput
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrInput
}

func (EscalationChainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*escalationChainArgs)(nil)).Elem()
}

type EscalationChainInput interface {
	pulumi.Input

	ToEscalationChainOutput() EscalationChainOutput
	ToEscalationChainOutputWithContext(ctx context.Context) EscalationChainOutput
}

func (*EscalationChain) ElementType() reflect.Type {
	return reflect.TypeOf((**EscalationChain)(nil)).Elem()
}

func (i *EscalationChain) ToEscalationChainOutput() EscalationChainOutput {
	return i.ToEscalationChainOutputWithContext(context.Background())
}

func (i *EscalationChain) ToEscalationChainOutputWithContext(ctx context.Context) EscalationChainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EscalationChainOutput)
}

// EscalationChainArrayInput is an input type that accepts EscalationChainArray and EscalationChainArrayOutput values.
// You can construct a concrete instance of `EscalationChainArrayInput` via:
//
//	EscalationChainArray{ EscalationChainArgs{...} }
type EscalationChainArrayInput interface {
	pulumi.Input

	ToEscalationChainArrayOutput() EscalationChainArrayOutput
	ToEscalationChainArrayOutputWithContext(context.Context) EscalationChainArrayOutput
}

type EscalationChainArray []EscalationChainInput

func (EscalationChainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EscalationChain)(nil)).Elem()
}

func (i EscalationChainArray) ToEscalationChainArrayOutput() EscalationChainArrayOutput {
	return i.ToEscalationChainArrayOutputWithContext(context.Background())
}

func (i EscalationChainArray) ToEscalationChainArrayOutputWithContext(ctx context.Context) EscalationChainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EscalationChainArrayOutput)
}

// EscalationChainMapInput is an input type that accepts EscalationChainMap and EscalationChainMapOutput values.
// You can construct a concrete instance of `EscalationChainMapInput` via:
//
//	EscalationChainMap{ "key": EscalationChainArgs{...} }
type EscalationChainMapInput interface {
	pulumi.Input

	ToEscalationChainMapOutput() EscalationChainMapOutput
	ToEscalationChainMapOutputWithContext(context.Context) EscalationChainMapOutput
}

type EscalationChainMap map[string]EscalationChainInput

func (EscalationChainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EscalationChain)(nil)).Elem()
}

func (i EscalationChainMap) ToEscalationChainMapOutput() EscalationChainMapOutput {
	return i.ToEscalationChainMapOutputWithContext(context.Background())
}

func (i EscalationChainMap) ToEscalationChainMapOutputWithContext(ctx context.Context) EscalationChainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EscalationChainMapOutput)
}

type EscalationChainOutput struct{ *pulumi.OutputState }

func (EscalationChainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EscalationChain)(nil)).Elem()
}

func (o EscalationChainOutput) ToEscalationChainOutput() EscalationChainOutput {
	return o
}

func (o EscalationChainOutput) ToEscalationChainOutputWithContext(ctx context.Context) EscalationChainOutput {
	return o
}

// The name of the escalation chain.
func (o EscalationChainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EscalationChain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the OnCall team (using the `onCall.getTeam` datasource).
func (o EscalationChainOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EscalationChain) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

type EscalationChainArrayOutput struct{ *pulumi.OutputState }

func (EscalationChainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EscalationChain)(nil)).Elem()
}

func (o EscalationChainArrayOutput) ToEscalationChainArrayOutput() EscalationChainArrayOutput {
	return o
}

func (o EscalationChainArrayOutput) ToEscalationChainArrayOutputWithContext(ctx context.Context) EscalationChainArrayOutput {
	return o
}

func (o EscalationChainArrayOutput) Index(i pulumi.IntInput) EscalationChainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EscalationChain {
		return vs[0].([]*EscalationChain)[vs[1].(int)]
	}).(EscalationChainOutput)
}

type EscalationChainMapOutput struct{ *pulumi.OutputState }

func (EscalationChainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EscalationChain)(nil)).Elem()
}

func (o EscalationChainMapOutput) ToEscalationChainMapOutput() EscalationChainMapOutput {
	return o
}

func (o EscalationChainMapOutput) ToEscalationChainMapOutputWithContext(ctx context.Context) EscalationChainMapOutput {
	return o
}

func (o EscalationChainMapOutput) MapIndex(k pulumi.StringInput) EscalationChainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EscalationChain {
		return vs[0].(map[string]*EscalationChain)[vs[1].(string)]
	}).(EscalationChainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EscalationChainInput)(nil)).Elem(), &EscalationChain{})
	pulumi.RegisterInputType(reflect.TypeOf((*EscalationChainArrayInput)(nil)).Elem(), EscalationChainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EscalationChainMapInput)(nil)).Elem(), EscalationChainMap{})
	pulumi.RegisterOutputType(EscalationChainOutput{})
	pulumi.RegisterOutputType(EscalationChainArrayOutput{})
	pulumi.RegisterOutputType(EscalationChainMapOutput{})
}
