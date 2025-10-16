package config 


const (
	DefaultScreenWidth  = 800
	DefaultScreenHeight = 600
	DefaultTitle        = "Vampire Survivors Clone"
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
	Title        string
	TPS          int
}

func Default() *Config {
	return &Config{
		ScreenWidth:  DefaultScreenWidth,
		ScreenHeight: DefaultScreenHeight,
		Title:        DefaultTitle,
		TPS:          60,
	}
}
