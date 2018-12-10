package infrastructure

import (
	"OutsourcingPlatform/ospf/infrastructure/cache"
	"OutsourcingPlatform/ospf/infrastructure/db"
	"OutsourcingPlatform/ospf/infrastructure/logger_core"
	"OutsourcingPlatform/ospf/infrastructure/option"
	"OutsourcingPlatform/ospf/infrastructure/session"
	"database/sql"
	"os"
	"sync/atomic"
	"time"

	"github.com/asdine/storm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"OutsourcingPlatform/internal/httpapi"
	"net/http"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat"
)

type Infrastructure struct {
	opts           atomic.Value
	Logger         *zap.Logger
	DB             *sql.DB
	stormDB        *storm.DB
	SessionManager *session.SessionManager
	CacheManager   *cache.CacheManager
	Wechat         *wechat.Wechat
}

func NewInfrastructure(opts *option.Options) (i *Infrastructure) {

	i = &Infrastructure{}
	i.setOpts(opts)
	i.DB = db.NewDB(opts.DBHost, opts.DBPort, opts.DBUser, opts.DBPassword, opts.DBName, opts.DBDriver)
	i.CacheManager = cache.NewCacheManager(opts.ProDir + "/my.db")
	i.InitLogger()

	i.SessionManager = new(session.SessionManager)
	i.SessionManager.Init(
		opts.SessionKey,
		opts.MaxLifeTime,
		opts.Path,
		opts.HTTPOnly,
		opts.MaxAge,
		i.CacheManager,
		i.Logger)

	wechat.InitWeChat(opts.AppID, opts.AppSecret, i.CacheManager)
	i.Wechat = wechat.NewWechat(&wechat.CFG)

	return
}

func (i *Infrastructure) GetOpts() *option.Options {
	return i.opts.Load().(*option.Options)
}

func (i *Infrastructure) AllowOrigin(next httpapi.APIHandler) httpapi.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) (response interface{}, err error) {
		w.Header().Set("Access-Control-Allow-Origin", i.GetOpts().AllowOrigin)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		w.Header().Set("content-type", "text/html;application/json;charset=utf-8")
		return next(w, r)
	}
}

func (i *Infrastructure) GetServerDomain() string {
	sdomain := i.GetOpts().ServerDomain + ":" + i.GetOpts().ServerPort
	return sdomain
}
func (i *Infrastructure) setOpts(opts *option.Options) {
	i.opts.Store(opts)
}

func (i *Infrastructure) InitLogger() {
	logLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	logEncoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "StacktraceKey",
		LineEnding:     zapcore.DefaultLineEnding, //add "\n" in line end automatically
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeTime:     logEncodeTime,             //show style of time
		EncodeCaller:   zapcore.FullCallerEncoder, //caller info
		EncodeName:     zapcore.FullNameEncoder,
	}
	logConsoleConfig := zap.Config{
		Level:            logLevel,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    logEncoderConfig,
		OutputPaths:      []string{"stdout", i.GetOpts().ProDir + string(os.PathSeparator) + i.GetOpts().ConsoleOutPutPath},
		ErrorOutputPaths: []string{"stderr", i.GetOpts().ProDir + string(os.PathSeparator) + i.GetOpts().ErrOutPutPath},
	}

	coreOption := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		jsonOutputPath := []string{i.GetOpts().ProDir + string(os.PathSeparator) + i.GetOpts().JsonOutPutPath}
		jsonFileWriteSyncer, _, err := zap.Open(jsonOutputPath...)
		if err != nil {
			panic(err)
		}

		DbLoggerWriteSyncer := dbloggercore.NewDbLoggerWriteSyncer(
			i.DB, i.GetOpts().LoggerBufferCap, i.GetOpts().LoggerBufDuration)

		logEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

		c := zapcore.NewCore(
			zapcore.NewJSONEncoder(logEncoderConfig),
			zapcore.NewMultiWriteSyncer(DbLoggerWriteSyncer, jsonFileWriteSyncer),
			logLevel,
		)
		return zapcore.NewTee(core, c)
	})

	var err error
	i.Logger, err = logConsoleConfig.Build(coreOption)
	if err != nil {
		panic(err)
	}

}

func logEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
