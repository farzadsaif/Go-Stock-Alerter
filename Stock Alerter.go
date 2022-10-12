package main

import (
	"bufio"
	"fmt"
	"math"
	"net/smtp"
	"os"
	"strings"
	"syscall"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/piquette/finance-go/quote"
	"golang.org/x/term"
)

// speech variable required by speech package
var speech = htgotts.Speech{Folder: "audio", Language: voices.English}

func main() {

	// Input variables
	// Email variables are set in the credentials function because password requires secure input

	var stock_threshold float64
	var ticker string

	var toEmailAddress string
	var to_text string

	// Get variables from user
	fmt.Println("Enter Stock Threshold:")
	fmt.Scanln(&stock_threshold)
	fmt.Println("Enter Stock Ticker:")
	fmt.Scanln(&ticker)
	from, password, err := credentials()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("Enter to Email:")
	fmt.Scanln(&toEmailAddress)
	toEmail := []string{toEmailAddress}

	fmt.Println("Enter SMS Number @ SMS gateway:")
	fmt.Scanln(&to_text)
	toSMS := []string{to_text}

	//Get quote from ticker

	q, err := quote.Get(ticker)
	if err != nil {
		panic(err)
	}
	cp := q.RegularMarketChangePercent

	// The actual stock alerter

	fmt.Print(cp, "%\n")
	for {
		if math.Abs(cp) > math.Abs(stock_threshold) {
			speak(cp, stock_threshold)
			email(from, password, toEmail, cp, stock_threshold)
			SMS(from, password, toSMS, cp, stock_threshold, email)
		}
		time.Sleep(30 * time.Minute)
	}

}

// speak() provides stock alerts via computer speakers

func speak(x, y float64) {
	if x < y {
		s := "Buy now"
		speech.Speak(s)
	} else {
		s := "Sell now"
		speech.Speak(s)
	}
}

//email() uses SMTP protocol to send emails
//Currently this is set to use Gmail, which is only possible with a Gmail buisness account with less-secure app access turned on
//To use a different email provider, change the host and port

func email(from, password string, to []string, x, y float64) {
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	var subject string
	var body string
	if x < y {
		subject = "Stock Alert – Buy Now\n"
		body = "Buy now"
	} else {
		subject = "Stock Alert – Sell now\n"
		body = "Sell now"
	}

	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}

//SMS uses SMS gateway to send texts via the given email address
//This is done via callback

func SMS(from, password string, toSMS []string, x, y float64, f func(from, password string, to []string, x, y float64)) {
	f(from, password, toSMS, x, y)
}

//Credentials function takes in email and password in a secure manner
//I can move the rest of the input variables here; code is a bit split up for now. Functionality wise it makes no difference

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter From Email: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter From Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
