package handler

import (
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
	pb "github.com/jackson6/gcic-service/user-service/proto/user"
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

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request, userService pb.UserServiceClient){
	defer r.Body.Close()

	var req pb.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := userService.Create(context.Background(), &req)
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

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request, user *pb.User, userService pb.UserServiceClient){
	defer r.Body.Close()

	var update pb.User
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	err := userService.Update(context.Background(), &update)
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