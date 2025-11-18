package kafka

type IMessageQueue interface {
	Publish(topic Topic, message string)
	Subscribe(topic Topic, groupID ConsumerGroup)
}
