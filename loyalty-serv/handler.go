// plan-service/handler.go

package main

import (
	"encoding/json"
	pb "github.com/jackson6/gcic-service/loyalty-serv/proto/loyalty"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	client *http.Client
	namespace string
}

func (s *service) GetRepo() Repository {
	return &LoyaltyRepository{s.client}
}

func (s *service) EarnPoints(ctx context.Context, req *pb.Transaction, res *pb.Response) error {
	repo := s.GetRepo()

	var transactionData Transaction
	response := new(pb.Transaction)

	transaction := new(Transaction)
	transaction.Class = "org.gcic.mynetwork.EarnPoints"
	transaction.Member = "resource:" + s.namespace + ".Member#" + req.MemberId
	transaction.Points = req.Points
	transaction.Timestamp = time.Now()

	data, err := repo.Add(s.namespace, transaction)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, transactionData)
	if err != nil {
		return err
	}

	response.MemberId = req.MemberId
	response.TransactionId = transactionData.TransactionId
	response.Timestamp = transactionData.Timestamp.String()
	response.Points = transactionData.Points

	res.Transaction = response
	return nil
}


func (s *service) UsePoints(ctx context.Context, req *pb.Transaction, res *pb.Response) error {
	repo := s.GetRepo()

	var transactionData Transaction
	response := new(pb.Transaction)

	transaction := new(Transaction)
	transaction.Class = "org.gcic.mynetwork.UsePoints"
	transaction.Member = "resource:" + s.namespace + ".Member#" + req.MemberId
	transaction.Partner = "resource:" + s.namespace + ".Partner#" + req.PartnerId
	transaction.Points = req.Points
	transaction.Timestamp = time.Now()

	data, err := repo.Add(s.namespace, transaction)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, transactionData)
	if err != nil {
		return err
	}

	response.MemberId = req.MemberId
	response.TransactionId = transactionData.TransactionId
	response.Timestamp = transactionData.Timestamp.String()
	response.Points = transactionData.Points

	res.Transaction = response
	return nil
}

func (s *service) AddMember(ctx context.Context, req *pb.Member, res *pb.Response) error {
	repo := s.GetRepo()

	var memberData Member
	response := new(pb.Member)

	member := &Member{
		Class: "org.gcic.mynetwork.Member",
		MemberId: req.MemberId,
		FirstName: req.FirstName,
		LastName: req.LastName,
		CardNumber: req.CardNumber,
		Points: 0.00,
	}

	data, err := repo.Add(s.namespace+ ".Member", member)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, memberData)
	if err != nil {
		return err
	}

	response.MemberId = member.MemberId
	response.FirstName = member.FirstName
	response.LastName = member.LastName
	response.CardNumber = member.CardNumber
	response.Points = member.Points

	res.Code = 200
	res.Member = response
	return nil
}

func (s *service) AddPartner(ctx context.Context, req *pb.Partner, res *pb.Response) error {
	repo := s.GetRepo()

	var partnerData Partner
	response := new(pb.Partner)

	partner := &Partner{
		Class: "org.gcic.mynetwork.Partner",
		PartnerId: req.PartnerId,
		Name: req.Name,
		Address: req.Address,
	}

	data, err := repo.Add(s.namespace+ ".Partner", partner)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, partnerData)
	if err != nil {
		return err
	}

	response.PartnerId = partner.PartnerId
	response.Name = partner.Name
	response.Address = partner.Address

	res.Code = 200
	res.Partner = response
	return nil
}

func (s *service) GetMember(ctx context.Context, req *pb.MemberRequestId, res *pb.Response) error {
	repo := s.GetRepo()

	var member Member
	response := new(pb.Member)

	data, err := repo.GetSingle(s.namespace+ ".Member", req.MemberId)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, member)
	if err != nil {
		return err
	}

	response.MemberId = member.MemberId
	response.FirstName = member.FirstName
	response.LastName = member.LastName
	response.CardNumber = member.CardNumber
	response.Points = member.Points

	res.Code = 200
	res.Member = response
	return nil
}

func (s *service) GetPartner(ctx context.Context, req *pb.PartnerRequestId, res *pb.Response) error {
	repo := s.GetRepo()

	var partner Partner
	response := new(pb.Partner)

	data, err := repo.GetSingle(s.namespace + ".Partner", req.PartnerId)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, partner)
	if err != nil {
		return err
	}

	response.PartnerId = partner.PartnerId
	response.Name = partner.Name
	response.Address = partner.Address

	res.Code = 200
	res.Partner = response
	return nil
}
