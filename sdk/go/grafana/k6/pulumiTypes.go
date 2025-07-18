// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package k6

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

var _ = internal.GetEnvOrDefault

type GetLoadTestsLoadTest struct {
	BaselineTestRunId string `pulumi:"baselineTestRunId"`
	Created           string `pulumi:"created"`
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ProjectId         string `pulumi:"projectId"`
	Script            string `pulumi:"script"`
	Updated           string `pulumi:"updated"`
}

// GetLoadTestsLoadTestInput is an input type that accepts GetLoadTestsLoadTestArgs and GetLoadTestsLoadTestOutput values.
// You can construct a concrete instance of `GetLoadTestsLoadTestInput` via:
//
//	GetLoadTestsLoadTestArgs{...}
type GetLoadTestsLoadTestInput interface {
	pulumi.Input

	ToGetLoadTestsLoadTestOutput() GetLoadTestsLoadTestOutput
	ToGetLoadTestsLoadTestOutputWithContext(context.Context) GetLoadTestsLoadTestOutput
}

type GetLoadTestsLoadTestArgs struct {
	BaselineTestRunId pulumi.StringInput `pulumi:"baselineTestRunId"`
	Created           pulumi.StringInput `pulumi:"created"`
	Id                pulumi.StringInput `pulumi:"id"`
	Name              pulumi.StringInput `pulumi:"name"`
	ProjectId         pulumi.StringInput `pulumi:"projectId"`
	Script            pulumi.StringInput `pulumi:"script"`
	Updated           pulumi.StringInput `pulumi:"updated"`
}

func (GetLoadTestsLoadTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadTestsLoadTest)(nil)).Elem()
}

func (i GetLoadTestsLoadTestArgs) ToGetLoadTestsLoadTestOutput() GetLoadTestsLoadTestOutput {
	return i.ToGetLoadTestsLoadTestOutputWithContext(context.Background())
}

func (i GetLoadTestsLoadTestArgs) ToGetLoadTestsLoadTestOutputWithContext(ctx context.Context) GetLoadTestsLoadTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetLoadTestsLoadTestOutput)
}

// GetLoadTestsLoadTestArrayInput is an input type that accepts GetLoadTestsLoadTestArray and GetLoadTestsLoadTestArrayOutput values.
// You can construct a concrete instance of `GetLoadTestsLoadTestArrayInput` via:
//
//	GetLoadTestsLoadTestArray{ GetLoadTestsLoadTestArgs{...} }
type GetLoadTestsLoadTestArrayInput interface {
	pulumi.Input

	ToGetLoadTestsLoadTestArrayOutput() GetLoadTestsLoadTestArrayOutput
	ToGetLoadTestsLoadTestArrayOutputWithContext(context.Context) GetLoadTestsLoadTestArrayOutput
}

type GetLoadTestsLoadTestArray []GetLoadTestsLoadTestInput

func (GetLoadTestsLoadTestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetLoadTestsLoadTest)(nil)).Elem()
}

func (i GetLoadTestsLoadTestArray) ToGetLoadTestsLoadTestArrayOutput() GetLoadTestsLoadTestArrayOutput {
	return i.ToGetLoadTestsLoadTestArrayOutputWithContext(context.Background())
}

func (i GetLoadTestsLoadTestArray) ToGetLoadTestsLoadTestArrayOutputWithContext(ctx context.Context) GetLoadTestsLoadTestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetLoadTestsLoadTestArrayOutput)
}

type GetLoadTestsLoadTestOutput struct{ *pulumi.OutputState }

func (GetLoadTestsLoadTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadTestsLoadTest)(nil)).Elem()
}

func (o GetLoadTestsLoadTestOutput) ToGetLoadTestsLoadTestOutput() GetLoadTestsLoadTestOutput {
	return o
}

func (o GetLoadTestsLoadTestOutput) ToGetLoadTestsLoadTestOutputWithContext(ctx context.Context) GetLoadTestsLoadTestOutput {
	return o
}

func (o GetLoadTestsLoadTestOutput) BaselineTestRunId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.BaselineTestRunId }).(pulumi.StringOutput)
}

func (o GetLoadTestsLoadTestOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.Created }).(pulumi.StringOutput)
}

func (o GetLoadTestsLoadTestOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLoadTestsLoadTestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetLoadTestsLoadTestOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetLoadTestsLoadTestOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.Script }).(pulumi.StringOutput)
}

func (o GetLoadTestsLoadTestOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadTestsLoadTest) string { return v.Updated }).(pulumi.StringOutput)
}

type GetLoadTestsLoadTestArrayOutput struct{ *pulumi.OutputState }

func (GetLoadTestsLoadTestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetLoadTestsLoadTest)(nil)).Elem()
}

func (o GetLoadTestsLoadTestArrayOutput) ToGetLoadTestsLoadTestArrayOutput() GetLoadTestsLoadTestArrayOutput {
	return o
}

func (o GetLoadTestsLoadTestArrayOutput) ToGetLoadTestsLoadTestArrayOutputWithContext(ctx context.Context) GetLoadTestsLoadTestArrayOutput {
	return o
}

