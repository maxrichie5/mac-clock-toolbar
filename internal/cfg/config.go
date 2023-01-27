package cfg

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	ActiveTimeZones []string
}

var configFile = strings.Join([]string{
	homeDir(),
	".config",
	"mac-clock-toolbar-settings.json",
}, string(os.PathSeparator))

var config *Config

func Start() {
	config = initConfig()
}

func initConfig() *Config {
	file, err := os.ReadFile(configFile)
	if os.IsNotExist(err) {
		f, err := os.Create(configFile)
		if err != nil {
			log.Fatalln(err)
		}

		c := &Config{}
		var b []byte
		b, err = json.Marshal(c)

		_, err = f.Write(b)
		if err != nil {
			log.Fatalln(err)
		}

		file = b
	} else if err != nil {
		log.Fatalln(err)
	}

	c := &Config{}
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}

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

func (c *Config) hasActiveTimeZone(zone string) bool {
	for _, z := range c.ActiveTimeZones {
		if z == zone {
			return true
		}
	}
	return false
}

func (c *Config) addActiveTimeZone(zone string) {
	c.ActiveTimeZones = append(c.ActiveTimeZones, zone)
}

func (c *Config) removeActiveTimeZone(zone string) {
	t := make([]string, 0, len(c.ActiveTimeZones))
	for _, tz := range c.ActiveTimeZones {
		if tz != zone {
			t = append(t, tz)
		}
	}
	c.ActiveTimeZones = t
}

func (c *Config) write() {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(configFile, b, 0777)
	if err != nil {
		log.Fatalln(err)
	}
}
