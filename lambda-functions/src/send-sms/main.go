package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Reservation struct {
	Id string`json:"id"`
	Name string`json:"name"`
	Email string`json:"email"`
	Telephone string`json:"telephone"`
	Party int64`json:"party"`
	Timestamp int64`json:"timestamp"`
	Message string `json:"message"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent)  {
	for _, message := range sqsEvent.Records {
		reservation := Reservation{}
		err := json.Unmarshal([]byte(message.Body), &reservation)
		if err != nil {
			panic(err)
		}
		fmt.Println(reservation)
	}
}

func main() {
	lambda.Start(handler)
}