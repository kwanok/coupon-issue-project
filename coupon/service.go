package coupon

import (
	"context"
	"github.com/kwanok/coupon-generate-study/db"
	"log"
)

func SetCoupon(coupon Coupon) {
	db.Client.RPush(db.Ctx, "coupons", coupon.Code)
}

func PopCoupon(ctx context.Context) string {
	coupon, err := db.Client.LPop(ctx, "coupons").Result()
	if err != nil {
		log.Fatalf("failed to pop coupon: %v", err)
	}

	return coupon
}
