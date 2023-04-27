package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	// request := LitleOnlineRequest{
	// 	Version:      "11.4",
	// 	XMLNamespace: "http://www.litle.com/schema",
	// 	MerchantID:   "100",
	// 	Authentication: Authentication{
	// 		User:     "User Name",
	// 		Password: "password",
	// 	},
	// 	Sale: &Sale{
	// 		ID:          "1",
	// 		ReportGroup: "ABC Division",
	// 		CustomerID:  "038945",
	// 		OrderID:     "5234234",
	// 		Amount:      40000,
	// 		OrderSource: "3dsAuthenticated",
	// 		BillToAddress: &Address{
	// 			Name:         "John Smith",
	// 			AddressLine1: "100 Main St",
	// 			AddressLine2: "100 Main St",
	// 			AddressLine3: "100 Main St",
	// 			City:         "Boston",
	// 			State:        "MA",
	// 			Zip:          "12345",
	// 			Country:      "US",
	// 			Email:        "jsmith@someaddress.com",
	// 			Phone:        "555-123-4567",
	// 		},
	// 		Card: &Card{
	// 			Type:              "VI",
	// 			Number:            "4005550000081019",
	// 			ExpDate:           "1210",
	// 			CardValidationNum: "555",
	// 		},
	// 	},
	// }

	request := LitleOnlineRequest{
		Version:      "11.4",
		XMLNamespace: "http://www.litle.com/schema",
		MerchantID:   "100",
		Authentication: Authentication{
			User:     "User Name",
			Password: "password",
		},
		Authorization: &Authorization{
			ID:          "834262",
			ReportGroup: "ABC Division",
			CustomerID:  "038945",
			OrderID:     "65347567",
			Amount:      40000,
			OrderSource: "3dsAuthenticated",
			BillToAddress: Address{
				Name:         "John Smith",
				AddressLine1: "100 Main St",
				City:         "Boston",
				Country:      "USA",
				State:        "MA",
				Zip:          "12345",
				Email:        "jsmith@someaddress.com",
				Phone:        "555-123-4567",
			},
			Card: Card{
				Type:              "VI",
				Number:            "4000000000000002",
				ExpDate:           "1209",
				CardValidationNum: "555",
			},
		},
	}

	c, _ := NewClient("https://www.testvantivcnp.com/sandbox/communicator/online")

	c.SetLog(os.Stdout)
	ctx := context.Background()
	r, _ := c.CreateOnlineSale(ctx, request)

	// if e != nil {
	// 	fmt.Println(e)
	// }

	fmt.Println(r)

	r.FooBar()

}
