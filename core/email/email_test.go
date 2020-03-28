package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)

	assert.True(sendMails)

	email := Email{}
	email.Init()

	// Testing sending of a simple text email
	err := email.SendSimpleEmail("TEST", smtpConfig.Str("email"))
	assert.NoError(err)

	err = email.SendSimpleEmail("TEST", "Invalid")
	assert.Error(err)

	sendMails = false
	err = email.SendSimpleEmail("TEST", "invalid")
	assert.NoError(err)

}
