package handler

import (
	"encoding/json"
	"github.com/zerosuxx/go-http-commander/utility"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type CommandHandler struct {
	Shell *utility.Shell
}

func (handler *CommandHandler) Handle(res http.ResponseWriter, req *http.Request) {
	inputReader, outputWriter := io.Pipe()
	reqBody, _ := ioutil.ReadAll(req.Body)

	if len(reqBody) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		log.Println("Request body can't be empty!")

		return
	}

	var commandWithArgs []string
	_ = json.Unmarshal(reqBody, &commandWithArgs)

	if len(commandWithArgs) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		log.Println("Command can't be empty!")

		return
	}

	log.Printf("Command: %v", commandWithArgs)

	go func() {
		if _, err := io.Copy(res, inputReader); err != nil {
			utility.LogError(err)
		}
	}()

	if err := handler.Shell.Run(commandWithArgs[0], commandWithArgs[1:], outputWriter); err != nil {
		_, _ = outputWriter.Write([]byte(err.Error()))
		utility.LogError(err)
	}

	if err := outputWriter.Close(); err != nil {
		log.Println(err)
		utility.LogError(err)
	}
}

func CreateCommandHandler() *CommandHandler {
	return &CommandHandler{
		Shell: &utility.Shell{},
	}
}
