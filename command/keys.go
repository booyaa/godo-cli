package command

import (
	"flag"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"strings"
)

type KeysCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *KeysCommand) Help() string {
	helpText := `
Usage: godo-cli keys [options]
Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *KeysCommand) List(args []string) int {
	opt := &godo.ListOptions{}
	keys, _, err := c.Client.Keys.List(opt)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, key := range keys {
		c.Ui.Output(fmt.Sprintf("id:%d name:%s", key.ID, key.Name))
	}

	return 0
}

func (c *KeysCommand) Create(args []string) int {

	var name = ""
	var publicKeyPath = ""

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.StringVar(&name, "name", "", "")
	cmdFlags.StringVar(&publicKeyPath, "public_key", "", "")

	err := cmdFlags.Parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	if name == "" && publicKeyPath == "" {
		c.Help()
		return -1
	}

	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to ReadFile %v", err))
		return -1
	}

	opt := &godo.KeyCreateRequest{
		Name:      name,
		PublicKey: string(publicKey),
	}

	key, _, err := c.Client.Keys.Create(opt)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to Create %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("id:%d name:%s", key.ID, key.Name))

	return 0
}

func (c *KeysCommand) Run(args []string) int {

	if len(args) < 1 {
		c.Ui.Output(c.Help())
		return 0
	}

	newArgs := args[1:]
	switch args[0] {
	case "list":
		return c.List(newArgs)
	case "create":
		return c.Create(newArgs)
	case "delete":
		//return c.Get(newArgs)
	}

	return 0
}

func (c *KeysCommand) Synopsis() string {
	return fmt.Sprintf("Show available SSH keys")
}
