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

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/personal_notification_rules/)
//
// **Note**: you must be running Grafana OnCall >= v1.8.0 to use this resource.
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
//			myUser, err := oncall.GetUser(ctx, &oncall.GetUserArgs{
//				Username: "my_username",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_step_1", &oncall.UserNotificationRuleArgs{
//				UserId:   pulumi.String(myUser.Id),
//				Position: pulumi.Int(0),
//				Type:     pulumi.String("notify_by_mobile_app"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_step_2", &oncall.UserNotificationRuleArgs{
//				UserId:   pulumi.String(myUser.Id),
//				Position: pulumi.Int(1),
//				Duration: pulumi.Int(600),
//				Type:     pulumi.String("wait"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_step_3", &oncall.UserNotificationRuleArgs{
//				UserId:   pulumi.String(myUser.Id),
//				Position: pulumi.Int(2),
//				Type:     pulumi.String("notify_by_phone_call"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_step_4", &oncall.UserNotificationRuleArgs{
//				UserId:   pulumi.String(myUser.Id),
//				Position: pulumi.Int(3),
//				Duration: pulumi.Int(300),
//				Type:     pulumi.String("wait"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_step_5", &oncall.UserNotificationRuleArgs{
//				UserId:   pulumi.String(myUser.Id),
//				Position: pulumi.Int(4),
//				Type:     pulumi.String("notify_by_slack"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_important_step_1", &oncall.UserNotificationRuleArgs{
//				UserId:    pulumi.String(myUser.Id),
//				Important: pulumi.Bool(true),
//				Position:  pulumi.Int(0),
//				Type:      pulumi.String("notify_by_mobile_app_critical"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_important_step_2", &oncall.UserNotificationRuleArgs{
//				UserId:    pulumi.String(myUser.Id),
//				Important: pulumi.Bool(true),
//				Position:  pulumi.Int(1),
//				Duration:  pulumi.Int(300),
//				Type:      pulumi.String("wait"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oncall.NewUserNotificationRule(ctx, "my_user_important_step_3", &oncall.UserNotificationRuleArgs{
//				UserId:    pulumi.String(myUser.Id),
//				Important: pulumi.Bool(true),
//				Position:  pulumi.Int(2),
//				Type:      pulumi.String("notify_by_mobile_app_critical"),
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
// $ pulumi import grafana:onCall/userNotificationRule:UserNotificationRule name "{{ id }}"
// ```
type UserNotificationRule struct {
	pulumi.CustomResourceState

	// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// Boolean value which indicates if a rule is “important”
	Important pulumi.BoolOutput `pulumi:"important"`
	// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
	Position pulumi.IntPtrOutput `pulumi:"position"`
	// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
	Type pulumi.StringOutput `pulumi:"type"`
	// User ID
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserNotificationRule registers a new resource with the given unique name, arguments, and options.
func NewUserNotificationRule(ctx *pulumi.Context,
	name string, args *UserNotificationRuleArgs, opts ...pulumi.ResourceOption) (*UserNotificationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/oncallUserNotificationRule:OncallUserNotificationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserNotificationRule
	err := ctx.RegisterResource("grafana:onCall/userNotificationRule:UserNotificationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserNotificationRule gets an existing UserNotificationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserNotificationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserNotificationRuleState, opts ...pulumi.ResourceOption) (*UserNotificationRule, error) {
	var resource UserNotificationRule
	err := ctx.ReadResource("grafana:onCall/userNotificationRule:UserNotificationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserNotificationRule resources.
type userNotificationRuleState struct {
	// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
	Duration *int `pulumi:"duration"`
	// Boolean value which indicates if a rule is “important”
	Important *bool `pulumi:"important"`
	// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
	Position *int `pulumi:"position"`
	// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
	Type *string `pulumi:"type"`
	// User ID
	UserId *string `pulumi:"userId"`
}

type UserNotificationRuleState struct {
	// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
	Duration pulumi.IntPtrInput
	// Boolean value which indicates if a rule is “important”
	Important pulumi.BoolPtrInput
	// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
	Position pulumi.IntPtrInput
	// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
	Type pulumi.StringPtrInput
	// User ID
	UserId pulumi.StringPtrInput
}

func (UserNotificationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*userNotificationRuleState)(nil)).Elem()
}

type userNotificationRuleArgs struct {
	// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
	Duration *int `pulumi:"duration"`
	// Boolean value which indicates if a rule is “important”
	Important *bool `pulumi:"important"`
	// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
	Position *int `pulumi:"position"`
	// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
	Type string `pulumi:"type"`
	// User ID
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserNotificationRule resource.
type UserNotificationRuleArgs struct {
	// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
	Duration pulumi.IntPtrInput
	// Boolean value which indicates if a rule is “important”
	Important pulumi.BoolPtrInput
	// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
	Position pulumi.IntPtrInput
	// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
	Type pulumi.StringInput
	// User ID
	UserId pulumi.StringInput
}

func (UserNotificationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userNotificationRuleArgs)(nil)).Elem()
}

type UserNotificationRuleInput interface {
	pulumi.Input

	ToUserNotificationRuleOutput() UserNotificationRuleOutput
	ToUserNotificationRuleOutputWithContext(ctx context.Context) UserNotificationRuleOutput
}

func (*UserNotificationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**UserNotificationRule)(nil)).Elem()
}

func (i *UserNotificationRule) ToUserNotificationRuleOutput() UserNotificationRuleOutput {
	return i.ToUserNotificationRuleOutputWithContext(context.Background())
}

func (i *UserNotificationRule) ToUserNotificationRuleOutputWithContext(ctx context.Context) UserNotificationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserNotificationRuleOutput)
}

