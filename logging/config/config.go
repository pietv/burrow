package config

import (
	"bytes"
	"fmt"

	"github.com/hyperledger/burrow/logging/structure"

	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/hyperledger/burrow/logging/loggers"
)

type LoggingConfig struct {
	RootSink     *SinkConfig `toml:",omitempty"`
	ExcludeTrace bool
	NonBlocking  bool
}

// For encoding a top-level '[logging]' TOML table
type LoggingConfigWrapper struct {
	Logging *LoggingConfig `toml:",omitempty"`
}

func DefaultNodeLoggingConfig() *LoggingConfig {
	return &LoggingConfig{
		RootSink: Sink().SetOutput(StderrOutput().SetFormat(loggers.JSONFormat)),
	}
}

func DefaultClientLoggingConfig() *LoggingConfig {
	return &LoggingConfig{
		// No output
		RootSink: Sink().
			SetTransform(FilterTransform(ExcludeWhenAnyMatches,
				structure.CapturedLoggingSourceKey, "")).
			SetOutput(StderrOutput()),
	}
}

// Returns the TOML for a top-level logging config wrapped with [logging]
func (lc *LoggingConfig) RootTOMLString() string {
	return TOMLString(LoggingConfigWrapper{lc})
}

func (lc *LoggingConfig) TOMLString() string {
	return TOMLString(lc)
}

func (lc *LoggingConfig) RootJSONString() string {
	return JSONString(LoggingConfigWrapper{lc})
}

func (lc *LoggingConfig) JSONString() string {
	return JSONString(lc)
}

func TOMLString(v interface{}) string {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	err := encoder.Encode(v)
	if err != nil {
		// Seems like a reasonable compromise to make the string function clean
		return fmt.Sprintf("Error encoding TOML: %s", err)
	}
	return buf.String()
}

func JSONString(v interface{}) string {
	bs, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error encoding JSON: %s", err)
	}
	return string(bs)
}
