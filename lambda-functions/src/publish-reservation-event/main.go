package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
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

func publishReservationEvent(reservation Reservation) {
	dynamoDbSession, _ := session.NewSession(&aws.Config{ Region: aws.String("eu-west-1")}, )
	svc := sns.New(dynamoDbSession)
	reservationJson, _ := json.Marshal(reservation)
	params := &sns.PublishInput{
		Message: aws.String(string(reservationJson)),
		TopicArn: aws.String("arn:aws:sns:eu-west-1:896764428848:test-topic"),
	}

	fmt.Printf("Publishing Reservation Event for object %s", string(reservationJson))
	result, err := svc.Publish(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func handler(ctx context.Context, e events.DynamoDBEvent) {
	for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		reservationId := record.Change.NewImage["Id"].String()
		name := record.Change.NewImage["Name"].String()
		email := record.Change.NewImage["Email"].String()
		party, _ := record.Change.NewImage["Party"].Integer()
		timestamp, _ := record.Change.NewImage["Timestamp"].Integer()
		telephone := record.Change.NewImage["Telephone"].String()
		message := record.Change.NewImage["Message"].String()

		var reservation = Reservation{
			Id: reservationId,
			Name: name,
			Email: email,
			Telephone: telephone,
			Party: party,
			Timestamp: timestamp,
			Message: message,
		}

		publishReservationEvent(reservation)
	}
}

func main() {
	lambda.Start(handler)
}