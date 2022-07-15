package grpc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"

	learningsv1 "github.com/blixenkrone/lea/proto/compiled/learnings/v1"
	"github.com/blixenkrone/lea/storage"

	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	learningsv1.RegisterFileUploadServiceServer(g, srv)
	return srv
}

func (s server) Serve(lis net.Listener) error {
	return s.srv.Serve(lis)
}

func (s server) GracefulStop() {
	s.srv.GracefulStop()
}

func (s server) UploadFile(stream learningsv1.FileUploadService_UploadFileServer) error {
	var fileSize int
	ctx := stream.Context()

	for {
		req, err := stream.Recv()
		if !errors.Is(err, io.EOF) {
			if err != nil {
				s.log.Errorf("error receiving data: %v", err)
				return status.Errorf(codes.Unknown, "error received: %w", err)
			}
		}
		if v, ok := req.Request.(*learningsv1.UploadFileRequest_File); ok {
			fileSize += len(v.File.Content)
			br := bytes.NewReader(v.File.Content)
		}

		if v, ok := req.Request.(*learningsv1.UploadFileRequest_Metadata); ok {
			_ = v.Metadata.FileType
			s.learningsDB.GetModule(ctx, uuid.Nil)
		}
	}
}

func (s server) AddCourse(ctx context.Context, req *learningsv1.AddCourseRequest) (*learningsv1.AddCourseResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name must be provided")
	}

	newCourse := learningsv1.Course{
		Id:        uuid.NewString(),
		IsActive:  req.IsActive,
		Name:      req.Name,
		ModuleIds: []string{},
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	c, err := s.learningsDB.AddCourse(ctx, &newCourse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.log.Infof("created course %s with id %s", c.CourseName, c.ID)

	return &learningsv1.AddCourseResponse{}, nil
}

func (s server) ListCourses(req *learningsv1.ListCoursesRequest, stream learningsv1.LearningsService_ListCoursesServer) error {
	stream.Context()
	panic("not implemented") // TODO: Implement
}

func (s server) Ping(ctx context.Context, empty *learningsv1.PingRequest) (*learningsv1.PingResponse, error) {
	return &learningsv1.PingResponse{
		Message: "Pong!",
	}, nil
}
