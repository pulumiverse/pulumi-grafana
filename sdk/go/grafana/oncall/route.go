// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/routes/)
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
//			exampleSlackChannel, err := onCall.GetSlackChannel(ctx, &oncall.GetSlackChannelArgs{
//				Name: "example_slack_channel",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = onCall.NewEscalationChain(ctx, "default", &onCall.EscalationChainArgs{
//				Name: pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleIntegration, err := onCall.NewIntegration(ctx, "example_integration", &onCall.IntegrationArgs{
//				Name:         pulumi.String("Grafana Integration"),
//				Type:         pulumi.String("grafana"),
//				DefaultRoute: nil,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = onCall.NewRoute(ctx, "example_route", &onCall.RouteArgs{
//				IntegrationId:     exampleIntegration.ID(),
//				EscalationChainId: _default.ID(),
//				RoutingRegex:      pulumi.String("us-(east|west)"),
//				Position:          pulumi.Int(0),
//				Slack: &oncall.RouteSlackArgs{
//					ChannelId: pulumi.String(exampleSlackChannel.SlackId),
//					Enabled:   pulumi.Bool(true),
//				},
//				Telegram: &oncall.RouteTelegramArgs{
//					Id:      pulumi.String("ONCALLTELEGRAMID"),
//					Enabled: pulumi.Bool(true),
//				},
//				Msteams: &oncall.RouteMsteamsArgs{
//					Id:      pulumi.String("ONCALLMSTEAMSID"),
//					Enabled: pulumi.Bool(false),
//				},
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
// $ pulumi import grafana:onCall/route:Route name "{{ id }}"
// ```
type Route struct {
	pulumi.CustomResourceState

	// The ID of the escalation chain.
	EscalationChainId pulumi.StringOutput `pulumi:"escalationChainId"`
	// The ID of the integration.
	IntegrationId pulumi.StringOutput `pulumi:"integrationId"`
	// MS teams-specific settings for a route.
	Msteams RouteMsteamsPtrOutput `pulumi:"msteams"`
	// The position of the route (starts from 0).
	Position pulumi.IntOutput `pulumi:"position"`
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex pulumi.StringOutput `pulumi:"routingRegex"`
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType pulumi.StringPtrOutput `pulumi:"routingType"`
	// Slack-specific settings for a route.
	Slack RouteSlackPtrOutput `pulumi:"slack"`
	// Telegram-specific settings for a route.
	Telegram RouteTelegramPtrOutput `pulumi:"telegram"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EscalationChainId == nil {
		return nil, errors.New("invalid value for required argument 'EscalationChainId'")
	}
	if args.IntegrationId == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationId'")
	}
	if args.Position == nil {
		return nil, errors.New("invalid value for required argument 'Position'")
	}
	if args.RoutingRegex == nil {
		return nil, errors.New("invalid value for required argument 'RoutingRegex'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/oncallRoute:OncallRoute"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("grafana:onCall/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("grafana:onCall/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The ID of the escalation chain.
	EscalationChainId *string `pulumi:"escalationChainId"`
	// The ID of the integration.
	IntegrationId *string `pulumi:"integrationId"`
	// MS teams-specific settings for a route.
	Msteams *RouteMsteams `pulumi:"msteams"`
	// The position of the route (starts from 0).
	Position *int `pulumi:"position"`
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex *string `pulumi:"routingRegex"`
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType *string `pulumi:"routingType"`
	// Slack-specific settings for a route.
	Slack *RouteSlack `pulumi:"slack"`
	// Telegram-specific settings for a route.
	Telegram *RouteTelegram `pulumi:"telegram"`
}

type RouteState struct {
	// The ID of the escalation chain.
	EscalationChainId pulumi.StringPtrInput
	// The ID of the integration.
	IntegrationId pulumi.StringPtrInput
	// MS teams-specific settings for a route.
	Msteams RouteMsteamsPtrInput
	// The position of the route (starts from 0).
	Position pulumi.IntPtrInput
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex pulumi.StringPtrInput
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType pulumi.StringPtrInput
	// Slack-specific settings for a route.
	Slack RouteSlackPtrInput
	// Telegram-specific settings for a route.
	Telegram RouteTelegramPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The ID of the escalation chain.
	EscalationChainId string `pulumi:"escalationChainId"`
	// The ID of the integration.
	IntegrationId string `pulumi:"integrationId"`
	// MS teams-specific settings for a route.
	Msteams *RouteMsteams `pulumi:"msteams"`
	// The position of the route (starts from 0).
	Position int `pulumi:"position"`
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex string `pulumi:"routingRegex"`
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType *string `pulumi:"routingType"`
	// Slack-specific settings for a route.
	Slack *RouteSlack `pulumi:"slack"`
	// Telegram-specific settings for a route.
	Telegram *RouteTelegram `pulumi:"telegram"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The ID of the escalation chain.
	EscalationChainId pulumi.StringInput
	// The ID of the integration.
	IntegrationId pulumi.StringInput
	// MS teams-specific settings for a route.
	Msteams RouteMsteamsPtrInput
	// The position of the route (starts from 0).
	Position pulumi.IntInput
	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	RoutingRegex pulumi.StringInput
	// The type of route. Can be jinja2, regex Defaults to `regex`.
	RoutingType pulumi.StringPtrInput
	// Slack-specific settings for a route.
	Slack RouteSlackPtrInput
	// Telegram-specific settings for a route.
	Telegram RouteTelegramPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//	RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//	RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

// The ID of the escalation chain.
func (o RouteOutput) EscalationChainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.EscalationChainId }).(pulumi.StringOutput)
}

// The ID of the integration.
func (o RouteOutput) IntegrationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.IntegrationId }).(pulumi.StringOutput)
}

// MS teams-specific settings for a route.
func (o RouteOutput) Msteams() RouteMsteamsPtrOutput {
	return o.ApplyT(func(v *Route) RouteMsteamsPtrOutput { return v.Msteams }).(RouteMsteamsPtrOutput)
}

// The position of the route (starts from 0).
func (o RouteOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *Route) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
func (o RouteOutput) RoutingRegex() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.RoutingRegex }).(pulumi.StringOutput)
}

// The type of route. Can be jinja2, regex Defaults to `regex`.
func (o RouteOutput) RoutingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.RoutingType }).(pulumi.StringPtrOutput)
}

// Slack-specific settings for a route.
func (o RouteOutput) Slack() RouteSlackPtrOutput {
	return o.ApplyT(func(v *Route) RouteSlackPtrOutput { return v.Slack }).(RouteSlackPtrOutput)
}

// Telegram-specific settings for a route.
func (o RouteOutput) Telegram() RouteTelegramPtrOutput {
	return o.ApplyT(func(v *Route) RouteTelegramPtrOutput { return v.Telegram }).(RouteTelegramPtrOutput)
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Route {
		return vs[0].([]*Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Route {
		return vs[0].(map[string]*Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteArrayInput)(nil)).Elem(), RouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapInput)(nil)).Elem(), RouteMap{})
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
