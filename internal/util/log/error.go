package log

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Error ...
func Error(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Error(content, zap.String("data", string(jsonData)))
}
