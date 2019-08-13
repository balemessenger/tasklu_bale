package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ConfYaml
}

func NewConfig(path string) *Config {
	conf, err := loadConf(path)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return nil
	}
	return &Config{
		ConfYaml: conf,
	}
}

var defaultConf = []byte(`
core:
  mode: "release" # release, debug, test
  work_pool_size: 1000

prometheus:
  port: 8080
log:
  level: debug
endpoints:

  grpc:
    address: "127.0.0.1:5050"

  http:
    address: ":4040"
    user: "test"
    pass: "test"
`)

type ConfYaml struct {
	Core SectionCore `yaml:"core"`

	Prometheus SectionPrometheus `yaml:"prometheus"`
	Log        SectionLog        `yaml:"log"`
	Endpoints  SectionEndpoints  `yaml:"endpoints"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Mode         string `yaml:"mode"`
	WorkPoolSize int    `yaml:"work_pool_size"`
}

// SectionPostgres is sub section of config.
type SectionPostgres struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DB         string `yaml:"db"`
	User       string `yaml:"user"`
	Pass       string `yaml:"pass"`
	BatchCount int    `yaml:"batch_count"`
}

// SectionKafka is sub section of config.
type SectionKafka struct {
	BootstrapServers string `yaml:"bootstrap_servers"`
	GroupId          string `yaml:"group_id"`
	AutoOffsetReset  string `yaml:"auto_offset_reset"`
	Topic            string `yaml:"topic"`
}

type SectionPrometheus struct {
	Port int `yaml:"port"`
}

type SectionLog struct {
	Level string `yaml:"level"`
}

type SectionEndpoints struct {
	Grpc SectionGrpc `yaml:"grpc"`

	Http SectionHttp `yaml:"http"`
}

type SectionGrpc struct {
	Address string `yaml:"address"`
}

type SectionHttp struct {
	Address string `yaml:"address"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
}

// LoadConf load config from file and read in environment variables that match
func loadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()          // read in environment variables that match
	viper.SetEnvPrefix("taskulu") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			log.Errorf("File does not exist : %s", confPath)
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".pkg" (without extension).
		viper.AddConfigPath("/etc/taskulu/")
		viper.AddConfigPath("$HOME/.taskulu")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.WorkPoolSize = viper.GetInt("core.work_pool_size")

	// Prometheus
	conf.Prometheus.Port = viper.GetInt("prometheus.port")

	//Log
	conf.Log.Level = viper.GetString("log.level")

	//Endpoints

	conf.Endpoints.Grpc.Address = viper.GetString("endpoints.grpc.address")

	conf.Endpoints.Http.Address = viper.GetString("endpoints.http.address")
	conf.Endpoints.Http.User = viper.GetString("endpoints.http.user")
	conf.Endpoints.Http.Pass = viper.GetString("endpoints.http.pass")

	return conf, nil
}
