package services

import (
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	proto "github.com/moveio/server/moveio/protobuf"
)

type RequestGestureServer struct {
	PipelineCol *mgo.Collection
	GestureCol  *mgo.Collection
}


func (s *RequestGestureServer) PostGesture(ctx context.Context, in *proto.GestureMessage) (*proto.GestureMessage, error) {


	return in, nil
}