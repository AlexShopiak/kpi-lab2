package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (s *MySuite) TestComputeHandler(c *C) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	handler := ComputeHandler{
		Input:  strings.NewReader("* 5 2"),
		Output: buffer,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(buffer.String(), Equals, "10")
}