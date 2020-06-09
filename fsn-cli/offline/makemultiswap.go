// Copyright 2019 The fsn-go-sdk Authors
// This file is part of the fsn-go-sdk library.
//
// The fsn-go-sdk library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fsn-go-sdk library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fsn-go-sdk library. If not, see <http://www.gnu.org/licenses/>.

package offline

import (
	"github.com/fsn-dev/fsn-go-sdk/efsn/cmd/utils"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/fsnapi"
	"gopkg.in/urfave/cli.v1"
)

var CommandMakeMultiSwap = cli.Command{
	Name:      "makemultiswap",
	Aliases:   []string{"makemswap"},
	Category:  "offline",
	Usage:     "build make multi-swap raw transaction",
	ArgsUsage: "",
	Description: `
build make multi-swap raw transaction`,
	Flags: append([]cli.Flag{
		multiSwapFromAssetIDFlag,
		multiSwapFromAmountFlag,
		multiSwapFromStartFlag,
		multiSwapFromEndFlag,
		multiSwapToAssetIDFlag,
		multiSwapToAmountFlag,
		multiSwapToStartFlag,
		multiSwapToEndFlag,
		swapSwapSizeFlag,
		swapTargetsFlag,
		descriptionFlag,
	}, commonFlags...),
	Action: makemultiswap,
}

func makemultiswap(ctx *cli.Context) error {
	setLogger(ctx)
	if len(ctx.Args()) != 3 {
		cli.ShowCommandHelpAndExit(ctx, "makemultiswap", 1)
	}

	fomeAssetID := getHashSlice(ctx, multiSwapFromAssetIDFlag.Name)
	fromAmount := getHexBigIntSlice(ctx, multiSwapFromAmountFlag.Name)
	fromStartTime := getHexUint64Slice(ctx, multiSwapFromStartFlag.Name)
	fromEndTime := getHexUint64Slice(ctx, multiSwapFromEndFlag.Name)
	toAssetID := getHashSlice(ctx, multiSwapToAssetIDFlag.Name)
	toAmount := getHexBigIntSlice(ctx, multiSwapToAmountFlag.Name)
	toStartTime := getHexUint64Slice(ctx, multiSwapToStartFlag.Name)
	toEndTime := getHexUint64Slice(ctx, multiSwapToEndFlag.Name)
	swapSize := getBigInt(ctx, swapSwapSizeFlag.Name)
	targets := getAddressSlice(ctx, swapTargetsFlag.Name)
	description := ctx.String(descriptionFlag.Name)

	// 1. construct corresponding arguments and options
	baseArgs, signOptions := getBaseArgsAndSignOptions(ctx)
	args := &common.MakeMultiSwapArgs{
		FusionBaseArgs: baseArgs,
		FromAssetID:    fomeAssetID,
		FromStartTime:  fromStartTime,
		FromEndTime:    fromEndTime,
		MinFromAmount:  fromAmount,
		ToAssetID:      toAssetID,
		ToStartTime:    toStartTime,
		ToEndTime:      toEndTime,
		MinToAmount:    toAmount,
		SwapSize:       swapSize,
		Targes:         targets,
		Description:    description,
	}

	// 2. check parameters
	now := getNowTime()
	args.Init(getBigIntFromUint64(now))
	if err := args.ToParam().Check(common.BigMaxUint64, now); err != nil {
		utils.Fatalf("check parameter failed: %v", err)
	}

	// 3. build and/or sign transaction through fsnapi
	tx, err := fsnapi.BuildFSNTx(common.MakeMultiSwapFunc, args, signOptions)
	if err != nil {
		utils.Fatalf("create tx error: %v", err)
	}

	return printTx(tx, false)
}
