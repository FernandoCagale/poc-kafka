package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	godotenv.Load()
}

func getKafkaReader(topic, groupID string) *kafka.Reader {
	brokers := strings.Split(os.Getenv("KAFKA_URL"), ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func main() {
	pflag.StringP("topic", "t", "topic", "topic name")
	pflag.StringP("group", "g", "group", "group name")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	reader := getKafkaReader(viper.GetString("topic"), viper.GetString("group"))

	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
