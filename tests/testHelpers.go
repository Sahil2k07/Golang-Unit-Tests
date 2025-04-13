package tests

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
)

func SetupContext(method, path, body string) echo.Context {
	e := echo.New()

	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}
