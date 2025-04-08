package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateScheduleRequest(w http.ResponseWriter, r *http.Request) {
	request, err := r.GetBody()
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Log(r).WithError(err).Error(request)

}
