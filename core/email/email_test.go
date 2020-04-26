package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	assert := assert.New(t)
	e := New()

	assert.True(sendMails, "Send mails needs to be true for testing")

	assert.NoError(e.Simple("TEST", email))
	assert.Error(e.Simple("TEST", "Invalid"))

	sendMails = false
	assert.NoError(e.Simple("TEST", "invalid"))
}
