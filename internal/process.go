package internal

import (
	"fmt"
	"sort"
	"strings"
)

func Process(workflow Workflow, config Config, logger Logger) error {
	docker, err := NewDocker(config)
	if err != nil {
		return err
	}
	err = docker.Build(".", logger)

	id := make([]int, 0)
	for k, _ := range workflow.Tasks {
		id = append(id, k)
	}
	sort.Ints(id)

	for _, k := range id {
		command := strings.Join(workflow.Tasks[k].Commands, ";")
		fmt.Println("===", workflow.Tasks[k].Name, "===")
		fmt.Println(command)
		fmt.Println()
		err = docker.Run(command, k)
		if err != nil {
			return err
		}
	}
	return nil
}
