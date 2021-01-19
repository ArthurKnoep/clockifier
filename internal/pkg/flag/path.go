package flag

import (
	"os"
	"strings"
)

type CfgPath string

func (c *CfgPath) Set(value string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	*c = CfgPath(strings.Replace(value, "~", homeDir, 1))
	return nil
}

func (c CfgPath) String() string {
	return string(c)
}
