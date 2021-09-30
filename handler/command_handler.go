package handler

import (
	"encoding/json"
	"github.com/zerosuxx/go-commander/utility"
	"io/ioutil"
	"log"
	"net/http"
)

type CommandHandler struct {
	Shell *utility.Shell
}

func (handler *CommandHandler) Handle(res http.ResponseWriter, req *http.Request) {
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
	commandOutput, err := handler.Shell.Exec(commandWithArgs[0], commandWithArgs[1:])
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_, _ = res.Write([]byte(err.Error()))

		return
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write(commandOutput)
	if err != nil {
		utility.LogError(err)
	}
}

func CreateCommandHandler() *CommandHandler {
	return &CommandHandler{
		Shell: &utility.Shell{},
	}
}
