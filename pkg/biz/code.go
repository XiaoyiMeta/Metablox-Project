package biz

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"net/http"
)

var (
	CodeOK               = gcode.New(http.StatusOK, "success", nil)
	CodeError            = gcode.New(http.StatusInternalServerError, "error", "system error, please contact administrator")
	CodeInvalidParameter = gcode.CodeInvalidParameter
	CodeBadRequest       = gcode.New(http.StatusBadRequest, "bad request", nil)
	CodeUnauthorized     = gcode.New(http.StatusUnauthorized, "unauthorized", nil)
	CodeForbidden        = gcode.New(http.StatusForbidden, "Forbidden", nil)
	CodeNotFound         = gcode.New(http.StatusNotFound, "not found", nil)
)
