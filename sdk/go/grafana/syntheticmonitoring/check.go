// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package syntheticmonitoring

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// Synthetic Monitoring checks are tests that run on selected probes at defined
// intervals and report metrics and logs back to your Grafana Cloud account. The
// target for checks can be a domain name, a server, or a website, depending on
// what information you would like to gather about your endpoint. You can define
// multiple checks for a single endpoint to check different capabilities.
//
// * [Official documentation](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/create-checks/checks/)
//
// ## Example Usage
//
// ### DNS Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "dns", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("DNS Defaults"),
//				Target:  pulumi.String("grafana.com"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Ohio),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Dns: &syntheticmonitoring.CheckSettingsDnsArgs{},
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
// ### DNS Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "dns", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("DNS Updated"),
//				Target:  pulumi.String("grafana.net"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Dns: &syntheticmonitoring.CheckSettingsDnsArgs{
//						IpVersion:  pulumi.String("Any"),
//						Server:     pulumi.String("8.8.4.4"),
//						Port:       pulumi.Int(8600),
//						RecordType: pulumi.String("CNAME"),
//						Protocol:   pulumi.String("TCP"),
//						ValidRCodes: pulumi.StringArray{
//							pulumi.String("NOERROR"),
//							pulumi.String("NOTAUTH"),
//						},
//						ValidateAnswerRrs: &syntheticmonitoring.CheckSettingsDnsValidateAnswerRrsArgs{
//							FailIfMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-bad-stuff*"),
//							},
//							FailIfNotMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-good-stuff*"),
//							},
//						},
//						ValidateAuthorityRrs: &syntheticmonitoring.CheckSettingsDnsValidateAuthorityRrsArgs{
//							FailIfMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-bad-stuff*"),
//							},
//							FailIfNotMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-good-stuff*"),
//							},
//						},
//						ValidateAdditionalRrs: syntheticmonitoring.CheckSettingsDnsValidateAdditionalRrArray{
//							&syntheticmonitoring.CheckSettingsDnsValidateAdditionalRrArgs{
//								FailIfMatchesRegexps: pulumi.StringArray{
//									pulumi.String(".+-bad-stuff*"),
//								},
//								FailIfNotMatchesRegexps: pulumi.StringArray{
//									pulumi.String(".+-good-stuff*"),
//								},
//							},
//						},
//					},
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
// ### HTTP Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "http", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("HTTP Defaults"),
//				Target:  pulumi.String("https://grafana.com"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Ohio),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Http: &syntheticmonitoring.CheckSettingsHttpArgs{},
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
// ### HTTP Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "http", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("HTTP Defaults"),
//				Target:  pulumi.String("https://grafana.org"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Mumbai),
//					pulumi.Int(main.Probes.Mumbai),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Http: &syntheticmonitoring.CheckSettingsHttpArgs{
//						IpVersion:                  pulumi.String("V6"),
//						Method:                     pulumi.String("TRACE"),
//						Body:                       pulumi.String("and spirit"),
//						NoFollowRedirects:          pulumi.Bool(true),
//						BearerToken:                pulumi.String("asdfjkl;"),
//						ProxyUrl:                   pulumi.String("https://almost-there"),
//						FailIfSsl:                  pulumi.Bool(true),
//						FailIfNotSsl:               pulumi.Bool(true),
//						Compression:                pulumi.String("deflate"),
//						CacheBustingQueryParamName: pulumi.String("pineapple"),
//						TlsConfig: &syntheticmonitoring.CheckSettingsHttpTlsConfigArgs{
//							ServerName: pulumi.String("grafana.org"),
//							ClientCert: pulumi.String(`-----BEGIN CERTIFICATE-----
//
// MIIEljCCAn4CCQCKJPUQQxeO0zANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJT
// RTAeFw0yMTA1MjkxOTIyNTdaFw0yNDAzMTgxOTIyNTdaMA0xCzAJBgNVBAYTAlNF
// MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnmbazDNUT0rSI4BpGZK+
// 0AJ+9FDkIYWJUtRLJoxw8CF+AobMFploYA2L2Myt80cTA1w8FrewjC8qlqdnrPWr
// h1ely2zsUljgi1/niH0ndjFzliL7UkinXQiAsTtYOrOQmzyd/o5PNdu7dz0m7stD
// BN/Sz5TlXZnA1/eJbqV/kqMau6b1MaBx8SbRfUG9+cSmUobFJwuktDrPuwJhcEkl
// iDmhEqu1GuZzmKvzPacLTVia1vSlmCTCu89NiHI8iGiiLtqNrapup7f8j5m3a3SL
// a+vXhplFj2piNl7Nc0dfuVgtEliTI+qUL2/+4A7gzRWZpHy21/LxMMXmBhdJW9En
// FWkev97VZLgb5TR3+qpSWmXcodjPy4dibvwsOMpdd+Q4AYulwvlDw5idRPVgGvk7
// qq03+w9ppZ5Fugws9k2CD9F/75JX2mCbRpkuPe8XXZ7bqrMaQgQMLOrs68HuiiCk
// FTklglq4DMKxnf/Y/T/MgIa9Q1o28YSevh6A7FnfPGARj2H2T4rToi+bC1Vf7qNB
// Z18bDpz99tRUTbyiRUSBMWLCGhU6c4HAqUrfrkpperOKFBQ3i38a79838oFdXHBW
// 6rx1t5cC3XwtEoUyeBKAygez8G1LDXbN3607MxVhAjhHKtPkYvuBfysSNU6JrR0z
// UV1IURJANt2UMuKgSEkG/IMCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAcipMhp/w
// yzfPy61faVAw9SPaMNRlnW9FCDC3N9CGOjo2knjXpObPzyzsJiUURTjrA9eFMpRA
// e2Rgn2j+nvm2XdLAlC4Kh8jqv/wCL0X6BTQMdN5aOhXdSiXtpXOMvXYY/dQ4ebRZ
// XeRCVWQD79JbV6/uyx0nCV3FVcU7L1P4UjxroefVr0soLPMirgxHmOxLnkoVgdcB
// tqufP5kJx9CIeJXPx3QQsk1XfEtxtUvuw4ZaZkQnNUqvGl7V+AZpur5Eqfv3zBi8
// QxxL7qGkARNssNWH2Ju+tqpM/UZRnjlFrDR4SXUgT0coTduBalUY6qHkciHmRpiP
// tf3SgpDeiCSOV2iVFGdaR1mz3muWoAYWFstcWN3a3HjjVugIi23yLN8Gv8CNeoH4
// prulinFCLrFgAh8SLAF8mOAZanT06LH8jOIFYrdUxH+ZeRBR0rLoFjUF+JB7UKD9
// 5TA+B4EBzQ1tMbGFU1DX79MjAejq0IV0Nzq+GMfBvLHxEf4+Oz8nqhDXQcJ6TdtY
// l3Lyw5zBvOL80SBK+Mr0UP7d9U3VXgbGHCYVJU6Ot1TwiGwahtWALRALA3TWeGkq
// 7kyD1H+nm+9lfKhuyBRQnRGBVyze2lAp7oxwshJuhBwEXosXFxq1Cy6QhPN77r6N
// vuhxvtppolNnyOgGxwG4zquqq2V5/+vKjKY=
// -----END CERTIFICATE-----
// `),
//
//						},
//						Headers: pulumi.StringArray{
//							pulumi.String("Content-Type: multipart/form-data; boundary=something"),
//						},
//						BasicAuth: &syntheticmonitoring.CheckSettingsHttpBasicAuthArgs{
//							Username: pulumi.String("open"),
//							Password: pulumi.String("sesame"),
//						},
//						ValidStatusCodes: pulumi.IntArray{
//							pulumi.Int(200),
//							pulumi.Int(201),
//						},
//						ValidHttpVersions: pulumi.StringArray{
//							pulumi.String("HTTP/1.0"),
//							pulumi.String("HTTP/1.1"),
//							pulumi.String("HTTP/2.0"),
//						},
//						FailIfBodyMatchesRegexps: pulumi.StringArray{
//							pulumi.String(".*bad stuff.*"),
//						},
//						FailIfBodyNotMatchesRegexps: pulumi.StringArray{
//							pulumi.String(".*good stuff.*"),
//						},
//						FailIfHeaderMatchesRegexps: syntheticmonitoring.CheckSettingsHttpFailIfHeaderMatchesRegexpArray{
//							&syntheticmonitoring.CheckSettingsHttpFailIfHeaderMatchesRegexpArgs{
//								Header:       pulumi.String("Content-Type"),
//								Regexp:       pulumi.String("application/soap*"),
//								AllowMissing: pulumi.Bool(true),
//							},
//						},
//					},
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
// ### Ping Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "ping", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("Ping Defaults"),
//				Target:  pulumi.String("grafana.com"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Ohio),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Ping: &syntheticmonitoring.CheckSettingsPingArgs{},
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
// ### Ping Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "ping", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("Ping Updated"),
//				Target:  pulumi.String("grafana.net"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Ping: &syntheticmonitoring.CheckSettingsPingArgs{
//						IpVersion:    pulumi.String("Any"),
//						PayloadSize:  pulumi.Int(20),
//						DontFragment: pulumi.Bool(true),
//					},
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
// ### TCP Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "tcp", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("TCP Defaults"),
//				Target:  pulumi.String("grafana.com:80"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Ohio),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Tcp: &syntheticmonitoring.CheckSettingsTcpArgs{},
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
// ### TCP Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "tcp", &syntheticmonitoring.CheckArgs{
//				Job:     pulumi.String("TCP Defaults"),
//				Target:  pulumi.String("grafana.com:443"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Tcp: &syntheticmonitoring.CheckSettingsTcpArgs{
//						IpVersion: pulumi.String("V6"),
//						Tls:       pulumi.Bool(true),
//						QueryResponses: syntheticmonitoring.CheckSettingsTcpQueryResponseArray{
//							&syntheticmonitoring.CheckSettingsTcpQueryResponseArgs{
//								Send:   pulumi.String("howdy"),
//								Expect: pulumi.String("hi"),
//							},
//							&syntheticmonitoring.CheckSettingsTcpQueryResponseArgs{
//								Send:     pulumi.String("like this"),
//								Expect:   pulumi.String("like that"),
//								StartTls: pulumi.Bool(true),
//							},
//						},
//						TlsConfig: &syntheticmonitoring.CheckSettingsTcpTlsConfigArgs{
//							ServerName: pulumi.String("grafana.com"),
//							CaCert: pulumi.String(`-----BEGIN CERTIFICATE-----
//
// MIIEljCCAn4CCQCKJPUQQxeO0zANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJT
// RTAeFw0yMTA1MjkxOTIyNTdaFw0yNDAzMTgxOTIyNTdaMA0xCzAJBgNVBAYTAlNF
// MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnmbazDNUT0rSI4BpGZK+
// 0AJ+9FDkIYWJUtRLJoxw8CF+AobMFploYA2L2Myt80cTA1w8FrewjC8qlqdnrPWr
// h1ely2zsUljgi1/niH0ndjFzliL7UkinXQiAsTtYOrOQmzyd/o5PNdu7dz0m7stD
// BN/Sz5TlXZnA1/eJbqV/kqMau6b1MaBx8SbRfUG9+cSmUobFJwuktDrPuwJhcEkl
// iDmhEqu1GuZzmKvzPacLTVia1vSlmCTCu89NiHI8iGiiLtqNrapup7f8j5m3a3SL
// a+vXhplFj2piNl7Nc0dfuVgtEliTI+qUL2/+4A7gzRWZpHy21/LxMMXmBhdJW9En
// FWkev97VZLgb5TR3+qpSWmXcodjPy4dibvwsOMpdd+Q4AYulwvlDw5idRPVgGvk7
// qq03+w9ppZ5Fugws9k2CD9F/75JX2mCbRpkuPe8XXZ7bqrMaQgQMLOrs68HuiiCk
// FTklglq4DMKxnf/Y/T/MgIa9Q1o28YSevh6A7FnfPGARj2H2T4rToi+bC1Vf7qNB
// Z18bDpz99tRUTbyiRUSBMWLCGhU6c4HAqUrfrkpperOKFBQ3i38a79838oFdXHBW
// 6rx1t5cC3XwtEoUyeBKAygez8G1LDXbN3607MxVhAjhHKtPkYvuBfysSNU6JrR0z
// UV1IURJANt2UMuKgSEkG/IMCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAcipMhp/w
// yzfPy61faVAw9SPaMNRlnW9FCDC3N9CGOjo2knjXpObPzyzsJiUURTjrA9eFMpRA
// e2Rgn2j+nvm2XdLAlC4Kh8jqv/wCL0X6BTQMdN5aOhXdSiXtpXOMvXYY/dQ4ebRZ
// XeRCVWQD79JbV6/uyx0nCV3FVcU7L1P4UjxroefVr0soLPMirgxHmOxLnkoVgdcB
// tqufP5kJx9CIeJXPx3QQsk1XfEtxtUvuw4ZaZkQnNUqvGl7V+AZpur5Eqfv3zBi8
// QxxL7qGkARNssNWH2Ju+tqpM/UZRnjlFrDR4SXUgT0coTduBalUY6qHkciHmRpiP
// tf3SgpDeiCSOV2iVFGdaR1mz3muWoAYWFstcWN3a3HjjVugIi23yLN8Gv8CNeoH4
// prulinFCLrFgAh8SLAF8mOAZanT06LH8jOIFYrdUxH+ZeRBR0rLoFjUF+JB7UKD9
// 5TA+B4EBzQ1tMbGFU1DX79MjAejq0IV0Nzq+GMfBvLHxEf4+Oz8nqhDXQcJ6TdtY
// l3Lyw5zBvOL80SBK+Mr0UP7d9U3VXgbGHCYVJU6Ot1TwiGwahtWALRALA3TWeGkq
// 7kyD1H+nm+9lfKhuyBRQnRGBVyze2lAp7oxwshJuhBwEXosXFxq1Cy6QhPN77r6N
// vuhxvtppolNnyOgGxwG4zquqq2V5/+vKjKY=
// -----END CERTIFICATE-----
// `),
//
//						},
//					},
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
// ### Traceroute Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "traceroute", &syntheticmonitoring.CheckArgs{
//				Job:       pulumi.String("Traceroute defaults"),
//				Target:    pulumi.String("grafana.com"),
//				Enabled:   pulumi.Bool(false),
//				Frequency: pulumi.Int(120000),
//				Timeout:   pulumi.Int(30000),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Ohio),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Traceroute: &syntheticmonitoring.CheckSettingsTracerouteArgs{},
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
// ### Traceroute Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/syntheticmonitoring"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := syntheticmonitoring.GetProbes(ctx, &syntheticmonitoring.GetProbesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = syntheticmonitoring.NewCheck(ctx, "traceroute", &syntheticmonitoring.CheckArgs{
//				Job:       pulumi.String("Traceroute complex"),
//				Target:    pulumi.String("grafana.net"),
//				Enabled:   pulumi.Bool(false),
//				Frequency: pulumi.Int(120000),
//				Timeout:   pulumi.Int(30000),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &syntheticmonitoring.CheckSettingsArgs{
//					Traceroute: &syntheticmonitoring.CheckSettingsTracerouteArgs{
//						MaxHops:        pulumi.Int(25),
//						MaxUnknownHops: pulumi.Int(10),
//						PtrLookup:      pulumi.Bool(false),
//					},
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
// $ pulumi import grafana:syntheticMonitoring/check:Check name "{{ id }}"
// ```
type Check struct {
	pulumi.CustomResourceState

	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/configure-alerts/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity pulumi.StringPtrOutput `pulumi:"alertSensitivity"`
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly pulumi.BoolPtrOutput `pulumi:"basicMetricsOnly"`
	// Whether to enable the check. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 1 hour (3600000 ms). Defaults to `60000`.
	Frequency pulumi.IntPtrOutput `pulumi:"frequency"`
	// Name used for job label.
	Job pulumi.StringOutput `pulumi:"job"`
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// List of probe location IDs where this target will be checked from.
	Probes pulumi.IntArrayOutput `pulumi:"probes"`
	// Check settings. Should contain exactly one nested block.
	Settings CheckSettingsOutput `pulumi:"settings"`
	// Hostname to ping.
	Target pulumi.StringOutput `pulumi:"target"`
	// The tenant ID of the check.
	TenantId pulumi.IntOutput `pulumi:"tenantId"`
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewCheck registers a new resource with the given unique name, arguments, and options.
func NewCheck(ctx *pulumi.Context,
	name string, args *CheckArgs, opts ...pulumi.ResourceOption) (*Check, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Job == nil {
		return nil, errors.New("invalid value for required argument 'Job'")
	}
	if args.Probes == nil {
		return nil, errors.New("invalid value for required argument 'Probes'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Check
	err := ctx.RegisterResource("grafana:syntheticMonitoring/check:Check", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheck gets an existing Check resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckState, opts ...pulumi.ResourceOption) (*Check, error) {
	var resource Check
	err := ctx.ReadResource("grafana:syntheticMonitoring/check:Check", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Check resources.
type checkState struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/configure-alerts/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity *string `pulumi:"alertSensitivity"`
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly *bool `pulumi:"basicMetricsOnly"`
	// Whether to enable the check. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 1 hour (3600000 ms). Defaults to `60000`.
	Frequency *int `pulumi:"frequency"`
	// Name used for job label.
	Job *string `pulumi:"job"`
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels map[string]string `pulumi:"labels"`
	// List of probe location IDs where this target will be checked from.
	Probes []int `pulumi:"probes"`
	// Check settings. Should contain exactly one nested block.
	Settings *CheckSettings `pulumi:"settings"`
	// Hostname to ping.
	Target *string `pulumi:"target"`
	// The tenant ID of the check.
	TenantId *int `pulumi:"tenantId"`
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout *int `pulumi:"timeout"`
}

type CheckState struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/configure-alerts/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity pulumi.StringPtrInput
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly pulumi.BoolPtrInput
	// Whether to enable the check. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 1 hour (3600000 ms). Defaults to `60000`.
	Frequency pulumi.IntPtrInput
	// Name used for job label.
	Job pulumi.StringPtrInput
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels pulumi.StringMapInput
	// List of probe location IDs where this target will be checked from.
	Probes pulumi.IntArrayInput
	// Check settings. Should contain exactly one nested block.
	Settings CheckSettingsPtrInput
	// Hostname to ping.
	Target pulumi.StringPtrInput
	// The tenant ID of the check.
	TenantId pulumi.IntPtrInput
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout pulumi.IntPtrInput
}

func (CheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkState)(nil)).Elem()
}

type checkArgs struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/configure-alerts/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity *string `pulumi:"alertSensitivity"`
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly *bool `pulumi:"basicMetricsOnly"`
	// Whether to enable the check. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 1 hour (3600000 ms). Defaults to `60000`.
	Frequency *int `pulumi:"frequency"`
	// Name used for job label.
	Job string `pulumi:"job"`
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels map[string]string `pulumi:"labels"`
	// List of probe location IDs where this target will be checked from.
	Probes []int `pulumi:"probes"`
	// Check settings. Should contain exactly one nested block.
	Settings CheckSettings `pulumi:"settings"`
	// Hostname to ping.
	Target string `pulumi:"target"`
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a Check resource.
type CheckArgs struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/configure-alerts/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity pulumi.StringPtrInput
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly pulumi.BoolPtrInput
	// Whether to enable the check. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 1 hour (3600000 ms). Defaults to `60000`.
	Frequency pulumi.IntPtrInput
	// Name used for job label.
	Job pulumi.StringInput
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels pulumi.StringMapInput
	// List of probe location IDs where this target will be checked from.
	Probes pulumi.IntArrayInput
	// Check settings. Should contain exactly one nested block.
	Settings CheckSettingsInput
	// Hostname to ping.
	Target pulumi.StringInput
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout pulumi.IntPtrInput
}

func (CheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkArgs)(nil)).Elem()
}

