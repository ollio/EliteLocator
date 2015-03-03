package main

type Config struct {
	EliteDangerous struct {
		Path string
		Lang string
	}
	EliteOCR struct {
		Enable bool
		Path string
		ScreenShotPath string
	}
}

