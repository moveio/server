package main

import (
	"log"
	"net"

	mgo "gopkg.in/mgo.v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"fmt"
	"github.com/moveio/server/services"
	proto "github.com/moveio/server/moveio/protobuf"
)

const (
	port = ":50051"
)

type Server struct {
}

const (
	user = "moveio"
	pass = "moveioteam"
)

func main() {
	fmt.Println("starting server")

	session, err := mgo.Dial("mongodb://" + user + ":" + pass + "@ds155150.mlab.com:55150/moveio")

	//defer session.Close()
	fmt.Print("connected to db")

	db := session.DB("moveio")

	/*db.C("pipelines").RemoveAll(nil)
	db.C("gesture").RemoveAll(nil)*/


	pipelinesCol := db.C("pipelines")
	gestureCol := db.C("gesture")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterPipelinesServServer(s, &services.PipelineServer{PipelineCol: pipelinesCol, GestureCol: gestureCol})
	proto.RegisterGesturesServServer(s, &services.GestureServer{PipelineCol: pipelinesCol, GestureCol: gestureCol})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
