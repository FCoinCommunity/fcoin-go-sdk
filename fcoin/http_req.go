package fcoin

// Public API call
func (c *Client) ServerTime() (ServerTimeRsp, error) {
	var rsp ServerTimeRsp
	err := c.request(GetServerTime, "GET", NoSignature, nil, &rsp)
	return rsp, err
}

func (c *Client) Currencies() (CurrenciesRsp, error) {
	var rsp CurrenciesRsp
	err := c.request(GetCurrencies, "GET", NoSignature, nil, &rsp)
	return rsp, err
}

func (c *Client) Symbols() (SymbolsRsp, error) {
	var rsp SymbolsRsp
	err := c.request(GetSymbols, "GET", NoSignature, nil, &rsp)
	return rsp, err
}

// Account API call
func (c *Client) AccountsBalance() (AccountsBalanceRsp, error) {
	var rsp AccountsBalanceRsp
	err := c.request(GetBalance, "GET", NeedSignature, nil, &rsp)
	return rsp, err
}

// Order API call
func (c *Client) CreateOrder(args *CreateOrderArgs) (CreateOrderRsp, error) {
	var rsp CreateOrderRsp
	err := c.request(OrdersBase, "POST", NeedSignature, args, &rsp)
	return rsp, err
}

func (c *Client) GetOrders(args *GetOrdersArgs) (GetOrdersRsp, error) {
	var rsp GetOrdersRsp
	err := c.request(OrdersBase, "GET", NeedSignature, args, &rsp)
	return rsp, err
}

func (c *Client) GetOrder(orderId string) (GetOrderRsp, error) {
	var rsp GetOrderRsp
	url := OrdersBase + "/" + orderId
	err := c.request(url, "GET", NeedSignature, nil, &rsp)
	return rsp, err
}

func (c *Client) SubmitCancelOrder(orderId string) (SubmitCancelOrderRsp, error) {
	var rsp SubmitCancelOrderRsp
	url := OrdersBase + "/" + orderId + "/submit-cancel"
	err := c.request(url, "POST", NeedSignature, nil, &rsp)
	return rsp, err
}
