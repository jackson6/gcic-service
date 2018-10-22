package handler

import (
	"encoding/json"
	pb "github.com/jackson6/gcic-service/plan-service/proto/plan"
	"github.com/jackson6/gcic-service/service-client/lib"
	"github.com/jackson6/gcic-service/service-client/client"
	"golang.org/x/net/context"
	"net/http"
)

func GetPlanEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	resp, err := service.Plan.All(context.Background(), &pb.Request{})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Plans,
	}
	RespondJSON(w, http.StatusOK, response)
}

func CreatePlanEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	plan := new(pb.Plan)
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}
	resp, err := service.Plan.Create(context.Background(), plan)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Plan,
	}
	RespondJSON(w, http.StatusOK, response)
}

func UpdatePlanEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client){
	defer r.Body.Close()

	var update pb.Plan
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := service.Plan.Get(context.Background(), &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	plan := lib.UpdateBuilder(resp.Plan, update)


	_, err = service.Plan.Update(context.Background(), plan.(*pb.Plan))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
	}
	RespondJSON(w, http.StatusOK, response)
}

func DeletePlanEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	plan := new(pb.Plan)
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	_, err := service.Plan.Create(context.Background(), plan)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
	}
	RespondJSON(w, http.StatusOK, response)
}
