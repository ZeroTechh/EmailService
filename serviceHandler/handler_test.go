package serviceHandler

import (
	"context"
	"testing"

	proto "github.com/ZeroTechh/VelocityCore/proto/EmailService"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	assert := assert.New(t)

	handler := Handler{}
	handler.Init()
	assert.True(true)

	assert.True(config.Map("service").Bool("sendMails"))

	_, err := handler.SendSimpleEmail(context.TODO(), &proto.EmailData{
		Email: config.Map("smtp").Str("email"),
		Text:  "test",
	})
	assert.NoError(err)
}
