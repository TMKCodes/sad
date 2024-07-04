package protowire

import (
	"github.com/sadnetwork/sad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SadMessage_Reject) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_Reject is nil")
	}
	return x.Reject.toAppMessage()
}

func (x *RejectMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RejectMessage is nil")
	}
	return &appmessage.MsgReject{
		Reason: x.Reason,
	}, nil
}

func (x *SadMessage_Reject) fromAppMessage(msgReject *appmessage.MsgReject) error {
	x.Reject = &RejectMessage{
		Reason: msgReject.Reason,
	}
	return nil
}
