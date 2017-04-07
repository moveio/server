package services

import (
	proto "github.com/moveio/server/moveio/protobuf"
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
)

type GestureServer struct {
	PipelineCol *mgo.Collection
	GestureCol *mgo.Collection
}


func (s * GestureServer) CreateGesture (ctx context.Context, in *proto.Gesture) (*proto.Gesture, error)  {
	return in, nil
}

func (s * GestureServer) GetGesture (ctx context.Context, in *proto.Gesture) (*proto.Gesture, error)  {
	return in, nil
}
