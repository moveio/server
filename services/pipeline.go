package services

import (
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	proto "github.com/moveio/server/moveio/protobuf"
	"gopkg.in/mgo.v2/bson"
)

type PipelineServer struct {
	PipelineCol *mgo.Collection
	GestureCol  *mgo.Collection
	HookCol     *mgo.Collection
}

func (s *PipelineServer) CreatePipeline(ctx context.Context, in *proto.Pipeline) (*proto.Pipeline, error) {
	in.Id = bson.NewObjectId().Hex()

	err := s.PipelineCol.Insert(in)
	if err != nil {
		return nil, err
	}
	return in, nil

}

func (s *PipelineServer) GetPipeline(ctx context.Context, in *proto.Pipeline) (*proto.Pipeline, error) {
	pipeline := proto.Pipeline{}
	err := s.PipelineCol.Find(bson.M{"id": in.Id}).One(&pipeline)
	if err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (s *PipelineServer) GetAllPipeline(ctx context.Context, in *proto.Pipeline) (*proto.Pipelines, error) {
	pipelines := proto.Pipelines{}
	err := s.PipelineCol.Find(bson.M{}).All(&pipelines.Pipelines)
	if err != nil {
		return nil, err
	}
	return &pipelines, nil
}

func (s *PipelineServer) DeletePipeline(ctx context.Context, in *proto.Pipeline) (*proto.Pipeline, error) {
	err := s.PipelineCol.Remove(bson.M{"id": in.Id})
	if err != nil {
		return nil, err
	}
	return &proto.Pipeline{}, nil
}

func (s *PipelineServer) PutPipeline(ctx context.Context, in *proto.Pipeline) (*proto.Pipeline, error) {

	err := s.PipelineCol.Update(bson.M{"id": in.Id}, bson.M{"$set": &in})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
