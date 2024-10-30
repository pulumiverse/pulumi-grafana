// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
//
// This resource represents an instance-scoped resource and uses Grafana's admin APIs.
// It does not work with API tokens or service accounts which are org-scoped.
// You must use basic auth.
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
//			_, err := oss.NewUser(ctx, "staff", &oss.UserArgs{
//				Email:    pulumi.String("staff.name@example.com"),
//				Name:     pulumi.String("Staff Name"),
//				Login:    pulumi.String("staff"),
//				Password: pulumi.String("my-password"),
//				IsAdmin:  pulumi.Bool(false),
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
// $ pulumi import grafana:index/user:User name "{{ id }}"
// ```
//
// Deprecated: grafana.index/user.User has been deprecated in favor of grafana.oss/user.User
type User struct {
	pulumi.CustomResourceState

	// The email address of the Grafana user.
	Email pulumi.StringOutput `pulumi:"email"`
	// Whether to make user an admin. Defaults to `false`.
	IsAdmin pulumi.BoolPtrOutput `pulumi:"isAdmin"`
	// The username for the Grafana user.
	Login pulumi.StringPtrOutput `pulumi:"login"`
	// The display name for the Grafana user.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the Grafana user.
	Password pulumi.StringOutput `pulumi:"password"`
	// The numerical ID of the Grafana user.
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/user:User"),
		},
	})
	opts = append(opts, aliases)
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("grafana:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("grafana:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The email address of the Grafana user.
	Email *string `pulumi:"email"`
	// Whether to make user an admin. Defaults to `false`.
	IsAdmin *bool `pulumi:"isAdmin"`
	// The username for the Grafana user.
	Login *string `pulumi:"login"`
	// The display name for the Grafana user.
	Name *string `pulumi:"name"`
	// The password for the Grafana user.
	Password *string `pulumi:"password"`
	// The numerical ID of the Grafana user.
	UserId *int `pulumi:"userId"`
}

type UserState struct {
	// The email address of the Grafana user.
	Email pulumi.StringPtrInput
	// Whether to make user an admin. Defaults to `false`.
	IsAdmin pulumi.BoolPtrInput
	// The username for the Grafana user.
	Login pulumi.StringPtrInput
	// The display name for the Grafana user.
	Name pulumi.StringPtrInput
	// The password for the Grafana user.
	Password pulumi.StringPtrInput
	// The numerical ID of the Grafana user.
	UserId pulumi.IntPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The email address of the Grafana user.
	Email string `pulumi:"email"`
	// Whether to make user an admin. Defaults to `false`.
	IsAdmin *bool `pulumi:"isAdmin"`
	// The username for the Grafana user.
	Login *string `pulumi:"login"`
	// The display name for the Grafana user.
	Name *string `pulumi:"name"`
	// The password for the Grafana user.
	Password string `pulumi:"password"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The email address of the Grafana user.
	Email pulumi.StringInput
	// Whether to make user an admin. Defaults to `false`.
	IsAdmin pulumi.BoolPtrInput
	// The username for the Grafana user.
	Login pulumi.StringPtrInput
	// The display name for the Grafana user.
	Name pulumi.StringPtrInput
	// The password for the Grafana user.
	Password pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// The email address of the Grafana user.
func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Whether to make user an admin. Defaults to `false`.
func (o UserOutput) IsAdmin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IsAdmin }).(pulumi.BoolPtrOutput)
}

// The username for the Grafana user.
func (o UserOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

// The display name for the Grafana user.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for the Grafana user.
func (o UserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The numerical ID of the Grafana user.
func (o UserOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *User) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
