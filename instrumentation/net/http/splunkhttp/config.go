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

package splunkhttp

import (
	"os"
	"strings"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Environmental variables used for configuration.
const (
	envVarServerTimingEnabled = "SPLUNK_CONTEXT_SERVER_TIMING_ENABLED" // Adds `Server-Timing` header to HTTP responses
)

// WithOTelOpts is use to pass the OpenTelemetry SDK options.
func WithOTelOpts(opts ...otelhttp.Option) Option {
	return optionFunc(func(cfg *config) {
		cfg.OTelOpts = opts
	})
}

// WithServerTiming enables or disables the ServerTimingMiddleware.
//
// The default is to enable the ServerTimingMiddleware if this option is not passed.
// Additionally, the SPLUNK_CONTEXT_SERVER_TIMING_ENABLED environment variable
// can be set to TRUE or FALSE to specify this option. This option value will be
// given precedence if both it and the environment variable are set.
func WithServerTiming(v bool) Option {
	return optionFunc(func(cfg *config) {
		cfg.ServerTimingEnabled = v
	})
}

// Option is used for setting optional config properties.
type Option interface {
	apply(*config)
}

// config represents the available configuration options.
type config struct {
	OTelOpts            []otelhttp.Option
	ServerTimingEnabled bool
}

// optionFunc provides a convenience wrapper for simple Options
// that can be represented as functions.
type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// newConfig creates a new config struct and applies opts to it.
func newConfig(opts ...Option) *config {
	serverTimingEnabled := true
	if v := os.Getenv(envVarServerTimingEnabled); strings.EqualFold(v, "false") {
		serverTimingEnabled = false
	}

	c := &config{
		ServerTimingEnabled: serverTimingEnabled,
	}
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}