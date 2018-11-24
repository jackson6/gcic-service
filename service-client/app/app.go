package app

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/context"
	"github.com/jackson6/gcic-service/service-client/client"
	"github.com/jackson6/gcic-service/service-client/config"
	"github.com/jackson6/gcic-service/service-client/handler"
	userPb "github.com/jackson6/gcic-service/user-service/proto/user"
	"github.com/micro/go-web"
	k8s "github.com/micro/kubernetes/go/web"
	netContext "golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	Name         = "go.micro.api"
	Address      = ":3000"
	Handler      = "meta"
	APIPath      = "/"
	Namespace    = "go.micro.api"
	HeaderPrefix = "X-Micro-"
	CORS         = map[string]bool{"*": true}
)

func GetUserStruct(data interface{}) (*userPb.User, error) {
	user := new(userPb.User)
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// app has router and firebase instances
type App struct {
	Router *restful.WebService
	SECRET string
	Firebase *firebase.App
	Client *client.Client
	BucketHandle *storage.BucketHandle
	BucketName string
	Service web.Service
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config, client *client.Client) {
	var err error

	a.Service = k8s.NewService(
		web.Name("go.micro.api.api"),
	)

	a.Service.Init()

	a.Router = new(restful.WebService)
	wc := restful.NewContainer()
	a.Router.Consumes(restful.MIME_XML, restful.MIME_JSON, "application/x-www-form-urlencoded", "multipart/form-data")
	a.Router.Produces(restful.MIME_JSON, restful.MIME_XML)
	a.Router.Path("/api")
	a.Router.Route(a.Router.GET("/pln").To(a.GetPlan))

	a.setRouters()
	a.Service.Handle("/", wc)
	wc.Add(a.Router)

	a.SECRET = config.SECRET
	a.Client = client.NewClient()


	opt := option.WithCredentialsFile("config/invest-ff3f4-firebase-adminsdk-zgkg5-ae79e82ab1.json")
	app, err := firebase.NewApp(netContext.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	a.BucketName = config.BUCKITID
	a.BucketHandle, err = configureStorage(config.BUCKITID)
	if err != nil {
		log.Fatalf("error initializing google storage: %v\n", err)
	}

	a.Firebase = app
}

func configureStorage(bucketID string) (*storage.BucketHandle, error) {
	ctx := netContext.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("config/gcic-82ebaca534ed.json"))
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketID), nil
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects

	a.Get("/user/", a.GetUser, a.ValidateMiddleware)
	a.Post("/user/", a.CreateUser, nil)
	a.Put("/user/", a.UpdateUser, a.ValidateMiddleware)
	a.Post("/user/pic/", a.UploadPrifilePicture, a.ValidateMiddleware)
	a.Get("/user/referral/", a.GetReferral, a.ValidateMiddleware)
	a.Get("/user/contact/", a.GetContact, a.ValidateMiddleware)
	a.Post("/user/message/", a.GetMessage, a.ValidateMiddleware)
	a.Get("/user/member/{id}/", a.GetUserEmail, nil)

	a.Get("/payment/history/", a.GetPaymentHistory, a.ValidateMiddleware)

	a.Get("/plan/{id}/", a.GetSinglePlan, nil)
	a.Get("/plan/", a.GetPlan, nil)
	a.Post("/plan/", a.CreatePlan, nil)
	a.Put("/plan/", a.UpdatePlan, nil)
	a.Delete("/plan/", a.DeletePlan, nil)

	a.Get("/partner/", a.GetPartner, nil)
	a.Post("/partner/", a.CreatePartner, nil)
	a.Put("/partner/", a.UpdatePartner, nil)
	a.Delete("/partner/", a.DeletePartner, nil)

	a.Get("/benefit/", a.GetBenefit, nil)
	a.Get("/benefit/{id}/", a.GetBenefitDetails, nil)

	a.Post("/benefit/", a.CreateBenefit, nil)
	a.Put("/benefit/", a.UpdateBenefit, nil)
	a.Delete("/benefit/", a.DeleteBenefit, nil)

	a.Post("/contact/us/", a.ContactUs, nil)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(req *restful.Request, rsp *restful.Response), auth restful.FilterFunction) {
	if auth != nil {
		a.Router.Route(a.Router.GET(path).Filter(auth).To(f))
	} else {
		a.Router.Route(a.Router.GET(path).To(f))
	}

}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(req *restful.Request, rsp *restful.Response), auth restful.FilterFunction) {
	if auth != nil {
		a.Router.Route(a.Router.POST(path).Filter(auth).To(f))
	} else {
		a.Router.Route(a.Router.POST(path).To(f))
	}
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(req *restful.Request, rsp *restful.Response), auth restful.FilterFunction) {
	if auth != nil {
		a.Router.Route(a.Router.PUT(path).Filter(auth).To(f))
	} else {
		a.Router.Route(a.Router.PUT(path).To(f))
	}
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(req *restful.Request, rsp *restful.Response), auth restful.FilterFunction) {
	if auth != nil {
		a.Router.Route(a.Router.DELETE(path).Filter(auth).To(f))
	} else {
		a.Router.Route(a.Router.DELETE(path).To(f))
	}
}

