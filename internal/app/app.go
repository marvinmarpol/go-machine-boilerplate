package app

import (
	"bufio"
	"go-machine-boilerplate/internal/app/kafka"
	"os"
	"strings"
)

func Run() error {

	k := kafka.NewKafkaCluster()
	publisher := kafka.NewPublisher(k)
	subscriber := kafka.NewSubscriber(k)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Fields(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		switch args[0] {
		case "publish":
			if len(args) >= 3 {
				publisher.Publish(kafka.Topic(args[1]), args[2])
			}
		case "subscribe":
			if len(args) >= 3 {
				go subscriber.Subscribe(kafka.Topic(args[1]), kafka.ConsumerGroup(args[2]))
			}
		}

	}

	return nil
}
