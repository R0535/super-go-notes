package main

import "fmt"

//Interfaces
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender //Regresa otra interfaz
}
type ISender interface { //Lo que regresa como interfaz
	GetSenderMethod() string
	GetSenderChannel() string
}

//Structs to implement Interfaces
type SMSNotification struct {
}

//Receiver Funcs as Notifications Senders
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via sms")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

//Nested Struct as Sender
type SMSNotificationSender struct{}

//Receiver Funcs as Senders
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct{}
type EmailNotificationSender struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Envinado por correo")
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Por Correo"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Gmail"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	} else if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No notification Type")
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)
}
