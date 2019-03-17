package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"os"
	"strconv"
	"time"
)

type Reservation struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Party     int  `json:"party"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

type ContactDetails struct {
	Name string
	Email string
}

func formatTimestamp(millis int64) string {
	unixTime := time.Unix(0, millis * int64(time.Millisecond))
	day := unixTime.Day()
	month := unixTime.Month()
	year := unixTime.Year()
	hour := unixTime.Hour()
	minute := unixTime.Minute()
	return fmt.Sprintf("%d/%d/%d at %d:%d", day, month, year, hour, minute)
}

func buildJson(from ContactDetails, to ContactDetails, reservation Reservation, subject string, templateId string) string {
	return `{
		"from": {
			"email": "` + from.Email + `", 
			"name": "` + from.Name + `"
		},
		"reply_to": {
			"email": "` + from.Email + `", 
			"name": "` + from.Name + `"
		},
		"personalizations": [{
			"subject": "Your recent reservation at \"Sto Psito\" Restaurant", 
			"dynamic_template_data": {
				"name": "` + reservation.Name + `", 
				"email": "` + reservation.Email + `",
				"telephone": "` + reservation.Telephone + `",
				"party": "` + strconv.Itoa(reservation.Party) + `",
				"timestamp": "` + formatTimestamp(reservation.Timestamp) + `",
				"message": "` + reservation.Message + `"
			}, 
			"to": [{
				"email": "` + to.Email + `", 
				"name": "` + to.Name + `"
			}]
		}],
		"subject": "` + subject + `", 
		"template_id": "` + templateId + `",
	}`
}

func sendConfirmationEmail(from ContactDetails, to ContactDetails, reservation Reservation, subject string, templateId string) {

	sendGridJson := buildJson(from, to, reservation, subject, templateId)

	apiKey := os.Getenv("API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", host)
	request.Method = "POST"
	request.Body = []byte(sendGridJson)

	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func handler(ctx context.Context, snsEvent events.SNSEvent) {

	customerTemplateId := os.Getenv("CUSTOMER_TEMPLATE_ID")
	adminTemplateId := os.Getenv("ADMIN_TEMPLATE_ID")
	administratorName := os.Getenv("ADMIN_NAME")
	administratorEmail := os.Getenv("ADMIN_EMAIL")

	for _, r := range snsEvent.Records {
		reservation := Reservation{}
		err := json.Unmarshal([]byte(r.SNS.Message), &reservation)

		if err != nil {
			panic(err)
		}

		adminDetails := ContactDetails {
			Name: administratorName,
			Email: administratorEmail,
		}

		customerDetails := ContactDetails {
			Name: reservation.Name,
			Email: reservation.Email,
		}

		customerSubject := "Your recent reservation at Sto Psito Restaurant"
		administratorSubject := "New Reservation - " + strconv.Itoa(int(reservation.Timestamp))

		sendConfirmationEmail(adminDetails, customerDetails, reservation, customerSubject, customerTemplateId)
		sendConfirmationEmail(customerDetails, adminDetails, reservation, administratorSubject, adminTemplateId)
	}
}

func main() {
	lambda.Start(handler)
}
