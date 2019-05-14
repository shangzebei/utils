package lb

import "errors"

type LBPoll struct {
	i int64
}

func (c *LBPoll) SelectOne(s []string) (string, error) {
	c.i++
	if c.i > 9223372036854700000 {
		c.i = 0
	}
	if len(s) != 0 {
		return s[c.i%int64(len(s))], nil
	} else {
		return "", errors.New("not servers to router")
	}
}
