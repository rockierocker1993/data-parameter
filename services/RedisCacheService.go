package services

import (
	"context"
	"data-parameter/config"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ReloadAllCache(c *gin.Context) {
	slog.Info("in method ReloadCache")
	ctx := context.Background()
	config.RDB.Set(ctx, "lookup_values", nil, 0)
}

func ReloadCacheLookupValue(c *gin.Context) {
	slog.Info("in method ReloadCacheLookupValue")
	ctx := context.Background()
	config.RDB.Set(ctx, "lookup_values", nil, 0)
}
