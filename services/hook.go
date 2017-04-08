package services

import (
	proto "github.com/moveio/server/moveio/protobuf"
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

type HookServer struct {
	PipelineCol *mgo.Collection
	GestureCol  *mgo.Collection
	HookCol     *mgo.Collection
}

func (s *HookServer) CreateHook(ctx context.Context, in *proto.Hook) (*proto.Hook, error) {
	in.Id = bson.NewObjectId().Hex()

	err := s.HookCol.Insert(in)
	if err != nil {
		return nil, err
	}
	return in, nil

}

func (s *HookServer) GetHook(ctx context.Context, in *proto.Hook) (*proto.Hook, error) {
	gesture := proto.Hook{}
	err := s.HookCol.Find(bson.M{"id": in.Id}).One(&gesture)
	if err != nil {
		return nil, err
	}
	return &gesture, nil
}

func (s *HookServer) PutHook(ctx context.Context, in *proto.Hook) (*proto.Hook, error) {
	err := s.HookCol.Update(bson.M{"id": in.Id}, bson.M{"$set": &in})
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *HookServer) GetAllHook(ctx context.Context, in *proto.Hook) (*proto.Hooks, error) {
	gestures := proto.Hooks{}
	err := s.HookCol.Find(bson.M{"userid": in.UserId}).All(&gestures.Hooks)
	if err != nil {
		return nil, err
	}
	return &gestures, nil
}


