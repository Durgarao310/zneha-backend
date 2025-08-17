package errors_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	commonErrors "github.com/Durgarao310/zneha-backend/internal/common/errors"
	appErrors "github.com/Durgarao310/zneha-backend/internal/errors"
	"github.com/Durgarao310/zneha-backend/utils"
	"github.com/gin-gonic/gin"
)

type createReq struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
}

func TestValidationErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(commonErrors.GlobalErrorHandler())
	r.POST("/test", func(c *gin.Context) {
		var req createReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err) // pass to middleware
			return
		}
		utils.SuccessResponse(c, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	// Missing both required fields triggers validation errors
	body := bytes.NewBufferString(`{"name":"ab","email":"not-an-email"}`)
	req := httptest.NewRequest(http.MethodPost, "/test", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		toFail(t, "expected status 400 got %d", w.Code)
	}
}

func TestAppError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(commonErrors.GlobalErrorHandler())
	r.GET("/app", func(c *gin.Context) {
		c.Error(appErrors.BadRequest("bad stuff", errors.New("root")))
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/app", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		toFail(t, "expected status 400 got %d", w.Code)
	}
}

func TestPanicRecovery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(commonErrors.GlobalErrorHandler())
	r.GET("/panic", func(c *gin.Context) {
		panic("boom")
	})
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		toFail(t, "expected status 500 got %d", w.Code)
	}
}

func toFail(t *testing.T, format string, args ...interface{}) {
	if t != nil {
		// Use Errorf to continue test execution contextually
		// but in these tests we can just fail
		// Use Fatalf for simplicity
		//nolint
		t.Fatalf(format, args...)
	}
}
