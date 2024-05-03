package test

import (
	"fmt"
	"gotest.tools/v3/assert"
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
			fmt.Println(err)
		}
	}
	err := os.RemoveAll(fmt.Sprintf("./out/%s", url))
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
}

func Test_scope(T *testing.T) {
	workflowFile := "workflow_test.yaml"
	configFile := "config_test.yaml"

	config, err := internal.NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
		T.Fail()
	}
	logger, err := internal.NewLogger(config.GetLogPath(), config.GetTarget())
	if err != nil {
		fmt.Println(err)
		T.Fail()
	}

	path := fmt.Sprintf("%s/%s", config.GetWorkPath(), config.GetTarget())
	scope, err := internal.NewScope(config.GetScope(), path)
	if err != nil {
		fmt.Println(err)
		T.Fail()
	}
	scopePath := fmt.Sprintf("%s/%s/%s", config.GetWorkPath(), config.GetTarget(), config.GetScopeFileName())
	err = scope.WriteScope(scopePath)
	if err != nil {
		fmt.Println(err)
		T.Fail()
	}

	workflow, err := internal.NewWorkflow(workflowFile, config)
	if err != nil {
		fmt.Println(err)
		T.Fail()
	}

	err = internal.Process(workflow, config, logger)
	if err != nil {
		//This error is caused by the task to being too fast. So the logs can't be read.
		if err.Error() != "Error response from daemon: can not get logs from container which is dead or marked for removal" {
			fmt.Println(err.Error())
			T.Fail()
		}
	}

	dat, err := os.ReadFile(fmt.Sprintf("./out/%s/validScope.txt", url))
	if err != nil {
		fmt.Println(err)
		T.Fail()
	}
	assert.Equal(T, string(dat), check)
}
