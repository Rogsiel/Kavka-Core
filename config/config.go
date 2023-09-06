package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

var ENV_ITEMS = []string{"devel", "prod"}
var ENV string

type (
	IConfig struct {
		App   App   `yaml:"APP"`
		Mongo Mongo `yaml:"MONGO"`
		Redis Redis `yaml:"REDIS"`
		SMS   `yaml:"SMS"`
		MinIOCredentials
	}
	App struct {
		Name   string `yaml:"NAME"`
		HTTP   HTTP   `yaml:"HTTP"`
		Server Server `yaml:"SERVER"`
		Auth   Auth   `yaml:"AUTH"`
	}
	HTTP struct {
		Host    string `yaml:"HOST"`
		Port    int    `yaml:"PORT"`
		Address string `yaml:"ADDRESS"`
	}
	Auth struct {
		SECRET             string
		OTP_EXPIRE_SECONDS time.Duration `json:"OTP_EXPIRE_SECONDS"`
	}
	Server struct {
		CORS CORS `yaml:"CORS"`
	}
	CORS struct {
		AllowOrigins string `yaml:"ALLOW_ORIGINS"`
	}
	Redis struct {
		Host     string `yaml:"HOST"`
		Username string `yaml:"USERNAME"`
		Password string `yaml:"PASSWORD"`
		Port     int    `yaml:"PORT"`
		DB       int    `yaml:"DB"`
	}
	Mongo struct {
		Host     string `yaml:"HOST"`
		Username string `yaml:"USERNAME"`
		Password string `yaml:"PASSWORD"`
		Port     int    `yaml:"PORT"`
		DBName   string `yaml:"DB_NAME"`
	}
	MinIOCredentials struct {
		Endpoint  string `json:"endpoint"`
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`
	}
	// TODO - Add sms-service's configs
	SMS struct{}
)

const defaultEnvPath = "/config/configs.yml"

var ProjectRootPath = ConfigsDirPath() + "/../"

func ConfigsDirPath() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error in generating env dir")
	}

	return filepath.Dir(f)
}

func Read() *IConfig {
	// Load ENV
	env := os.Getenv("ENV")
	if len(strings.TrimSpace(env)) == 0 {
		ENV = ENV_ITEMS[0]
	} else if slices.Contains(ENV_ITEMS, env) {
		ENV = env
	} else {
		panic(errors.New("Invalid ENV key: " + env))
	}

	// Load YAML configs
	var cfg *IConfig

	data, readErr := os.ReadFile(ConfigsDirPath() + "/configs.yml")
	if readErr != nil {
		panic(readErr)
	}

	parseErr := yaml.Unmarshal(data, &cfg)
	if parseErr != nil {
		panic(parseErr)
	}

	// Load JwtSecret keys
	secretData, secretErr := os.ReadFile(ConfigsDirPath() + "/jwt_secret.pem")
	if secretErr != nil {
		panic(secretErr)
	}

	cfg.App.Auth.SECRET = strings.TrimSpace(string(secretData))

	// Load MinIO credentials
	filename := "minio-credentials.json"
	credFile, credErr := os.ReadFile(ConfigsDirPath() + "/" + filename)
	if credErr != nil {
		panic(credErr)
	}

	var cred MinIOCredentials

	jsonErr := json.Unmarshal(credFile, &cred)
	if jsonErr != nil {
		panic(jsonErr)
	}

	cfg.MinIOCredentials = cred

	return cfg
}