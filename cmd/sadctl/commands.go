package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/sadnetwork/sad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.SadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.SadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.SadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.SadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.SadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.SadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.SadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.SadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.SadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.SadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.SadMessage_BanRequest{}),
	reflect.TypeOf(protowire.SadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
