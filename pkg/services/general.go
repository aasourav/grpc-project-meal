package services

import (
	"context"

	"aas.dev/pkg/models/types"
	mailsvc "github.com/aasourav/proto/mail-service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GeneralService struct {
	repo any
}

func NewGeneralService(repo any) *GeneralService {
	return &GeneralService{repo: repo}
}

func (s *GeneralService) EmailVerify(c *gin.Context, emailVerify types.EmailVerifyTypes) (*mailsvc.MailServiceResponse, error) {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := mailsvc.NewEmailServiceClient(conn)
	req := &mailsvc.MailServiceRequest{Email: emailVerify.Email, Name: emailVerify.Name}

	mailSvcRes, err := client.SendVerificationEmail(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	// token, _ := utils.GenerateJWT(emailVerify.Email)
	// c.SetCookie(types.VERIFIED_EAMAIL_COOKIE, token, 3600, "/", "", false, true)
	return mailSvcRes, nil
}
