package log

import (
	"testing"
)

func TestInfow(t *testing.T) {
	Infow("infow", "name", "xiaowang", "age", 13)
	Info("info", "name", "xiaowa", "age", 18)
}

/*
https://github.com/uber-go/zap/issues/604
*/
func TestConcurrentSafe(t *testing.T) {

}
