package structs

type Device struct {
	Id string
	Type string
	Traits string
	Name DeviceName
	WillReportState bool
	RoomHint string
	DeviceInfo DeviceInfo
}

type DeviceName struct {
	DefaultNames []string
	Name []string
	Nicknames []string
}

type DeviceInfo struct {
	Manufacturer string
	Model string
	HwVersion string
	SwVersion string
}
