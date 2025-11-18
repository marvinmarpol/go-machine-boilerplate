package kafka

type Subscriber struct {
	MessageQueue IMessageQueue
}

func NewSubscriber(MessageQueue IMessageQueue) *Subscriber {
	return &Subscriber{MessageQueue}
}

func (s *Subscriber) Subscribe(topic Topic, groupID ConsumerGroup) {
	s.MessageQueue.Subscribe(topic, groupID)
}
