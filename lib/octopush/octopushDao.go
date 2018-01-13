package octopush

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/lzientek/octopush-middleware/config"
)

type OctopushSms struct {
	Userlogin     string `json:"user_login"`
	APIKey        string `json:"api_key"`
	SmsRecipients string `json:"sms_recipients"`
	SmsText       string `json:"sms_text"`
	SmsType       string `json:"sms_type"`
	SmsSender     string `json:"sms_sender"`
	RequestMode   string `json:"request_mode"`
}

type OctopushResult struct {
	Cost             string `json:"cost"`
	Ticket           string `json:"ticket"`
	SendingDate      string `json:"sending_date"`
	NumberOfSendings string `json:"number_of_sendings"`
	CurrencyCode     string `json:"currency_code"`
	Successes        string `json:"successs"`
	Recipients       string `json:"recipient"`
	CountryCode      string `json:"country_code"`
	Failures         string `json:"failures"`
}

func (o *OctopushSms) Send() (OctopushResult, error) {
	c := config.GetConfig()
	url := c.GetString("app.octpush.url")
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(o)
	if o.RequestMode == "" {
		o.RequestMode = c.GetString("app.octpush.request_mode")
	}
	resp, err := http.Post(url, "application/json", buffer)

	var result OctopushResult

	if err != nil {
		err = json.NewDecoder(resp.Body).Decode(&result)
		return result, err
	}

	return result, err
}
