# WARNING: This file is autogenerated - changes will be overwritten when regenerated by https://github.com/pulumi/ci-mgmt

PACK := grafana
ORG := pulumiverse
PROJECT := github.com/$(ORG)/pulumi-$(PACK)
PROVIDER_PATH := provider
VERSION_PATH := $(PROVIDER_PATH)/pkg/version.Version
TFGEN := pulumi-tfgen-$(PACK)
PROVIDER := pulumi-resource-$(PACK)
JAVA_GEN := pulumi-java-gen
TESTPARALLELISM := 10
WORKING_DIR := $(shell pwd)
PULUMI_PROVIDER_BUILD_PARALLELISM ?=
PULUMI_CONVERT := 1
PULUMI_MISSING_DOCS_ERROR := true

# Override during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local & branch builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 0.0.0-alpha.0+dev
# Use this normalised version everywhere rather than the raw input to ensure consistency.
VERSION_GENERIC = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")

# Strips debug information from the provider binary to reduce its size and speed up builds
LDFLAGS_STRIP_SYMBOLS=-s -w
LDFLAGS_PROJ_VERSION=-X $(PROJECT)/$(VERSION_PATH)=$(VERSION_GENERIC)
LDFLAGS_UPSTREAM_VERSION=
LDFLAGS_EXTRAS=
LDFLAGS=$(LDFLAGS_PROJ_VERSION) $(LDFLAGS_UPSTREAM_VERSION) $(LDFLAGS_EXTRAS) $(LDFLAGS_STRIP_SYMBOLS)

# Create a `.make` directory for tracking targets which don't generate a single file output. This should be ignored by git.
# For targets which either don't generate a single file output, or the output file is committed, we use a "sentinel"
# file within `.make/` to track the staleness of the target and only rebuild when needed. 
# For each phony target, we create an internal target with the same name, but prefixed with `.make/` where the work is performed.
# At the end of each internal target we run `@touch $@` to update the file which is the name of the target.

# Ensure all directories exist before evaluating targets to avoid issues with `touch` creating directories.
_ := $(shell mkdir -p .make bin .pulumi/bin)

# Build the provider and all SDKs and install ready for testing
build: install_plugins provider build_sdks install_sdks
# Keep aliases for old targets to ensure backwards compatibility
development: build
only_build: build
# Prepare the workspace for building the provider and SDKs
# Importantly this is run by CI ahead of restoring the bin directory and resuming SDK builds
prepare_local_workspace: install_plugins upstream
# Creates all generated files which need to be committed
generate: generate_sdks schema
generate_sdks: generate_dotnet generate_go generate_nodejs generate_python
build_sdks: build_dotnet build_go build_nodejs build_python
install_sdks: install_dotnet_sdk install_go_sdk install_nodejs_sdk install_python_sdk
.PHONY: development only_build build generate generate_sdks build_sdks install_sdks

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Main Targets"
	@echo "  build (default)     Build the provider and all SDKs and install for testing"
	@echo "  generate            Generate all SDKs, documentation and schema"
	@echo "  provider            Build the local provider binary"
	@echo "  lint_provider<.fix> Run the linter on the provider (& optionally fix)"
	@echo "  test_provider       Run the provider tests"
	@echo "  test                Run the example tests (must run 'build' first)"
	@echo "  clean               Clean up generated files"
	@echo ""
	@echo "More Precise Targets"
	@echo "  schema        Generate the schema"
	@echo "  generate_sdks Generate all SDKs"
	@echo "  build_sdks    Build all SDKs"
	@echo "  install_sdks  Install all SDKs"
	@echo "  provider_dist Build and package the provider for all platforms"
	@echo ""
	@echo "Tool Targets"
	@echo "  ci-mgmt     Re-generate CI configuration from .ci-mgmt.yaml"
	@echo "  debug_tfgen Start a debug server for tfgen"
	@echo ""
	@echo "Internal Targets (automatically run as dependencies of other targets)"
	@echo "  prepare_local_workspace  Prepare for building"
	@echo "  install_plugins          Install plugin dependencies"
	@echo "  upstream                 Initialize the upstream submodule, if present"
	@echo ""
	@echo "Language-Specific Targets"
	@echo "  generate_[language]    Generate the SDK files ready for committing"
	@echo "  build_[language]       Build the SDK to check correctness"
	@echo "  install_[language]_sdk Install the SDK ready for testing"
	@echo ""
	@echo "  [language] = dotnet go nodejs python"
	@echo ""
.PHONY: help

