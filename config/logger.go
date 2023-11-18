package config

type Logger struct {
	Path       string `mapstructure:"path" json:"path" yaml:"path"`
	Name       string `mapstructure:"name" json:"name" yaml:"name"`
	TimeFormat string `mapstructure:"time_format" json:"time_format" yaml:"time_format"`
	Ext        string `mapstructure:"ext" json:"ext" yaml:"ext"`
}
