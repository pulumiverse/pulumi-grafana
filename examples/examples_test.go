// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getUrl(t *testing.T) string {
	name := os.Getenv("GRAFANA_URL")
	if name == "" {
		t.Skipf("Skipping test due to missing GRAFANA_URL environment variable")
	}

	return name
}

func getAuth(t *testing.T) string {
	name := os.Getenv("GRAFANA_AUTH")
	if name == "" {
		t.Skipf("Skipping test due to missing GRAFANA_AUTH environment variable")
	}

	return name
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	url := getUrl(t)
	auth := getAuth(t)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"url": url,
			"auth": auth,
		},
	}
}