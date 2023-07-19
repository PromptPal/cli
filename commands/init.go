package commands

import (
	"errors"
	"os"

	"github.com/PromptPal/cli/commands/internal"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var InitCommand *cli.Command = &cli.Command{
	Name:  "init",
	Usage: "init your PromptPal config",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Value:   false,
			Usage:   "force overwrite existing config file",
		},
	},
	Action: commandInit,
}

func commandInit(c *cli.Context) error {
	isForce := c.Bool("force")

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	configFilePath := cwd + "/" + defaultConfigFileName
	_, err = os.Stat(configFilePath)

	// the file exist
	if err == nil {
		if !isForce {
			return errors.New(" ⛔ config file already exists")
		}
	}

	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		if err != nil {
			return err
		}
	}

	defaultConfig := internal.Configuration{
		Input: internal.ConfigurationInput{
			Http: internal.ConfigurationInputHttp{
				Url:      "http://localhost:8080",
				APIToken: "@env.PROMPTPAL_API_TOKEN",
			},
		},
		Output: internal.ConfigurationOutput{
			Schema: "./schema.g.json",
			GoTypes: &internal.ConfigurationOutputGo{
				Prefix:      "PPExample",
				PackageName: "main",
				Output:      "./example/types.g.go",
			},
			TypeScriptTypes: &internal.ConfigurationOutputTS{
				Prefix: "PPExample",
				Output: "./example/types.g.ts",
			},
		},
	}

	configBuf, err := yaml.Marshal(defaultConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, configBuf, 0644)
	if err != nil {
		return err
	}

	color.Green.Println(" ✅ config file created")
	return nil
}
