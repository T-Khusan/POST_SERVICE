package postgres

import (
	"context"
	"fmt"
	"post_service/pkg/helper"

	pb "post_service/genproto/post_service"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/saidamir98/udevs_pkg/logger"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) PostI {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Reload(ctx context.Context, req *pb.PostReloadReq) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				logger.Error(err)
			}
		} else {
			err = tx.Commit()
			if err != nil {
				logger.Error(err)
			}
		}
	}()

	queryDel := `DELETE FROM posts`

	_, err = tx.Exec(queryDel)
	if err != nil {
		return fmt.Errorf("del post %s", err.Error())
	}

	query := `INSERT INTO posts (id, user_id, title, body) VALUES `

	values := []interface{}{}

	for _, data := range req.Data {
		query += "(?, ?, ?, ?),"
		values = append(values, data.Id, data.UserId, data.Title, data.Body)
		fmt.Println("VALUES", values)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	_, err = tx.Exec(query, values...)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("insert post err: %v", err)
	}

	return nil
}
