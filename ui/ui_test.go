package ui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessageLog(t *testing.T) {
	messageLog := NewMessageLog(100)

	assert.NotNil(t, messageLog)
	assert.Equal(t, 0, len(messageLog.messages))
}

func TestMessageLog_SendMessage(t *testing.T) {
	messageLog := NewMessageLog(10)

	messageLog.SendMessage("first message")
	assert.Equal(t, 1, len(messageLog.messages))
	assert.Equal(t, "first message", messageLog.messages[0])

	// Fill up the message log with nine more messages
	for i := 0; i < 9; i++ {
		messageLog.SendMessage("test message")
	}

	// New messages are pre-pended to the messages list (the order they will be displayed to the user), check that our
	// first message is now the last element in the list
	assert.Equal(t, "first message", messageLog.messages[9])
	assert.Equal(t, 10, len(messageLog.messages))

	// Add another message. This will cause the first message to be truncated
	messageLog.SendMessage("newest message")
	assert.Equal(t, 10, len(messageLog.messages))
	assert.Equal(t, "test message", messageLog.messages[9])
	assert.Equal(t, "newest message", messageLog.messages[0])

}

func TestMessageLog_PrintMessages(t *testing.T) {
	messageLog := NewMessageLog(10)

	messageLog.SendMessage("first message")

	for i := 0; i < 4; i++ {
		messageLog.SendMessage("test message")
	}

	messageLog.SendMessage("newest message")
	assert.Equal(t, 6, len(messageLog.messages))

	printedMessages := messageLog.PrintMessages(0, 0, 0 ,0, 5)
	assert.Equal(t, 5, len(printedMessages))
	assert.Equal(t, "newest message", printedMessages[4])

	printedMessages = messageLog.PrintMessages(0, 0, 0, 0, 100)
	assert.Equal(t, 6, len(printedMessages))
	assert.Equal(t, "newest message", printedMessages[5])
	assert.Equal(t, "first message", printedMessages[0])
}
