package cliargstojson

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

func Sum(a, b int) int {
	if a < 0 || b < 0 {
		aStr := strconv.Itoa(a)
		bStr := strconv.Itoa(b)
		logrus.Infof("a: %s, b: %s", aStr, bStr)
	}
	return a + b
}
