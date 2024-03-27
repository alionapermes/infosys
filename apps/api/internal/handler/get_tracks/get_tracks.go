package get_tracks

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"

	"infosys-api/internal/handler"
)

type track struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Album string `json:"album"`
	Band  string `json:"band"`
}

type Response struct {
	Tracks []track `json:"tracks"`
}

func Handler(ctx echo.Context) error {
	rows, err := handler.PgConn.Query(
		context.Background(),
		`select * from "tracks"`,
	)
	if err != nil {
		return handler.ResponseExecutionFailed(ctx, "failed to select tracks: %s", err)
	}

	tracks, err := pgx.CollectRows(rows, pgx.RowToStructByName[track])
	if err != nil {
		return handler.ResponseExecutionFailed(ctx, "failed to parse tracks: %s", err)
	}

	return handler.ResponseOK(ctx, Response{Tracks: tracks})
}
