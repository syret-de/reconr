package main

import (
	"fmt"
	"log"
	"reconr/internal"
)

func main() {
	workflowFile := "workflow.yaml"
	configFile := "config.yaml"
	scopeFile := "scope.txt"
	//err := os.Setenv("NO_COLOR", "1")
	//if err != nil {
	//	log.Fatal(err)
	//}

	config, err := internal.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := internal.NewLogger(config.GetLogPath(), config.GetTarget())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("=========================")
	log.Println("[*] Starting new scan [*]")
	log.Println("=========================")

	scope, err := internal.NewScope(config.GetScope(), config.GetTarget())
	if err != nil {
		log.Fatal(err)
	}
	scopePath := fmt.Sprintf("out/%s/%s", config.GetTarget(), scopeFile)
	fmt.Println(scopePath)
	err = scope.WriteScope(scopePath)
	if err != nil {
		log.Fatal(err)
	}

	workflow, err := internal.NewWorkflow(workflowFile, config)
	if err != nil {
		log.Fatal(err)
	}

	err = internal.Process(workflow, config, logger)
	if err != nil {
		log.Fatal(err)
	}
}
