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

var stock_threshold float64
var ticker string
var speech = htgotts.Speech{Folder: "audio", Language: voices.English}

var toEmailAddress string
var to_text string

func main() {

	// Get variables from user
	fmt.Println("Enter stock threshold:")
	fmt.Scanln(&stock_threshold)
	fmt.Println("Enter stock ticker:")
	fmt.Scanln(&ticker)
	from, password, err := credentials()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("Enter to email:")
	fmt.Scanln(&toEmailAddress)
	toEmail := []string{toEmailAddress}

	fmt.Println("Enter SMS number @ SMS gateway:")
	fmt.Scanln(&to_text)
	toSMS := []string{to_text}

	q, err := quote.Get(ticker)
	if err != nil {
		panic(err)
	}
	cp := q.RegularMarketChangePercent

	fmt.Print(cp, "%\n")
	for true {
		if math.Abs(cp) > math.Abs(stock_threshold) {
			speak(cp, stock_threshold)
			email(from, password, toEmail, cp, stock_threshold)
			SMS(from, password, toSMS, cp, stock_threshold, email)
		}
		time.Sleep(30 * time.Minute)
	}

}

func speak(x, y float64) {
	if x < y {
		s := "Buy now"
		speech.Speak(s)
	} else {
		s := "Sell now"
		speech.Speak(s)
	}
}

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

func SMS(from, password string, toSMS []string, x, y float64, f func(from, password string, to []string, x, y float64)) {
	f(from, password, toSMS, x, y)
}

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
