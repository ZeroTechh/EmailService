package main

import (
	"go.uber.org/zap"

	"github.com/ZeroTechh/EmailService/handler"
	"github.com/ZeroTechh/VelocityCore/logger"
	proto "github.com/ZeroTechh/VelocityCore/proto/EmailService"
	"github.com/ZeroTechh/VelocityCore/services"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/ZeroTechh/hades"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config"})
	log    = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)

func main() {
	defer utils.HandlePanic(log)
	handler := handler.New()

	grpcServer, listner := utils.CreateGRPCServer(
		services.EmailVerificationSrv,
		log,
	)

	proto.RegisterEmailServer(grpcServer, handler)
	if err := grpcServer.Serve(*listner); err != nil {
		log.Fatal("Service Failed With Error", zap.Error(err))
	}
}
