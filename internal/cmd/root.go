// Copyright Mia srl
// SPDX-License-Identifier: Apache-2.0
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

package cmd

import (
	"context"
	"os"

	"github.com/go-logr/logr"
	"github.com/mia-platform/mipy/internal/logger"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mipy",
		Short: "Mia-Platform Infrastructure Provisioning Helper",
		Long:  `Find more information at: https://github.com/mia-platform/mipy`,
	}

	// initialize clioptions and setup during initialization
	// options := clioptions.NewCLIOptions()

	logger := logger.NewLogger(os.Stderr)
	rootCmd.SetContext(logr.NewContext(context.Background(), logger))

	// add sub commands
	rootCmd.AddCommand(
		ConfigCmd(),
		InitCmd(),
		LaunchCmd(),
	)

	return rootCmd
}
