package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"encoding/hex"
	"crypto/rand"

)

// GetProDir 用于获取项目根目录
func GetProDir() (string, error) {

	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	path = filepath.Dir(path)
	return strings.Replace(path, "\\", "/", -1), nil
}

func round(num float64) int {
	//留整数部分 四舍五入
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	//将要保留的小数通过 x10 放到整数部分 去小数,再/10
	return float64(round(num*output)) / output
}

func HandlePanic(proDir string) {
	if err := recover(); err != nil {
		formatErr := fmt.Sprintf("%s\t%v", time.Now().Local(), err)
		stack := make([]byte, 1024*10)
		if len(stack) == 0 {
			fmt.Println("stack length = 0")
			return
		}
		runtime.Stack(stack, true)
		stack = bytes.Replace(stack, []byte("\u0000"), []byte(""), -1)
		fmt.Printf("%s\n%s", formatErr, stack)
		ioutil.WriteFile(proDir+"/panic", bytes.Join([][]byte{[]byte(formatErr), stack}, []byte("\n")), 0644)
	}
}

// 获取多个url参数 map版本
func GetURLParams(r *http.Request, key ...string) map[string]string {
	urlParams := make(map[string]string)
	for _, param := range key {
		urlParams[param] = r.URL.Query().Get(param)
	}
	return urlParams
}

func ToOpenURL(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

// Query returns the keyed url query value if it exists
func QueryUrl(key string, r *http.Request) string {
	value, _ := GetQuery(key, r)
	return value

}

// GetQuery is like Query(), it returns the keyed url query value
func GetQuery(key string, r *http.Request) (string, bool) {
	if values, ok := r.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

func SetCOOKIE(w http.ResponseWriter, key, value string) {

	Cookie := http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(value),
		Path:     "/",
		HttpOnly: false,
		MaxAge:   7200,
	}
	http.SetCookie(w, &Cookie)
}
func GetRemoteIP(req *http.Request) string {

	ip := req.Header.Get("Remote_addr")
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	if ip == "" {
		return host
	}
	return ip
}

func GetMd5Time() string {
	CurrentTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(CurrentTime, 10))
	md5Time := fmt.Sprintf("%x", h.Sum(nil))
	return md5Time
}

func GenerateRandomID() string{
sid := make([]byte, 16)
rand.Read(sid)
return hex.EncodeToString(sid)
}


/*
作用:将时间戳转换为本地时间
输入:time类型时间戳
输出:字符串
*/
func ShowTime(t time.Time)  {
	localLocation, err := time.LoadLocation("Local")//服务器设置的时区
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println("转化为:",t.In(localLocation))
}

/*
作用:判断请求方的referer是否包含domain
输入:domain域名,r请求
输出:是或否
*/
func IsClient(r *http.Request,domain string) bool{
	referer:=r.Header.Get("referer")
	if strings.Contains(referer,domain){
		return true
	}
	return false
}


/*
作用:判断请求方的referer是否为空
输入:domain域名,r请求
输出:是或否
*/
func IsRefererEmpty(r *http.Request) bool{
	referer:=r.Header.Get("referer")
	if referer==""{
		return true
	}
	return false
}


