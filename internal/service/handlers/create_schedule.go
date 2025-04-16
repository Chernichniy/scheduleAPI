package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type temporary struct { //TEMPORARY. DELETE IN FUTURE
	data string
}

func CreateScheduleRequest(w http.ResponseWriter, r *http.Request) {

	var request temporary
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Log(r).WithError(err).Error(request)

	if request.data == "" {
		Log(r).WithError(err).Error("didn't provide telegram @username")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	//chatID, err := requests.getChatID(request.data)
	if err != nil {
		Log(r).WithError(err).Error("failed to get chat ID with provided @username")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Further an message structuring is needed

}
