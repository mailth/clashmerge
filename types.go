package main

type ConfigItem struct {
	Operation string `yaml:"operation"`
	Type      string `yaml:"type"`
	Data      map[string]any `yaml:"data"`
	URL       string `yaml:"url"`
}
