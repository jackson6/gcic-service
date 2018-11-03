package handler

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"github.com/jackson6/gcic-service/service-client/lib"
	"log"
	"net/http"
	"golang.org/x/net/context"
	pb "github.com/jackson6/gcic-service/partner-service/proto/partner"
	"github.com/jackson6/gcic-service/service-client/client"
	"strconv"
)

func GetPartnerEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	resp, err := service.Partner.All(context.Background(), &pb.Request{})
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

func CreatePartnerEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client, bucket *storage.BucketHandle, bucketName string) {
	defer r.Body.Close()

	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 32)
	lon, _ := strconv.ParseFloat(r.FormValue("lon"), 32)

	coords := pb.Coord{
		Lat: float32(lat),
		Lon: float32(lon),
	}

	r.ParseMultipartForm(32000 << 20)
	partner := pb.Partner{
		Name: r.FormValue("name"),
		Address: r.FormValue("address"),
		Parish: r.FormValue("parish"),
		Country: r.FormValue("country"),
		Contact: r.FormValue("contact"),
		Coord: &coords,
	}

	log.Println(partner)

	imgUrl, err := lib.UploadFileFromForm(r, bucket, bucketName)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	partner.Logo = imgUrl

	resp, err := service.Partner.Create(context.Background(), &partner)
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

func UpdatePartnerEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client){
	defer r.Body.Close()

	var update pb.Partner
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	resp, err := service.Partner.Get(context.Background(), &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, InternalError, err)
		return
	}

	updated := lib.UpdateBuilder(resp.Partner, update)
	newPartner := updated.(pb.Partner)

	_, err = service.Partner.Update(context.Background(), &newPartner)
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

func DeletePartnerEndPoint(w http.ResponseWriter, r *http.Request, service *client.Client) {
	defer r.Body.Close()
	partner := new(pb.Partner)
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		RespondError(w, http.StatusBadRequest, BadRequest, err)
		return
	}

	_, err := service.Partner.Delete(context.Background(), partner)
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