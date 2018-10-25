package handler

import (
	"net/http"
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/jackson6/gcic-service/service-client/lib"
	"github.com/jackson6/gcic-service/service-client/client"
	pb "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
)

func GetBenefitEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	resp, err := service.Benefit.All(context.Background(), &pb.Request{})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Benefits,
	}
	RespondJSON(w, http.StatusOK, response)
}

func CreateBenefitEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	benefit := new(pb.Benefit)
	if err := json.NewDecoder(r.Body).Decode(&benefit); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}
	resp, err := service.Benefit.Create(context.Background(), benefit)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Benefit,
	}
	RespondJSON(w, http.StatusOK, response)
}

func UpdateBenefitEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client){
	defer r.Body.Close()

	var update pb.Benefit
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := service.Benefit.Get(context.Background(), &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	updated := lib.UpdateBuilder(resp.Benefit, update)
	newBenefit := updated.(pb.Benefit)

	_, err = service.Benefit.Update(context.Background(), &newBenefit)
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

func DeleteBenefitEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	benefit := new(pb.Benefit)
	if err := json.NewDecoder(r.Body).Decode(&benefit); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	_, err := service.Benefit.Delete(context.Background(), benefit)
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