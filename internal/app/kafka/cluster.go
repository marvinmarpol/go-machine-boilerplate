package kafka

import (
	"fmt"
	"sync"
)

type Message struct {
	Value string
}

type Topic string
type ConsumerGroup string

type ConsumerMetadata struct {
	Offset int
}

type KafkaCluster struct {
	mu          sync.Mutex
	MapTopic    map[Topic][]Message
	MapConsumer map[Topic]map[ConsumerGroup]*ConsumerMetadata
}

func NewKafkaCluster() *KafkaCluster {
	return &KafkaCluster{
		MapTopic:    make(map[Topic][]Message),
		MapConsumer: make(map[Topic]map[ConsumerGroup]*ConsumerMetadata),
	}
}

func (k *KafkaCluster) Publish(topic Topic, message string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.MapTopic[topic] = append(k.MapTopic[topic], Message{Value: message})
}

func (k *KafkaCluster) Subscribe(topic Topic, groupID ConsumerGroup) {
	k.mu.Lock()
	if k.MapConsumer[topic] == nil {
		k.MapConsumer[topic] = make(map[ConsumerGroup]*ConsumerMetadata)
	}
	k.mu.Unlock()

	for {
		k.mu.Lock()
		metadata := k.MapConsumer[topic][groupID]
		if metadata == nil {
			k.MapConsumer[topic][groupID] = &ConsumerMetadata{
				Offset: 0,
			}
			metadata = k.MapConsumer[topic][groupID]
		}

		messages := k.MapTopic[topic]
		for metadata.Offset < len(messages) {
			fmt.Printf("subscriber of topic %s with group %s retrieve value %s with offset %d\n", topic, groupID, messages[metadata.Offset], metadata.Offset)
			metadata.Offset++
		}
		k.mu.Unlock()
	}

}
