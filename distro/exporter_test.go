// Copyright Splunk Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package distro

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	noneRealm    = "none"
	invalidRealm = "not-a-valid-realm"
	fakeEndpoint = "some non-zero value"
)

func TestOTLPTracesEndpoint(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assert.Equal(t, "", otlpTracesEndpoint())
	})

	t.Run("none realm", func(t *testing.T) {
		t.Setenv(splunkRealmKey, noneRealm)

		assert.Equal(t, "", otlpTracesEndpoint())
	})

	t.Run("realm", func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)

		want := fmt.Sprintf(otlpRealmEndpointFormat, invalidRealm)
		assert.Equal(t, want, otlpTracesEndpoint())
	})

	t.Run(otelExporterOTLPEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterOTLPEndpointKey, fakeEndpoint)

		// SPLUNK_REALM is set, make sure it does not take precedence.
		assert.Equal(t, "", otlpTracesEndpoint())
	})

	t.Run(otelExporterOTLPTracesEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterOTLPTracesEndpointKey, "some non-zero value")

		// SPLUNK_REALM is set, make sure it does not take precedence.
		assert.Equal(t, "", otlpTracesEndpoint())
	})

	t.Run(otelExporterOTLPMetricsEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterOTLPMetricsEndpointKey, "some non-zero value")

		// OTEL_EXPORTER_OTLP_METRICS_ENDPOINT is ignored for traces exporter.
		want := fmt.Sprintf(otlpRealmEndpointFormat, invalidRealm)
		assert.Equal(t, want, otlpTracesEndpoint())
	})
}

func TestOTLPMetricsEndpoint(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assert.Equal(t, "", otlpMetricsEndpoint())
	})

	t.Run("none realm", func(t *testing.T) {
		t.Setenv(splunkRealmKey, noneRealm)

		assert.Equal(t, "", otlpMetricsEndpoint())
	})

	t.Run("realm", func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)

		want := fmt.Sprintf(otlpRealmEndpointFormat, invalidRealm)
		assert.Equal(t, want, otlpMetricsEndpoint())
	})

	t.Run(otelExporterOTLPEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterOTLPEndpointKey, fakeEndpoint)

		// SPLUNK_REALM is set, make sure it does not take precedence.
		assert.Equal(t, "", otlpMetricsEndpoint())
	})

	t.Run(otelExporterOTLPMetricsEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterOTLPMetricsEndpointKey, "some non-zero value")

		// SPLUNK_REALM is set, make sure it does not take precedence.
		assert.Equal(t, "", otlpMetricsEndpoint())
	})

	t.Run(otelExporterOTLPTracesEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterOTLPTracesEndpointKey, "some non-zero value")

		// OTEL_EXPORTER_OTLP_TRACES_ENDPOINT is ignored for metrics exporter.
		want := fmt.Sprintf(otlpRealmEndpointFormat, invalidRealm)
		assert.Equal(t, want, otlpMetricsEndpoint())
	})
}

func TestJaegerEndpoint(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assert.Equal(t, defaultJaegerEndpoint, jaegerEndpoint())
	})

	t.Run("none realm", func(t *testing.T) {
		t.Setenv(splunkRealmKey, noneRealm)

		assert.Equal(t, defaultJaegerEndpoint, jaegerEndpoint())
	})

	t.Run("realm", func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)

		want := fmt.Sprintf(realmEndpointFormat, invalidRealm)
		assert.Equal(t, want, jaegerEndpoint())
	})

	t.Run(otelExporterJaegerEndpointKey, func(t *testing.T) {
		t.Setenv(splunkRealmKey, invalidRealm)
		t.Setenv(otelExporterJaegerEndpointKey, fakeEndpoint)

		// SPLUNK_REALM is still set, make sure it does not take precedence.
		assert.Equal(t, "", jaegerEndpoint())
	})
}
