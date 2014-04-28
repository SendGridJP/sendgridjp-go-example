package main

import (
    "fmt"
    "github.com/sendgrid/sendgrid-go"
)

func main() {
    sg := sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")
    message := sendgrid.NewMail()
    message.AddTo("yamil@sendgrid.com")
    message.AddToName("Yamil Asusta")
    message.AddSubject("SendGrid Testing")
    message.AddText("WIN")
    message.AddFrom("yamil@sendgrid.com")
    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}
