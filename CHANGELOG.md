# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.8.10

- bump golangci-lint to v2.11.4
- bump osv-scanner to v2.3.5
- bump go-modtool to v0.7.1
- add opencontainers/runtime-spec replace directive
- update indirect deps

## v1.8.9

- downgrade fxamacker/cbor/v2 v2.9.1 → v2.9.0
- downgrade go.yaml.in/yaml/v2 v2.4.4 → v2.4.3
- downgrade k8s.io/klog/v2 v2.140.0 → v2.130.1
- downgrade k8s.io/kube-openapi and k8s.io/utils to earlier versions

## v1.8.8

- Enabled parallel golangci-lint runners

## v1.8.7

- chore: enable golangci-lint in check target and fix all lint violations (depguard, prealloc, revive, staticcheck)
- chore: update .golangci.yml to standard project config with nestif, errname, unparam, bodyclose, forcetypeassert, asasalint, prealloc linters

## v1.8.6

- standardize Makefile: multiline trivy format

## v1.8.5

- chore: verify project health — all tests pass, linting clean, precommit succeeds

## v1.8.3

- Update Go to 1.26.0

## v1.8.2

- Update Go from 1.25.5 to 1.25.7
- Update bborbe dependencies (errors, k8s, validation)
- Update Kubernetes client libraries to v0.33.7
- Update testing frameworks (ginkgo v2.28.1, gomega v1.39.1)
- Update various indirect dependencies

## v1.8.1

- update go to 1.25.5 and dependencies
- add go-modtool to tools.go
- cleanup go.mod exclude blocks

## v1.8.0

- update go and deps

## v1.7.0

- Add GitHub workflows for CI, Claude Code, and Claude Code review
- Add golangci-lint configuration
- Update Makefile with new targets (lint, gosec, trivy, osv-scanner)
- Update Go version from 1.24.1 to 1.25.2
- Add new security scanning and linting tools
- Code formatting improvements

## v1.6.0

- remove vendor
- go mod update

## v1.5.2

- add namespace to Alert.Identifier

## v1.5.1

- fix Alert.Identifier
- add test for Alert

## v1.5.0

- add AlertEventHandler
- add ResourceEventHandler

## v1.4.1

- enable applyconfig

## v1.4.0

- use new k8s code generator

## v1.3.4

- go mod update
 
## v1.3.3

- update to k8s v0.31.0
- go mod update

## v1.3.2

- use interface instead clientset struct

## v1.3.1

- refactor create clientset

## v1.3.0

- add clientset creator
- add alerts contains

## v1.2.0

- add AlertDeployer
- go mod update

## v1.1.0

- add sort
- go mod update

## v1.0.1

- downgrade k8s to v0.28.6

## v1.0.0

- Initial Version
