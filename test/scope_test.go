package test

import (
	"fmt"
	"gotest.tools/v3/assert"
	"log"
	"os"
	"reconr/internal"
	"testing"
	"time"
)

var check = "192.168.10.10\n192.168.20.10\n"
var url = "whatever.com"

func init() {
	if _, err := os.Stat(fmt.Sprintf("./log/%s.log", url)); err == nil {
		err := os.Remove(fmt.Sprintf("./log/%s.log", url))
		if err != nil {
			log.Fatal(err)
		}
	}
	err := os.RemoveAll(fmt.Sprintf("./out/%s", url))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
}

func Test_scope(T *testing.T) {
	workflowFile := "workflow_test.yaml"
	configFile := "config_test.yaml"

	config, err := internal.NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
	}
	logger, err := internal.NewLogger(config.GetLogPath(), config.GetTarget())
	if err != nil {
		fmt.Println(err)
	}

	path := fmt.Sprintf("%s/%s", config.GetWorkPath(), config.GetTarget())
	scope, err := internal.NewScope(config.GetScope(), path)
	if err != nil {
		fmt.Println(err)
	}
	scopePath := fmt.Sprintf("%s/%s/%s", config.GetWorkPath(), config.GetTarget(), config.GetScopeFileName())
	err = scope.WriteScope(scopePath)
	if err != nil {
		fmt.Println(err)
	}

	workflow, err := internal.NewWorkflow(workflowFile, config)
	if err != nil {
		fmt.Println(err)
	}

	err = internal.Process(workflow, config, logger)
	if err != nil {
		//This error is caused by the task to be to fast. So the logs can not benn read.
		if err.Error() != "Error response from daemon: can not get logs from container which is dead or marked for removal" {
			fmt.Println(err.Error())
		}
	}

	dat, err := os.ReadFile(fmt.Sprintf("./out/%s/validScope.txt", url))
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(T, string(dat), check)
}
