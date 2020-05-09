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

func init() {
	log.Println("[Firebase]", "Init cloud messaging service")
	var err error

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	file := filepath.Join(filepath.Dir(basepath), "firebase-admin-key.json")
	_, ferr := os.Stat(file)

	if os.IsNotExist(ferr) {
		log.Println("[Firebase]", "Cannot found the key:", file)
	}

	opt := option.WithCredentialsFile(file)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println("[Firebase]", err.Error())
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
