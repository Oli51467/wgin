package config

type Configuration struct {
	Environment Environment `mapstructure:"app" json:"app" yaml:"app"`
	Logger      Logger      `mapstructure:"logger" json:"logger" yaml:"logger"`
	Database    Database    `mapstructure:"database" json:"database" yaml:"database"`
	Jwt         Jwt         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis       Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
}
