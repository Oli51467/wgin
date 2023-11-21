package config

type Jwt struct {
	Secret             string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl             int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`                                        // token 有效期（秒）
	JwtBlacklistPeriod int64  `mapstructure:"jwt_blacklist_period" json:"jwt_blacklist_period" yaml:"jwt_blacklist_period"` // 黑名单宽限时间（秒）
	RefreshPeriod      int64  `mapstructure:"refresh_period" json:"refresh_period" yaml:"refresh_period"`                   // token 自动刷新宽限时间（秒）
}