GEN_PULUMI_HOME := $(WORKING_DIR)/.pulumi
GEN_PULUMI_CONVERT_EXAMPLES_CACHE_DIR := $(GEN_PULUMI_HOME)/examples-cache
GEN_ENVS := PULUMI_HOME=$(GEN_PULUMI_HOME) PULUMI_CONVERT_EXAMPLES_CACHE_DIR=$(GEN_PULUMI_CONVERT_EXAMPLES_CACHE_DIR) PULUMI_CONVERT=$(PULUMI_CONVERT) PULUMI_DISABLE_AUTOMATIC_PLUGIN_ACQUISITION=$(PULUMI_CONVERT)

generate_dotnet: .make/generate_dotnet
build_dotnet: .make/build_dotnet
.make/generate_dotnet: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/generate_dotnet: .make/install_plugins bin/$(TFGEN)
	$(GEN_ENVS) $(WORKING_DIR)/bin/$(TFGEN) dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		printf "module fake_dotnet_module // Exclude this directory from Go tools\n\ngo 1.17\n" > go.mod && \
		echo "$(VERSION_GENERIC)" >version.txt
	@touch $@
.make/build_dotnet: .make/generate_dotnet
	cd sdk/dotnet/ && dotnet build
	@touch $@
.PHONY: generate_dotnet build_dotnet

generate_go: .make/generate_go
build_go: .make/build_go
.make/generate_go: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/generate_go: .make/install_plugins bin/$(TFGEN)
	$(GEN_ENVS) $(WORKING_DIR)/bin/$(TFGEN) go --out sdk/go/
	@touch $@
.make/build_go: .make/generate_go
	cd sdk && go list "$$(grep -e "^module" go.mod | cut -d ' ' -f 2)/go/..." | xargs -I {} bash -c 'go build {} && go clean -i {}'
	@touch $@
.PHONY: generate_go build_go

generate_java: .make/generate_java
build_java: .make/build_java
.make/generate_java: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/generate_java: PACKAGE_VERSION := $(VERSION_GENERIC)
.make/generate_java: .make/install_plugins bin/pulumi-java-gen .make/schema
	PULUMI_HOME=$(GEN_PULUMI_HOME) PULUMI_CONVERT_EXAMPLES_CACHE_DIR=$(GEN_PULUMI_CONVERT_EXAMPLES_CACHE_DIR) bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java  --build gradle-nexus
	printf "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17\n" > sdk/java/go.mod
	@touch $@
.make/build_java: PACKAGE_VERSION := $(VERSION_GENERIC)
.make/build_java: .make/generate_java
	cd sdk/java/ && \
		gradle --console=plain build && \
		gradle --console=plain javadoc
	@touch $@
.PHONY: generate_java build_java

generate_nodejs: .make/generate_nodejs
build_nodejs: .make/build_nodejs
.make/generate_nodejs: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/generate_nodejs: .make/install_plugins bin/$(TFGEN)
	$(GEN_ENVS) $(WORKING_DIR)/bin/$(TFGEN) nodejs --out sdk/nodejs/
	printf "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.17\n" > sdk/nodejs/go.mod
	@touch $@
.make/build_nodejs: .make/generate_nodejs
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE* package.json yarn.lock ./bin/
	@touch $@
.PHONY: generate_nodejs build_nodejs

generate_python: .make/generate_python
build_python: .make/build_python
.make/generate_python: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/generate_python: .make/install_plugins bin/$(TFGEN)
	$(GEN_ENVS) $(WORKING_DIR)/bin/$(TFGEN) python --out sdk/python/
	printf "module fake_python_module // Exclude this directory from Go tools\n\ngo 1.17\n" > sdk/python/go.mod
	cp README.md sdk/python/
	@touch $@
.make/build_python: .make/generate_python
	cd sdk/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		rm ./bin/go.mod && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build==1.2.1 && \
		cd ./bin && \
		../venv/bin/python -m build .
	@touch $@
.PHONY: generate_python build_python

