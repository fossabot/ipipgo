package ipipgo

import (
	"fmt"
	"testing"
)

func TestGetGeo(t *testing.T) {
	geo, err := GetGeo("60.221.218.191")
	if err != nil {
		t.FailNow()
	}
	if geo.Country == "" {
		t.FailNow()
	}
	if geo.Province == "" {
		t.FailNow()
	}
	if geo.City == "" {
		t.FailNow()
	}
	if geo.ISP == "" {
		t.FailNow()
	}
	s := fmt.Sprintf("%v", geo)
	if s == "" {
		t.FailNow()
	}
}
