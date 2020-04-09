package ui

// MessageLog keeps track of a list messages, and defines how many messages to keep track of before truncating the list
type MessageLog struct {
	messages  []string
	MaxLength int
}

// NewMessageLog creates a new MessageLog with a maxLength
func NewMessageLog(maxLength int) *MessageLog {
	messageLog := MessageLog{MaxLength: maxLength}
	messageLog.messages = []string{}
	return &messageLog
}

// SendMessage adds a new message to the MessageLog. If the new message would exceed the total number of messages this
// MessageLog can hold, the oldest message will be truncated from the log. New messages are pre-pended onto the messages
// slice
func (ml *MessageLog) SendMessage(message string) {
	// Prepend the message onto the messageLog slice
	if len(ml.messages) >= ml.MaxLength {
		// Throw away any messages that exceed our total queue size
		ml.messages = ml.messages[:len(ml.messages)-1]
	}
	ml.messages = append([]string{message}, ml.messages...)
}

// PrintMessages prints messages, up to displayNum, in reverse order (newest messages get printed first). Any messges
// in the messages slice will not be printed
func (ml *MessageLog) PrintMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, displayNum int) []string {
	// Print the latest five messages from the messageLog. These will be printed in reverse order (newest at the top),
	// to make it appear they are scrolling down the screen
	clearMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, 1)

	toShow := 0

	if len(ml.messages) <= displayNum {
		// Just loop through the messageLog, printing them in reverse order
		toShow = len(ml.messages)
	} else {
		// If we have more than {displayNum} messages stored, just show the {displayNum} most recent
		toShow = displayNum
	}

	printedMessages := []string{}

	for i := toShow; i > 0; i-- {
		PrintText(1, (viewAreaY-1)+i, 0, 0, ml.messages[i-1], "white", "", 1)
		printedMessages = append(printedMessages, ml.messages[i-1])
	}

	return printedMessages
}

// ClearMessage clears the defined message area, starting at viewAreaX and Y, and ending at the width and height of
// the message area
func clearMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, layer int) {
	ClearArea(viewAreaX, viewAreaY, windowSizeX, windowSizeY-viewAreaY, 1)
}

// PrintToMessageArea clears the message area, and print a single message at the top
func PrintToMessageArea(message string, viewAreaX, viewAreaY, windowSizeX, windowSizeY, layer int) {
	clearMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, layer)
	PrintText(1, viewAreaY, 0, 0, message, "white", "", 1)
}
