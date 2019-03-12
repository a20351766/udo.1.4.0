/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/hyperledger/udo/bccsp/factory"
	"github.com/hyperledger/udo/cmd/common"
	"github.com/hyperledger/udo/discovery/cmd"
)

func main() {
	factory.InitFactories(nil)
	cli := common.NewCLI("discover", "Command line client for udo discovery service")
	discovery.AddCommands(cli)
	cli.Run(os.Args[1:])
}
