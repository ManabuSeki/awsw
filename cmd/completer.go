package awsw

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"gopkg.in/yaml.v2"
)

type Data struct {
	Projects []Project `yaml:"projects`
}

type Project struct {
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	RoleName    string        `yaml:"role_name"`
	Envs        []Environment `yaml:"environments"`
}

type Environment struct {
	Name string `yaml:"displayName"`
	Env  string `yaml:"env"`
	ID   string `yaml:"accound_id"`
}

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := strings.Split(d.TextBeforeCursor(), " ")

	return argumentsCompleter(excludeOptions(args))
}

func argumentsCompleter(args []string) []prompt.Suggest {
	d, err := LoadYamlFile()
	if err != nil {
		panic(err)
	}

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(createFirstSuggest(d.Projects), args[0], true)
	}

	project := Project{}

	for _, p := range d.Projects {
		if p.Name == args[0] {
			project = p
			break
		}
	}
	second := args[1]
	if len(args) == 2 {
		return prompt.FilterHasPrefix(createSecondSuggest(project), second, true)
	}

	third := args[2]
	if len(args) == 3 {
		subcommands := []prompt.Suggest{
			prompt.Suggest{Text: "red", Description: ""},
			prompt.Suggest{Text: "orange", Description: ""},
			prompt.Suggest{Text: "yellow", Description: ""},
			prompt.Suggest{Text: "green", Description: ""},
			prompt.Suggest{Text: "blue", Description: ""},
		}
		return prompt.FilterHasPrefix(subcommands, third, true)
	}
	return []prompt.Suggest{}
}

func createFirstSuggest(p []Project) []prompt.Suggest {
	suggests := make([]prompt.Suggest, len(p))

	for i, project := range p {
		suggests[i] = prompt.Suggest{Text: project.Name, Description: project.Description}
	}
	return suggests
}

func createSecondSuggest(p Project) []prompt.Suggest {
	suggests := make([]prompt.Suggest, len(p.Envs))

	for i, env := range p.Envs {
		suggests[i] = prompt.Suggest{Text: env.Env, Description: ""}
	}
	return suggests
}

func LoadYamlFile() (Data, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filePath := homePath + "/.aws/account.yaml"

	if _, err := os.Stat(filePath); err != nil {
		fmt.Println("No such file: " + filePath)
	}

	// buf, err := ioutil.ReadFile("account.yaml")
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("buf: %+v\n", string(buf))

	var d Data
	err = yaml.Unmarshal(buf, &d)

	if err != nil {
		return Data{}, err
	}
	return d, nil
}