clean:
	rm -rf sdk/{dotnet,nodejs,go,python}
	rm -rf bin/*
	rm -rf .make/*
	if dotnet nuget list source | grep "$(WORKING_DIR)/nuget"; then \
		dotnet nuget remove source "$(WORKING_DIR)/nuget" \
	; fi
.PHONY: clean

install_dotnet_sdk: .make/install_dotnet_sdk
.make/install_dotnet_sdk: .make/build_dotnet
	mkdir -p nuget
	find sdk/dotnet/bin -name '*.nupkg' -print -exec cp -p "{}" ${WORKING_DIR}/nuget \;
	if ! dotnet nuget list source | grep "${WORKING_DIR}/nuget"; then \
		dotnet nuget add source "${WORKING_DIR}/nuget" --name "${WORKING_DIR}/nuget" \
	; fi
	@touch $@
install_go_sdk:
install_java_sdk:
install_nodejs_sdk: .make/install_nodejs_sdk
.make/install_nodejs_sdk: .make/build_nodejs
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin
	@touch $@
install_python_sdk:
.PHONY: install_dotnet_sdk install_go_sdk install_java_sdk install_nodejs_sdk install_python_sdk

# Install Pulumi plugins required for TFGen to resolve references
install_plugins: .make/install_plugins
.make/install_plugins: export PULUMI_HOME := $(WORKING_DIR)/.pulumi
.make/install_plugins: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/install_plugins: .pulumi/bin/pulumi
	.pulumi/bin/pulumi plugin install converter terraform 1.0.19
	.pulumi/bin/pulumi plugin install resource random 4.16.7
	.pulumi/bin/pulumi plugin install resource aws 6.56.1
	@touch $@
.PHONY: install_plugins

lint_provider: provider
	cd provider && golangci-lint run --path-prefix provider -c ../.golangci.yml
# `lint_provider.fix` is a utility target meant to be run manually
# that will run the linter and fix errors when possible.
lint_provider.fix:
	cd provider && golangci-lint run --path-prefix provider -c ../.golangci.yml --fix
.PHONY: lint_provider lint_provider.fix

# `make provider_no_deps` builds the provider binary directly, without ensuring that
# `cmd/pulumi-resource-grafana/schema.json` is valid and up to date.
# To create a release ready binary, you should use `make provider`.
build_provider_cmd = cd provider && go build $(PULUMI_PROVIDER_BUILD_PARALLELISM) -o $(WORKING_DIR)/bin/$(PROVIDER) -ldflags "$(LDFLAGS)" $(PROJECT)/$(PROVIDER_PATH)/cmd/$(PROVIDER)
provider: bin/$(PROVIDER)
provider_no_deps:
	$(call build_provider_cmd)
bin/$(PROVIDER): .make/schema
	$(call build_provider_cmd)
.PHONY: provider provider_no_deps

test: export PATH := $(WORKING_DIR)/bin:$(PATH)
test:
	cd examples && go test -v -tags=all -parallel $(TESTPARALLELISM) -timeout 2h
.PHONY: test

test_provider:
	cd provider && go test -v -short ./... -parallel $(TESTPARALLELISM)
.PHONY: test_provider

tfgen: schema
schema: .make/schema 
# This does actually have dependencies, but we're keeping it around for backwards compatibility for now
tfgen_no_deps: .make/schema
.make/schema: export PULUMI_HOME := $(WORKING_DIR)/.pulumi
.make/schema: export PATH := $(WORKING_DIR)/.pulumi/bin:$(PATH)
.make/schema: export PULUMI_CONVERT := $(PULUMI_CONVERT)
.make/schema: export PULUMI_CONVERT_EXAMPLES_CACHE_DIR := $(WORKING_DIR)/.pulumi/examples-cache
.make/schema: export PULUMI_DISABLE_AUTOMATIC_PLUGIN_ACQUISITION := $(PULUMI_CONVERT)
.make/schema: export PULUMI_MISSING_DOCS_ERROR := $(PULUMI_MISSING_DOCS_ERROR)
.make/schema: bin/$(TFGEN) .make/install_plugins .make/upstream
	$(WORKING_DIR)/bin/$(TFGEN) schema --out provider/cmd/$(PROVIDER)
	(cd provider && VERSION=$(VERSION_GENERIC) go generate cmd/$(PROVIDER)/main.go)
	@touch $@
tfgen_build_only: bin/$(TFGEN)
bin/$(TFGEN): provider/*.go provider/go.* .make/upstream
	(cd provider && go build $(PULUMI_PROVIDER_BUILD_PARALLELISM) -o $(WORKING_DIR)/bin/$(TFGEN) -ldflags "$(LDFLAGS_PROJ_VERSION) $(LDFLAGS_EXTRAS)" $(PROJECT)/$(PROVIDER_PATH)/cmd/$(TFGEN))
.PHONY: tfgen schema tfgen_no_deps tfgen_build_only

# Apply patches to the upstream submodule, if it exists
upstream: .make/upstream
# Re-run if the upstream commit or the patches change
.make/upstream: $(wildcard patches/*) $(wildcard .git/modules/upstream/HEAD)
ifneq ("$(wildcard upstream)","")
	./upstream.sh init
endif
	@touch $@
.PHONY: upstream

bin/pulumi-java-gen: .pulumi-java-gen.version
	pulumictl download-binary -n pulumi-language-java -v v$(shell cat .pulumi-java-gen.version) -r pulumi/pulumi-java

# To make an immediately observable change to .ci-mgmt.yaml:
#
# - Edit .ci-mgmt.yaml
# - Run make ci-mgmt to apply the change locally.
#
ci-mgmt: .ci-mgmt.yaml
	go run github.com/pulumi/ci-mgmt/provider-ci@master generate
.PHONY: ci-mgmt

# Because some codegen depends on the version of the CLI used, we install a local CLI
# version pinned to the same version as `provider/go.mod`.
#
# This logic compares the version of .pulumi/bin/pulumi already installed. If it matches
# the desired version, we just print. Otherwise we (re)install pulumi at the desired
# version.
.pulumi/bin/pulumi: .pulumi/version
	@if [ -x .pulumi/bin/pulumi ] && [ "v$$(cat .pulumi/version)" = "$$(.pulumi/bin/pulumi version)" ]; then \
		echo "pulumi/bin/pulumi version: v$$(cat .pulumi/version)"; \
		touch $@; \
	else \
		curl -fsSL https://get.pulumi.com | \
			HOME=$(WORKING_DIR) sh -s -- --version "$$(cat .pulumi/version)"; \
	fi

# Compute the version of Pulumi to use by inspecting the Go dependencies of the provider.
.pulumi/version: provider/go.mod
	cd provider && go list -f "{{slice .Version 1}}" -m github.com/pulumi/pulumi/pkg/v3 | tee ../$@

# Start debug server for tfgen
debug_tfgen:
	dlv  --listen=:2345 --headless=true --api-version=2  exec $(WORKING_DIR)/bin/$(TFGEN) -- schema --out provider/cmd/$(PROVIDER)
.PHONY: debug_tfgen

# Provider cross-platform build & packaging

# These targets assume that the schema-embed.json exists - it's generated by tfgen.
# We disable CGO to ensure that the binary is statically linked.
bin/linux-amd64/$(PROVIDER): TARGET := linux-amd64
bin/linux-arm64/$(PROVIDER): TARGET := linux-arm64
bin/darwin-amd64/$(PROVIDER): TARGET := darwin-amd64
bin/darwin-arm64/$(PROVIDER): TARGET := darwin-arm64
bin/windows-amd64/$(PROVIDER).exe: TARGET := windows-amd64
bin/%/$(PROVIDER) bin/%/$(PROVIDER).exe:
	@# check the TARGET is set
	test $(TARGET)
	cd provider && \
		export GOOS=$$(echo "$(TARGET)" | cut -d "-" -f 1) && \
		export GOARCH=$$(echo "$(TARGET)" | cut -d "-" -f 2) && \
		export CGO_ENABLED=0 && \
		go build -o "${WORKING_DIR}/$@" $(PULUMI_PROVIDER_BUILD_PARALLELISM) -ldflags "$(LDFLAGS)" "$(PROJECT)/$(PROVIDER_PATH)/cmd/$(PROVIDER)"

bin/$(PROVIDER)-v$(VERSION_GENERIC)-linux-amd64.tar.gz: bin/linux-amd64/$(PROVIDER)
bin/$(PROVIDER)-v$(VERSION_GENERIC)-linux-arm64.tar.gz: bin/linux-arm64/$(PROVIDER)
bin/$(PROVIDER)-v$(VERSION_GENERIC)-darwin-amd64.tar.gz: bin/darwin-amd64/$(PROVIDER)
bin/$(PROVIDER)-v$(VERSION_GENERIC)-darwin-arm64.tar.gz: bin/darwin-arm64/$(PROVIDER)
bin/$(PROVIDER)-v$(VERSION_GENERIC)-windows-amd64.tar.gz: bin/windows-amd64/$(PROVIDER).exe
bin/$(PROVIDER)-v$(VERSION_GENERIC)-%.tar.gz:
	@mkdir -p dist
	@# $< is the last dependency (the binary path from above) e.g. bin/linux-amd64/pulumi-resource-xyz
	@# $@ is the current target e.g. bin/pulumi-resource-xyz-v1.2.3-linux-amd64.tar.gz
	tar --gzip -cf $@ README.md LICENSE -C $$(dirname $<) .

provider_dist-linux-amd64: bin/$(PROVIDER)-v$(VERSION_GENERIC)-linux-amd64.tar.gz
provider_dist-linux-arm64: bin/$(PROVIDER)-v$(VERSION_GENERIC)-linux-arm64.tar.gz
provider_dist-darwin-amd64: bin/$(PROVIDER)-v$(VERSION_GENERIC)-darwin-amd64.tar.gz
provider_dist-darwin-arm64: bin/$(PROVIDER)-v$(VERSION_GENERIC)-darwin-arm64.tar.gz
provider_dist-windows-amd64: bin/$(PROVIDER)-v$(VERSION_GENERIC)-windows-amd64.tar.gz
provider_dist: provider_dist-linux-amd64 provider_dist-linux-arm64 provider_dist-darwin-amd64 provider_dist-darwin-arm64 provider_dist-windows-amd64
.PHONY: provider_dist-linux-amd64 provider_dist-linux-arm64 provider_dist-darwin-amd64 provider_dist-darwin-arm64 provider_dist-windows-amd64 provider_dist

# Permit providers to extend the Makefile with provider-specific Make includes.
include $(wildcard .mk/*.mk)
