package utils

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

const (
	Ip2RegionSearchMethodMemory = 1
	Ip2RegionSearchMethodBinary = 2
	Ip2RegionSearchMethodBtree  = 3
)

func IpToRegion(ip string, method int) (*ip2region.IpInfo, error) {
	region, err := ip2region.New("./infrastructure/utils/ip2region.db")
	if err != nil {
		return nil, err
	}
	var ipInfo ip2region.IpInfo
	switch method {
	case Ip2RegionSearchMethodMemory:
		ipInfo, err = region.MemorySearch(ip)
		if err != nil {
			return nil, err
		}
	case Ip2RegionSearchMethodBinary:
		ipInfo, err = region.BinarySearch(ip)
		if err != nil {
			return nil, err
		}
	case Ip2RegionSearchMethodBtree:
		ipInfo, err = region.BtreeSearch(ip)
		if err != nil {
			return nil, err
		}
	default:
		ipInfo, err = region.BtreeSearch(ip)
		if err != nil {
			return nil, err
		}
	}
	defer region.Close()
	return &ipInfo, nil
}
