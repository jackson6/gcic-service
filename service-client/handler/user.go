package handler

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	paymentProto "github.com/jackson6/gcic-service/payment-service/proto/payment"
	planProto "github.com/jackson6/gcic-service/plan-service/proto/plan"
	"github.com/jackson6/gcic-service/service-client/client"
	"github.com/jackson6/gcic-service/service-client/lib"
	pb "github.com/jackson6/gcic-service/user-service/proto/user"
	"golang.org/x/net/context"
	"net/http"
)

func GetUserEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User){
	defer r.Body.Close()
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: user,
	}
	RespondJSON(w, http.StatusOK, response)
}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client){
	defer r.Body.Close()

	var req pb.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := service.User.Create(context.Background(), &req)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.User,
	}
	RespondJSON(w, http.StatusOK, response)
}

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User, service *client.Client){
	defer r.Body.Close()

	var update pb.User
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	newUser := lib.UpdateBuilder(user, update)

	_, err := service.User.Update(context.Background(), newUser.(*pb.User))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: newUser,
	}
	RespondJSON(w, http.StatusOK, response)
}

func RefreshMembershipEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User, service *client.Client) {
	defer r.Body.Close()

	charge := new(paymentProto.Charge)
	if err := json.NewDecoder(r.Body).Decode(&charge); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	planResp, err := service.Plan.Get(context.Background(), &planProto.Plan{Id:charge.Id})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	newCharge := &paymentProto.Charge{
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

func UploadProfilePicEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User, service *client.Client, bucket *storage.BucketHandle, bucketName string){

	imgUrl, err := lib.UploadFileFromForm(r, bucket, bucketName)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	user = &pb.User{Id:"1234"}
	update := &pb.User{ProfilePic:imgUrl}

	updated := lib.UpdateBuilder(*user, update)
	newUser := updated.(pb.User)

	_, err = service.User.Update(context.Background(), &newUser)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: newUser,
	}
	RespondJSON(w, http.StatusOK, response)
}