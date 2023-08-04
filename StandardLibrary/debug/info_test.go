package __

import (
	"log"
	"runtime/debug"
	"testing"
)

func TestBuildInfo(t *testing.T) {
	if info, ok := debug.ReadBuildInfo(); ok {
		log.Println(info.GoVersion)
		log.Println(info.Path)
		for _, setting := range info.Settings {
			log.Println(setting.Key, setting.Value)
		}
	}
}
