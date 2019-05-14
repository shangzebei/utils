package lb

import (
	"errors"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type LbRand struct {
}

func init() {

	rand.Seed(time.Now().UnixNano())
}

func (c *LbRand) SelectOne(s []string) (string, error) {
	i := len(s)
	if i != 0 {
		intn := rand.Intn(len(s))
		logrus.Debugf("lb select %s", s[intn])
		return s[intn], nil
	} else {
		logrus.Warnf("not servers to router")
		return "", errors.New("not servers to router")
	}

}
