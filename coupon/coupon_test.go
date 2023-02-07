package coupon

import (
	"bufio"
	"fmt"
	"github.com/kwanok/coupon-issue-project/rdb"
	"github.com/stretchr/testify/suite"
	"os"
	"sync"
	"testing"
)

type CouponTestSuite struct {
	suite.Suite
	wr *bufio.Writer
}

func (suite *CouponTestSuite) SetupTest() {
	queue.Client = rdb.NewClient("localhost:6379")
	suite.wr = bufio.NewWriter(os.Stdout)
}

func (suite *CouponTestSuite) TearDownTest() {
	suite.wr.Flush()
	queue.Client.Close()
}

func (suite *CouponTestSuite) TestPush() {
	err := queue.push(&Coupon{Code: "test"})
	if err != nil {
		suite.T().Errorf("Error push coupon: %v", err)
	}
}

func (suite *CouponTestSuite) TestPush500000Coupon() {
	cnt := 500000
	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func(i int) {
			err := queue.push(&Coupon{Code: fmt.Sprintf("Test_%d", i)})
			if err != nil {
				suite.T().Errorf("Error push coupon: %v", err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func (suite *CouponTestSuite) TestPop500000Coupons() {
	cnt := 500000

	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			_, err := queue.pop()
			if err != nil {
				suite.T().Errorf("Error pop coupon: %v", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestCouponTestSuite(t *testing.T) {
	suite.Run(t, new(CouponTestSuite))
}
