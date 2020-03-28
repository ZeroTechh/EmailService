package handler

import (
	"context"

	"github.com/ZeroTechh/VelocityCore/logger"
	proto "github.com/ZeroTechh/VelocityCore/proto/EmailService"
	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/hades"
	"go.uber.org/zap"

	"github.com/ZeroTechh/EmailService/email"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config", "../config"})
	log    = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
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
	funcLog := blaze.NewFuncLog(
		"EmailService.Handler.SendSimpleEmail",
		log,
		zap.String("email", emailData.Email),
		zap.String("text", emailData.Text),
	)

	err := handler.email.SendSimpleEmail(emailData.Text, emailData.Email)
	if err != nil {
		err = funcLog.Error(
			err,
			zap.String("email", emailData.Email),
			zap.String("text", emailData.Text),
		)
		return &proto.Empty{}, err
	}

	funcLog.Completed(
		zap.String("email", emailData.Email),
		zap.String("text", emailData.Text),
	)
	return &proto.Empty{}, nil
}
