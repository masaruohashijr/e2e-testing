package secondary

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/mms"
)

type mmsAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewMmsApi(config *config.ConfigType) mms.SecondaryPort {
	return &mmsAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *mmsAPI) SendMMS(to, from, message string) error {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/MMS/Messages.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("To", to)
	values.Add("From", from)
	values.Add("Body", message)
	values.Add("StatusCallback", services.BaseUrl+"/MmsStatus")

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a *mmsAPI) ViewMMS(smsSid string) (domains.Mms, error) {
	s := domains.Mms{}
	return s, nil
}

func (a *mmsAPI) ListMMS(from, to string) ([]domains.Mms, error) {
	s := domains.Mms{}
	ss := []domains.Mms{s}
	return ss, nil
}
