# fcoin-go-sdk
FCoin Go SDK

## Dependencies

| Name | Url | HEAD tested |
| ---- | --- | ----------- |
| websocket | https://github.com/gorilla/websocket | 5ed622c449da6d44c3c8329331ff47a9e5844f71 |

## Usage

```
// Auth
t := time.Now().Unix() * 1000  // Unix time(milliseconds)
timeStr := strconv.FormatInt(t, 10)
api := fcoin.Authorize("key", "secret", timeStr)

// Usage
rsp, err := api.AccountsBalance()
if err != nil || rsp.Status != 0 {
    // non-zero indicates API call failed
    fmt.Printf(rsp.Msg)
}
fmt.Printf(rsp.Data)

// API with args
args := &CreateOrderArgs{
	Symbol: symbol,
	Side:   side,
	Type:   t,
	Price:  price,
	Amount: amount,
}
rsp, err := api.GetOrders(args)
```

## Websocket Usage

Example:

```
// Initialize websocket client(market data only)
api := fcoin.Client{}
if err := api.InitWS(); err != nil {
    //
}
rsp, err := api.WSPing()

// Subscribe to topic and receive(see ws_test.go)
api.WSSubscribe("message-id", "topic1", "topic2")  // or more topics

// message-id can be a empty string
// api.WSSubscribe("", "topic3")

// subscription result will contain same "message-id" if the field is not empty
var rsp2 WSSubRsp
if err := api.ws.ReadJSON(&rsp2) {
    // Suppose subscription result returns faster than data
    rsp2.ID == "message-id"
}

```

**Caution**: websocket provides bi-directional communication, `Ping()` in example does not work when you are subscribing to topics because next message would not necessarily be Ping-Response from server.

**Recommended**: subscribe to expected topics, pass client to your own message dispatcher and enjoy data.


