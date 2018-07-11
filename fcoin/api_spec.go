package fcoin

const (
	// API base
	BaseUrl   = "https://api.fcoin.com/v2"
	WSBaseUrl = "wss://api.fcoin.com/v2/ws"

	// Public endpoints
	GetServerTime = "/public/server-time"
	GetCurrencies = "/public/currencies"
	GetSymbols    = "/public/symbols"

	// Account endpoints
	GetBalance = "/accounts/balance"

	// Order endpoints
	OrdersBase = "/orders"

	// API Spec
	MaxTimeDiffMs = 30 * 1000
	NeedSignature = true
	NoSignature   = false
)

type APIResponse struct {
	// Status code, (error) Message
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// Public API
type ServerTimeRsp struct {
	APIResponse
	Data int64 `json:"data"`
}

type CurrenciesRsp struct {
	APIResponse
	Data []string `json:"data"`
}

type SymbolPair struct {
	Name          string `json:"name"`
	BaseCurrenty  string `json:"base_currency"`
	QuoteCurrenty string `json:"quote_currency"`
	PriceDecimal  int    `json:"price_decimal"`
	AmountDecimal int    `json:"amount_decimal"`
}

type SymbolsRsp struct {
	APIResponse
	Data []SymbolPair `json:"data"`
}

// Account API
type CurrencyBalance struct {
	Currency  string `json:"currency"`
	Category  string `json:"category"`
	Available string `json:"available"`
	Frozen    string `json:"frozen"`
	Balance   string `json:"balance"`
}

type AccountsBalanceRsp struct {
	APIResponse
	Data []CurrencyBalance `json:"data"`
}

// Orders API
type CreateOrderArgs struct {
	Amount string `json:"amount"`
	Price  string `json:"price"`
	Type   string `json:"type"`
	Side   string `json:"side"`
	Symbol string `json:"symbol"`
}

type CreateOrderRsp struct {
	APIResponse
	Data string `json:"data"`
}

type GetOrdersArgs struct {
	// after/before certain page
	After  string `json:"After,omitempty"`
	Before string `json:"before,omitempty"`
	Limit  string `json:"limit,omitempty"`
	States string `json:"states,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

type OrderDetail struct {
	ID            string `json:"id"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Amount        string `json:"amount"`
	State         string `json:"state"`
	ExecutedValue string `json:"executed_value"`
	FillFees      string `json:"fill_fees"`
	FilledAmount  string `json:"filled_amount"`
	CreatedAt     int    `json:"created_at"`
	Source        string `json:"source"`
}

type GetOrdersRsp struct {
	APIResponse
	Data []OrderDetail `json:"data"`
}

type GetOrderRsp struct {
	APIResponse
	Data OrderDetail `json:"data"`
}

type SubmitCancelOrderRsp struct {
	APIResponse
	Data bool `json:"data"`
}

// Websocket
type WSHello struct {
	Type string `json:"type"`
	TS   int64  `json:"ts"`
}

type WSArgs struct {
	// Message id, server returns this for accounting
	ID   string        `json:"id,omitempty"`
	Args []interface{} `json:"args,omitempty"`
	Cmd  string        `json:"cmd,omitempty"`
}

type WSPingRsp struct {
	Type     string `json:"type"`
	RemoteTs int64  `json:"ts"`
	Gap      int64  `json:"gap"`
}

type WSSubRsp struct {
	ID     string   `json:"id"`
	Type   string   `json:"type"`
	Topics []string `json:"topics"`
}

// Websocket Message structs for ReadJSON
type WSSymbolList struct {
	// helper for sub "all-tickers", topic is always "all-tickers"
	Topic   string `json:"topic"`
	Tickers []struct {
		Symbol string    `json:"symbol"`
		Ticker []float64 `json:"ticker"`
	} `json:"tickers"`
}

func (l *WSSymbolList) Symbols() []string {
	var r []string
	for _, v := range l.Tickers {
		r = append(r, v.Symbol)
	}
	return r
}

type WSTick struct {
	// ticker.$symbol
	Type   string      `json:"type"`
	Seq    int         `json:"seq"`
	Ticker [11]float64 `json:"ticker"`
}

type WSTickDepth struct {
	// depth.$level.$symbol
	Bids []float64 `json:"bids"`
	Asks []float64 `json:"asks"`
	Ts   int64     `json:"ts"`
	Seq  int       `json:"seq"`
	Type string    `json:"type"`
}

type WSTrade struct {
	// trade.$symbol
	Type   string  `json:"type"`
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Ts     int64   `json:"ts"`
	Side   string  `json:"side"`
	Price  float64 `json:"price"`
}

type WSCandle struct {
	// candle.$resolution.$symbol
	Type     string  `json:"type"`
	ID       int     `json:"id"`
	Seq      int     `json:"seq"`
	Open     float64 `json:"open"`
	Close    float64 `json:"close"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Count    int     `json:"count"`
	BaseVol  float64 `json:"base_vol"`
	QuoteVol float64 `json:"quote_vol"`
}
