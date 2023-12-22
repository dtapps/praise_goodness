package praise_goodness

func (c *Client) SetApiURL(apiURL string) {
	c.config.apiURL = apiURL
}

func (c *Client) SetMchID(mchID int64) {
	c.config.mchID = mchID
}

func (c *Client) SetAppKey(key string) {
	c.config.Key = key
}
