package ggservice

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

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

func init() {
	log.Println("[Firebase]", "Init cloud messaging service")
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	var err error
	opt := option.WithCredentialsFile(filepath.Join(filepath.Dir(basepath), "firebase-admin-key.json"))
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println("[Firebase]", err.Error())
	}
}
