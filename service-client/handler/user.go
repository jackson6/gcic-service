package handler

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"github.com/emicklei/go-restful"
	paymentProto "github.com/jackson6/gcic-service/payment-service/proto/payment"
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

func GetUserEmailEndPoint(w http.ResponseWriter, r *restful.Request, service *client.Client){
	defer r.Request.Body.Close()

	id := r.PathParameter("id")

	user := new(pb.User)

	req := pb.User{MemberId: id}
	resp, err := service.User.GetByMemberId(context.Background(), &req)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	user.Email = resp.User.Email
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: user,
	}
	RespondJSON(w, http.StatusOK, response)
}

func GetUserReferralEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User, service *client.Client){
	defer r.Body.Close()

	resp, err := service.User.GetUserReferral(context.Background(), user)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: resp.Users,
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

	update := new(pb.User)
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	update.Id = user.Id

	//updated := lib.UpdateBuilder(user, update)
	//newUser := updated.(pb.User)

	_, err := service.User.Update(context.Background(), update)
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

func UploadProfilePicEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User, service *client.Client, bucket *storage.BucketHandle, bucketName string){

	imgUrl, err := lib.UploadFileFromForm(r, bucket, bucketName)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	update := &pb.User{ProfilePic:imgUrl}
	update.Id = user.Id

	//updated := lib.UpdateBuilder(user, update)
	//newUser := updated.(pb.User)

	_, err = service.User.Update(context.Background(), update)
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

func Messages(w http.ResponseWriter, r *http.Request, user *pb.User, service *client.Client){
	defer r.Body.Close()

	req := paymentProto.Transaction{UserId: user.Id}
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