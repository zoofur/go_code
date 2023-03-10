package main

import "fmt"

func caseBeiKeApp() {
	appVersion := AppVersion{
		Version:  "2",
		Business: BEIKE,
		Client:   CLIENT_APP,
	}

	flag, err := appVersion.compareApp(
		NewBeiKeAppVersion("2", GreatThan),
		NewLianJiaAppVersion("9", GreatThan),
		)
	if !flag && err != nil {
		fmt.Println(err)
	}
	if flag {
		fmt.Println("caseBeiKeApp logic")
	}
}

func caseAllClient() {
	appVersion := AppVersion{
		Version:  "2",
		Business: BEIKE,
		Client:   CLIENT_APP,
	}

	if appVersion.compareAll(
		NewBeiKeAppVersion("2", GreatThan),
		NewLianJiaAppVersion("9", GreatThan),
		NewXcxAppVersion("9", GreatThan),
	) {
		fmt.Println("caseAllClient logic")
	}
}

func caseBeiKeAppError() {
	appVersion := AppVersion{
		Version:  "2",
		Business: BEIKE,
		Client:   CLIENT_APP,
	}

	flag, err := appVersion.compareApp(
		NewXcxAppVersion("9", GreatThan),
		NewBeiKeAppVersion("2", GreatThan),
		)
	if !flag && err != nil {
		fmt.Println(err)
	}
}

func caseExtend() {
	beike := NewBeiKeAppVersion("2", GreatThan)
	extend(beike)
}

func CaseRecommend() {
	appVersion := AppVersion{}
	flag, err := appVersion.compareApp(
		NewBeiKeAppVersion("2", GreatThan),
		NewLianJiaAppVersion("9", GreatThan),
		)
	if !flag && err != nil {
		// 对错误进行处理
		fmt.Println(err)
	}
	if flag {
		// 进入版本控制的逻辑
		fmt.Println("recommend logic")
	}
}