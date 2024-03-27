package add_track

import (
	"context"
	"infosys-api/internal/handler"
	"log"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Title string `json:"title"`
	Album string `json:"album"`
	Band  string `json:"band"`
}

type Response struct {
	ID int `json:"id"`
}

func Handler(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return handler.ResponseInvalidParams(ctx, err)
	}

	if handler.PgConn == nil {
		return nil
	}

	row := handler.PgConn.QueryRow(
		context.Background(),
		`insert into "tracks"("title", "album", "band") values($1, $2, $3) returning "id"`,
		req.Title,
		req.Album,
		req.Band,
	)

	var res Response
	if err := row.Scan(&res.ID); err != nil {
		log.Printf("add_track: failed to insert %s", err)
		return nil
	}

	return handler.ResponseOK(ctx, res)
}
