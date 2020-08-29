package api

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// gRPC server
type Server struct {

}

func (s *Server) SayDate(ctx context.Context, input *Date) (*Date, error){
	logrus.Infof("Received message %s", input.CurrDate)
	return &Date{CurrDate: "29"}, nil
}