func (a *App) ValidateMiddleware(req *restful.Request, w *restful.Response, chain *restful.FilterChain) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	authorizationHeader := req.Request.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			client, err := a.Firebase.Auth(netContext.Background())
			if err != nil {
				handler.RespondError(w, http.StatusOK, handler.UnauthorizedError, err)
				return
			}
			token, err := client.VerifyIDToken(netContext.Background(), bearerToken[1])
			if err != nil {
				log.Println(bearerToken[1])
				handler.RespondError(w, http.StatusOK, handler.UnauthorizedError, err)
				return
			}
			resp, err := a.Client.User.Get(netContext.Background(), &userPb.User{Id:token.UID})
			if err != nil {
				handler.RespondError(w, http.StatusOK, handler.NotFound, err)
				return
			}
			context.Set(req.Request, "user", resp.User)
			chain.ProcessFilter(req, w)
		} else {
			handler.RespondError(w, http.StatusOK, handler.UnauthorizedError, fmt.Errorf("Invalid token"))
			return
		}
	} else {
		handler.RespondError(w, http.StatusOK, handler.UnauthorizedError, fmt.Errorf("No token provided"))
	}
}

func (a *App) Run(host string) {
	rand.Seed(time.Now().Unix())
	fmt.Printf("Starting gcic Server Client...%v\n", time.Now().String())
	//log.Fatal(http.ListenAndServe(host, &CORSRouterDecorator{a.Router}))

	if err := a.Service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (a *App) GetContact(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.GetContacts(w, r.Request, user, a.Client)
}

func (a *App) GetMessage(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.GetMessages(w, r.Request, user, a.Client)
}

func (a *App) GetReferral(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.GetUserReferralEndPoint(w, r.Request, user, a.Client)
}

func (a *App) GetUser(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.GetUserEndPoint(w, r.Request, user)
}

func (a *App) CreateUser(r *restful.Request, w *restful.Response) {
	handler.CreateUserEndPoint(w, r.Request, a.Client)
}

func (a *App) RefreshMembership(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.RefreshMembershipEndPoint(w, r.Request, user, a.Client)
}

func (a *App) UpdateUser(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.UpdateUserEndPoint(w, r.Request, user, a.Client)
}

func (a *App) UploadPrifilePicture(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.UploadProfilePicEndPoint(w, r.Request, user, a.Client, a.BucketHandle, a.BucketName)
}

func (a *App) GetPaymentHistory(r *restful.Request, w *restful.Response) {
	user := new(userPb.User)
	userContext := context.Get(r.Request, "user")
	user, err := GetUserStruct(userContext)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, handler.InternalError, err)
	}
	handler.PaymentHistoryEndPoint(w, r.Request, user, a.Client)
}

func (a *App) GetSinglePlan(r *restful.Request, w *restful.Response) {
	handler.GetSingleEndPoint(w, r.Request, a.Client)
}

func (a *App) GetPlan(r *restful.Request, w *restful.Response) {
	handler.GetPlanEndPoint(w, r.Request, a.Client)
}

func (a *App) CreatePlan(r *restful.Request, w *restful.Response) {
	handler.CreatePlanEndPoint(w, r.Request, a.Client)
}

func (a *App) UpdatePlan(r *restful.Request, w *restful.Response) {
	handler.UpdatePlanEndPoint(w, r.Request, a.Client)
}

func (a *App) DeletePlan(r *restful.Request, w *restful.Response) {
	handler.DeletePlanEndPoint(w, r.Request, a.Client)
}

func (a *App) GetPartner(r *restful.Request, w *restful.Response) {
	handler.GetPartnerEndPoint(w, r.Request, a.Client)
}

func (a *App) CreatePartner(r *restful.Request, w *restful.Response) {
	handler.CreatePartnerEndPoint(w, r.Request, a.Client, a.BucketHandle, a.BucketName)
}

func (a *App) UpdatePartner(r *restful.Request, w *restful.Response) {
	handler.UpdatePartnerEndPoint(w, r.Request, a.Client)
}

func (a *App) DeletePartner(r *restful.Request, w *restful.Response) {
	handler.DeletePartnerEndPoint(w, r.Request, a.Client)
}

func (a *App) GetBenefit(r *restful.Request, w *restful.Response) {
	handler.GetBenefitEndPoint(w, r.Request, a.Client)
}

func (a *App) GetBenefitDetails(r *restful.Request, w *restful.Response) {
	handler.GetBenefitDetailsEndPoint(w, r, a.Client)
}

func (a *App) CreateBenefit(r *restful.Request, w *restful.Response) {
	handler.CreateBenefitEndPoint(w, r.Request, a.Client, a.BucketHandle, a.BucketName)
}

func (a *App) UpdateBenefit(r *restful.Request, w *restful.Response) {
	handler.UpdateBenefitEndPoint(w, r.Request, a.Client)
}

func (a *App) DeleteBenefit(r *restful.Request, w *restful.Response) {
	handler.DeleteBenefitEndPoint(w, r.Request, a.Client)
}

func (a *App) GetUserEmail(r *restful.Request, w *restful.Response) {
	handler.GetUserEmailEndPoint(w, r, a.Client)
}

func (a *App) ContactUs(r *restful.Request, w *restful.Response) {
	handler.ContactUsEndPoint(w, r.Request, a.Client)
}