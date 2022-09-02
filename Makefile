apps = 'products' 'details' 'ratings' 'reviews'
.PHONY: run
run: proto wire
	for app in $(apps) ;\
	do \
		 go run ./cmd/$$app -f configs/$$app.yml  & \
	done
.PHONY: wire
wire:
	wire ./...
.PHONY: gqlgen
gqlgen:
	gqlgen --config configs/gqlgen.yml
.PHONY: test
test: mock
	for app in $(apps) ;\
	do \
		go test -v ./internal/app/$$app/... -f `pwd`/configs/$$app.yml -covermode=count -coverprofile=dist/cover-$$app.out ;\
	done
.PHONY: build
build:
	for app in $(apps) ;\
	do \
		GOOS=linux GOARCH="amd64" go build -o dist/$$app-linux-amd64 ./cmd/$$app/; \
		GOOS=darwin GOARCH="amd64" go build -o dist/$$app-darwin-amd64 ./cmd/$$app/; \
	done
.PHONY: cover
cover: test
	for app in $(apps) ;\
	do \
		go tool cover -html=dist/cover-$$app.out; \
	done
.PHONY: mock
mock:
	mockery --all
.PHONY: lint
lint:
	golint ./...
.PHONY: proto
proto:
	rm -f api/proto/*.go
	protoc -I api/proto -I ${GOPATH}/src/github.com/protocolbuffers/protobuf/src ./api/proto/* --go_out=api/proto --go-grpc_out=require_unimplemented_servers=false:api/proto/
.PHONY: dash
dash: # create grafana dashboard
	 for app in $(apps) ;\
	 do \
	 	jsonnet -J ./grafana/grafonnet-lib   -o ./configs/grafana/dashboards/$$app.json  --ext-str app=$$app ./scripts/grafana/dashboard.jsonnet ;\
	 done
.PHONY: pubdash
pubdash:
	 for app in $(apps) ;\
	 do \
	 	jsonnet -J ./grafana/grafonnet-lib  -o ./configs/grafana/dashboards-api/$$app-api.json  --ext-str app=$$app  ./scripts/grafana/dashboard-api.jsonnet ; \
	 	curl -X DELETE --user admin:admin  -H "Content-Type: application/json" 'http://localhost:3000/api/dashboards/db/$$app'; \
	 	curl -x POST --user admin:admin  -H "Content-Type: application/json" --data-binary "@./configs/grafana/dashboards-api/$$app-api.json" http://localhost:3000/api/dashboards/db ; \
	 done
.PHONY: rules
rules:
	for app in $(apps) ;\
	do \
	 	jsonnet  -o ./configs/prometheus/rules/$$app.yml --ext-str app=$$app  ./scripts/prometheus/rules.jsonnet ; \
	done
.PHONY: docker
docker-compose: build dash rules
	docker-compose -f deployments/docker-compose.yml up --build -d
all: lint cover docker
