package octopush

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/lzientek/octopush-middleware/config"
)

type OctopushSms struct {
	Userlogin     string `form:"user_login"`
	APIKey        string `form:"api_key"`
	SmsRecipients string `form:"sms_recipients"`
	SmsText       string `form:"sms_text"`
	SmsType       string `form:"sms_type"`
	SmsSender     string `form:"sms_sender"`
	RequestMode   string `form:"request_mode"`
}

type OctopushResult struct {
	Cost             float32 `json:"cost"`
	Balance          float32 `json:"balance"`
	Ticket           string  `json:"ticket"`
	SendingDate      int     `json:"sending_date"`
	NumberOfSendings int     `json:"number_of_sendings"`
	CurrencyCode     string  `json:"currency_code"`
}

func (o *OctopushSms) Send() (OctopushResult, error) {
	c := config.GetConfig()
	octoUrl := c.GetString("app.octopush.url")

	if o.RequestMode == "" {
		o.RequestMode = c.GetString("app.octopush.request_mode")
	}

	resp, err := http.PostForm(octoUrl, url.Values{
		"user_login":     {o.Userlogin},
		"api_key":        {o.APIKey},
		"sms_recipients": {o.SmsRecipients},
		"sms_text":       {o.SmsText},
		"sms_type":       {o.SmsType},
		"sms_sender":     {o.SmsSender},
		"request_mode":   {o.RequestMode},
	})

	var result OctopushResult

	if err == nil {

		err = json.NewDecoder(resp.Body).Decode(&result)
		return result, err
	}

	return result, err
}
