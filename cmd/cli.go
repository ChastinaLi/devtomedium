package cmd

import (
	"fmt"
	"log"
	"os"

	"devtomedium/pkg"

	"github.com/urfave/cli"
)

func Run() {
	app := cli.NewApp()
	app.Name = "devtomedium"
	app.Usage = "Upload dev.to posts to Medium"
	app.Action = func(c *cli.Context) error {
		url := "https://dev.to/li_chastina/auto-refresh-aws-tokens-using-iam-role-and-boto3-2cjf"
		result := pkg.Download_markdown(url)
		fmt.Println(result)
		fmt.Printf("Hello %s", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
