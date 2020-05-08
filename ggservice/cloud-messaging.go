package ggservice

import (
	"context"
	"log"
	"os"
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
	var err error

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	file := filepath.Join(filepath.Dir(basepath), "firebase-admin-key.json")
	info, err := os.Stat(file)
	if os.IsNotExist(err) || info.IsDir() {
		log.Println("[Firebase]", "Cannot found the key:", file)
	}

	opt := option.WithCredentialsFile(file)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println("[Firebase]", err.Error())
	}
}
