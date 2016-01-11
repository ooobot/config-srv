package config

import (
	"github.com/micro/go-platform/config"
)

var (
	// We need a path splitter since its structured
	// in go-platform
	PathSplitter = "/"

	reader config.Reader
)

func Init() error {
	reader = config.NewReader()
	return nil
}

func Parse(ch ...*config.ChangeSet) (*config.ChangeSet, error) {
	return reader.Parse(ch...)
}

func Values(ch *config.ChangeSet) (config.Values, error) {
	return reader.Values(ch)
}
