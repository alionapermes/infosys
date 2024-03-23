package handler

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

const databaseURL = "postgres://postgres:postgres@db:5432/postgres"

var PgConn *pgx.Conn

type Result struct {
	IsConnected bool `json:"is_connected"`
}

type Response struct {
	Result Result `json:"result"`
}

func ConnectDb(ctx echo.Context) error {
	if PgConn == nil {
		// urlExample := "postgres://username:password@localhost:5432/database_name"
		conn, err := pgx.Connect(context.Background(), databaseURL)
		if err != nil {
			return nil
		}

		ctx.JSON(http.StatusOK, Response{
			Result: Result{IsConnected: true},
		})
		PgConn = conn
	} else {
		PgConn.Close(context.Background())
		ctx.JSON(http.StatusOK, Response{
			Result: Result{IsConnected: false},
		})
	}

	return nil
}
