package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {
	godotenv.Load()
}

func getKafkaWriter(topic string) *kafka.Writer {
	brokers := strings.Split(os.Getenv("KAFKA_URL"), ",")
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func main() {
	pflag.StringP("topic", "t", "topic", "topic name")
	pflag.StringP("message", "m", "message", "message of publish")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	kafkaWriter := getKafkaWriter(viper.GetString("topic"))

	defer kafkaWriter.Close()

	err := kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(viper.GetString("message")),
		},
	)

	if err != nil {
		fmt.Println(err.Error())
	}
}
