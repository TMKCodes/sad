package main

import (
	"context"
	"fmt"

	"github.com/sadnetwork/sad/cmd/sadwallet/daemon/client"
	"github.com/sadnetwork/sad/cmd/sadwallet/daemon/pb"
	"github.com/sadnetwork/sad/cmd/sadwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatSAD(addressBalance.Available), utils.FormatSad(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, Sad %s %s%s\n", utils.FormatSad(response.Available), utils.FormatSad(response.Pending), pendingSuffix)

	return nil
}