type CheckInput interface {
	pulumi.Input

	ToCheckOutput() CheckOutput
	ToCheckOutputWithContext(ctx context.Context) CheckOutput
}

func (*Check) ElementType() reflect.Type {
	return reflect.TypeOf((**Check)(nil)).Elem()
}

func (i *Check) ToCheckOutput() CheckOutput {
	return i.ToCheckOutputWithContext(context.Background())
}

func (i *Check) ToCheckOutputWithContext(ctx context.Context) CheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckOutput)
}

// CheckArrayInput is an input type that accepts CheckArray and CheckArrayOutput values.
// You can construct a concrete instance of `CheckArrayInput` via:
//
//	CheckArray{ CheckArgs{...} }
type CheckArrayInput interface {
	pulumi.Input

	ToCheckArrayOutput() CheckArrayOutput
	ToCheckArrayOutputWithContext(context.Context) CheckArrayOutput
}

type CheckArray []CheckInput

func (CheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Check)(nil)).Elem()
}

func (i CheckArray) ToCheckArrayOutput() CheckArrayOutput {
	return i.ToCheckArrayOutputWithContext(context.Background())
}

func (i CheckArray) ToCheckArrayOutputWithContext(ctx context.Context) CheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckArrayOutput)
}

// CheckMapInput is an input type that accepts CheckMap and CheckMapOutput values.
// You can construct a concrete instance of `CheckMapInput` via:
//
//	CheckMap{ "key": CheckArgs{...} }
type CheckMapInput interface {
	pulumi.Input

	ToCheckMapOutput() CheckMapOutput
	ToCheckMapOutputWithContext(context.Context) CheckMapOutput
}

