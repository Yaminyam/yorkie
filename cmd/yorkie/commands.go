/*
 * Copyright 2020 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package main is the entry point of the Yorkie CLI.
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/yorkie-team/yorkie/cmd/yorkie/config"
	"github.com/yorkie-team/yorkie/cmd/yorkie/document"
	"github.com/yorkie-team/yorkie/cmd/yorkie/project"
)

var rootCmd = &cobra.Command{
	Use:   "yorkie",
	Short: "Document store for collaborative applications based on CRDT",
}

// Run executes CLI.
func Run() int {
	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}

func init() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
	rootCmd.AddCommand(project.SubCmd)
	rootCmd.AddCommand(document.SubCmd)
	// TODO(chacha912): set rpcAddr from env using viper.
	// https://github.com/spf13/cobra/blob/main/user_guide.md#bind-flags-with-config
	rootCmd.PersistentFlags().StringVar(&config.RPCAddr, "rpc-addr", "localhost:11101", "Address of the rpc server")
	rootCmd.PersistentFlags().BoolVar(&config.IsInsecure, "insecure", false, "Skip the TLS connection of the client")
}
