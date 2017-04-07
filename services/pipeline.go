package services

import (
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	proto "github.com/moveio/server/moveio/protobuf"
)

type PipelineServer struct {
	PipelineCol *mgo.Collection
	GestureCol *mgo.Collection
}

func (s * PipelineServer) CreatePipeline (ctx context.Context, in *proto.Pipeline) (*proto.Pipeline, error)  {
	return in, nil
}

func (s * PipelineServer) GetPipeline (ctx context.Context, in *proto.Pipeline) (*proto.Pipeline, error)  {
	return in, nil
}
