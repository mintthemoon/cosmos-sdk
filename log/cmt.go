package log

import (
	cmtlog "github.com/cometbft/cometbft/libs/log"
	"github.com/rs/zerolog"
)

var _ cmtlog.Logger = (*CometZerologWrapper)(nil)

// CometZerologWrapper provides a wrapper around a zerolog.Logger instance. It implements
// CometBFT's Logger interface.
type CometZerologWrapper struct {
	zerolog.Logger
}

// Info implements CometBFT's Logger interface and logs with level INFO. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z CometZerologWrapper) Info(msg string, keyVals ...interface{}) {
	z.Logger.Info().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Error implements CometBFT's Logger interface and logs with level ERR. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z CometZerologWrapper) Error(msg string, keyVals ...interface{}) {
	z.Logger.Error().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Debug implements CometBFT's Logger interface and logs with level DEBUG. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z CometZerologWrapper) Debug(msg string, keyVals ...interface{}) {
	z.Logger.Debug().Fields(getLogFields(keyVals...)).Msg(msg)
}

// With returns a new wrapped logger with additional context provided by a set
// of key/value tuples. The number of tuples must be even and the key of the
// tuple must be a string.
func (z CometZerologWrapper) With(keyVals ...interface{}) cmtlog.Logger {
	return CometZerologWrapper{z.Logger.With().Fields(getLogFields(keyVals...)).Logger()}
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals)%2 != 0 {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}
