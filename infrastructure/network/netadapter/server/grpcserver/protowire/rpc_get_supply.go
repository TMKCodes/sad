package protowire

import (
	"github.com/sadnetwork/sad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SadMessage_GetCoinSupplyRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.GetCoinSupplyRequestMessage{}, nil
}

func (x *SadMessage_GetCoinSupplyRequest) fromAppMessage(_ *appmessage.GetCoinSupplyRequestMessage) error {
	x.GetCoinSupplyRequest = &GetCoinSupplyRequestMessage{}
	return nil
}

func (x *SadMessage_GetCoinSupplyResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_GetCoinSupplyResponse is nil")
	}
	return x.GetCoinSupplyResponse.toAppMessage()
}

func (x *SadMessage_GetCoinSupplyResponse) fromAppMessage(message *appmessage.GetCoinSupplyResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetCoinSupplyResponse = &GetCoinSupplyResponseMessage{
		MaxSeep:         message.MaxSeep,
		CirculatingSeep: message.CirculatingSeep,

		Error: err,
	}
	return nil
}

func (x *GetCoinSupplyResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetCoinSupplyResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	return &appmessage.GetCoinSupplyResponseMessage{
		MaxSeep:         x.MaxSeep,
		CirculatingSeep: x.CirculatingSeep,

		Error: rpcErr,
	}, nil
}
