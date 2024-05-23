package main

import (
	"flag"
	"fmt"
	"log"
	"reconr/internal"
)

func main() {
	// Parse command line flags
	workflowFile := flag.String("workflow", "workflowProxy.yaml", "Path to the workflow file")
	configFile := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	// Initialize configuration
	config, err := internal.NewConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize logger
	logger, err := internal.NewLogger(config.GetLogPath(), config.GetTarget())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("=========================")
	log.Println("[*] Starting new scan [*]")
	log.Println("=========================")

	path := fmt.Sprintf("%s/%s", config.GetWorkPath(), config.GetTarget())

	// Initialize scope
	scope, err := internal.NewScope(config.GetScope(), path)
	if err != nil {
		log.Fatal(err)
	}
	scopePath := fmt.Sprintf("%s/%s", path, config.GetScopeFileName())
	err = scope.WriteScope(scopePath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize workflow
	workflow, err := internal.NewWorkflow(*workflowFile, config)
	if err != nil {
		log.Fatal(err)
	}

	// Process workflow
	err = internal.Process(workflow, config, logger)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("=====================")
	log.Println("[*] Scan finished [*]")
	log.Println("=====================")
}
