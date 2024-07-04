package protowire

import (
	"github.com/sadnetwork/sad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *SadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
