package handler

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

var PgConn *pgx.Conn

func ResponseOK(ctx echo.Context, res any) error {
	return ctx.JSON(http.StatusOK, res)
}

func ResponseConnRequired(ctx echo.Context) error {
	return ctx.String(http.StatusBadRequest, "db connection required")
}

func ResponseInvalidParams(ctx echo.Context, err error) error {
	msg := fmt.Sprintf("invalid params: %s", err)
	// log.Errorf("inva bind req: %s; body: %+v", err, ctx.Request().Body)

	return ctx.String(http.StatusBadRequest, msg)
}

func ResponseExecutionFailed(ctx echo.Context, format string, args ...any) error {
	msg := fmt.Sprintf(format, args...)

	return ctx.String(http.StatusBadRequest, msg)
}
