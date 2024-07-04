package protowire

import (
	"github.com/sadnetwork/sad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SadMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest is nil")
	}
	return &appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage{}, nil
}

func (x *SadMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest) fromAppMessage(_ *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage) error {
	x.NotifyVirtualSelectedParentBlueScoreChangedRequest = &NotifyVirtualSelectedParentBlueScoreChangedRequestMessage{}
	return nil
}

func (x *SadMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse is nil")
	}
	return x.NotifyVirtualSelectedParentBlueScoreChangedResponse.toAppMessage()
}

func (x *SadMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse) fromAppMessage(message *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyVirtualSelectedParentBlueScoreChangedResponse = &NotifyVirtualSelectedParentBlueScoreChangedResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyVirtualSelectedParentBlueScoreChangedResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyVirtualSelectedParentBlueScoreChangedResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *SadMessage_VirtualSelectedParentBlueScoreChangedNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SadMessage_VirtualSelectedParentBlueScoreChangedNotification is nil")
	}
	return x.VirtualSelectedParentBlueScoreChangedNotification.toAppMessage()
}

func (x *SadMessage_VirtualSelectedParentBlueScoreChangedNotification) fromAppMessage(message *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage) error {
	x.VirtualSelectedParentBlueScoreChangedNotification = &VirtualSelectedParentBlueScoreChangedNotificationMessage{
		VirtualSelectedParentBlueScore: message.VirtualSelectedParentBlueScore,
	}
	return nil
}

func (x *VirtualSelectedParentBlueScoreChangedNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "VirtualSelectedParentBlueScoreChangedNotificationMessage is nil")
	}
	return &appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage{
		VirtualSelectedParentBlueScore: x.VirtualSelectedParentBlueScore,
	}, nil
}
