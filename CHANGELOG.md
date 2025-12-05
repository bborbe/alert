# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

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
