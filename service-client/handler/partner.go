package handler

import (
	"encoding/json"
	"github.com/jackson6/gcic-service/service-client/lib"
	"net/http"
	"golang.org/x/net/context"
	pb "github.com/jackson6/gcic-service/partner-service/proto/partner"
)

func GetPartnerEndPoint(w http.ResponseWriter, r *http.Request, partnerService pb.PartnerServiceClient) {
	defer r.Body.Close()
	resp, err := partnerService.All(context.Background(), &pb.Request{})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
	return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Partners,
	}
	RespondJSON(w, http.StatusOK, response)
}

func CreatePartnerEndPoint(w http.ResponseWriter, r *http.Request, partnerService pb.PartnerServiceClient) {
	defer r.Body.Close()
	partner := new(pb.Partner)
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}
	resp, err := partnerService.Create(context.Background(), partner)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Partner,
	}
	RespondJSON(w, http.StatusOK, response)
}

func UpdatePartnerEndPoint(w http.ResponseWriter, r *http.Request, partnerService pb.PartnerServiceClient){
	defer r.Body.Close()

	var update pb.Partner
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := partnerService.Get(context.Background(), &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	newUser := lib.UpdateBuilder(resp.Partner, update)

	_, err = partnerService.Update(context.Background(), newUser.(*pb.Partner))
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

func DeletePartnerEndPoint(w http.ResponseWriter, r *http.Request, partnerService pb.PartnerServiceClient) {
	defer r.Body.Close()
	partner := new(pb.Partner)
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	_, err := partnerService.Delete(context.Background(), partner)
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