package delete_track

import (
	"context"
	"infosys-api/internal/handler"

	"github.com/labstack/echo/v4"
)

type Request struct {
	ID int `json:"id"`
}

type Response struct {
	OK bool `json:"ok"`
}

func Handler(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return handler.ResponseInvalidParams(ctx, err)
	}

	if handler.PgConn == nil {
		return handler.ResponseConnRequired(ctx)
	}

	_, err := handler.PgConn.Exec(context.Background(), `delete from "tracks" where "id" = $1`, req.ID)

	return handler.ResponseOK(ctx, Response{OK: err == nil})
}
