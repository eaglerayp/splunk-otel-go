module github.com/signalfx/splunk-otel-go/instrumentation/internal

go 1.18

require (
	github.com/signalfx/splunk-otel-go v1.1.0
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/otel v1.11.1
	go.opentelemetry.io/otel/trace v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/signalfx/splunk-otel-go => ../../
