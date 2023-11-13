package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Pingye007/godoing/config"
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

const (
	red    = 31
	green  = 32
	yellow = 33
	cyan   = 36
	grey   = 37
	ldebug = "debug"
	ltrace = "trace"
	linfo  = "info"
	lwarn  = "warn"
	lerror = "error"
	lfatal = "fatal"
	lpanic = "panic"
)

type LogFormatter struct{}

func (lm *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.TraceLevel, logrus.DebugLevel:
		levelColor = green
	case logrus.InfoLevel:
		levelColor = cyan
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = grey
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:06")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s \033[%dm[%s]\033[0m (%s %s): %s \n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)

	} else {
		fmt.Fprintf(b, "%s \033[%dm[%s]\033[0m %s \n", timestamp, levelColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

func initLog() {
	level := strings.ToLower(config.Cfg.Log.Level)
	time := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("log/%s_%s.log", time, config.Cfg.SysConfig.Version)
	var w io.Writer
	if level == linfo || level == lwarn {
		w = os.Stdout
	} else {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println("create log file failed:", err)
		}
		if level == ldebug || level == ltrace {
			w = io.MultiWriter(os.Stdout, file)
		} else {
			w = io.MultiWriter(os.Stderr, file)
		}
	}

	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetFormatter(&LogFormatter{})
	Log.SetOutput(w)

	switch level {
	case ldebug:
		Log.SetLevel(logrus.DebugLevel)
	case ltrace:
		Log.SetLevel(logrus.TraceLevel)
	case linfo:
		Log.SetLevel(logrus.InfoLevel)
	case lwarn:
		Log.SetLevel(logrus.WarnLevel)
	case lerror:
		Log.SetLevel(logrus.ErrorLevel)
	case lfatal:
		Log.SetLevel(logrus.FatalLevel)
	case lpanic:
		Log.SetLevel(logrus.PanicLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}

func init() {
	initLog()
}
