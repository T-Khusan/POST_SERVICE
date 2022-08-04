package service

import (
	"context"
	"fmt"
	"post_service/pkg/helper"
	"post_service/storage"

	pb "post_service/genproto/post_service"

	"post_service/pkg/logger"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type postService struct {
	logger  logger.Logger
	storage storage.StorageI
	pb.UnimplementedPostServiceServer
}

func NewPostService(db *sqlx.DB, log logger.Logger) *postService {
	return &postService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *postService) Reload(ctx context.Context, req *pb.PostReloadReq) (*emptypb.Empty, error) {
	fmt.Println("SERVICE ---->>>")
	s.logger.Info("---Reload Posts--->", logger.Any("req", req))

	err := s.storage.Post().Reload(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "Error reloading posts", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
