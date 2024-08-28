package lib

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/francoispqt/onelog"
)

// https://documentation.onesignal.com/reference/create-notification#sms-content

type onesignalCfg struct {
	AppID      string `json:"app_id"`
	RestAppKey string `json:"rest_api_key"`
	Log        bool   `json:"log"`
}

type onesignalMessenger struct {
	cfg    onesignalCfg
	logger *onelog.Logger
	ctx    context.Context
}

func (o onesignalMessenger) Name() string {
	return "onesignal"
}

// Push sends the email through onesignal API.
func (p onesignalMessenger) Push(msg Message) error {
	email, ok := msg.Subscriber.Attribs["email"].(string)
	if !ok {
		return fmt.Errorf("could not find subscriber phone")
	}
	p.logger.InfoWith("sending onesignal message" + email).Write()
	//
	return nil
}

func (p onesignalMessenger) Flush() error {
	return nil
}

func (p onesignalMessenger) Close() error {
	return nil
}

// NewOneSignal creates new instance of pinpoint
func NewOneSignal(cfg []byte, l *onelog.Logger) (Messenger, error) {
	var c onesignalCfg
	if err := json.Unmarshal(cfg, &c); err != nil {
		return nil, err
	}

	if c.AppID == "" {
		return nil, fmt.Errorf("invalid app_id")
	}
	ctx := context.WithValue(context.Background(), onesignal.AppAuth, c.AppID)
	//
	return onesignalMessenger{
		cfg:    c,
		logger: l,
		ctx:    ctx,
	}, nil
}
