package main

import (
	"log"
	"fmt"
	"bytes"
	"net/smtp"
	"html/template"
)


func SendLink (data CommentInfo) error {

	fmt.Println("SendLinkga keldi")
	Sender := "gobootcamp075@gmail.com"
	Password := "1501069Mm"

	receivers := []string {
		data.Email,
	}

	host := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("",Sender, Password, host)

	t, err := template.ParseFiles("main.html")

	if err != nil {
		log.Fatal(err)
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test \n%s\n\n", mimeHeaders)))

	t.Execute(&body, data)
	
	err = smtp.SendMail(host + ":" + port, auth, Sender, receivers, body.Bytes())	
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(data)
	return nil
}