package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/qingconglaixueit/wechatbot/pkg/logger"
)

// Configuration 项目配置
type Configuration struct {
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
	// 会话超时时间
	SessionTimeout time.Duration `json:"session_timeout"`
	// GPT模型
	Model string `json:"model"`
	// 回复前缀
	ReplyPrefix string `json:"reply_prefix"`
	// 超时时间
	TimeOut time.Duration `json:"timeout"`
	// url
	URL string `json:"url"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			AutoPass: false,
			Model:    "forefront",
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				logger.Danger(fmt.Sprintf("open config error: %v", err))
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				logger.Danger(fmt.Sprintf("decode config error: %v", err))
				return
			}
		}
		// 有环境变量使用环境变量
		AutoPass := os.Getenv("AUTO_PASS")
		SessionTimeout := os.Getenv("SESSION_TIMEOUT")
		Model := os.Getenv("MODEL")
		ReplyPrefix := os.Getenv("REPLY_PREFIX")
		TimeOut := os.Getenv("TIMEOUT")
		URL := os.Getenv("URL")
		if AutoPass == "true" {
			config.AutoPass = true
		}
		if SessionTimeout != "" {
			duration, err := time.ParseDuration(SessionTimeout)
			if err != nil {
				logger.Danger(fmt.Sprintf("config session timeout error: %v, get is %v", err, SessionTimeout))
				return
			}
			config.SessionTimeout = duration
		}
		if Model != "" {
			config.Model = Model
		}
		if ReplyPrefix != "" {
			config.ReplyPrefix = ReplyPrefix
		}
		if TimeOut != "" {
			duration, err := time.ParseDuration(TimeOut)
			if err != nil {
				logger.Danger(fmt.Sprintf("config session timeout error: %v, get is %v", err, TimeOut))
				return
			}
			config.TimeOut = duration
		}
		if URL != "" {
			config.URL = URL
		}
	})

	return config
}
