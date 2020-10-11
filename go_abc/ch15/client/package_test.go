package client

import (
	"ch15/series"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(series.GetFibSeries(5))
}
