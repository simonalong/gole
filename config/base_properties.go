package config

var GoleCfg GoleConfig

// GoleConfig base前缀
type GoleConfig struct {
	Application GoleApplication `yaml:"application"`
	Logger      GoleLogger      `yaml:"logger"`
}

type GoleApplication struct {
	Name       string `yaml:"name"`        // 应用名字
	Version    string `yaml:"version"`     // 应用版本
	BannerShow bool   `yaml:"banner-show"` // 是否展示banner
}

type GoleLogger struct {
	Level string // 日志root级别：trace/debug/info/warn/error/fatal/panic，默认：info
	Path  string
	Time  LoggerTime  // 时间配置
	Color LoggerColor // 日志颜色
	Split LoggerSplit // 日志切分
	Dir   string
	Max   struct {
		History int
	}
	Console struct {
		WriteFile bool
	}
}

type LoggerTime struct {
	Format string `yaml:"format"` // 时间格式，time包中的内容，比如：time.RFC3339
}

type LoggerColor struct {
	Enable bool `yaml:"enable"` // 是否启用
}

type LoggerSplit struct {
	Enable bool  `yaml:"enable"` // 日志是否启用切分：true/false，默认false
	Size   int64 `yaml:"size"`   // 日志拆分的单位：MB
}

type GoleProfile struct {
	Active string `yaml:"active"`
}

type StorageConnectionConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Parameters string `yaml:"parameters"`
}

// ---------------------------- gole.logger ----------------------------

type LoggerConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"level"`
	Time  struct {
		Format string `yaml:"format"`
	} `yaml:"time"`
	Color struct {
		Enable bool `yaml:"enable"`
	} `yaml:"color"`
	Split struct {
		Enable bool  `yaml:"enable"`
		Size   int64 `yaml:"size"`
	} `yaml:"split"`
	Dir string `yaml:"dir"`
	Max struct {
		History int `yaml:"history"`
	} `yaml:"max"`
	Console struct {
		WriteFile bool `yaml:"writeFile"`
	} `yaml:"console"`
}
