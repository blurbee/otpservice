package util

import (
	"io"
	"log"
	"os"

	"github.com/blurbee/otpserver/api"
	"gopkg.in/yaml.v3"
)

var debug *log.Logger
var info *log.Logger
var warn *log.Logger
var erur *log.Logger

var isDebug = true

type MongoStoreConfig struct {
	Id               string `yaml:"id"`
	Database         string `yaml:"database"`
	Collection       string `yaml:"collection"`
	ConnectionUrlEnv string `yaml:"connectionurlenv"`
}

type PostgresStoreConfig struct {
	Id               string `yaml:"id"`
	Database         string `yaml:"database"`
	Table            string `yaml:"table"`
	Column           string `yaml:"column"`
	ConnectionUrlEnv string `yaml:"connectionurlenv"`
}

type EmailStoreConfig struct {
	StoreType string `yaml:"storetype"`
	// UserIdField string `yaml:"useridfieldname"` // For mongo field name is "email"
	StoreId string `yaml:"id"` // id of the data store (mongo or postgres) config
}

type PhoneStoreConfig struct {
	StoreType string `yaml:"storetype"`
	// UserIdField string `yaml:"useridfieldname"` // for mongo field name is "phone"
	StoreId string `yaml:"storeid"` // id of the data store (mongo or postgres) config
}

type WhatsappStoreConfig struct {
	StoreType string `yaml:"storetype"`
	// UserIdField string `yaml:"useridfieldname"` // for mongo field name is "whatsapp"
	StoreId string `yaml:"storeid"` // id of the data store (mongo or postgres) config
}

type TextStoreConfig struct {
	StoreType string `yaml:"storetype"`
	// UserIdField string `yaml:"useridfieldname"` // for mongo field name is "text"
	StoreId string `yaml:"storeid"` // id of the data store (mongo or postgres) config
}

type Scenario struct {
	Id               string              `yaml:"id"`
	Ttl              int                 `yaml:"ttl"`
	KeyAlpha         bool                `yaml:"keywithalpha"`
	NumAttempts      int                 `yaml:"numattempts"`
	AllowEmail       bool                `yaml:"allowemail"`
	AllowText        bool                `yaml:"allowtext"`
	AllowWhatsapp    bool                `yaml:"allowwhatsapp"`
	KeyLeadingText   string              `yaml:"leadingtext"`
	PhoneStoreCfg    PhoneStoreConfig    `yaml:"phonestorecfg"`
	TextStoreCfg     TextStoreConfig     `yaml:"textstoreconfig"`
	EmailStoreCfg    EmailStoreConfig    `yaml:"emailstorecfg"`
	WhatsappStoreCfg WhatsappStoreConfig `yaml:"whatsappstorecfg"`
}

type TwilioConfig struct {
	AccountIdEnv string `yaml:"accountidenv"`
	AuthTokenEnv string `yaml:"authtokenenv"`
	PhoneNumber  string `yaml:"phonenumber"`
}

type EmailConfig struct {
	MailServer  string `yaml:"mailserver"`
	Port        int    `yaml:"port"`
	User        string `yaml:"username"`
	PasswordEnv string `yaml:"passwordenv"`
	Email       string `yaml:"email"`
}

type Config struct {
	SecretsFile       string                `yaml:"secretsfile"`
	MongoStoreCfgs    []MongoStoreConfig    `yaml:"mongostores"`
	PostgresStoreCfgs []PostgresStoreConfig `yaml:"postgresstores"`
	TwilioCfg         TwilioConfig          `yaml:"twilioconfig"`
	EmailCfg          EmailConfig           `yaml:"emailserverconfig"`
	Scenarios         map[string]Scenario   `yaml:"scenarios"`
	secrets           map[string]string
}

type Secrets struct {
	Secrets map[string]string `yaml:"secrets"`
}

func InitLogs() (err api.StatusCode) {
	// logfile, er := os.Create("app.log")

	// if er != nil {
	// 	log.Fatal(er)
	// }

	// defer logfile.Close()
	debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	//	debug.SetOutput(logfile)

	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	//	info.SetOutput(logfile)

	warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	//	warn.SetOutput(logfile)

	erur = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime)
	//	erur.SetOutput(logfile)

	return api.OK
}

func Info(v ...any) {
	info.Println(v...)
}

func SetDebug(doDebug bool) {
	isDebug = doDebug
}

func Debug(v ...any) {
	if isDebug {
		debug.Println(v...)
	}
}

func Warn(v ...any) {
	warn.Println(v...)
}

func Error(v ...any) {
	erur.Println(v...)
}

func LoadConfig(file io.Reader, cfg *Config) (status api.StatusCode) {

	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(cfg)
	if err != nil {
		Error("Loading configuartions failed", err)
		status = api.CONFIG_ERROR
		return
	}

	// load secrets
	if cfg.SecretsFile == "" {
		Error("Secrets file information missing.")
		return api.CONFIG_LOAD_FAILED
	}

	secFile, er := os.Open(cfg.SecretsFile)
	if er != nil {
		Error("Secrets file cannot be opened.", er)
		return api.CONFIG_LOAD_FAILED
	}

	var sec Secrets

	decoder = yaml.NewDecoder(secFile)
	err = decoder.Decode(&sec)
	if err != nil {
		Error("Loading secrets failed: ", err)
		status = api.CONFIG_ERROR
		return
	}
	cfg.secrets = sec.Secrets

	Info("Config load success.")
	return api.OK
}

func (c *Config) GetScenario(name string) (scenario Scenario, err api.StatusCode) {
	scenario, prs := c.Scenarios[name]
	if !prs {
		return scenario, api.INVALID_INPUT
	}
	return scenario, api.OK
}

func (c *Config) GetMongoConfigs() *[]MongoStoreConfig {
	return &c.MongoStoreCfgs
}

func (c *Config) GetTwilioConfig() *TwilioConfig {
	return &c.TwilioCfg
}

func (c *Config) GetSecret(key string) string {
	return c.secrets[key]
}