type CheckMap map[string]CheckInput

func (CheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Check)(nil)).Elem()
}

func (i CheckMap) ToCheckMapOutput() CheckMapOutput {
	return i.ToCheckMapOutputWithContext(context.Background())
}

func (i CheckMap) ToCheckMapOutputWithContext(ctx context.Context) CheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckMapOutput)
}

type CheckOutput struct{ *pulumi.OutputState }

func (CheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Check)(nil)).Elem()
}

func (o CheckOutput) ToCheckOutput() CheckOutput {
	return o
}

func (o CheckOutput) ToCheckOutputWithContext(ctx context.Context) CheckOutput {
	return o
}

// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/testing/synthetic-monitoring/configure-alerts/synthetic-monitoring-alerting/). Defaults to `none`.
func (o CheckOutput) AlertSensitivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Check) pulumi.StringPtrOutput { return v.AlertSensitivity }).(pulumi.StringPtrOutput)
}

// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
func (o CheckOutput) BasicMetricsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Check) pulumi.BoolPtrOutput { return v.BasicMetricsOnly }).(pulumi.BoolPtrOutput)
}

// Whether to enable the check. Defaults to `true`.
func (o CheckOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Check) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 1 hour (3600000 ms). Defaults to `60000`.
func (o CheckOutput) Frequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Check) pulumi.IntPtrOutput { return v.Frequency }).(pulumi.IntPtrOutput)
}

