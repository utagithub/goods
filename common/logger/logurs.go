package logger

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strconv"
	"time"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		_, _ = fmt.Fprintf(b, "[%s] [%d] [%s]  %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		_, _ = fmt.Fprintf(b, "[%s] [%d] [%s] %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// Logrus 项目中没有使用Logrus
var (
	Logrus *logrus.Logger
)

func InitLogurs() {
	Logrus = logrus.New() //新建一个实例

	Logrus.SetReportCaller(true) //开启返回函数名和行号

	//设置日志输出到文件
	logFileName := strconv.Itoa(time.Now().Year()) + "-" + strconv.Itoa(int(time.Now().Month())) + "-" + strconv.Itoa(time.Now().Day())
	writer, _ := os.OpenFile("/opt/homebrew/var/www/web_go/goods/storage/log/"+logFileName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Logrus.Out = writer
	gin.DefaultWriter = Logrus.Out //设置gin框架日志输出

	Logrus.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter

	Logrus.SetLevel(logrus.InfoLevel) //设置最低的Level
}
