package config

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Operations int

const (
	MakeCall Operations = iota + 1
	ListCall
	ViewCall
)

var ConfigPath string

type ConfigType struct {
	ConfigPath     string
	HttpClient     http.Client
	AccountSid     string
	AuthToken      string
	From           string
	FromSid        string
	To             string
	ToSid          string
	AvayaNumBR     string
	AvayaNumCA     string
	PhoneNumbers   map[string]PhoneNumber
	Parameters     map[string]string
	NumberA        string
	NumberB        string
	NumberASid     string
	NumberBSid     string
	NumberC        string
	NumberD        string
	NumberCSid     string
	NumberDSid     string
	NumberE        string
	NumberF        string
	NumberESid     string
	NumberFSid     string
	NumberBR1      string
	NumberBR2      string
	NumberBR3      string
	ActionUrl      string
	ApiUrl         string
	VoiceUrl       string
	FriendlyName   string
	ApiVersion     string
	StatusCallback string
	Fallback       string
	Timeout        int
	BaseUrl        string
	Logger         string
}

func (c ConfigType) GetApiURL() string {
	return c.ApiUrl + "/" + c.ApiVersion
}

func NewConfig() (config ConfigType) {
	config.ConfigPath = ConfigPath
	config = ReadConfig(config)
	config.HttpClient = http.Client{
		Timeout: 180 * time.Second,
	}
	return config
}

func ReadConfig(config ConfigType) ConfigType {
	_, err := os.Stat(config.ConfigPath)
	if err != nil {
		log.Fatal("File configuration "+config.ConfigPath+" missing: ", config.ConfigPath)
	}
	if _, err := toml.DecodeFile(config.ConfigPath, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *ConfigType) SelectNumber(option string) (string, string) {
	if phone, ok := c.PhoneNumbers[option]; ok {
		return phone.Number, phone.NumberSid
	}
	switch option {
	case "NumberA":
		return c.NumberA, c.NumberASid
	case "NumberB":
		return c.NumberB, c.NumberBSid
	case "NumberC":
		return c.NumberC, c.NumberCSid
	case "NumberD":
		return c.NumberD, c.NumberDSid
	case "NumberE":
		return c.NumberE, c.NumberESid
	case "NumberF":
		return c.NumberF, c.NumberFSid
	case "NumberBR1":
		return c.NumberBR1, ""
	case "NumberBR2":
		return c.NumberBR2, ""
	}
	return option, ""
}

type Context struct {
	Numbers    []PhoneNumber `json:"phoneNumbers"`
	Parameters []Parameter   `json:"parameters"`
}

type PhoneNumber struct {
	Number    string `json:"phoneNumber"`
	NumberSid string `json:"sid"`
	Alias     string `json:"alias"`
}

type Parameter struct {
	Key   string `json:"name"`
	Value string `json:"value"`
}
