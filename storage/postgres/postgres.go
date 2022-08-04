package postgres

import (
	"context"
	pb "post_service/genproto/post_service"
)

type PostI interface {
	Reload(ctx context.Context, req *pb.PostReloadReq) error
}
