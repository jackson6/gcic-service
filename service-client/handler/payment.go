package handler

import (
	"context"
	"encoding/json"
	"github.com/jackson6/gcic-service/service-client/client"
	"net/http"
	pb "github.com/jackson6/gcic-service/payment-service/proto/payment"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
	planProto "github.com/jackson6/gcic-service/plan-service/proto/plan"
)

func PaymentHistoryEndPoint(w http.ResponseWriter, r *http.Request, user *userProto.User, service *client.Client){
	defer r.Body.Close()

	req := pb.Transaction{UserId: user.Id}
	resp, err := service.Payment.History(context.Background(), &req)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Transactions,
	}
	RespondJSON(w, http.StatusOK, response)
}

func RefreshMembershipEndPoint(w http.ResponseWriter, r *http.Request, user *userProto.User, service *client.Client) {
	defer r.Body.Close()

	charge := new(pb.Charge)
	if err := json.NewDecoder(r.Body).Decode(&charge); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	planResp, err := service.Plan.Get(context.Background(), &planProto.Plan{Id:charge.Id})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	newCharge := &pb.Charge{
		Amount: planResp.Plan.Amount,
		Description: planResp.Plan.Description,
		Currency: charge.Currency,
		Token: charge.Token,
		UserId: user.UserId,
	}

	resp, err := service.Payment.CreateCharge(context.Background(), newCharge)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp,
	}
	RespondJSON(w, http.StatusOK, response)
}