package entities

type Client struct {
	name         ClientName
	governmentId ClientGovernmentId
}

func (c *Client) Name() string {
	return c.name.value
}

func (c *Client) GovernmentId() string {
	return c.governmentId.value
}

func NewClient(name ClientName, governmentId ClientGovernmentId) (*Client, error) {
	return &Client{name, governmentId}, nil
}
