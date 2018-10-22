package handler

import (
	"encoding/json"
	"github.com/jackson6/gcic-service/service-client/client"
	pb "github.com/jackson6/gcic-service/email-service/proto/email"
	"golang.org/x/net/context"
	"net/http"
)

func ContactUsEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	email := new(pb.ContactUs)
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}
	_, err := service.Email.Contact(context.Background(), email)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success"
	}
	RespondJSON(w, http.StatusOK, response)
}
