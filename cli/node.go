package cli

import (
	"fmt"
)

func (c *Cli) NodeInfo() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	nodeState, err := c.sdk.NodeInfo()
	if err != nil {
		return err
	}

	c.PrettyPrint(nodeState)
	return nil
}
