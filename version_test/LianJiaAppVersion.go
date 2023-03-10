package main

type LianJiaAppVersion struct {
	AppVersion
}

func NewLianJiaAppVersion(version string, fun uint8) Version {
	var appVersion Version
	appVersion = LianJiaAppVersion{
		AppVersion: AppVersion{
			Business: LIANJIA,
			Version:  version,
			Client:   CLIENT_APP,
			Fun:      fun,
		},
	}
	return appVersion
}