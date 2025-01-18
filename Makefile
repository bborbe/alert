
default: precommit

precommit: ensure format generate test check addlicense
	@echo "ready to commit"

ensure:
	go mod tidy
	go mod verify
	go mod vendor

format:
	go run -mod=vendor github.com/incu6us/goimports-reviser/v3 -project-name github.com/bborbe/alert -format -excludes vendor ./...

generate:
	rm -rf mocks avro
	go generate -mod=vendor ./...

test:
	go test -mod=vendor -p=$${GO_TEST_PARALLEL:-1} -cover -race $(shell go list -mod=vendor ./... | grep -v /vendor/)

check: vet errcheck vulncheck

vet:
	go vet -mod=vendor $(shell go list -mod=vendor ./... | grep -v /vendor/)

errcheck:
	go run -mod=vendor github.com/kisielk/errcheck -ignore '(Close|Write|Fprint)' $(shell go list -mod=vendor ./... | grep -v /vendor/ | grep -v k8s/client)

addlicense:
	go run -mod=vendor github.com/google/addlicense -c "Benjamin Borbe" -y $$(date +'%Y') -l bsd $$(find . -name "*.go" -not -path './vendor/*')

vulncheck:
	go run -mod=vendor golang.org/x/vuln/cmd/govulncheck $(shell go list -mod=vendor ./... | grep -v /vendor/)

generatek8s:
	rm -rf k8s/client ${GOPATH}/src/github.com/bborbe/alert
	chmod a+x vendor/k8s.io/code-generator/*.sh
	bash vendor/k8s.io/code-generator/generate-groups.sh applyconfiguration,client,deepcopy,informer,lister \
	github.com/bborbe/alert/k8s/client github.com/bborbe/alert/k8s/apis \
	monitoring.benjamin-borbe.de:v1
	cp -R ${GOPATH}/src/github.com/bborbe/alert/k8s .
