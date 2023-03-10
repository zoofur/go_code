package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	BEIKE   = "beike"
	LIANJIA = "lianjia"

	CLIENT_APP = "app"
	CLIENT_XCX = "xcx"

	GreatThan = 1
	LessThan  = 2
)

type AppVersion struct {
	Version  string
	Business string
	Client   string
	Fun      uint8
}


// compareAll 默认比较所有端，对于不传入的端即默认不做版本控制，下发全部内容
// todo:方法有争议，实习结束还用不到这个方法就直接删除
func (a AppVersion) compareAll(versions ...Version) bool {
	var flag bool
	for _, v := range versions {
		if v != nil {
			switch v.GetAppVersion().Fun {
			case GreatThan:
				flag = a.GreatThan(v.GetAppVersion())
			case LessThan:
				flag = a.LessThan(v.GetAppVersion())
			default:
				flag = true
			}
		}
		if flag {
			return flag
		}
	}
	return flag
}

// compareApp 只比较 Client 为 App 的端，对于不可比的端抛出 error
func (a AppVersion) compareApp(versions ...Version) (bool, error) {
	var flag bool
	for _, v := range versions {
		if v != nil {
			if v.GetAppVersion().Client != CLIENT_APP {
				return false, errors.New("compareApp compare a Invalid client")
			}
			switch v.GetAppVersion().Fun {
			case GreatThan:
				flag = a.GreatThan(v.GetAppVersion())
			case LessThan:
				flag = a.LessThan(v.GetAppVersion())
			}
			if flag {
				return flag, nil
			}
		}
	}
	return flag, nil
}

func (a AppVersion) compareXcx(versions ...Version) (bool, error) {
	var flag bool
	for _, v := range versions {
		if v != nil {
			if v.GetAppVersion().Client != CLIENT_XCX {
				return false, errors.New("compareClient compare a Invalid client")
			}
			switch v.GetAppVersion().Fun {
			case GreatThan:
				flag = a.GreatThan(v.GetAppVersion())
			case LessThan:
				flag = a.LessThan(v.GetAppVersion())
			}
			if flag {
				return flag, nil
			}
		}
	}
	return flag, nil
}

// 对原有方法改动较小，只修改入参即可
func (a AppVersion) GreatThan(version AppVersion) bool {
	if a.Business != version.Business {
		return false
	}
	client, _ := strconv.Atoi(a.Version)
	compare, _ := strconv.Atoi(version.Version)
	if client >= compare {
		return true
	}
	return false
}

func (a AppVersion) LessThan(version AppVersion) bool {
	if a.Business != version.Business {
		return false
	}
	client, _ := strconv.Atoi(a.Version)
	compare, _ := strconv.Atoi(version.Version)
	if client < compare {
		return true
	}
	return false
}

func (a AppVersion) GetAppVersion() AppVersion {
	return a
}

func extend(version Version) {
	fmt.Println(BeiKeExtend(version))
}
