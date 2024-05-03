package main

import (
	"flag"
	"fmt"
	"log"
	"reconr/internal"
)

func main() {
	workflowFile := flag.String("workflow", "workflow3.yaml", "Path to the workflow file")
	configFile := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	config, err := internal.NewConfig(*configFile)
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

	path := fmt.Sprintf("%s/%s", config.GetWorkPath(), config.GetTarget())
	scope, err := internal.NewScope(config.GetScope(), path)
	if err != nil {
		log.Fatal(err)
	}
	scopePath := fmt.Sprintf("%s/%s/%s", config.GetWorkPath(), config.GetTarget(), config.GetScopeFileName())
	err = scope.WriteScope(scopePath)
	if err != nil {
		log.Fatal(err)
	}

	workflow, err := internal.NewWorkflow(*workflowFile, config)
	if err != nil {
		log.Fatal(err)
	}

	err = internal.Process(workflow, config, logger)
	if err != nil {
		//This error is caused by the task to be too fast. So the logs can not benn read.
		if err.Error() != "Error response from daemon: can not get logs from container which is dead or marked for removal" {
			log.Fatal(err)
		}
	}
}
