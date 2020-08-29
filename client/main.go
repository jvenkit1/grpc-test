package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc-test/api"
)

func main() {
	var conn *grpc.ClientConn
	port := 3000
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to dial server running at port %d", port)
	}
	defer conn.Close()

	c := api.NewGreetClient(conn)

	response, err := c.SayDate(context.Background(), &api.Date{CurrDate: "28"})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to call function SayDate")
	}
	logrus.WithFields(logrus.Fields{
		"Date": response.CurrDate,
	}).Info("Received response from server")
}