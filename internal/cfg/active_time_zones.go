package cfg

import (
	"log"
	"os"
)

func AddActiveTimeZone(zone string) {
	if !config.hasActiveTimeZone(zone) {
		config.addActiveTimeZone(zone)
		config.write()
	}
}

func RemoveActiveTimeZone(zone string) {
	if config.hasActiveTimeZone(zone) {
		config.removeActiveTimeZone(zone)
		config.write()
	}
}

func TimeZomeActive(zone string) bool {
	return config.hasActiveTimeZone(zone)
}

func GetActiveTimeZones() []string {
	return config.ActiveTimeZones
}

func homeDir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	return h
}

func (c *appConfig) removeActiveTimeZone(zone string) {
	t := make([]string, 0, len(c.ActiveTimeZones))
	for _, tz := range c.ActiveTimeZones {
		if tz != zone {
			t = append(t, tz)
		}
	}
	c.ActiveTimeZones = t
}

func (c *appConfig) hasActiveTimeZone(zone string) bool {
	for _, z := range c.ActiveTimeZones {
		if z == zone {
			return true
		}
	}
	return false
}

func (c *appConfig) addActiveTimeZone(zone string) {
	c.ActiveTimeZones = append(c.ActiveTimeZones, zone)
}
