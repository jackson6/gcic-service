package handler

import (
	"context"
	"encoding/json"
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	"github.com/jackson6/gcic-service/service-client/client"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
	"net/http"
)

func GetContacts(w http.ResponseWriter, r *http.Request, user *userProto.User, service *client.Client) {
	defer r.Body.Close()

	contactResponse, err := service.Chat.Contacts(context.Background(), &pb.User{Id: user.Id, ReferralCode: user.ReferralCode})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: contactResponse.Contacts,
	}
	RespondJSON(w, http.StatusOK, response)
}

func GetMessages(w http.ResponseWriter, r *http.Request, user *userProto.User, service *client.Client) {
	defer r.Body.Close()

	var req pb.MessageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	req.From = user.Id

	messageResponse, err := service.Chat.Messages(context.Background(), &req)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: messageResponse.Messages,
	}
	RespondJSON(w, http.StatusOK, response)
}