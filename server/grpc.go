package server

import (
	"context"
	"fmt"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage"
	"github.com/google/uuid"
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
	newCourse := learningsv1.Course{
		Id:        uuid.New().String(),
		IsActive:  req.IsActive,
		Name:      req.Name,
		ModuleIds: nil,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	c, err := s.learningsDB.AddCourse(ctx, &newCourse)
	if err != nil {
		return nil, fmt.Errorf("error adding course: %err", err)
	}

	return &learningsv1.Course{
		Id:        c.ID.String(),
		IsActive:  c.IsActive,
		Name:      c.CourseName,
		ModuleIds: nil,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}, nil
}

func (s server) ListCourses(req *learningsv1.ListCoursesRequest, stream learningsv1.LearningsService_ListCoursesServer) error {
	stream.Context()
	panic("not implemented") // TODO: Implement
}
