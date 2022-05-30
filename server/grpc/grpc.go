package grpc

import (
	"context"
	"fmt"
	"net"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/v1"
	"github.com/blixenkrone/lea/storage"

	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	log         logrus.FieldLogger
	srv         *grpc.Server
	learningsDB storage.LearningsStore
}

func NewServer(log *logrus.Entry, ldb storage.LearningsStore) server {
	unaryMw := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(log),
	))
	g := grpc.NewServer(unaryMw)
	srv := server{log, g, ldb}
	learningsv1.RegisterLearningsServiceServer(g, srv)
	return srv
}

func (s server) Serve(lis net.Listener) error {
	return s.srv.Serve(lis)
}

func (s server) GracefulStop() {
	s.srv.GracefulStop()
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

func (s server) Ping(ctx context.Context, empty *emptypb.Empty) (*learningsv1.PingResponse, error) {
	return &learningsv1.PingResponse{
		Message: "Pong!",
	}, nil
}
