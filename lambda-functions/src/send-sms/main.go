package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
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

func formatTimestamp(millis int64) string {
	unixTime := time.Unix(0, millis * int64(time.Millisecond))
	return unixTime.Format(time.RFC850)
}

func sendConfirmationSms(reservation Reservation) {
	timestamp := formatTimestamp(reservation.Timestamp)

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_ACCOUNT_AUTH_TOKEN")
	from := os.Getenv("TWILIO_ACCOUNT_FROM")
	to := os.Getenv("TWILIO_ACCOUNT_TO")
	body := fmt.Sprintf("%s - %d άτομα στο όνομα %s", timestamp, reservation.Party, reservation.Name)

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}

func handler(ctx context.Context, snsEvent events.SNSEvent)  {
	for _, r := range snsEvent.Records{
		reservation := Reservation{}
		err := json.Unmarshal([]byte(r.SNS.Message), &reservation)
		if err != nil {
			panic(err)
		}
		sendConfirmationSms(reservation)
	}
}

func main() {
	lambda.Start(handler)
}