package coupon

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestSetCoupon(t *testing.T) {
	SetCoupon(Coupon{
		ID:   0,
		Code: "Test1",
	})

	SetCoupon(Coupon{
		ID:   0,
		Code: "Test2",
	})
}

func TestPopCoupon(t *testing.T) {
	coupon := PopCoupon(context.Background())
	if coupon != "Test1" {
		t.Errorf("PopCoupon() = %s; want Test1", coupon)
	}

	coupon = PopCoupon(context.Background())
	if coupon != "Test2" {
		t.Errorf("PopCoupon() = %s; want Test2", coupon)
	}
}

func TestSet100000Coupons(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go func(i int) {
			defer wg.Done()
			SetCoupon(Coupon{
				ID:   0,
				Code: fmt.Sprintf("Test_%d", i),
			})
		}(i)
	}

	wg.Wait()
}

func TestPop100000Coupons(t *testing.T) {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	wg := sync.WaitGroup{}
	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go func() {
			defer wg.Done()
			PopCoupon(context.Background())
		}()
	}

	wg.Wait()
}
