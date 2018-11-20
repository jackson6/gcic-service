package client

import (
	microclient "github.com/micro/go-micro/client"
	userService "github.com/jackson6/gcic-service/user-service/proto/user"
	emailService "github.com/jackson6/gcic-service/email-service/proto/email"
	chatService "github.com/jackson6/gcic-service/chat-service/proto/chat"
	partnerService "github.com/jackson6/gcic-service/partner-service/proto/partner"
	paymentService "github.com/jackson6/gcic-service/payment-service/proto/payment"
	planService "github.com/jackson6/gcic-service/plan-service/proto/plan"
	benefitService "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
	loyaltyService "github.com/jackson6/gcic-service/loyalty-service/proto/loyalty"
)

// app has router and firebase instances
type Client struct {
	User userService.UserServiceClient
	Chat chatService.ChatServiceClient
	Email emailService.EmailServiceClient
	Partner partnerService.PartnerServiceClient
	Payment paymentService.PaymentServiceClient
	Plan planService.PlanServiceClient
	Benefit benefitService.BenefitServiceClient
	Loyalty loyaltyService.LoyaltyServiceClient
}

func (c *Client) NewClient() *Client {
	c.User = userService.NewUserServiceClient("user", microclient.DefaultClient)
	c.Chat = chatService.NewChatServiceClient("chat", microclient.DefaultClient)
	c.Email = emailService.NewEmailServiceClient("email", microclient.DefaultClient)
	c.Partner = partnerService.NewPartnerServiceClient("partner", microclient.DefaultClient)
	c.Payment = paymentService.NewPaymentServiceClient("payment", microclient.DefaultClient)
	c.Plan = planService.NewPlanServiceClient("plan", microclient.DefaultClient)
	c.Benefit = benefitService.NewBenefitServiceClient("benefit", microclient.DefaultClient)
	c.Loyalty = loyaltyService.NewLoyaltyServiceClient("loyalty", microclient.DefaultClient)
	return c
}