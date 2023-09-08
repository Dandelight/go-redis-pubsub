package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	sub := client.Subscribe("chat")
	defer sub.Close()

	go func() {
		for msg := range sub.Channel() {
			fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
		}
	}()

	fmt.Println("Up and running!")

	for {
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		n, err := client.Publish("chat", s).Result()

		fmt.Printf("%d client received the message", n)
	}
}
