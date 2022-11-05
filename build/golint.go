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

package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var golint = goyek.Define(goyek.Task{
	Name:  "golint",
	Usage: "golangci-lint --fix",
	Action: func(tf *goyek.TF) {
		if !cmd.Exec(tf, "go install github.com/golangci/golangci-lint/cmd/golangci-lint", cmd.Dir(dirBuild)) {
			return
		}
		ForGoModules(tf, func(tf *goyek.TF) {
			cmd.Exec(tf, "golangci-lint run --fix")
		})
	},
})