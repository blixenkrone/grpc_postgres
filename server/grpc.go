package server

import (
	"context"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	log         logrus.FieldLogger
	srv         *grpc.Server
	learningsDB storage.LearningsStore
}

func New(log logrus.FieldLogger, ldb storage.LearningsStore) server {
	g := grpc.NewServer()
	srv := server{log, g, ldb}
	learningsv1.RegisterLearningsServiceServer(g, srv)
	return srv
}

func (s server) AddCourse(ctx context.Context, req *learningsv1.AddCourseRequest) (*learningsv1.Course, error) {
	return &learningsv1.Course{
		Id:        "",
		IsActive:  false,
		Name:      "",
		Modules:   []*learningsv1.Module{},
		CreatedAt: timestamppb.Now(),
		UpdatedAt: &timestamppb.Timestamp{},
	}, nil
}

func (s server) ListCourses(req *learningsv1.ListCoursesRequest, stream learningsv1.LearningsService_ListCoursesServer) error {
	stream.Context()
	panic("not implemented") // TODO: Implement
}
