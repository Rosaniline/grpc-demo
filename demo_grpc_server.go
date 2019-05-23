package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/Rosaniline/grpc-demo/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ppl = []proto.Person{
	{
		Name: "唐伯虎",
		Id: 9527,
		Email: "9527@huafu.com",
	},
	{
		Name: "祝枝山",
		Id: 9487,
		Email: "9487@huafu.com",
	},
}

type server struct{

}

func (s *server) QueryPerson(srv proto.HuaFu_QueryPersonServer) error {

	for {
		req, err := srv.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		found := false

		for _, p := range ppl {
			if req.GetId() == p.Id {
				found = true
				srv.Send(&p)
			}
		}


		if !found {
			return status.Errorf(codes.NotFound, "ppl with id %v not found", req.Id)
		}
	}

	return nil
}

func (s *server) ListPeople(req *empty.Empty, srv proto.HuaFu_ListPeopleServer) error {
	for _, p := range ppl {
		time.Sleep(1*time.Second)
		if err := srv.Send(&p); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetPerson(ctx context.Context, req *proto.GetPersonRequest) (*proto.Person, error) {
	for _, p := range ppl {
		if p.Id == req.GetId() {
			return &p, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "ppl with id %v not found", req.Id)
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	s := grpc.NewServer()
	proto.RegisterHuaFuServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}