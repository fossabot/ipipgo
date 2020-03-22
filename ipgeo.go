package ipipgo

import (
	"encoding/json"
	"errors"
	"net"
	"strings"
)

var (
	ErrInvalidIP = errors.New("invalid IP address")
	ErrDecode    = errors.New("json decode failed")
)

const (
	responseLen = 5
)

type IPGeo struct {
	IP *net.IP

	Country  string
	Province string
	City     string

	ISP string
}

func (geo *IPGeo) String() string {
	var ls []string
	if geo.Country != "" {
		ls = append(ls, geo.Country)
	}
	if geo.Province != "" {
		ls = append(ls, geo.Province)
	}
	if geo.City != "" {
		ls = append(ls, geo.City)
	}
	if geo.ISP != "" {
		ls = append(ls, geo.ISP)
	}
	return strings.Join(ls, "，")
}

func GetGeo(ipStr string) (*IPGeo, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, ErrInvalidIP
	}
	res, err := httpGet("http://freeapi.ipip.net/" + ipStr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp := make([]string, responseLen)
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, ErrDecode
	}
	if len(resp) != responseLen {
		return nil, ErrDecode
	}
	return &IPGeo{
		IP:       &ip,
		Country:  resp[0],
		Province: resp[1],
		City:     resp[2],
		ISP:      resp[4],
	}, nil
}
