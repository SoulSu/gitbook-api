// Copyright © 2019 NAME HERE <soul.sxd@gmail.com>
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
	"github.com/SoulSu/gitbook-api/cmd/server"
	"github.com/SoulSu/gitbook-api/libs/log"

	"github.com/spf13/cobra"
)

var port string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `run server model`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("start server on:%s", port)
		if err := server.RunServer(port); err != nil {
			panic("Start server err:" + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&port, "port", "p", "8888", "the server listen port")
}
