// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"

	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var appVersion = "1.0.0"

var app *cli.App

func init() {
	app = utils.NewApp(appVersion, "a Fusion account manager")
	app.Commands = []cli.Command{
		commandNew,
		commandImport,
		commandList,
		commandUpdate,
		commandShowkey,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
