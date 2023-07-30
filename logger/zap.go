package logger

import (
	"cld/settings"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(cfg *settings.LogConfig) (err error) {
	// 获取日志写入器
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
		cfg.Compress,
	)
	// 获取日志编码器
	encoder := getEncoder()
	// 创建日志核心 高于Debug级别会被记录到日志中
	var level = new(zapcore.Level)
	if err = level.UnmarshalText([]byte((cfg.Level))); err != nil {
		fmt.Println("looger umarshltext falied" + err.Error())
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// 创建日志记录器，使用 zap.AddCaller() 添加调用者信息(zap/main.go:27)
	logger := zap.New(core, zap.AddCaller())
	// 使用zap.L()替代
	zap.ReplaceGlobals(logger)
	return
}

func getEncoder() zapcore.Encoder {
	// 创建一个生产环境下的编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 使用 ISO8601 格式编码时间
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 使用大写字母编码日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 创建并返回控制台编码器
	//日志保存格式 console / json
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxsize int, maxbackups int, maxage int, compress bool) zapcore.WriteSyncer {
	// 以追加模式打开日志文件，如果文件不存在则创建
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxAge:     maxage,
		MaxBackups: maxbackups,
		Compress:   compress,
	}
	// 将文件包装为 WriteSyncer，用于日志输出
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic stack是否输出堆栈信息
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
