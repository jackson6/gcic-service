package handler

import (
	"net/http"
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/jackson6/gcic-service/service-client/lib"
	pb "github.com/jackson6/gcic-service/plan-service/proto/plan"
)

func GetPlanEndPoint(w http.ResponseWriter, r *http.Request, planService pb.PlanServiceClient) {
	defer r.Body.Close()
	resp, err := planService.All(context.Background(), &pb.Request{})
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

func CreatePlanEndPoint(w http.ResponseWriter, r *http.Request, planService pb.PlanServiceClient) {
	defer r.Body.Close()
	plan := new(pb.Plan)
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}
	resp, err := planService.Create(context.Background(), plan)
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

func UpdatePlanEndPoint(w http.ResponseWriter, r *http.Request, planService pb.PlanServiceClient){
	defer r.Body.Close()

	var update pb.Plan
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := planService.Get(context.Background(), &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	plan := lib.UpdateBuilder(resp.Plan, update)


	_, err = planService.Update(context.Background(), plan.(*pb.Plan))
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

func DeletePlanEndPoint(w http.ResponseWriter, r *http.Request, planService pb.PlanServiceClient) {
	defer r.Body.Close()
	plan := new(pb.Plan)
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	_, err := planService.Create(context.Background(), plan)
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
