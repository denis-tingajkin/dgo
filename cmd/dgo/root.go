// Copyright (c) 2020 Andrey Sobolev.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cmd - cobra commands for forwarder
package dgo

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dgo tooling for docker",
	Short: "This tool `D-go` (aka docker go) is intended to help with this approach and provide smooth and fast experience.",
	Long: `
		Docker go tooling.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute - execute the command
func Execute() error {
	return rootCmd.Execute()
}
