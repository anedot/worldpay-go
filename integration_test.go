package worldpay

import (
	"context"
	"testing"
)

var sandboxUrl = "https://www.testvantivcnp.com/sandbox/communicator/online"

func TestCreateOnlineSale(t *testing.T) {
	c, _ := NewClient(sandboxUrl)

	c.CreateOnlineSale(context.Background(), request * )
}
