package ggservice

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

//FirebaseKeyPath determine the location where the firebase admin key is stored
var FirebaseKeyPath string
var app *firebase.App

//SendToToken send a message to a specified device
func SendMessageToDevice(message *messaging.Message) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Println("[Firebase]", "Error getting Messaging client: %v\n", err)
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		log.Println("[Firebase]", err)
	}
}

//SendMessageToDevices send a multicast message to specified devices
func SendMessageToDevices(message *messaging.MulticastMessage) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Println("[Firebase]", "Error getting Messaging client: %v\n", err)
	}

	_, err = client.SendMulticast(ctx, message)

	if err != nil {
		log.Println("[Firebase]", err)
	}
}

//Setup prepare the application for working with firebase service.
func Setup() (e error) {

	_, e = os.Stat(FirebaseKeyPath)
	if os.IsNotExist(e) {
		log.Println("[Firebase]", "Cannot found the key:", FirebaseKeyPath)
		return e
	}

	opt := option.WithCredentialsFile(FirebaseKeyPath)
	app, e = firebase.NewApp(context.Background(), nil, opt)
	if e != nil {
		log.Println("[Firebase]", e.Error())
	}

	return
}

//SetupWithPath set the FilebaseKeyPath by specified path then run Setup
func SetupWithPath(path string) (e error) {
	FirebaseKeyPath = path
	return Setup()
}
