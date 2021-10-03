// Package config loads configuration data from an external file.
package config

// YML represents the first level of the YML file.
type YML struct {
	Generated Generated              `yaml:"generated"`
	Templates Templates              `yaml:"templates"`
	Options   map[string]interface{} `yaml:"custom"`
}

// Generated represents generated properties of the YML file.
type Generated struct {
	Setup  string `yaml:"setup"`
	Output string `yaml:"output"`
}

// Templates represent template properties of the YML file.
type Templates struct {
	Function string `yaml:"function"`
}
