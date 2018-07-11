package fcoin

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	skipTestErr = fmt.Errorf("Missing FC_ACCESS_KEY/FC_ACCESS_SECRET environment args")
)

func getTestClient() (*Client, error) {
	fkey, ok := os.LookupEnv("FC_ACCESS_KEY")
	if !ok {
		return nil, skipTestErr
	}
	fsec, ok := os.LookupEnv("FC_ACCESS_SECRET")
	if !ok {
		return nil, skipTestErr
	}

	// unix ms
	time := time.Now().Unix() * 1000
	api, err := Authorize(fkey, fsec, time)
	if err != nil {
		return nil, err
	}
	return api, nil
}

func TestPublicAPI(t *testing.T) {
	now := time.Now().Unix() * 1000
	api, err := Authorize("invalid-key", "empty-seret", now)
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := api.ServerTime()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Server time %v", rsp.Data)

	rsp2, err := api.Currencies()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Currencies list %v", rsp2.Data)

	rsp3, err := api.Symbols()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Symbols list %v", rsp3.Data)
}

func TestClient(t *testing.T) {
	api, err := getTestClient()
	if err == skipTestErr {
		t.Log(err)
		return
	} else if err != nil {
		t.Fatal(err)
	}

	rsp, err := api.AccountsBalance()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got Account balance %v", rsp.Data)

	orderArgs := &CreateOrderArgs{
		Symbol: "btcusdt",
		Side:   "buy",
		Type:   "limit",
		Price:  "1.0",
		Amount: "1.0",
	}
	rsp1, err := api.CreateOrder(orderArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Created Order %v", rsp1.Data)

	arg := &GetOrdersArgs{
		After: "0",
	}
	rsp2, err := api.GetOrders(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got Orders %v", rsp2.Data)

	rsp3, err := api.GetOrder("1123")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got Order %v", rsp3.Data)

	rsp4, err := api.SubmitCancelOrder("1123")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Cancel order result: %v", rsp4.Data)
}
