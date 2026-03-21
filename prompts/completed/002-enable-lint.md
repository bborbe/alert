---
status: completed
summary: 'Enabled golangci-lint in check target, updated .golangci.yml to standard config, and fixed all 4 lint violations (depguard: replaced pkg/errors with bborbe/errors, prealloc: preallocate AlertSpecs slice, revive: renamed builtin-shadowing parameter ''new'' to ''updated'', staticcheck: removed unnecessary ObjectMeta embedding in selector).'
container: alert-002-enable-lint
dark-factory-version: v0.59.5-dirty
created: "2026-03-21T09:12:33Z"
queued: "2026-03-21T10:12:33Z"
started: "2026-03-21T10:12:41Z"
completed: "2026-03-21T10:19:14Z"
---

<summary>
- The `check` Make target now includes `lint` as a dependency
- The TODO comment about enabling lint is removed from the Makefile
- All golangci-lint violations are fixed in the source code
- The golangci-lint config is updated to match the standard v2 config used across bborbe projects
- `make precommit` passes cleanly
</summary>

<objective>
Enable the golangci-lint linter in the `check` Make target, fix all lint violations, and update the golangci-lint config to the standard project template. The lint target already exists but is excluded from the `check` dependency chain.
</objective>

<context>
Read CLAUDE.md for project conventions and build commands.

Read these files before making changes:
- `Makefile` — the `check` target currently excludes `lint`, with a TODO comment above
- `.golangci.yml` — current linter config (missing several linters vs the standard)

Understand the codebase:
- `*.go` files in the project root — the source files that will need lint fixes
- `k8s/` — generated K8s client code (likely needs exclusions)
- `mocks/` — generated mock code (likely needs exclusions)

Reference standard `.golangci.yml` config to match (from the kv repo):
```yaml
version: "2"

run:
  timeout: 5m
  tests: true

linters:
  enable:
    - govet
    - errcheck
    - staticcheck
    - unused
    - revive
    - gosec
    - gocyclo
    - depguard
    - dupl
    - nestif
    - errname
    - unparam
    - bodyclose
    - forcetypeassert
    - asasalint
    - prealloc
  settings:
    depguard:
      rules:
        Main:
          deny:
            - pkg: "github.com/pkg/errors"
              desc: "use github.com/bborbe/errors instead"
            - pkg: "github.com/bborbe/argument"
              desc: "use github.com/bborbe/argument/v2 instead"
            - pkg: "golang.org/x/net/context"
              desc: "use context from standard library instead"
            - pkg: "golang.org/x/lint/golint"
              desc: "deprecated, use revive or staticcheck instead"
            - pkg: "io/ioutil"
              desc: "deprecated since Go 1.16, use io and os packages instead"
    funlen:
      lines: 80
      statements: 50
    gocognit:
      min-complexity: 20
    nestif:
      min-complexity: 4
    maintidx:
      min-maintainability-index: 20
  exclusions:
    presets:
      - comments
      - std-error-handling
      - common-false-positives
    rules:
      - linters:
          - staticcheck
        text: "SA1019"
      - linters:
          - revive
        path: "_test\\.go$"
        text: "dot-imports"
      - linters:
          - revive
        text: "unused-parameter"
      - linters:
          - revive
        text: "exported"
      - linters:
          - dupl
        path: "_test\\.go$"
      - linters:
          - unparam
        path: "_test\\.go$"
      - linters:
          - dupl
        path: "-test-suite\\.go$"
      - linters:
          - revive
        path: "-test-suite\\.go$"
        text: "dot-imports"

formatters:
  enable:
    - gofmt
    - goimports
```
</context>

<requirements>
1. Update `.golangci.yml` to match the reference standard config shown in the context section:
   - Add these missing linters to the `enable` list: `nestif`, `errname`, `unparam`, `bodyclose`, `forcetypeassert`, `asasalint`, `prealloc`
   - Add the `settings` block for `depguard`, `funlen`, `gocognit`, `nestif`, `maintidx` matching the kv reference config
   - Update the `depguard` deny rules to match the kv reference (fix the typo "erros" -> "errors" in pkg paths, add all deny entries from kv)
   - Add exclusion rules for `errname`, `unparam`, `dupl`, `revive` on test files and test-suite files matching the kv reference
   - Keep any alert-specific exclusions that make sense (e.g., SA1019 staticcheck suppression)

2. Update the `Makefile`:
   - Remove the two comment lines (the `# TODO: enable lint` line and the `# check: lint vet errcheck vulncheck osv-scanner gosec trivy` line)
   - Change `check: vet errcheck vulncheck osv-scanner gosec trivy` to `check: lint vet errcheck vulncheck osv-scanner gosec trivy`

   The result should look like:
   ```makefile
   .PHONY: check
   check: lint vet errcheck vulncheck osv-scanner gosec trivy
   ```

3. Run `make lint` to identify all golangci-lint violations

4. Fix all lint violations in the Go source files:
   - Fix each violation at the source — do not add nolint directives unless the violation is a false positive in generated code
   - For generated code in `k8s/` and `mocks/`, add path exclusions to `.golangci.yml` rather than modifying generated files
   - Common fixes: unused parameters, missing error checks, unchecked type assertions, unclosed response bodies

5. Run `make lint` again to confirm zero violations

6. Run `make precommit` to confirm everything passes end-to-end
</requirements>

<constraints>
- Do NOT commit — dark-factory handles git
- Do NOT modify generated files in `k8s/` or `mocks/` — add `.golangci.yml` exclusions instead
- Do NOT add `//nolint` directives unless absolutely necessary for false positives — fix the actual code
- Do NOT change any behavior or logic — only fix lint violations
- Do NOT remove or weaken existing linter rules — only add missing ones
- Existing tests must still pass
</constraints>

<verification>
Run `make lint` — must pass with exit code 0.
Run `make precommit` — must pass with exit code 0.
</verification>
