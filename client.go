package praise_goodness

import (
	"errors"
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string // 接口地址
	MchID  int64
	Key    string
}

// Client 实例
type Client struct {
	config struct {
		apiURL string
		mchID  int64
		Key    string
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.apiURL = config.ApiURL
	c.config.mchID = config.MchID
	c.config.Key = config.Key

	if c.config.apiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	return c, nil
}
