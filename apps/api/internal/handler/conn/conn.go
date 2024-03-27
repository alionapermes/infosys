package conn

import (
	"context"
	"fmt"
	"infosys-api/internal/handler"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type Request struct {
	DbName string `json:"dbname"`
	DbUser string `json:"dbuser"`
	DbPass string `json:"dbpass"`
	DbPort int    `json:"dbport"`
}

type Response struct {
	IsConnected bool `json:"is_connected"`
}

func Handler(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return handler.ResponseInvalidParams(ctx, err)
	}

	if handler.PgConn == nil || handler.PgConn.IsClosed() {
		connString := fmt.Sprintf(
			"postgres://%s:%s@db:%d/%s",
			req.DbUser,
			req.DbPass,
			req.DbPort,
			req.DbName,
		)

		return openConn(ctx, connString)
	}

	return closeConn(ctx)
}

func openConn(ctx echo.Context, connString string) error {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil
	}

	handler.PgConn = conn

	return handler.ResponseOK(ctx, Response{IsConnected: true})
}

func closeConn(ctx echo.Context) error {
	handler.PgConn.Close(context.Background())

	return handler.ResponseOK(ctx, Response{IsConnected: false})
}
