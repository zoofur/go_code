package main

type XcxAppVersion struct {
	AppVersion
}

func NewXcxAppVersion(version string, fun uint8) Version {
	var appVersion Version
	appVersion = BeiKeAppVersion{
		AppVersion: AppVersion{
			Business: BEIKE,
			Version:  version,
			Client:   CLIENT_XCX,
			Fun:      fun,
		},
	}
	return appVersion
}
