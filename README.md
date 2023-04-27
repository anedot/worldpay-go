# worldpay-cnp

> Go interface to the [WorldPay cnpAPI](http://support.worldpay.com/support/CNP-API/content/introduction.htm)

## Usage
```go
import (
    "context"
    "os"

    "github.com/anedot/worldpay-cnp"
)

func main() {
	request := LitleOnlineRequest{
		Version:      "11.4",
		XMLNamespace: "http://www.litle.com/schema",
		MerchantID:   "100",
		Authentication: Authentication{
			User:     "User Name",
			Password: "password",
		},
		Sale: Sale{
			ID:          "1",
			ReportGroup: "ABC Division",
			CustomerID:  "038945",
			OrderID:     "5234234",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: &Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				AddressLine2: "100 Main St",
				AddressLine3: "100 Main St",
				City:         "Boston",
				State:        "MA",
				Zip:          "12345",
				Country:      "US",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: &Card{
				Type:              "VI",
				Number:            "4005550000081019",
				ExpDate:           "1210",
				CardValidationNum: "555",
			},
		},
	}

	c, _ := NewClient("https://www.testvantivcnp.com/sandbox/communicator/online")

	ctx := context.Background()
	r, _ := c.CreateOnlineSale(ctx, request)
}
```

## Online Transactions
### Authorization

### Capture

### Refund Reversal

### Sale

### Void

