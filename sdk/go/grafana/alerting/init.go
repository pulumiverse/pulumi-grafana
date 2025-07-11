// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alerting

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "grafana:alerting/contactPoint:ContactPoint":
		r = &ContactPoint{}
	case "grafana:alerting/messageTemplate:MessageTemplate":
		r = &MessageTemplate{}
	case "grafana:alerting/muteTiming:MuteTiming":
		r = &MuteTiming{}
	case "grafana:alerting/notificationPolicy:NotificationPolicy":
		r = &NotificationPolicy{}
	case "grafana:alerting/ruleGroup:RuleGroup":
		r = &RuleGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"grafana",
		"alerting/contactPoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"alerting/messageTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"alerting/muteTiming",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"alerting/notificationPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"grafana",
		"alerting/ruleGroup",
		&module{version},
	)
}
