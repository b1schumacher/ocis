package config

// App defines the available app configuration.
type App struct {
	Name        string `yaml:"name" env:"COLLABORATION_APP_NAME" desc:"The name of the app" introductionVersion:"5.1"`
	Description string `yaml:"description" env:"COLLABORATION_APP_DESCRIPTION" desc:"App description" introductionVersion:"5.1"`
	Icon        string `yaml:"icon" env:"COLLABORATION_APP_ICON" desc:"Icon for the app" introductionVersion:"5.1"`
	LockName    string `yaml:"lockname" env:"COLLABORATION_APP_LOCKNAME" desc:"Name for the app lock" introductionVersion:"5.1"`

	Addr     string `yaml:"addr" env:"COLLABORATION_APP_ADDR" desc:"The URL where the WOPI app is located, such as https://127.0.0.1:8080." introductionVersion:"5.1"`
	Insecure bool   `yaml:"insecure" env:"COLLABORATION_APP_INSECURE" desc:"Skip TLS certificate verification when connecting to the WOPI app" introductionVersion:"5.1"`
}
