package secondary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"zarbat_test/internal/config"
	"zarbat_test/pkg/domains"
	usage "zarbat_test/pkg/ports/usages"
)

type usageAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewUsageApi(config *config.ConfigType) usage.SecondaryPort {
	return &usageAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *usageAPI) ViewUsage(usageSid string) (domains.Usage, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Usages/%s.json", a.config.AccountSid, usageSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	println(apiEndpoint)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	dummyUsage := domains.Usage{}
	if err != nil {
		return dummyUsage, err
	}
	defer resp.Body.Close()
	// Print Response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyUsage, err
	}
	b := string(body)
	fmt.Println("response Body:", b)
	usage := domains.Usage{}
	json.Unmarshal(body, &usage)
	return usage, nil
}

func (a *usageAPI) ListUsage() ([]domains.Usage, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Usages.json", a.config.AccountSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	println(apiEndpoint)
	q := req.URL.Query()
	today := time.Now()
	q.Add("Day", strconv.Itoa(today.Day()))
	m := int(today.Month())
	q.Add("Month", strconv.Itoa(m))
	q.Add("Year", strconv.Itoa(today.Year()))
	q.Add("Product", "1")
	q.Add("Page", "0")
	q.Add("PageSize", "10")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Print Response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := string(body)
	fmt.Println("response Body:", b)
	usageResponse := domains.UsageResponse{}
	json.Unmarshal(body, &usageResponse)
	for _, usage := range usageResponse.Usages {
		fmt.Println(usage.Product, usage.ProdutctId, usage.Quantity)
	}
	return usageResponse.Usages, nil
}
