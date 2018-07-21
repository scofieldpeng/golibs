package log

import (
	"testing"
	"github.com/scofieldpeng/golibs/filepath"
	"time"
	"io/ioutil"
	"github.com/sirupsen/logrus"
)

func TestFileStdWriter(t *testing.T) {
	runFileDir,_:= filepath.GetRunDir()
	writer.Init()
	if runFileDir != writer.dirPath {
		t.Fatal("run dir is wrong, get:" + writer.dirPath)
	}
	t.Log("run dir:" + runFileDir)
	if writer.GenerateFileName() != time.Now().Format("2006-01-02") + ".log" {
		t.Fatal("log filename wrong, get:",writer.fileName)
	}
	t.Log("log filename:",writer.fileName)
	writer.Close()

	writer.Init("/tmp")
	if writer.dirPath != "/tmp" {
		t.Fatal("run dir is wrong, want /tmp, get:" + writer.dirPath)
	}
	// for log file write is ready
	time.Sleep(1 * time.Second)
	logContent1 := "hello world"
	if _,err := writer.Write([]byte(logContent1));err != nil {
		t.Fatal("write log faild,err:" + err.Error())
	}
	t.Log("log file path:" + writer.GetLogFilePath())
	b,err := ioutil.ReadFile(writer.GetLogFilePath())
	if err != nil {
		t.Fatal("read " + writer.GetLogFilePath() + " failed!,err:" + err.Error())
	}
	if string(b) != logContent1 {
		t.Fatal("log file not equal " + logContent1 + ",get: " + string(b))
	}
	writer.Close()
}

func TestLogger(t *testing.T) {
	Init(true)
	time.Sleep(1*time.Second)
	GetLogger().Debug("hello world")

	Init(false,"/tmp")
	time.Sleep(time.Second * 1)
	GetLogger().WithFields(logrus.Fields{
		"name":"scofield",
		"age":"26",
	}).Debug("hello logger")
}
