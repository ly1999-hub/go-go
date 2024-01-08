package log

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Warn ...
func Warn(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Warn(content, zap.String("data", string(jsonData)))
}
