package cli

import (
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

	return nil
}
