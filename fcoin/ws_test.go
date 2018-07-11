package fcoin

import (
	"testing"
	"time"
)

func TestWebsocket(t *testing.T) {
	api := Client{}
	if err := api.InitWS(); err != nil {
		t.Fatal(err)
	}

	rsp, err := api.WSPing()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ping from websocket: %v", rsp)

	if err := api.WSSubscribe("some-test-id", "depth.L20.ethusdt"); err != nil {
		t.Fatal(err)
	}

	// Receiving data for a while
	quit := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		quit <- 1
	}()

	for {
		select {
		case <-quit:
			return
		default:
			_, rsp2, _ := api.WS.ReadMessage()
			t.Log(string(rsp2))
		}
	}

}
