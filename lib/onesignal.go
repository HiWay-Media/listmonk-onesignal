package lib

import (

	"github.com/francoispqt/onelog"
)

// https://documentation.onesignal.com/reference/create-notification#sms-content

type onesignalCfg struct {
	AppID       string `json:"app_id"`
    RestAppKey   string `json:"rest_api_key"`
	Log         bool   `json:"log"`
}

type onesignalMessenger struct {
	cfg    onesignalCfg

	logger *onelog.Logger
}

func (o onesignalMessenger) Name() string {
	return "onesignal"
}


// Push sends the sms through onesignal API.
func (p onesignalMessenger) Push(msg Message) error {


}


func (p onesignalMessenger) Flush() error {
	return nil
}

func (p onesignalMessenger) Close() error {
	return nil
}


// NewOneSignal creates new instance of pinpoint
func NewOneSignal(cfg []byte, l *onelog.Logger) (Messenger, error) {
    var c pinpointCfg
	if err := json.Unmarshal(cfg, &c); err != nil {
		return nil, err
	}
    //
    return onesignalMessenger{
		cfg:    c,
		logger: l,
	}, nil
}