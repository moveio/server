package services

import (
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	proto "github.com/moveio/server/moveio/protobuf"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type RequestGestureServer struct {
	PipelineCol *mgo.Collection
	GestureCol  *mgo.Collection
}

func (s *RequestGestureServer) PostGesture(ctx context.Context, in *proto.GestureMessage) (*proto.GestureMessage, error) {
	gesture := proto.Gesture{}
	err := s.GestureCol.Find(bson.M{"meta": in.Meta, "userid": in.UserId}).One(&gesture)
	if err != nil {
		return nil, err
	}

	fmt.Println(gesture)

	return in, nil
}
