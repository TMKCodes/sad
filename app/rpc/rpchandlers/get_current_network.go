package rpchandlers

import (
	"github.com/sadnetwork/sad/app/appmessage"
	"github.com/sadnetwork/sad/app/rpc/rpccontext"
	"github.com/sadnetwork/sad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
