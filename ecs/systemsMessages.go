package ecs

type SystemMessageType struct {
	Name string
}

type SystemMessage struct {
	MessageType SystemMessageType
	Originator System
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
	Messages map[System][]SystemMessage
	Subscriptions map[System][]SystemMessageType
}

func InitializeSystemMessageQueue() *SystemMessageQueue {
	smq := SystemMessageQueue{}
	smq.Messages = make(map[System][]SystemMessage)
	smq.Subscriptions = make(map[System][]SystemMessageType)
	return &smq
}

// BroadcastMessage appends a system message onto the games SystemMessageQueue, allowing it to consumed by a service
// subscribes to the MessageType.
func (smq *SystemMessageQueue) BroadcastMessage(messageType SystemMessageType, messageContent map[string]string, originator System) {
	newMessage := SystemMessage{MessageType: messageType, MessageContent: messageContent, Originator: originator}

	// Find all subscriptions to this message type, and add this message to the subscribers message queue
	for subscribedSystem, typeList := range smq.Subscriptions {
		if MessageTypeInSlice(messageType, typeList) {
			smq.Messages[subscribedSystem] = append(smq.Messages[subscribedSystem], newMessage)
		}
	}
}

// GetSubscribedMessages returns a list of SystemMessages that have messageType. Can return an empty list
func (smq *SystemMessageQueue) GetSubscribedMessages(system System) []SystemMessage {
	messages := []SystemMessage{}

	for _, message := range smq.Messages[system] {
		messages = append(messages, message)
	}

	return messages
}

// DeleteMessages deletes a processed message from the queue (for example, if the event has been processed)
func (smq *SystemMessageQueue) DeleteMessages(messageName string, system System) {
	modifiedQueue := smq.Messages[system]
	for index, message := range smq.Messages[system] {
		if message.MessageType.Name == messageName {
			modifiedQueue[index] = modifiedQueue[len(modifiedQueue)-1]
			modifiedQueue = modifiedQueue[:len(modifiedQueue)-1]
		}
	}

	smq.Messages[system] = modifiedQueue
}

//MessageTypeInSlice will return true if the MessageType provided is present in the slice provided, false otherwise
func MessageTypeInSlice(a SystemMessageType, list []SystemMessageType) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}