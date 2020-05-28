package main

import (
	"go.uber.org/zap"
"time"
)

func main() {
    logger, _ := zap.NewDevelopment()
		logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()))
		
}