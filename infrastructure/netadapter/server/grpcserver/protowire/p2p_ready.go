package protowire

import (
	"github.com/sadnetwork/sad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *SadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
