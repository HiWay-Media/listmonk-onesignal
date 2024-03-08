package lib

// https://documentation.onesignal.com/reference/create-notification#sms-content

type onesignalCfg struct {
	AppID       string `json:"app_id"`
    RestAppKey   string `json:"rest_api_key"`
	Log         bool   `json:"log"`
}