package utils

import "go.uber.org/zap"

var ZapLogger *zap.SugaredLogger

func InitLogger() {
	ZapLogger = newZapLogger()
}
