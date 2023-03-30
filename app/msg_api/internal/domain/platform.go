package domain

type Platform int

var (
	PlatformPC  Platform = 1
	PlatformAPP Platform = 2
	PlatformH5  Platform = 3

	PlatformUnknown Platform = -1
)

var platformTxtMapping = make(map[Platform]string, 3)

func init() {
	platformTxtMapping = make(map[Platform]string, 3)
	platformTxtMapping[PlatformPC] = "PC"
	platformTxtMapping[PlatformAPP] = "APP"
	platformTxtMapping[PlatformH5] = "H5"

	platformTxtMapping[PlatformUnknown] = "UNKNOWN"
}

func PlatformFrom(v int) Platform {
	p := Platform(v)
	_, has := platformTxtMapping[p]
	if has {
		return p
	}
	return PlatformUnknown
}

func PlatformText(platform Platform) string {
	txt, has := platformTxtMapping[platform]
	if has {
		return txt
	}
	return platformTxtMapping[PlatformUnknown]
}
