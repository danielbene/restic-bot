package util

import (
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Config struct {
	State       string   `mapstructure:"good"`
	Orientation string   `mapstructure:"left"`
	Look        []string `mapstructure:"color"`
}

func ParseYml(f []byte) *Config {
	var conf Config
	var raw interface{}

	if err := yaml.Unmarshal(f, &raw); err != nil {
		panic(err)
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &conf})
	if err := decoder.Decode(raw); err != nil {
		panic(err)
	}

	return &conf
}
