package services

import (
	proto "github.com/moveio/server/moveio/protobuf"
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

type GestureServer struct {
	PipelineCol *mgo.Collection
	GestureCol  *mgo.Collection
}

func (s *GestureServer) CreateGesture(ctx context.Context, in *proto.Gesture) (*proto.Gesture, error) {
	// insert to ProjectsCol
	err := s.GestureCol.Insert(in)
	if err != nil {
		return nil, err
	}
	return in, nil

}

func (s *GestureServer) GetGesture(ctx context.Context, in *proto.Gesture) (*proto.Gesture, error) {
	gesture := proto.Gesture{}
	err := s.GestureCol.Find(bson.M{"id": in.Id}).One(&gesture)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *GestureServer) GetAllGesture(ctx context.Context, in *proto.Gesture) (*proto.Gesture, error) {
	gestures := []proto.Gesture{}
	err := s.GestureCol.Find(bson.M{}).All(&gestures)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *GestureServer) DeleteGesture(ctx context.Context, in *proto.Gesture) (*proto.Gesture, error) {
	err := s.GestureCol.Remove(bson.M{"id": in.Id})
	if err != nil {
		return nil, err
	}
	return &proto.Gesture{}, nil
}
