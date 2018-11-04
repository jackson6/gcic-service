package handler

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"github.com/emicklei/go-restful"
	pb "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
	partnerProto "github.com/jackson6/gcic-service/partner-service/proto/partner"
	"github.com/jackson6/gcic-service/service-client/client"
	"github.com/jackson6/gcic-service/service-client/lib"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type PartnerBenefit struct {
	Benefit *pb.Benefit `json:"benefit"`
	Partner *partnerProto.Partner `json:"partner"`
}

func GetBenefitDetailsEndPoint(w http.ResponseWriter, r *restful.Request, service *client.Client) {
	defer r.Request.Body.Close()

	var data PartnerBenefit
	
	id := r.PathParameter("id")

	benefitResp, err := service.Benefit.Get(context.Background(), &pb.Benefit{Id:id})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	log.Println(benefitResp.Benefit)

	partnerResp, err := service.Partner.Get(context.Background(), &partnerProto.Partner{Id:benefitResp.Benefit.PartnerId})
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	data.Benefit = benefitResp.Benefit
	data.Partner = partnerResp.Partner

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: data,
	}
	RespondJSON(w, http.StatusOK, response)
}

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

func CreateBenefitEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client, bucket *storage.BucketHandle, bucketName string) {
	defer r.Body.Close()

	r.ParseMultipartForm(32000 << 20)
	benefit := pb.Benefit{
		Id: r.FormValue("id"),
		Title:  r.FormValue("title"),
		Description: r.FormValue("description"),
		PartnerId: r.FormValue("partner_id"),
	}

	imgUrls, err := lib.UploadMultipleFileFromForm(r, bucket, bucketName)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	benefit.Img = imgUrls

	resp, err := service.Benefit.Create(context.Background(), &benefit)
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