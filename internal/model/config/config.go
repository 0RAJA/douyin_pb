package config

import (
	"time"
)

type All struct {
	Mysql  Mysql  `yaml:"Mysql"`
	Log    Log    `yaml:"Log"`
	Redis  Redis  `yaml:"Redis"`
	Email  Email  `yaml:"Email"`
	Token  Token  `yaml:"Token"`
	Rule   Rule   `yaml:"Rule"`
	Server Server `yaml:"Server"`
	App    App    `yaml:"App"`
}

type App struct {
	Name      string `yaml:"Name"`
	Version   string `yaml:"Version"`
	StartTime string `yaml:"StartTime"`
	Format    string `yaml:"Format"`
}

type Redis struct {
	Address         string        `yaml:"Address"`
	DB              int           `yaml:"DB"`
	Password        string        `yaml:"Password"`
	PoolSize        int           `yaml:"PoolSize"`
	PostInfoTimeout time.Duration `yaml:"PostInfoTimeout"`
}

type Log struct {
	Level         string `yaml:"Level"`
	LogSavePath   string `yaml:"LogSavePath"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
	HighLevelFile string `yaml:"HighLevelFile"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
}

type Mysql struct {
	DriverName string `yaml:"DriverName"`
	SourceName string `yaml:"SourceName"`
}

type Token struct {
	Key                 string        `yaml:"Key"`
	AssessTokenDuration time.Duration `yaml:"AssessTokenDuration"`
}

type Email struct {
	Password string   `yaml:"Password"`
	IsSSL    bool     `yaml:"IsSSL"`
	From     string   `yaml:"From"`
	To       []string `yaml:"To"`
	Host     string   `yaml:"Host"`
	Port     int      `yaml:"Port"`
	UserName string   `yaml:"UserName"`
}

type AliyunOSS struct {
	BasePath        string `yaml:"Base_path"`
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyId     string `yaml:"Access_key_id"`
	AccessKeySecret string `yaml:"Access_key_secret"`
	BucketName      string `yaml:"Bucket_name"`
	BucketUrl       string `yaml:"Bucket_url"`
}

type Server struct {
	RunMode               string        `yaml:"RunMode"`
	Address               string        `yaml:"Address"`
	ReadTimeout           time.Duration `yaml:"ReadTimeout"`
	WriteTimeout          time.Duration `yaml:"WriteTimeout"`
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"`
}

type Rule struct {
	DefaultCoverURL string `yaml:"DefaultCoverURL"`
	UsernameLenMax  int    `yaml:"UsernameLenMax"`
	UsernameLenMin  int    `yaml:"UsernameLenMin"`
	PasswordLenMax  int    `yaml:"PasswordLenMax"`
	PasswordLenMin  int    `yaml:"PasswordLenMin"`
	CommentLenMax   int    `yaml:"CommentLenMax"`
	CommentLenMin   int    `yaml:"CommentLenMin"`
	TitlesLenMax    int    `json:"TitlesLenMax"`
	TitlesLenMin    int    `json:"TitlesLenMin"`
}
