package main

import (
	"flag"
	"fmt"
	"log"
	"reconr/internal"
)

func main() {
	// Parse command line flags
	workflowFile := flag.String("workflow", "workflowFinal.yaml", "Path to the workflow file")
	configFile := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	// Initialize configuration
	config, err := internal.NewConfig(*configFile)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	// Initialize logger
	logger, err := internal.NewLogger(config.GetLogPath(), config.GetTarget())
	if err != nil {
		fmt.Println(err)
	}

	log.Println("=========================")
	log.Println("[*] Starting new scan [*]")
	log.Println("=========================")

	path := fmt.Sprintf("%s/%s", config.GetWorkPath(), config.GetTarget())

	// Initialize scope
	scope, err := internal.NewScope(config.GetScope(), path)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
	scopePath := fmt.Sprintf("%s/%s", path, config.GetScopeFileName())
	err = scope.WriteScope(scopePath)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	// Initialize workflow
	workflow, err := internal.NewWorkflow(*workflowFile, config)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	// Process workflow
	err = internal.Process(workflow, config, logger)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	log.Println("=====================")
	log.Println("[*] Scan finished [*]")
	log.Println("=====================")
}
