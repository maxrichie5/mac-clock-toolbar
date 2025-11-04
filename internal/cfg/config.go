package cfg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type appConfig struct {
	// These all need to be exported so json can unmarshal into them
	ActiveTimeZones   []string `json:"activeTimeZones"`
	ActiveTimeZoneFmt string   `json:"activeTimeZoneFmt"`
	TimeZoneMenuFmt   string   `json:"timeZoneMenuFmt"`
}

var configFile = strings.Join([]string{
	homeDir(),
	".config",
	"mac-clock-toolbar-settings.json",
}, string(os.PathSeparator))

var config *appConfig

func Start() {
	config = initConfig()
}

func initConfig() *appConfig {
	file, err := os.ReadFile(configFile)
	if os.IsNotExist(err) {
		f, err := os.Create(configFile)
		if err != nil {
			log.Fatalln(err)
		}

		c := &appConfig{}
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

	c := &appConfig{}
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(c)

	return c
}

func (c *appConfig) write() {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(configFile, b, 0777)
	if err != nil {
		log.Fatalln(err)
	}
}
