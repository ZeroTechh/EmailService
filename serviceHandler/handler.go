package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/EmailService"
	"github.com/ZeroTechh/hades"

	"github.com/ZeroTechh/EmailService/core/email"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config", "../config"})
)

// Handler is used to handle all email service functions
type Handler struct {
	email email.Email
}

// Init initializes
func (handler *Handler) Init() {
	handler.email.Init()
}

// SendSimpleEmail is used to send a simple text email
func (handler Handler) SendSimpleEmail(
	ctx context.Context,
	emailData *proto.EmailData) (*proto.Empty, error) {

	err := handler.email.SendSimpleEmail(emailData.Text, emailData.Email)
	if err != nil {
		return &proto.Empty{}, err
	}

	return &proto.Empty{}, nil
}
