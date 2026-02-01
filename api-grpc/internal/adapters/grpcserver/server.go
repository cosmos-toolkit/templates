package grpcserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	entityv1 "github.com/your-org/your-app/api/pb/entity/v1"
	"github.com/your-org/your-app/internal/app"
	"github.com/your-org/your-app/internal/domain"
)

// Server implements entityv1.EntityServiceServer and delegates to the application service.
type Server struct {
	entityv1.UnimplementedEntityServiceServer
	app *app.App
}

// NewServer creates a gRPC server adapter.
func NewServer(a *app.App) *Server {
	return &Server{app: a}
}

// GetEntity returns an entity by ID.
func (s *Server) GetEntity(ctx context.Context, req *entityv1.GetEntityRequest) (*entityv1.Entity, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id required")
	}
	entity, err := s.app.Service.GetEntity(ctx, req.Id)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return domainEntityToProto(entity), nil
}

// CreateEntity creates an entity.
func (s *Server) CreateEntity(ctx context.Context, req *entityv1.CreateEntityRequest) (*entityv1.Entity, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id required")
	}
	e := &domain.Entity{ID: req.Id}
	if err := s.app.Service.CreateEntity(ctx, e); err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return domainEntityToProto(e), nil
}

func domainEntityToProto(e *domain.Entity) *entityv1.Entity {
	return &entityv1.Entity{
		Id:        e.ID,
		CreatedAt: timestamppb.New(e.CreatedAt),
		UpdatedAt: timestamppb.New(e.UpdatedAt),
	}
}
