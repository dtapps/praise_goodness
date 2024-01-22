package praise_goodness

import (
	"go.dtapp.net/golog"
)

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}

func (c *Client) ConfigClient(config *ClientConfig) {
	c.config.apiURL = config.ApiURL
	c.config.mchID = config.MchID
	c.config.Key = config.Key
}

// ConfigApiMongoFun 接口日志配置
func (c *Client) ConfigApiMongoFun(apiClientFun golog.ApiMongoFun) {
	client := apiClientFun()
	if client != nil {
		c.mongoLog.client = client
		c.mongoLog.status = true
	}
}