// Name used for job label.
func (o CheckOutput) Job() pulumi.StringOutput {
	return o.ApplyT(func(v *Check) pulumi.StringOutput { return v.Job }).(pulumi.StringOutput)
}

// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
func (o CheckOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Check) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// List of probe location IDs where this target will be checked from.
func (o CheckOutput) Probes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Check) pulumi.IntArrayOutput { return v.Probes }).(pulumi.IntArrayOutput)
}

// Check settings. Should contain exactly one nested block.
func (o CheckOutput) Settings() CheckSettingsOutput {
	return o.ApplyT(func(v *Check) CheckSettingsOutput { return v.Settings }).(CheckSettingsOutput)
}

// Hostname to ping.
func (o CheckOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *Check) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

// The tenant ID of the check.
func (o CheckOutput) TenantId() pulumi.IntOutput {
	return o.ApplyT(func(v *Check) pulumi.IntOutput { return v.TenantId }).(pulumi.IntOutput)
}

// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
func (o CheckOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Check) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type CheckArrayOutput struct{ *pulumi.OutputState }

func (CheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Check)(nil)).Elem()
}

func (o CheckArrayOutput) ToCheckArrayOutput() CheckArrayOutput {
	return o
}

func (o CheckArrayOutput) ToCheckArrayOutputWithContext(ctx context.Context) CheckArrayOutput {
	return o
}

func (o CheckArrayOutput) Index(i pulumi.IntInput) CheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Check {
		return vs[0].([]*Check)[vs[1].(int)]
	}).(CheckOutput)
}

type CheckMapOutput struct{ *pulumi.OutputState }

func (CheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Check)(nil)).Elem()
}

func (o CheckMapOutput) ToCheckMapOutput() CheckMapOutput {
	return o
}

func (o CheckMapOutput) ToCheckMapOutputWithContext(ctx context.Context) CheckMapOutput {
	return o
}

func (o CheckMapOutput) MapIndex(k pulumi.StringInput) CheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Check {
		return vs[0].(map[string]*Check)[vs[1].(string)]
	}).(CheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckInput)(nil)).Elem(), &Check{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckArrayInput)(nil)).Elem(), CheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckMapInput)(nil)).Elem(), CheckMap{})
	pulumi.RegisterOutputType(CheckOutput{})
	pulumi.RegisterOutputType(CheckArrayOutput{})
	pulumi.RegisterOutputType(CheckMapOutput{})
}
