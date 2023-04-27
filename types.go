package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type (
	Client struct {
		ApiBase string
		Client  *http.Client
		Log     io.Writer
		mu      sync.Mutex
	}

	LitleOnlineRequest struct {
		XMLName        xml.Name        `xml:"litleOnlineRequest"`
		Version        string          `xml:"version,attr"`
		XMLNamespace   string          `xml:"xmlns,attr"`
		MerchantID     string          `xml:"merchantId,attr"`
		Authentication Authentication  `xml:"authentication"`
		Authorization  *Authorization  `xml:"authorization"`
		Capture        *Capture        `xml:"capture"`
		Sale           *Sale           `xml:"sale"`
		RefundReversal *RefundReversal `xml:"refundReversal"`
		Void           *Void           `xml:"void"`
	}

	LitleOnlineResponse struct {
		XMLName                xml.Name                `xml:"litleOnlineResponse"`
		Version                string                  `xml:"version,attr"`
		XMLNS                  string                  `xml:"xmlns,attr"`
		Response               string                  `xml:"response,attr"`
		Message                string                  `xml:"message,attr"`
		AuthorizationResponse  *AuthorizationResponse  `xml:"authorizationResponse"`
		CaptureResponse        *CaptureResponse        `xml:"captureResponse"`
		SaleResponse           *SaleResponse           `xml:"saleResponse"`
		RefundReversalResponse *RefundReversalResponse `xml:"refundReversalResponse"`
		VoidResponse           *VoidResponse           `xml:"voidResponse"`
	}

	Authentication struct {
		User     string `xml:"user"`
		Password string `xml:"password"`
	}

	Authorization struct {
		XMLName                  xml.Name                  `xml:"authorization"`
		ID                       string                    `xml:"id,attr"`
		ReportGroup              string                    `xml:"reportGroup,attr"`
		CustomerID               string                    `xml:"customerId,attr"`
		OrderID                  string                    `xml:"orderId"`
		Amount                   int                       `xml:"amount"`
		OrderSource              string                    `xml:"orderSource"`
		BillToAddress            Address                   `xml:"billToAddress"`
		Card                     Card                      `xml:"card"`
		CardholderAuthentication *CardholderAuthentication `xml:"cardholderAuthentication"`
	}

	Capture struct {
		XMLName      xml.Name      `xml:"capture"`
		ID           string        `xml:"id,attr"`
		ReportGroup  string        `xml:"reportGroup,attr"`
		CustomerID   string        `xml:"customerId,attr"`
		Partial      bool          `xml:"partial,attr"`
		CnpTxnID     string        `xml:"cnpTxnId"`
		EnhancedData *EnhancedData `xml:"enhancedData"`
	}

	RefundReversal struct {
		XMLName                xml.Name `xml:"refundReversal"`
		ID                     string   `xml:"id,attr"`
		ReportGroup            string   `xml:"reportGroup,attr"`
		CustomerID             string   `xml:"customerId,attr"`
		LitleTxnID             string   `xml:"litleTxnId"`
		Card                   Card     `xml:"card"`
		OriginalRefCode        string   `xml:"originalRefCode"`
		OriginalAmount         float64  `xml:"originalAmount"`
		OriginalTxnTime        string   `xml:"originalTxnTime"`
		OriginalSystemTraceID  string   `xml:"originalSystemTraceId"`
		OriginalSequenceNumber string   `xml:"originalSequenceNumber"`
	}

	Sale struct {
		XMLName                  xml.Name                  `xml:"sale"`
		ID                       string                    `xml:"id,attr"`
		ReportGroup              string                    `xml:"reportGroup,attr"`
		CustomerID               string                    `xml:"customerId,attr"`
		OrderID                  string                    `xml:"orderId"`
		Amount                   int                       `xml:"amount"`
		OrderSource              string                    `xml:"orderSource"`
		BillToAddress            *Address                  `xml:"billToAddress"`
		Card                     *Card                     `xml:"card"`
		CardholderAuthentication *CardholderAuthentication `xml:"cardholderAuthentication"`
		CustomBilling            *CustomBilling            `xml:"customBilling"`
		EnhancedData             *EnhancedData             `xml:"enhancedData"`
	}

	Void struct {
		XMLName     xml.Name `xml:"void"`
		ID          string   `xml:"id,attr"`
		ReportGroup string   `xml:"reportGroup,attr"`
		LitleTxnID  string   `xml:"litleTxnId"`
	}

	Address struct {
		Name         string `xml:"name"`
		AddressLine1 string `xml:"addressLine1"`
		AddressLine2 string `xml:"addressLine2"`
		AddressLine3 string `xml:"addressLine3"`
		City         string `xml:"city"`
		State        string `xml:"state"`
		Zip          string `xml:"zip"`
		Country      string `xml:"country"`
		Email        string `xml:"email"`
		Phone        string `xml:"phone"`
	}

	Card struct {
		Type              string `xml:"type"`
		Number            string `xml:"number"`
		ExpDate           string `xml:"expDate"`
		CardValidationNum string `xml:"cardValidationNum"`
	}

	CardholderAuthentication struct {
		AuthenticationValue         string `xml:"authenticationValue"`
		AuthenticationTransactionID string `xml:"authenticationTransactionId"`
	}

	CustomBilling struct {
		Phone      string `xml:"phone"`
		Descriptor string `xml:"descriptor"`
	}

	EnhancedData struct {
		CustomerReference      string         `xml:"customerReference"`
		SalesTax               int            `xml:"salesTax"`
		TaxExempt              bool           `xml:"taxExempt"`
		DiscountAmount         int            `xml:"discountAmount"`
		ShippingAmount         int            `xml:"shippingAmount"`
		DutyAmount             int            `xml:"dutyAmount"`
		ShipFromPostalCode     string         `xml:"shipFromPostalCode"`
		DestinationPostalCode  string         `xml:"destinationPostalCode"`
		DestinationCountryCode string         `xml:"destinationCountryCode"`
		InvoiceReferenceNumber string         `xml:"invoiceReferenceNumber"`
		OrderDate              string         `xml:"orderDate"`
		DetailTax              DetailTax      `xml:"detailTax"`
		LineItemData           []LineItemData `xml:"lineItemData"`
	}

	DetailTax struct {
		TaxIncludedInTotal bool   `xml:"taxIncludedInTotal"`
		TaxAmount          int    `xml:"taxAmount"`
		TaxRate            string `xml:"taxRate"`
		TaxTypeIdentifier  string `xml:"taxTypeIdentifier"`
		CardAcceptorTaxID  string `xml:"cardAcceptorTaxId"`
	}

	LineItemData struct {
		ItemSequenceNumber   int       `xml:"itemSequenceNumber"`
		ItemDescription      string    `xml:"itemDescription"`
		ProductCode          string    `xml:"productCode"`
		Quantity             int       `xml:"quantity"`
		UnitOfMeasure        string    `xml:"unitOfMeasure"`
		TaxAmount            int       `xml:"taxAmount"`
		LineItemTotal        int       `xml:"lineItemTotal"`
		LineItemTotalWithTax int       `xml:"lineItemTotalWithTax"`
		ItemDiscountAmount   int       `xml:"itemDiscountAmount"`
		CommodityCode        string    `xml:"commodityCode"`
		UnitCost             float64   `xml:"unitCost"`
		DetailTax            DetailTax `xml:"detailTax"`
	}

	AuthorizationResponse struct {
		XMLName              xml.Name `xml:"authorizationResponse"`
		ID                   string   `xml:"id,attr"`
		ReportGroup          string   `xml:"reportGroup,attr"`
		CustomerID           string   `xml:"customerId,attr"`
		LitleTxnID           string   `xml:"litleTxnId"`
		OrderID              string   `xml:"orderId"`
		Response             string   `xml:"response"`
		ResponseTime         string   `xml:"responseTime"`
		PostDate             string   `xml:"postDate"`
		Message              string   `xml:"message"`
		AuthCode             string   `xml:"authCode"`
		ApprovedAmount       string   `xml:"approvedAmount"`
		NetworkTransactionID string   `xml:"networkTransactionId"`
	}

	CaptureResponse struct {
		XMLName      xml.Name `xml:"captureResponse"`
		ID           string   `xml:"id,attr"`
		ReportGroup  string   `xml:"reportGroup,attr"`
		CustomerID   string   `xml:"customerId,attr"`
		CnpTxnID     string   `xml:"cnpTxnId"`
		Response     string   `xml:"response"`
		ResponseTime string   `xml:"responseTime"`
		PostDate     string   `xml:"postDate"`
		Message      string   `xml:"message"`
	}

	SaleResponse struct {
		XMLName      xml.Name     `xml:"saleResponse"`
		ID           string       `xml:"id,attr"`
		ReportGroup  string       `xml:"reportGroup,attr"`
		CustomerID   string       `xml:"customerId,attr"`
		CnpTxnID     string       `xml:"cnpTxnId"`
		Response     string       `xml:"response"`
		OrderID      string       `xml:"orderId"`
		ResponseTime string       `xml:"responseTime"`
		PostDate     string       `xml:"postDate"`
		Message      string       `xml:"message"`
		AuthCode     string       `xml:"authCode"`
		FraudResult  *FraudResult `xml:"fraudResult"`
	}

	RefundReversalResponse struct {
		XMLName      xml.Name `xml:"refundReversalResponse"`
		ID           string   `xml:"id,attr"`
		ReportGroup  string   `xml:"reportGroup,attr"`
		CustomerID   string   `xml:"customerId,attr"`
		LitleTxnID   string   `xml:"litleTxnId"`
		Response     string   `xml:"response"`
		ResponseTime string   `xml:"responseTime"`
		PostDate     string   `xml:"postDate"`
		Message      string   `xml:"message"`
	}

	VoidResponse struct {
		XMLName      xml.Name `xml:"voidResponse"`
		ID           string   `xml:"id,attr"`
		ReportGroup  string   `xml:"reportGroup,attr"`
		LitleTxnID   string   `xml:"litleTxnId"`
		Response     string   `xml:"response"`
		ResponseTime string   `xml:"responseTime"`
		PostDate     string   `xml:"postDate"`
		Message      string   `xml:"message"`
	}

	FraudResult struct {
		AVSResult            string `xml:"avsResult"`
		CardValidationResult string `xml:"cardValidationResult"`
		AuthenticationResult string `xml:"authenticationResult"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *LitleOnlineResponse) FooBar() {
	fmt.Println("TEST THIS IS WORKING SWEET")
}
