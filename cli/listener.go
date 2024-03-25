package cli

import (
	"github.com/breez/breez-sdk-go/breez_sdk"
	"github.com/sirupsen/logrus"
)

func (c *Cli) OnEvent(breezEvent breez_sdk.BreezEvent) {
	c.log.Infof("Received Breez event: %#v", breezEvent)

	switch event := breezEvent.(type) {
	case breez_sdk.BreezEventInvoicePaid:
		c.PrintSuccess("Payment received")
		if event.Details.Payment != nil {
			c.PrettyPrint(event.Details.Payment)
		}
	}
}

func (c *Cli) Log(logEntry breez_sdk.LogEntry) {
	level, err := logrus.ParseLevel(logEntry.Level)
	if err != nil {
		level = logrus.DebugLevel
	}

	c.log.Log(level, logEntry.Line)
}
