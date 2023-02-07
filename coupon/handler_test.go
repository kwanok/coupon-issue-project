package coupon

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kwanok/coupon-issue-project/rdb"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"

	"testing"
)

func TestSendHandler(t *testing.T) {
	app := fiber.New()
	Init(app, rdb.NewClient("localhost:6379"))

	request := httptest.NewRequest("GET", "/coupons", nil)
	resp, err := app.Test(request)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}
