package cli

import (
	"github.com/dangeross/breez-lightning-summit-2024/internal/util"
	"github.com/desertbit/grumble"
)

func (c *Cli) load() error {
	c.AddCommand(&grumble.Command{
		Name: "set_api_key",
		Help: "configure the Breez API key",
		Args: func(a *grumble.Args) {
			a.String("api_key", "api key")
		},
		Run: func(ctx *grumble.Context) (err error) {
			apiKey := ctx.Args.String("api_key")

			return c.SetApiKey(apiKey)
		},
	})

	/**
	 * Connection
	 */

	c.AddCommand(&grumble.Command{
		Name: "connect",
		Help: "initialize an SDK instance",
		Flags: func(f *grumble.Flags) {
			f.String("i", "invite_code", "", "optional greenlight invite code")
		},
		Run: func(ctx *grumble.Context) (err error) {
			inviteCode := util.NilString(ctx.Flags.String("invite_code"))

			return c.Connect(inviteCode)
		},
	})

	/**
	 * Node
	 */

	c.AddCommand(&grumble.Command{
		Name: "node_info",
		Help: "get the latest node state",
		Run: func(ctx *grumble.Context) (err error) {
			return c.NodeInfo()
		},
	})

	/**
	 * Lightning
	 */

	c.AddCommand(&grumble.Command{
		Name: "receive_payment",
		Help: "generate a bolt11 invoice",
		Args: func(a *grumble.Args) {
			a.Uint64("amount_msat", "amount to receive in millisatoshis")
			a.String("description", "payment description", grumble.Default(""))
		},
		Run: func(ctx *grumble.Context) (err error) {
			amountMsat := ctx.Args.Uint64("amount_msat")
			description := ctx.Args.String("description")

			return c.ReceivePayment(amountMsat, description)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "send_payment",
		Help: "send a lightning payment",
		Args: func(a *grumble.Args) {
			a.String("bolt11", "bolt11 lightning invoice")
			a.Uint64("amount_msat", "amount to send in millisatoshis", grumble.Default(uint64(0)))
		},
		Run: func(ctx *grumble.Context) (err error) {
			bolt11 := ctx.Args.String("bolt11")
			amountMsat := util.NilUint64(ctx.Args.Uint64("amount_msat"))

			return c.SendPayment(bolt11, amountMsat)
		},
	})

	return nil
}
