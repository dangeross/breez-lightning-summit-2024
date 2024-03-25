package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
	qrcode "github.com/mdp/qrterminal/v3"
)

func (c *Cli) ReceivePayment(amountMsat uint64, description string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	request := breez_sdk.ReceivePaymentRequest{
		AmountMsat:  amountMsat,
		Description: description,
	}

	response, err := c.sdk.ReceivePayment(request)

	if err != nil {
		return err
	}

	qrcode.GenerateHalfBlock(response.LnInvoice.Bolt11, qrcode.L, c.App)
	return nil
}
