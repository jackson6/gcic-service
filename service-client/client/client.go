package client

import (
	microclient "github.com/micro/go-micro/client"
	userService "github.com/jackson6/gcic-service/user-service/proto/user"
	emailService "github.com/jackson6/gcic-service/email-service/proto/email"
	partnerService "github.com/jackson6/gcic-service/partner-service/proto/partner"
	paymentService "github.com/jackson6/gcic-service/payment-service/proto/payment"
	planService "github.com/jackson6/gcic-service/plan-service/proto/plan"
	benefitService "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
	loyaltyService "github.com/jackson6/gcic-service/loyalty-service/proto/loyalty"
)

// app has router and firebase instances
type Client struct {
	User userService.UserServiceClient
	Email emailService.EmailServiceClient
	Partner partnerService.PartnerServiceClient
	Payment paymentService.PaymentServiceClient
	Plan planService.PlanServiceClient
	Benefit benefitService.BenefitServiceClient
	Loyalty loyaltyService.LoyaltyServiceClient
}

func (c *Client) NewClient() *Client {
	c.User = userService.NewUserServiceClient("gcic.user", microclient.DefaultClient)
	c.Email = emailService.NewEmailServiceClient("gcic.email", microclient.DefaultClient)
	c.Partner = partnerService.NewPartnerServiceClient("gcic.partner", microclient.DefaultClient)
	c.Payment = paymentService.NewPaymentServiceClient("gcic.payment", microclient.DefaultClient)
	c.Plan = planService.NewPlanServiceClient("gcic.plan", microclient.DefaultClient)
	c.Benefit = benefitService.NewBenefitServiceClient("gcic.benefit", microclient.DefaultClient)
	c.Loyalty = loyaltyService.NewLoyaltyServiceClient("gcic.loyalty", microclient.DefaultClient)
	return c
}