package service

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"kingsmith.com.cn/ks-horizon/infrastructure/persistence"
	"kingsmith.com.cn/ks-horizon/infrastructure/utils"
)

type userService struct {
}

var _ userServiceInterface = &userService{}

type userServiceInterface interface {
	UpsertRegion(xjId uint) (err error)
}

func (us *userService) UpsertRegion(xjId uint) (err error) {
	// 查询用户注册IP信息
	var regIp string
	var xjUser *persistence.XJUser
	whereStr := fmt.Sprintf("%v =?", persistence.XJUserColumns.XJID)
	xjUser, err = persistence.XJUsers(qm.Select(persistence.XJUserColumns.RegIP), qm.Where(whereStr, xjId)).OneG()
	if err != nil {
		return err
	}
	regIp = xjUser.RegIP.String
	// 根据IP信息转化地区信息
	var ipInfo *ip2region.IpInfo
	ipInfo, err = utils.IpToRegion(regIp, utils.Ip2RegionSearchMethodBtree)
	if err != nil {
		return err
	}
	fmt.Println(ipInfo)
	// 更新用户地区信息
	userRegion := &persistence.XJUserRegion{
		XJID:       xjId,
		RegIP:      regIp,
		CityID:     null.Int{Int: int(ipInfo.CityId), Valid: true},
		Country:    null.String{String: ipInfo.Country, Valid: true},
		Region:     null.String{String: ipInfo.Region, Valid: true},
		Province:   null.String{String: ipInfo.Province, Valid: true},
		City:       null.String{String: ipInfo.City, Valid: true},
		Isp:        null.String{String: ipInfo.ISP, Valid: true},
		DecodeBody: null.String{String: ipInfo.String(), Valid: true},
	}
	fmt.Println(userRegion)
	return userRegion.UpsertG(
		boil.Blacklist(
			persistence.XJUserRegionColumns.XJID,
			persistence.XJUserRegionColumns.RegIP,
		),
		boil.Infer(),
		//	boil.Whitelist(
		//	persistence.XJUserRegionColumns.XJID,
		//	persistence.XJUserRegionColumns.RegIP,
		//	persistence.XJUserRegionColumns.CityID,
		//	persistence.XJUserRegionColumns.Country,
		//	persistence.XJUserRegionColumns.Region,
		//	persistence.XJUserRegionColumns.Province,
		//	persistence.XJUserRegionColumns.City,
		//	persistence.XJUserRegionColumns.Isp,
		//	persistence.XJUserRegionColumns.DecodeBody,
		//),
	)
}
