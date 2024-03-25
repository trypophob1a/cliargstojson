package cliargstojson

import (
	"log/slog"
	"strconv"
)

func Sum(a, b int) int {
	if a < 0 || b < 0 {
		aStr := strconv.Itoa(a)
		bStr := strconv.Itoa(b)
		slog.Info("a = %v and b = %v must be non-negative", aStr, bStr)
	}
	return a + b
}
