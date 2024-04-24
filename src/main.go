package main

import (
	_ "embed"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

const (
	ExecutedScriptsTableName = "ExecutedScripts"
)

var (
	dataBase LocalStorage
)

func worker(command string, ch chan Answer, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		ch <- Answer{
			Error:         true,
			ErrorReason:   err.Error(),
			Command:       command,
			CommandResult: string(output),
		}
		return
	}
	ch <- Answer{
		Error:         false,
		ErrorReason:   "",
		Command:       command,
		CommandResult: string(output),
	}
}

func executeCommands(commands []string) []Answer {
	ch := make(chan Answer, len(commands))
	var wg sync.WaitGroup
	wg.Add(len(commands))

	for _, command := range commands {
		go worker(command, ch, &wg)
	}

	wg.Wait()
	var answers []Answer
	for answer := range ch {
		answers = append(answers, answer)
	}
	return answers
}

//go:embed db_config.yml
var rawDBConfig []byte

func main() {

	var dbConfig DBConfig
	var err error
	if err := yaml.Unmarshal(rawDBConfig, &dbConfig); err != nil {
		panic(err)
	}
	dataBase, err = NewDatabase(dbConfig,
		TablesNames{
			ExecutedScripts: ExecutedScriptsTableName,
		},
	)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
