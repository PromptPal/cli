package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/PromptPal/cli/commands/internal"
	"github.com/go-resty/resty/v2"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigFileName = "promptpal.yml"
)

type PromptSchemaResponse struct {
	Count int                     `json:"count"`
	Data  []internal.PromptSchema `json:"data"`
}

var GenerateCommand *cli.Command = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generate types",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Value:   false,
			Usage:   "force overwrite existing schema file",
		},
	},
	Action: commandGenerate,
}

func commandGenerate(c *cli.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	configFilePath := cwd + "/" + defaultConfigFileName
	_, err = os.Stat(configFilePath)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("config file not found. please run 'promptpal init' first")
		}
		return err
	}

	configBuf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	var config internal.Configuration
	err = yaml.Unmarshal(configBuf, &config)
	if err != nil {
		return err
	}

	client := resty.
		New().
		SetTimeout(10 * time.Second)

	// fetch prompts list
	schema, err := loadSchema(client, config)

	if err != nil {
		return err
	}

	// generate go types
	if config.Output.GoTypes != "" {
		goTypesBuf, err := internal.GenerateGoTypes(schema)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(config.Output.GoTypes, goTypesBuf, 0644)
		if err != nil {
			return err
		}
	}

	// generate typescript types
	if config.Output.TypeScriptTypes != "" {
		typeScriptTypesBuf, err := internal.GenerateTypeScriptTypes(schema)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(config.Output.TypeScriptTypes, typeScriptTypesBuf, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadSchema(client *resty.Client, config internal.Configuration) ([]internal.PromptSchema, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	schemaFilePath := path.Join(cwd, config.Output.Schema)

	// TODO: load from cache

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("API").
		SetAuthToken(config.Input.Http.APIToken).
		SetError(internal.ServerErrorResponse{}).
		SetQueryParam("limit", "100").
		SetQueryParam("cursor", strconv.Itoa(1<<30)).
		SetResult(PromptSchemaResponse{}).
		Get(config.Input.Http.Url + "/api/v1/public/prompts")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		errMsg := resp.Error().(*internal.ServerErrorResponse)
		return nil, fmt.Errorf("error: %d %s", errMsg.ErrorCode, errMsg.ErrorMessage)
	}

	schema, ok := resp.Result().(*PromptSchemaResponse)
	if !ok {
		return nil, errors.New("invalid prompt schema type")
	}

	// TODO:
	// handle count greater than 100
	if schema.Count == 0 {
		return nil, errors.New("no prompts found")
	}

	if config.Output.Schema != "" {
		schemaBuf, err := json.MarshalIndent(schema.Data, "", "  ")
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(schemaFilePath, schemaBuf, 0644)
		if err != nil {
			return nil, err
		}
	}

	return schema.Data, nil
}
