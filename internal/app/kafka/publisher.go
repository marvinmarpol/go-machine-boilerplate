package kafka

type Publisher struct {
	MessageQueue IMessageQueue
}

func NewPublisher(MessageQueue IMessageQueue) *Publisher {
	return &Publisher{
		MessageQueue,
	}
}

func (p *Publisher) Publish(topic Topic, message string) {
	p.MessageQueue.Publish(topic, message)
}
