package base

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"net/http"
)

var (
	CodeOK      = gcode.New(http.StatusOK, "success", nil)
	CodeUnknown = gcode.New(999, "unknown", nil)
)
