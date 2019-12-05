package common

import "testing"
import "github.com/google/uuid"
import (
	"crypto/rand"
	"math/big"
	"github.com/kataras/iris/core/errors"
	"time"
	"fmt"
)

func TestRetry(t *testing.T) {
	var s string
	var err error
	f := func() (error) {
		s, err = ff()
		return err
	}
	err = Retry(f, 3, 10*time.Millisecond)
	fmt.Println(s, err)
}

func ff() (string, error) {
	n, _ := rand.Int(rand.Reader, big.NewInt(10))
	if n.Int64() < 6 {
		return "error", errors.New("fail")
	}
	return uuid.New().String(), nil
}
