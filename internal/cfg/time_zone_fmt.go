package cfg

func SetActiveTimeZoneFmt(fmt string) {
	if fmt != config.ActiveTimeZoneFmt {
		config.ActiveTimeZoneFmt = fmt
		config.write()
	}
}

func GetActiveTimeZoneFmt() string {
	if config.ActiveTimeZoneFmt == "" {
		config.ActiveTimeZoneFmt = "15:04:05 MST"
		config.write()
	}
	return config.ActiveTimeZoneFmt
}

func SetTimeZoneMenuFmt(fmt string) {
	if fmt != config.TimeZoneMenuFmt {
		config.TimeZoneMenuFmt = fmt
		config.write()
	}
}

func GetTimeZoneMenuFmt() string {
	if config.TimeZoneMenuFmt == "" {
		config.TimeZoneMenuFmt = "15:04:05 MST"
		config.write()
	}
	return config.TimeZoneMenuFmt
}
