package handler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/EmailService"
	"github.com/ZeroTechh/hades"

	"github.com/ZeroTechh/EmailService/core/email"
)

var config = hades.GetConfig("main.yaml", []string{"config", "../config"})

// New creates a new handler.
func New() *Handler {
	h := Handler{}
	h.init()
	return &h
}

// Handler handles all email service functions.
type Handler struct {
	email *email.Email
}

// init initializes.
func (h *Handler) init() {
	h.email = email.New()
}

// SendSimpleEmail sends a simple text email.
func (h Handler) SendSimpleEmail(
	ctx context.Context, data *proto.EmailData) (*proto.Empty, error) {

	err := h.email.Simple(data.Text, data.Email)
	return &proto.Empty{}, err
}
