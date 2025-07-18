// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)
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
//			_, err = oncall.NewOutgoingWebhook(ctx, "test-acc-outgoing_webhook", &oncall.OutgoingWebhookArgs{
//				Name:   pulumi.String("my outgoing webhook"),
//				Url:    pulumi.String("https://example.com/"),
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
// $ pulumi import grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook name "{{ id }}"
// ```
//
// Deprecated: grafana.index/oncalloutgoingwebhook.OncallOutgoingWebhook has been deprecated in favor of grafana.oncall/outgoingwebhook.OutgoingWebhook
type OncallOutgoingWebhook struct {
	pulumi.CustomResourceState

	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader pulumi.StringPtrOutput `pulumi:"authorizationHeader"`
	// The data of the webhook.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// Toggle to send the entire webhook payload instead of using the values in the Data field.
	ForwardWholePayload pulumi.BoolPtrOutput `pulumi:"forwardWholePayload"`
	// Headers to add to the outgoing webhook request.
	Headers pulumi.StringPtrOutput `pulumi:"headers"`
	// The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
	HttpMethod pulumi.StringPtrOutput `pulumi:"httpMethod"`
	// Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
	IntegrationFilters pulumi.StringArrayOutput `pulumi:"integrationFilters"`
	// Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
	IsWebhookEnabled pulumi.BoolPtrOutput `pulumi:"isWebhookEnabled"`
	// The name of the outgoing webhook.
	Name pulumi.StringOutput `pulumi:"name"`
	// The auth data of the webhook. Used for Basic authentication
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// A template used to dynamically determine whether the webhook should execute based on the content of the payload.
	TriggerTemplate pulumi.StringPtrOutput `pulumi:"triggerTemplate"`
	// The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
	TriggerType pulumi.StringPtrOutput `pulumi:"triggerType"`
	// The webhook URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// Username to use when making the outgoing webhook request.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewOncallOutgoingWebhook registers a new resource with the given unique name, arguments, and options.
func NewOncallOutgoingWebhook(ctx *pulumi.Context,
	name string, args *OncallOutgoingWebhookArgs, opts ...pulumi.ResourceOption) (*OncallOutgoingWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.AuthorizationHeader != nil {
		args.AuthorizationHeader = pulumi.ToSecret(args.AuthorizationHeader).(pulumi.StringPtrInput)
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authorizationHeader",
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OncallOutgoingWebhook
	err := ctx.RegisterResource("grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOncallOutgoingWebhook gets an existing OncallOutgoingWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOncallOutgoingWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OncallOutgoingWebhookState, opts ...pulumi.ResourceOption) (*OncallOutgoingWebhook, error) {
	var resource OncallOutgoingWebhook
	err := ctx.ReadResource("grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OncallOutgoingWebhook resources.
type oncallOutgoingWebhookState struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader *string `pulumi:"authorizationHeader"`
	// The data of the webhook.
	Data *string `pulumi:"data"`
	// Toggle to send the entire webhook payload instead of using the values in the Data field.
	ForwardWholePayload *bool `pulumi:"forwardWholePayload"`
	// Headers to add to the outgoing webhook request.
	Headers *string `pulumi:"headers"`
	// The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
	HttpMethod *string `pulumi:"httpMethod"`
	// Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
	IntegrationFilters []string `pulumi:"integrationFilters"`
	// Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
	IsWebhookEnabled *bool `pulumi:"isWebhookEnabled"`
	// The name of the outgoing webhook.
	Name *string `pulumi:"name"`
	// The auth data of the webhook. Used for Basic authentication
	Password *string `pulumi:"password"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId *string `pulumi:"teamId"`
	// A template used to dynamically determine whether the webhook should execute based on the content of the payload.
	TriggerTemplate *string `pulumi:"triggerTemplate"`
	// The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
	TriggerType *string `pulumi:"triggerType"`
	// The webhook URL.
	Url *string `pulumi:"url"`
	// Username to use when making the outgoing webhook request.
	User *string `pulumi:"user"`
}

type OncallOutgoingWebhookState struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader pulumi.StringPtrInput
	// The data of the webhook.
	Data pulumi.StringPtrInput
	// Toggle to send the entire webhook payload instead of using the values in the Data field.
	ForwardWholePayload pulumi.BoolPtrInput
	// Headers to add to the outgoing webhook request.
	Headers pulumi.StringPtrInput
	// The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
	HttpMethod pulumi.StringPtrInput
	// Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
	IntegrationFilters pulumi.StringArrayInput
	// Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
	IsWebhookEnabled pulumi.BoolPtrInput
	// The name of the outgoing webhook.
	Name pulumi.StringPtrInput
	// The auth data of the webhook. Used for Basic authentication
	Password pulumi.StringPtrInput
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrInput
	// A template used to dynamically determine whether the webhook should execute based on the content of the payload.
	TriggerTemplate pulumi.StringPtrInput
	// The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
	TriggerType pulumi.StringPtrInput
	// The webhook URL.
	Url pulumi.StringPtrInput
	// Username to use when making the outgoing webhook request.
	User pulumi.StringPtrInput
}

func (OncallOutgoingWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallOutgoingWebhookState)(nil)).Elem()
}

type oncallOutgoingWebhookArgs struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader *string `pulumi:"authorizationHeader"`
	// The data of the webhook.
	Data *string `pulumi:"data"`
	// Toggle to send the entire webhook payload instead of using the values in the Data field.
	ForwardWholePayload *bool `pulumi:"forwardWholePayload"`
	// Headers to add to the outgoing webhook request.
	Headers *string `pulumi:"headers"`
	// The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
	HttpMethod *string `pulumi:"httpMethod"`
	// Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
	IntegrationFilters []string `pulumi:"integrationFilters"`
	// Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
	IsWebhookEnabled *bool `pulumi:"isWebhookEnabled"`
	// The name of the outgoing webhook.
	Name *string `pulumi:"name"`
	// The auth data of the webhook. Used for Basic authentication
	Password *string `pulumi:"password"`
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId *string `pulumi:"teamId"`
	// A template used to dynamically determine whether the webhook should execute based on the content of the payload.
	TriggerTemplate *string `pulumi:"triggerTemplate"`
	// The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
	TriggerType *string `pulumi:"triggerType"`
	// The webhook URL.
	Url string `pulumi:"url"`
	// Username to use when making the outgoing webhook request.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a OncallOutgoingWebhook resource.
type OncallOutgoingWebhookArgs struct {
	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	AuthorizationHeader pulumi.StringPtrInput
	// The data of the webhook.
	Data pulumi.StringPtrInput
	// Toggle to send the entire webhook payload instead of using the values in the Data field.
	ForwardWholePayload pulumi.BoolPtrInput
	// Headers to add to the outgoing webhook request.
	Headers pulumi.StringPtrInput
	// The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
	HttpMethod pulumi.StringPtrInput
	// Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
	IntegrationFilters pulumi.StringArrayInput
	// Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
	IsWebhookEnabled pulumi.BoolPtrInput
	// The name of the outgoing webhook.
	Name pulumi.StringPtrInput
	// The auth data of the webhook. Used for Basic authentication
	Password pulumi.StringPtrInput
	// The ID of the OnCall team (using the `onCall.getTeam` datasource).
	TeamId pulumi.StringPtrInput
	// A template used to dynamically determine whether the webhook should execute based on the content of the payload.
	TriggerTemplate pulumi.StringPtrInput
	// The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
	TriggerType pulumi.StringPtrInput
	// The webhook URL.
	Url pulumi.StringInput
	// Username to use when making the outgoing webhook request.
	User pulumi.StringPtrInput
}

func (OncallOutgoingWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallOutgoingWebhookArgs)(nil)).Elem()
}

type OncallOutgoingWebhookInput interface {
	pulumi.Input

	ToOncallOutgoingWebhookOutput() OncallOutgoingWebhookOutput
	ToOncallOutgoingWebhookOutputWithContext(ctx context.Context) OncallOutgoingWebhookOutput
}

func (*OncallOutgoingWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallOutgoingWebhook)(nil)).Elem()
}

func (i *OncallOutgoingWebhook) ToOncallOutgoingWebhookOutput() OncallOutgoingWebhookOutput {
	return i.ToOncallOutgoingWebhookOutputWithContext(context.Background())
}

func (i *OncallOutgoingWebhook) ToOncallOutgoingWebhookOutputWithContext(ctx context.Context) OncallOutgoingWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallOutgoingWebhookOutput)
}

// OncallOutgoingWebhookArrayInput is an input type that accepts OncallOutgoingWebhookArray and OncallOutgoingWebhookArrayOutput values.
// You can construct a concrete instance of `OncallOutgoingWebhookArrayInput` via:
//
//	OncallOutgoingWebhookArray{ OncallOutgoingWebhookArgs{...} }
type OncallOutgoingWebhookArrayInput interface {
	pulumi.Input

	ToOncallOutgoingWebhookArrayOutput() OncallOutgoingWebhookArrayOutput
	ToOncallOutgoingWebhookArrayOutputWithContext(context.Context) OncallOutgoingWebhookArrayOutput
}

type OncallOutgoingWebhookArray []OncallOutgoingWebhookInput

func (OncallOutgoingWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallOutgoingWebhook)(nil)).Elem()
}

func (i OncallOutgoingWebhookArray) ToOncallOutgoingWebhookArrayOutput() OncallOutgoingWebhookArrayOutput {
	return i.ToOncallOutgoingWebhookArrayOutputWithContext(context.Background())
}

func (i OncallOutgoingWebhookArray) ToOncallOutgoingWebhookArrayOutputWithContext(ctx context.Context) OncallOutgoingWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallOutgoingWebhookArrayOutput)
}

// OncallOutgoingWebhookMapInput is an input type that accepts OncallOutgoingWebhookMap and OncallOutgoingWebhookMapOutput values.
// You can construct a concrete instance of `OncallOutgoingWebhookMapInput` via:
//
//	OncallOutgoingWebhookMap{ "key": OncallOutgoingWebhookArgs{...} }
type OncallOutgoingWebhookMapInput interface {
	pulumi.Input

	ToOncallOutgoingWebhookMapOutput() OncallOutgoingWebhookMapOutput
	ToOncallOutgoingWebhookMapOutputWithContext(context.Context) OncallOutgoingWebhookMapOutput
}

type OncallOutgoingWebhookMap map[string]OncallOutgoingWebhookInput

func (OncallOutgoingWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallOutgoingWebhook)(nil)).Elem()
}

func (i OncallOutgoingWebhookMap) ToOncallOutgoingWebhookMapOutput() OncallOutgoingWebhookMapOutput {
	return i.ToOncallOutgoingWebhookMapOutputWithContext(context.Background())
}

func (i OncallOutgoingWebhookMap) ToOncallOutgoingWebhookMapOutputWithContext(ctx context.Context) OncallOutgoingWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallOutgoingWebhookMapOutput)
}

type OncallOutgoingWebhookOutput struct{ *pulumi.OutputState }

func (OncallOutgoingWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallOutgoingWebhook)(nil)).Elem()
}

func (o OncallOutgoingWebhookOutput) ToOncallOutgoingWebhookOutput() OncallOutgoingWebhookOutput {
	return o
}

func (o OncallOutgoingWebhookOutput) ToOncallOutgoingWebhookOutputWithContext(ctx context.Context) OncallOutgoingWebhookOutput {
	return o
}

// The auth data of the webhook. Used in Authorization header instead of user/password auth.
func (o OncallOutgoingWebhookOutput) AuthorizationHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.AuthorizationHeader }).(pulumi.StringPtrOutput)
}

// The data of the webhook.
func (o OncallOutgoingWebhookOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.Data }).(pulumi.StringPtrOutput)
}

// Toggle to send the entire webhook payload instead of using the values in the Data field.
func (o OncallOutgoingWebhookOutput) ForwardWholePayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.BoolPtrOutput { return v.ForwardWholePayload }).(pulumi.BoolPtrOutput)
}

// Headers to add to the outgoing webhook request.
func (o OncallOutgoingWebhookOutput) Headers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.Headers }).(pulumi.StringPtrOutput)
}

// The HTTP method used in the request made by the outgoing webhook. Defaults to `POST`.
func (o OncallOutgoingWebhookOutput) HttpMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.HttpMethod }).(pulumi.StringPtrOutput)
}

// Restricts the outgoing webhook to only trigger if the event came from a selected integration. If no integrations are selected the outgoing webhook will trigger for any integration.
func (o OncallOutgoingWebhookOutput) IntegrationFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringArrayOutput { return v.IntegrationFilters }).(pulumi.StringArrayOutput)
}

// Controls whether the outgoing webhook will trigger or is ignored. Defaults to `true`.
func (o OncallOutgoingWebhookOutput) IsWebhookEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.BoolPtrOutput { return v.IsWebhookEnabled }).(pulumi.BoolPtrOutput)
}

// The name of the outgoing webhook.
func (o OncallOutgoingWebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The auth data of the webhook. Used for Basic authentication
func (o OncallOutgoingWebhookOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The ID of the OnCall team (using the `onCall.getTeam` datasource).
func (o OncallOutgoingWebhookOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

// A template used to dynamically determine whether the webhook should execute based on the content of the payload.
func (o OncallOutgoingWebhookOutput) TriggerTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.TriggerTemplate }).(pulumi.StringPtrOutput)
}

// The type of event that will cause this outgoing webhook to execute. The types of triggers are: `escalation`, `alert group created`, `acknowledge`, `resolve`, `silence`, `unsilence`, `unresolve`, `unacknowledge`. Defaults to `escalation`.
func (o OncallOutgoingWebhookOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.TriggerType }).(pulumi.StringPtrOutput)
}

// The webhook URL.
func (o OncallOutgoingWebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Username to use when making the outgoing webhook request.
func (o OncallOutgoingWebhookOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallOutgoingWebhook) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type OncallOutgoingWebhookArrayOutput struct{ *pulumi.OutputState }

func (OncallOutgoingWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallOutgoingWebhook)(nil)).Elem()
}

func (o OncallOutgoingWebhookArrayOutput) ToOncallOutgoingWebhookArrayOutput() OncallOutgoingWebhookArrayOutput {
	return o
}

func (o OncallOutgoingWebhookArrayOutput) ToOncallOutgoingWebhookArrayOutputWithContext(ctx context.Context) OncallOutgoingWebhookArrayOutput {
	return o
}

func (o OncallOutgoingWebhookArrayOutput) Index(i pulumi.IntInput) OncallOutgoingWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OncallOutgoingWebhook {
		return vs[0].([]*OncallOutgoingWebhook)[vs[1].(int)]
	}).(OncallOutgoingWebhookOutput)
}

type OncallOutgoingWebhookMapOutput struct{ *pulumi.OutputState }

func (OncallOutgoingWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallOutgoingWebhook)(nil)).Elem()
}

func (o OncallOutgoingWebhookMapOutput) ToOncallOutgoingWebhookMapOutput() OncallOutgoingWebhookMapOutput {
	return o
}

func (o OncallOutgoingWebhookMapOutput) ToOncallOutgoingWebhookMapOutputWithContext(ctx context.Context) OncallOutgoingWebhookMapOutput {
	return o
}

func (o OncallOutgoingWebhookMapOutput) MapIndex(k pulumi.StringInput) OncallOutgoingWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OncallOutgoingWebhook {
		return vs[0].(map[string]*OncallOutgoingWebhook)[vs[1].(string)]
	}).(OncallOutgoingWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OncallOutgoingWebhookInput)(nil)).Elem(), &OncallOutgoingWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallOutgoingWebhookArrayInput)(nil)).Elem(), OncallOutgoingWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallOutgoingWebhookMapInput)(nil)).Elem(), OncallOutgoingWebhookMap{})
	pulumi.RegisterOutputType(OncallOutgoingWebhookOutput{})
	pulumi.RegisterOutputType(OncallOutgoingWebhookArrayOutput{})
	pulumi.RegisterOutputType(OncallOutgoingWebhookMapOutput{})
}
