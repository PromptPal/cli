package internal

import "time"

type ConfigurationInputHttp struct {
	Url      string `yaml:"url"`
	APIToken string `yaml:"token"`
}

type ConfigurationInput struct {
	Http ConfigurationInputHttp `yaml:"http"`
}

type ConfigurationOutputGo struct {
	Prefix      string `yaml:"prefix"`
	PackageName string `yaml:"package_name"`
	Output      string `yaml:"output"`
}

type ConfigurationOutputTS struct {
	Prefix string `yaml:"prefix"`
	Output string `yaml:"output"`
}

type ConfigurationOutput struct {
	Schema          string                 `yaml:"schema"`
	GoTypes         *ConfigurationOutputGo `yaml:"go_types"`
	TypeScriptTypes *ConfigurationOutputTS `yaml:"typescript_types"`
}

type Configuration struct {
	Input  ConfigurationInput  `yaml:"input"`
	Output ConfigurationOutput `yaml:"output"`
}

// MARK: Server

type ServerErrorResponse struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"error"`
}

// MARK: PromptSchema

type PromptVariable struct {
	Name string `json:"name"`
	// string, number, bool
	Type string `json:"type"`
}

type PromptSchema struct {
	HashID      string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Variables   []PromptVariable `json:"variables"`
	TokenCount  int              `json:"tokenCount"`
	CreatedAt   time.Time        `json:"createdAt"`
}
