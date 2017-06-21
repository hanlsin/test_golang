// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/hanlsin/test_golang/cobra/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mainCmd = &cobra.Command{
	Use: "cobra_test",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("mainCmd: PersistentPreRunE")
		return nil
	},
	/*
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("mainCmd: Run")
		},
	*/
}

var config string

func init() {
	mainCmd.Flags().StringVarP(&config, "config", "c", "(S) config file", "(L) config file")
	viper.BindPFlag("config", mainCmd.Flags().Lookup("config"))
}

func main() {
	cobra.OnInitialize(func() {
		fmt.Println("cobra.OnInitialize~!!")
	})

	mainCmd.AddCommand(cmd.VersionCmd)
	mainCmd.AddCommand(cmd.GetServerCmd())

	if err := mainCmd.Execute(); err != nil {
		panic(err)
	}

	fmt.Println("-------------------------------------------")

	cmd.RootCmd.Execute()

	fmt.Println("Bye~!")
}
