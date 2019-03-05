package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-lambda-go/lambda"
)

type ReservationRequest struct {
	Name string`json:"name"`
	Email string`json:"email"`
	Telephone string`json:"telephone"`
	Party int8`json:"party"`
	Timestamp int64`json:"timestamp"`
	Message string `json:"message"`
}

type ReservationResponse struct {
	Id string`json:"id"`
	Timestamp int64`json:"timestamp"`
}

type Reservation struct {
	Id string`json:"Id"`
	Name string`json:"Name"`
	Email string`json:"Email"`
	Telephone string`json:"Telephone"`
	Party int8`json:"Party"`
	Timestamp int64`json:"Timestamp"`
	Message string `json:"Message"`
}

func isFieldEmpty(s string) bool {
	return len(s) == 0
}

func isRequestMissingRequiredFields(request ReservationRequest) bool {
	return isFieldEmpty(request.Name) || isFieldEmpty(request.Email) || request.Timestamp == 0 || request.Party == 0
}

func saveReservationToDynamo(request ReservationRequest) (Reservation, error) {
	sess, err := session.NewSession(&aws.Config{ Region: aws.String("eu-west-1")}, )
	dynamoDb := dynamodb.New(sess)

	reservation := Reservation {
		Id: fmt.Sprintf("%s", uuid.Must(uuid.NewV4())),
		Name: request.Name,
		Email: request.Email,
		Telephone: request.Telephone,
		Party: request.Party,
		Timestamp: request.Timestamp,
		Message: request.Message,
	}

	marsalMap, err := dynamodbattribute.MarshalMap(reservation)
	input := &dynamodb.PutItemInput{
		Item: marsalMap,
		TableName: aws.String("Reservations"),
	}

	_, err = dynamoDb.PutItem(input)

	return reservation, err
}

func Handler(request ReservationRequest) (ReservationResponse, error) {
	if isRequestMissingRequiredFields(request) {
		panic("request is missing mandatory fields")
	}

	reservation, saveError := saveReservationToDynamo(request)

	if saveError != nil {
		panic(fmt.Errorf("could not save reservation to Dyanmo %s", saveError))
	}

	return ReservationResponse{Id: fmt.Sprintf("%s", reservation.Id), Timestamp: reservation.Timestamp,}, nil
}

func main() {
	lambda.Start(Handler)
}