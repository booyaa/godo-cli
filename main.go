package main

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/command"
	"github.com/mitchellh/cli"
	//"log"
	"os"
)

var GitCommit string

const ApplicationName = "godo-cli"
const Version = "0.0.1"
const VersionPrerelease = ""

func getClinet(accessToken string) *godo.Client {

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}
	return godo.NewClient(t.Client())

}

func main() {

	ui := &cli.BasicUi{Writer: os.Stdout}
	godoCli := getClinet(os.Getenv("DIGITALOCEAN_API_TOKEN"))

	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Ui:      ui,
				AppName: ApplicationName,
				Version: Version,
			}, nil
		},
		"sizes": func() (cli.Command, error) {
			return &command.SizesCommand{
				Ui:  ui,
				Cli: godoCli,
			}, nil
		},
		"regions": func() (cli.Command, error) {
			return &command.RegionsCommand{
				Ui:  ui,
				Cli: godoCli,
			}, nil
		},
		"images": func() (cli.Command, error) {
			return &command.ImagesCommand{
				Ui:  ui,
				Cli: godoCli,
			}, nil
		},
		//"bar": barCommandFactory,
	}

	exitStatus, _ := c.Run()
	/*
		if err != nil {
			log.Println(exitStatus, err, "ssss")
		}
	*/
	os.Exit(exitStatus)
}