func (o GetLoadTestsLoadTestArrayOutput) Index(i pulumi.IntInput) GetLoadTestsLoadTestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetLoadTestsLoadTest {
		return vs[0].([]GetLoadTestsLoadTest)[vs[1].(int)]
	}).(GetLoadTestsLoadTestOutput)
}

type GetProjectsProject struct {
	Created          string `pulumi:"created"`
	GrafanaFolderUid string `pulumi:"grafanaFolderUid"`
	Id               string `pulumi:"id"`
	IsDefault        bool   `pulumi:"isDefault"`
	Name             string `pulumi:"name"`
	Updated          string `pulumi:"updated"`
}

// GetProjectsProjectInput is an input type that accepts GetProjectsProjectArgs and GetProjectsProjectOutput values.
// You can construct a concrete instance of `GetProjectsProjectInput` via:
//
//	GetProjectsProjectArgs{...}
type GetProjectsProjectInput interface {
	pulumi.Input

	ToGetProjectsProjectOutput() GetProjectsProjectOutput
	ToGetProjectsProjectOutputWithContext(context.Context) GetProjectsProjectOutput
}

type GetProjectsProjectArgs struct {
	Created          pulumi.StringInput `pulumi:"created"`
	GrafanaFolderUid pulumi.StringInput `pulumi:"grafanaFolderUid"`
	Id               pulumi.StringInput `pulumi:"id"`
	IsDefault        pulumi.BoolInput   `pulumi:"isDefault"`
	Name             pulumi.StringInput `pulumi:"name"`
	Updated          pulumi.StringInput `pulumi:"updated"`
}

func (GetProjectsProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProject)(nil)).Elem()
}

func (i GetProjectsProjectArgs) ToGetProjectsProjectOutput() GetProjectsProjectOutput {
	return i.ToGetProjectsProjectOutputWithContext(context.Background())
}

func (i GetProjectsProjectArgs) ToGetProjectsProjectOutputWithContext(ctx context.Context) GetProjectsProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectOutput)
}

// GetProjectsProjectArrayInput is an input type that accepts GetProjectsProjectArray and GetProjectsProjectArrayOutput values.
// You can construct a concrete instance of `GetProjectsProjectArrayInput` via:
//
//	GetProjectsProjectArray{ GetProjectsProjectArgs{...} }
type GetProjectsProjectArrayInput interface {
	pulumi.Input

	ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput
	ToGetProjectsProjectArrayOutputWithContext(context.Context) GetProjectsProjectArrayOutput
}

type GetProjectsProjectArray []GetProjectsProjectInput

func (GetProjectsProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectsProject)(nil)).Elem()
}

func (i GetProjectsProjectArray) ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput {
	return i.ToGetProjectsProjectArrayOutputWithContext(context.Background())
}

func (i GetProjectsProjectArray) ToGetProjectsProjectArrayOutputWithContext(ctx context.Context) GetProjectsProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectArrayOutput)
}

type GetProjectsProjectOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProject)(nil)).Elem()
}

func (o GetProjectsProjectOutput) ToGetProjectsProjectOutput() GetProjectsProjectOutput {
	return o
}

func (o GetProjectsProjectOutput) ToGetProjectsProjectOutputWithContext(ctx context.Context) GetProjectsProjectOutput {
	return o
}

func (o GetProjectsProjectOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Created }).(pulumi.StringOutput)
}

func (o GetProjectsProjectOutput) GrafanaFolderUid() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.GrafanaFolderUid }).(pulumi.StringOutput)
}

func (o GetProjectsProjectOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProjectsProjectOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectsProject) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

func (o GetProjectsProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetProjectsProjectOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Updated }).(pulumi.StringOutput)
}

type GetProjectsProjectArrayOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectsProject)(nil)).Elem()
}

func (o GetProjectsProjectArrayOutput) ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput {
	return o
}

func (o GetProjectsProjectArrayOutput) ToGetProjectsProjectArrayOutputWithContext(ctx context.Context) GetProjectsProjectArrayOutput {
	return o
}

func (o GetProjectsProjectArrayOutput) Index(i pulumi.IntInput) GetProjectsProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProjectsProject {
		return vs[0].([]GetProjectsProject)[vs[1].(int)]
	}).(GetProjectsProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetLoadTestsLoadTestInput)(nil)).Elem(), GetLoadTestsLoadTestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetLoadTestsLoadTestArrayInput)(nil)).Elem(), GetLoadTestsLoadTestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectInput)(nil)).Elem(), GetProjectsProjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectArrayInput)(nil)).Elem(), GetProjectsProjectArray{})
	pulumi.RegisterOutputType(GetLoadTestsLoadTestOutput{})
	pulumi.RegisterOutputType(GetLoadTestsLoadTestArrayOutput{})
	pulumi.RegisterOutputType(GetProjectsProjectOutput{})
	pulumi.RegisterOutputType(GetProjectsProjectArrayOutput{})
}
