package services

import (
	"gopkg.in/mgo.v2"
	"golang.org/x/net/context"
	proto "github.com/moveio/server/moveio/protobuf"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"net/http"
	"bytes"
)

type RequestGestureServer struct {
	PipelineCol *mgo.Collection
	GestureCol  *mgo.Collection
	HookCol     *mgo.Collection
}

func (s *RequestGestureServer) PostGesture(ctx context.Context, in *proto.GestureMessage) (*proto.GestureMessage, error) {
	gesture := proto.Gesture{}
	err := s.GestureCol.Find(bson.M{"meta": in.Meta, "userid": in.UserId}).One(&gesture)
	if err != nil {
		return nil, err
	}

	pipelane := proto.Pipeline{}
	err = s.PipelineCol.Find(bson.M{"gestureid": gesture.Id, "userid": in.UserId}).One(&pipelane)
	if err != nil {
		return nil, err
	}

	fmt.Println(pipelane)

	hook := proto.Hook{}
	err = s.HookCol.Find(bson.M{"id": pipelane.HookId, "userid": in.UserId}).One(&hook)
	if err != nil {
		return nil, err
	}

	fmt.Println(hook)

	res, err := http.Post("http://"+hook.Address, "", bytes.NewBufferString(hook.Message))

	fmt.Println(res)

	return in, err
}