// UserNotificationRuleArrayInput is an input type that accepts UserNotificationRuleArray and UserNotificationRuleArrayOutput values.
// You can construct a concrete instance of `UserNotificationRuleArrayInput` via:
//
//	UserNotificationRuleArray{ UserNotificationRuleArgs{...} }
type UserNotificationRuleArrayInput interface {
	pulumi.Input

	ToUserNotificationRuleArrayOutput() UserNotificationRuleArrayOutput
	ToUserNotificationRuleArrayOutputWithContext(context.Context) UserNotificationRuleArrayOutput
}

type UserNotificationRuleArray []UserNotificationRuleInput

func (UserNotificationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserNotificationRule)(nil)).Elem()
}

func (i UserNotificationRuleArray) ToUserNotificationRuleArrayOutput() UserNotificationRuleArrayOutput {
	return i.ToUserNotificationRuleArrayOutputWithContext(context.Background())
}

func (i UserNotificationRuleArray) ToUserNotificationRuleArrayOutputWithContext(ctx context.Context) UserNotificationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserNotificationRuleArrayOutput)
}

// UserNotificationRuleMapInput is an input type that accepts UserNotificationRuleMap and UserNotificationRuleMapOutput values.
// You can construct a concrete instance of `UserNotificationRuleMapInput` via:
//
//	UserNotificationRuleMap{ "key": UserNotificationRuleArgs{...} }
type UserNotificationRuleMapInput interface {
	pulumi.Input

	ToUserNotificationRuleMapOutput() UserNotificationRuleMapOutput
	ToUserNotificationRuleMapOutputWithContext(context.Context) UserNotificationRuleMapOutput
}

type UserNotificationRuleMap map[string]UserNotificationRuleInput

func (UserNotificationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserNotificationRule)(nil)).Elem()
}

func (i UserNotificationRuleMap) ToUserNotificationRuleMapOutput() UserNotificationRuleMapOutput {
	return i.ToUserNotificationRuleMapOutputWithContext(context.Background())
}

func (i UserNotificationRuleMap) ToUserNotificationRuleMapOutputWithContext(ctx context.Context) UserNotificationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserNotificationRuleMapOutput)
}

type UserNotificationRuleOutput struct{ *pulumi.OutputState }

func (UserNotificationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserNotificationRule)(nil)).Elem()
}

func (o UserNotificationRuleOutput) ToUserNotificationRuleOutput() UserNotificationRuleOutput {
	return o
}

func (o UserNotificationRuleOutput) ToUserNotificationRuleOutputWithContext(ctx context.Context) UserNotificationRuleOutput {
	return o
}

// A time in seconds to wait (when `type=wait`). Can be 60, 300, 900, 1800, 3600
func (o UserNotificationRuleOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserNotificationRule) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// Boolean value which indicates if a rule is “important”
func (o UserNotificationRuleOutput) Important() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserNotificationRule) pulumi.BoolOutput { return v.Important }).(pulumi.BoolOutput)
}

// Personal notification rules execute one after another starting from position=0. A new escalation policy created with a position of an existing escalation policy will move the old one (and all following) down on the list.
func (o UserNotificationRuleOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserNotificationRule) pulumi.IntPtrOutput { return v.Position }).(pulumi.IntPtrOutput)
}

// The type of notification rule. Can be wait, notify*by*slack, notify*by*msteams, notify*by*sms, notify*by*phone*call, notify*by*telegram, notify*by*email, notify*by*mobile*app, notify*by*mobile*app*critical. NOTE: `notifyByMsteams` is only available for Grafana Cloud customers.
func (o UserNotificationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserNotificationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// User ID
func (o UserNotificationRuleOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserNotificationRule) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserNotificationRuleArrayOutput struct{ *pulumi.OutputState }

func (UserNotificationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserNotificationRule)(nil)).Elem()
}

func (o UserNotificationRuleArrayOutput) ToUserNotificationRuleArrayOutput() UserNotificationRuleArrayOutput {
	return o
}

func (o UserNotificationRuleArrayOutput) ToUserNotificationRuleArrayOutputWithContext(ctx context.Context) UserNotificationRuleArrayOutput {
	return o
}

func (o UserNotificationRuleArrayOutput) Index(i pulumi.IntInput) UserNotificationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserNotificationRule {
		return vs[0].([]*UserNotificationRule)[vs[1].(int)]
	}).(UserNotificationRuleOutput)
}

type UserNotificationRuleMapOutput struct{ *pulumi.OutputState }

func (UserNotificationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserNotificationRule)(nil)).Elem()
}

func (o UserNotificationRuleMapOutput) ToUserNotificationRuleMapOutput() UserNotificationRuleMapOutput {
	return o
}

func (o UserNotificationRuleMapOutput) ToUserNotificationRuleMapOutputWithContext(ctx context.Context) UserNotificationRuleMapOutput {
	return o
}

func (o UserNotificationRuleMapOutput) MapIndex(k pulumi.StringInput) UserNotificationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserNotificationRule {
		return vs[0].(map[string]*UserNotificationRule)[vs[1].(string)]
	}).(UserNotificationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserNotificationRuleInput)(nil)).Elem(), &UserNotificationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserNotificationRuleArrayInput)(nil)).Elem(), UserNotificationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserNotificationRuleMapInput)(nil)).Elem(), UserNotificationRuleMap{})
	pulumi.RegisterOutputType(UserNotificationRuleOutput{})
	pulumi.RegisterOutputType(UserNotificationRuleArrayOutput{})
	pulumi.RegisterOutputType(UserNotificationRuleMapOutput{})
}
