package ecs

type SystemMessageType struct {
	Name string
	CreatingSystem System
}

type SystemMessage struct {
	MessageType SystemMessageType
	MessageContent map[string]string
}

// SystemMessageQueue is a super simple way of messaging between systems. Essentially, it is nothing more than a list of
// messages. Each message has a type, and an originator. Each system can "subscribe" to a type of message, which
// basically just means that it will check the queue for any messages of that type before it does anything else.
// Messages can contain a map of information, which each system that creates messages of that type, and those that
// subscribe to it should know how to handle any information contained in the message. Ideally, the message queue will
// be cleared out occasionally, either by the subscribing systems, or the game loop. Pretty simple for now, but should
// solve a subset of problems nicely.
type SystemMessageQueue struct {
	Messages []SystemMessage
}

// BroadcastMessage appends a system message onto the games SystemMessageQueue, allowing it to consumed by a service
// subscribes to the MessageType.
// TODO: This name is a little misleading, as it doesn't actually broadcast, so much as append to a list for consumption
func (smq *SystemMessageQueue) BroadcastMessage(messageType SystemMessageType, messageContent map[string]string) {
	newMessage := SystemMessage{MessageType: messageType, MessageContent: messageContent}

	smq.Messages = append(smq.Messages, newMessage)
}

// GetMessagesOfType returns a list of SystemMessages that have messageType. Can return an empty list
func (smq *SystemMessageQueue) GetMessagesOfType(messageType SystemMessageType) []SystemMessage {
	messages := []SystemMessage{}

	for _, message := range smq.Messages {
		if message.MessageType == messageType {
			messages = append(messages, message)
		}
	}

	return messages
}