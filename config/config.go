package config

type Configuration struct {
	Environment Environment `mapstructure:"app" json:"app" yaml:"app"`
}
