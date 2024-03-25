package util

import (
	"fmt"
	"strings"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func AsPaymentTypeFilter(filter string) (breez_sdk.PaymentTypeFilter, error) {
	switch strings.ToLower(filter) {
	case "closedchannel":
		return breez_sdk.PaymentTypeFilterClosedChannel, nil
	case "sent":
		return breez_sdk.PaymentTypeFilterSent, nil
	case "received":
		return breez_sdk.PaymentTypeFilterReceived, nil
	}
	return breez_sdk.PaymentTypeFilterSent, fmt.Errorf("invalid PaymentTypeFilter")
}

func AsPaymentTypeFilterList(filters *[]string) (*[]breez_sdk.PaymentTypeFilter, error) {
	if filters != nil {
		list := []breez_sdk.PaymentTypeFilter{}

		for _, f := range *filters {
			if len(f) > 0 {
				filter, err := AsPaymentTypeFilter(f)
				if err != nil {
					return nil, err
				}
				list = append(list, filter)
			}
		}

		if len(list) == 0 {
			return nil, nil
		}

		return &list, nil
	}

	return nil, nil
}

func AsProvider(str string) (breez_sdk.BuyBitcoinProvider, error) {
	switch strings.ToLower(str) {
	case "moonpay":
		return breez_sdk.BuyBitcoinProviderMoonpay, nil
	}

	return 0, fmt.Errorf("invalid provider")
}
