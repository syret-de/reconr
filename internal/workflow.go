package internal

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"text/template"
)

type Task struct {
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
}

type Workflow struct {
	Tasks map[int]Task `yaml:"tasks"`
}

type tmpl struct {
	Target string
	Proxy  string
}

func NewWorkflow(file string, config Config) (Workflow, error) {
	input, err := template.New(file).ParseFiles(file)
	if err != nil {
		return Workflow{}, err
	}

	parsed := bytes.Buffer{}
	err = input.Execute(&parsed, tmpl{Target: config.GetTarget(), Proxy: config.GetProxy()})
	if err != nil {
		return Workflow{}, err
	}

	var workflow Workflow
	if err := yaml.Unmarshal(parsed.Bytes(), &workflow); err != nil {
		return Workflow{}, err
	}

	return workflow, nil
}
