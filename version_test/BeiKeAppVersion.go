package main

type BeiKeAppVersion struct {
	AppVersion
	extend string
}

func NewBeiKeAppVersion(version string, fun uint8) Version {
	var appVersion Version
	appVersion = BeiKeAppVersion{
		AppVersion: AppVersion{
			Business: BEIKE,
			Version:  version,
			Client:   CLIENT_APP,
			Fun:      fun,
		},
	}
	return appVersion
}

func BeiKeExtend(version Version) BeiKeAppVersion {
	beike := version.(BeiKeAppVersion)
	beike.extend = "AAA"
	return beike
}

