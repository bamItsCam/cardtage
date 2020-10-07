package cardtage

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func round2Uint(inFloat float64) uint {
	return uint(math.Round(inFloat))
}

func Res2WxH(inRes string) (float64, float64, error) {
	s := strings.Split(inRes, "x")
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("resolution '%s' is not a valid format, must be 'WxH'", inRes)
	}
	w, err := strconv.ParseFloat(s[0], 64)
	if err != nil {
		return 0, 0, fmt. Errorf("couldn't convert %s to an int", s[0])
	}
	h, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return 0, 0, fmt. Errorf("couldn't convert %s to an int", s[1])
	}
	if w < 0 || h < 0 {
		return 0, 0, fmt.Errorf("width or height was negative - don't do that")
	}
	return w, h, nil
}
