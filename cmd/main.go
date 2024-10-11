// Copyright 2024-2024 Mipy Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"

		"github.com/mia-platform/mipy/pkg/controller"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "my-cli",
        Short: "My CLI tool",
        Long:  `A longer description of My CLI tool.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Hello, CLI!")
        },
    }

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}