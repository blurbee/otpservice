package comms

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/blurbee/otpserver/api"
	"github.com/blurbee/otpserver/util"
)

type twilioSender struct {
	accountId   string
	authToken   string
	phoneNumber string
}

func TwilioInit(cfg *util.Config) (txtSender api.TextSender, err api.StatusCode) {
	t := new(twilioSender)
	err = t.Init(cfg)
	return t, err
}

func (t *twilioSender) Init(cfg *util.Config) (err api.StatusCode) {
	tcfg := cfg.GetTwilioConfig()

	t.accountId = cfg.GetSecret(tcfg.AccountIdEnv)
	t.authToken = cfg.GetSecret(tcfg.AuthTokenEnv)
	t.phoneNumber = tcfg.PhoneNumber
	if t.accountId == "" || t.authToken == "" || t.phoneNumber == "" {
		util.Error("Invalid Twilio account configuration. Account, AuthToken and Phone number must be present and valid.")
		err = api.CONN_FAILED
		return
	}
	return api.OK
}

/*
 */
func (t *twilioSender) SendText(dest string, message string) (err api.StatusCode) {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + t.accountId + "/Messages.json"
	v := url.Values{}
	v.Set("To", dest)
	v.Set("From", t.phoneNumber)
	v.Set("Body", message)
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(t.accountId, t.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, er := client.Do(req)
	if er != nil {
		util.Error("Error sending text to ", dest, " with error:", er.Error())
		err = api.SEND_ERROR
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, er := io.ReadAll(resp.Body)
		if er != nil {
			util.Debug("Error reading response from twilio sender:", er.Error())
			err = api.SEND_ERROR
		} else {
			er = json.Unmarshal(bodyBytes, &data)
			if er == nil {
				util.Info("Message sent:", data["sid"])
				util.Debug("Message sent to ", dest, " with response: ", data)
				err = api.OK
			} else {
				util.Debug("Error processing response from twilio sender:", er.Error())
				err = api.SEND_ERROR
			}
		}
	} else {
		util.Warn("Message send error:", resp.Status)
		err = api.SEND_ERROR
	}
	return
}
