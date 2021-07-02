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

type ConfigType struct {
	HttpClient     http.Client
	AccountSid     string
	AuthToken      string
	From           string
	FromSid        string
	To             string
	ToSid          string
	AvayaNumBR     string
	AvayaNumCA     string
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
	ApiVersion     string
	StatusCallback string
	Fallback       string
	Timeout        int
}

func (c ConfigType) GetBaseURL() string {
	return c.ApiUrl + "/" + c.ApiVersion
}

func NewConfig() (config ConfigType) {
	config = ReadConfig(config)
	config.HttpClient = http.Client{
		Timeout: 60 * time.Second,
	}
	return config
}

func ReadConfig(config ConfigType) ConfigType {
	var configfile = "../../internal/config/config.ini"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("File configuration "+configfile+" missing: ", configfile)
	}
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *ConfigType) SelectNumber(option string) (string, string) {
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
	}
	return option, ""
}
