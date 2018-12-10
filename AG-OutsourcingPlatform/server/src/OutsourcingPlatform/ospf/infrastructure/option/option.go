package option

import (
	"OutsourcingPlatform/ospf/constant"
	"os"
	"github.com/BurntSushi/toml"
	"fmt"
)

type Options struct {
	BuildType string
	BuildVer  string

	FrontDir string
	ProDir   string

	AppName string `toml:"AppName"`

	DBDriver   string `toml:"DBDriver"`
	DBHost     string `toml:"DBHost"`
	DBPort     string `toml:"DBPort"`
	DBUser     string `toml:"DBUser"`
	DBPassword string `toml:"DBPassword"`
	DBName     string `toml:"DBName"`
	DBTestName string `toml:"DBTestName"`

	RedisHost string `toml:"RedisHost"`
	RedisPort string `toml:"RedisPort"`
	RedisPwd  string `toml:"RedisPwd"`
	RedisDB   int    `toml:"RedisDB"`

	ServerPort string `toml:"serverPort"`

	//log
	ConsoleOutPutPath string `toml:"ConsoleOutPutPath"`
	JsonOutPutPath    string `toml:"JsonOutPutPath"`
	ErrOutPutPath     string `toml:"ErrOutPutPath"`

	LoggerBufferCap int `toml:"LoggerBufferCap"`

	LoggerBufDuration int64 `toml:"LoggerBufDuration"`

	//use for session
	SessionKey  string `toml:"SessionKey"`
	MaxLifeTime int    `toml:"MaxLifeTime"`
	Path        string `toml:"Path"`
	HTTPOnly    bool   `toml:"HTTPOnly"`
	MaxAge      int    `toml:"MaxAge"`

	FileSuffixes []string `toml:"suffixes"`

	//use for wechat
	ClientDomain string `toml:"ClientDomain"`
	AppID        string `toml:"AppID"`
	AppSecret    string `toml:"AppSecret"`

	//AllowOrigin
	AllowOrigin string `toml:"origin"`

//	qrcode maxlifetime

	QrcodeLifeTime int `toml:"QrcodeLifeTime"`

	ServerDomain string `toml:"ServerDomain"`

	//proDir
}

func NewOpts(proDir, BUILDTYPE string) (opts *Options) {

	opts = &Options{}

	var confFileName, weChatConfFileName string
	if BUILDTYPE == constant.BUILD_TYPE_DEPLOY {
		confFileName = "app-deploy.toml"
		weChatConfFileName = "wechat-deploy.toml"
	} else {
		confFileName = "app-dev.toml"
		weChatConfFileName = "wechat-dev.toml"
	}

	_, err := toml.DecodeFile(proDir+"/option-file/"+confFileName, opts)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile(proDir+"/option-file/"+weChatConfFileName, opts)
	if err != nil {
		panic(err)
	}

	opts.ProDir = proDir
	opts.FrontDir = proDir + "/dist/ospFrontend/"
	fmt.Println(opts.FrontDir)
	err = mkdirProDir(proDir+"/log", proDir+"/file/")
	if err != nil {
		panic("mkdirProDir failed " + err.Error())
	}

	return opts
}

func mkdirProDir(logDir, fileDir string) error {
	//创建log目录在项目目录下
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0770)
	} else {
		return err
	}

	//创建file目录
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		os.Mkdir(fileDir, 0770)
	} else {
		return err
	}



	return nil
}
