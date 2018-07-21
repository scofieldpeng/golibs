package log

import (
	"github.com/sirupsen/logrus"
	"fmt"
	"os"
	"time"
	"errors"
	"github.com/scofieldpeng/golibs/filepath"
	"strings"
	sysFilePath "path/filepath"
)

var (
	logger = &logrus.Logger{}
	writer = newFileStdWriter()
)

// log初始化
func Init(isDebug bool, filePath ...string) {
	logger = &logrus.Logger{}
	logger.Formatter = &logrus.JSONFormatter{}
	if isDebug {
		logger.Formatter = &logrus.TextFormatter{}
		logger.Out = os.Stdout
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(logrus.DebugLevel)
		writer.Init(filePath...)
		logger.Out = writer
	}
}

// 获取logrus.Logger
func GetLogger() *logrus.Logger {
	return logger
}

// 获取writer
func GetWriter() *fileStdWriter {
	return writer
}

func Close() {
	if writer != nil {
		writer.Close()
	}
}

type fileStdWriter struct {
	// 日志文件夹路径
	dirPath string
	// 日志文件名
	fileName string
	file     *os.File
}

func newFileStdWriter() *fileStdWriter {
	runDir, _ := filepath.GetRunDir()
	return &fileStdWriter{
		dirPath: runDir,
	}
}

// 初始化
func (f *fileStdWriter) Init(dirPath ...string) {
	var err error
	if len(dirPath) == 0 {
		dirPath = make([]string, 1)
		dirPath[0], err = filepath.GetRunDir()
		if err != nil {
			fmt.Println("获取应用当前目录失败,err:" + err.Error())
			os.Exit(1)
		}
	}

	if dirPath[0] != string(os.PathSeparator) {
		dirPath[0] = strings.TrimRight(dirPath[0], string(os.PathSeparator))
	}

	f.dirPath = dirPath[0] + string(os.PathSeparator) + "log"
	f.fileName = f.GenerateFileName()


	if _,err := os.Stat(f.dirPath);os.IsNotExist(err) {
		fmt.Printf("日志目录(%s)不存在,自动创建",f.dirPath)
		if err := os.MkdirAll(f.dirPath,0655);err != nil {
			fmt.Printf("\t|- 创建失败,error: %s \n",err.Error())
		} else {
			fmt.Printf("\t|- 创建成功")
		}
	}

	go f.Rotate()
}

// 设置文件名
func (f *fileStdWriter) GenerateFileName() string {
	return time.Now().Format("2006-01-02") + ".log"
}

// LoadFile
func (f *fileStdWriter) LoadFile() (err error) {
	if f.dirPath == "" || f.fileName == "" {
		return errors.New("没有找到合法的路径")
	}

	if f.file, err = os.OpenFile(f.GetLogFilePath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
		return err
	}

	return nil
}

// 获取日志文件路径
func (f *fileStdWriter) GetLogFilePath() (path string) {
	return f.dirPath + string(os.PathSeparator) + f.fileName
}

func (f *fileStdWriter) Rotate() {
	nowTime := time.Now()
	todayStartTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local).Unix()
	sleepTime := 86400 - (nowTime.Unix() - todayStartTime)
	for {
		f.GenerateFileName()
		if err := f.LoadFile(); err != nil {
			fmt.Println("无法载入日志文件,err:" + err.Error() + ",filename:" + f.dirPath + string(os.PathSeparator) + f.fileName)
			time.Sleep(5 * time.Second)
			continue
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
		sleepTime = 86400
	}
}

// 关闭logger
func (f *fileStdWriter) Close() {
	f.file.Close()
}

// 写入内容
func (f *fileStdWriter) Write(data []byte) (n int, err error) {
	n, err = os.Stderr.Write(data)
	if f.file != nil {
		n, err = f.file.Write(data)
	}

	return
}
