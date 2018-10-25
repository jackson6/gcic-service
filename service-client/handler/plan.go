package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	pb "github.com/jackson6/gcic-service/plan-service/proto/plan"
	"github.com/jackson6/gcic-service/service-client/client"
	"github.com/jackson6/gcic-service/service-client/lib"
	"golang.org/x/net/context"
	"net/http"
)

func GetSingleEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()

	params := mux.Vars(r)
	id := params["id"]

	resp, err := service.Plan.Get(context.Background(), &pb.Plan{Id:id})
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

	update := new(pb.Plan)
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := service.Plan.Get(context.Background(), update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	updated := lib.UpdateBuilder(resp.Plan, update)
	plan := updated.(pb.Plan)

	_, err = service.Plan.Update(context.Background(), &plan)
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

	delete := new(pb.Plan)
	if err := json.NewDecoder(r.Body).Decode(&delete); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	_, err := service.Plan.Delete(context.Background(), delete)
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
