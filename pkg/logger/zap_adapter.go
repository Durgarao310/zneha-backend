package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapAdapter adapts the existing Logger to implement the zap-like interface
// required by the error middleware
type ZapAdapter struct {
	logger *Logger
}

// NewZapAdapter creates a new adapter for the existing logger
func NewZapAdapter(logger *Logger) *ZapAdapter {
	return &ZapAdapter{logger: logger}
}

// Info implements the zap-like Info method
func (z *ZapAdapter) Info(msg string, fields ...zap.Field) {
	// Convert zap fields to a simple format for the existing logger
	if len(fields) > 0 {
		z.logger.Infof("%s - %v", msg, fieldsToMap(fields))
	} else {
		z.logger.Info(msg)
	}
}

// Error implements the zap-like Error method
func (z *ZapAdapter) Error(msg string, fields ...zap.Field) {
	// Convert zap fields to a simple format for the existing logger
	if len(fields) > 0 {
		z.logger.Errorf("%s - %v", msg, fieldsToMap(fields))
	} else {
		z.logger.Error(msg)
	}
}

// fieldsToMap converts zap fields to a simple map for logging
func fieldsToMap(fields []zap.Field) map[string]interface{} {
	result := make(map[string]interface{})
	for _, field := range fields {
		switch field.Type {
		case zapcore.StringType:
			result[field.Key] = field.String
		case zapcore.Int64Type:
			result[field.Key] = field.Integer
		case zapcore.BoolType:
			result[field.Key] = field.Integer == 1
		default:
			result[field.Key] = field.Interface
		}
	}
	return result
}
