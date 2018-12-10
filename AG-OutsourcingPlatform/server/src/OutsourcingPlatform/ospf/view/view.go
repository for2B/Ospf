package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
	"OutsourcingPlatform/ospf/constant"
	"bytes"
	"compress/gzip"
	"net/url"
)

var (
	indexTemplate *template.Template
	templates     *template.Template
)

func Init(templateDir string) {
	var allfile []string
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".html") {
			allfile = append(allfile, templateDir+fileName)
		}
	}
	// 这里要判断是否没有html文件
	if 0 == len(allfile) {
		fmt.Println("dist dir has no html file")
		return
	}
	templates = template.Must(template.ParseFiles(allfile...))
	indexTemplate = templates.Lookup("index.html")
	//log.Println("view init finish")
}

func LoadTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.Method == "GET" {
		indexTemplate.Execute(w, nil)
		return
	}
	http.Redirect(w, r, constant.RE_PROXY_SCHEME, http.StatusFound)
}

func GzipServeFile(w http.ResponseWriter, r *http.Request, Path string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	isGzip := strings.HasPrefix(r.Header.Get("Accept-Encoding"), "gzip")
	if isGzip {
		GzipFile(w, r, Path)
		return
	}
	http.ServeFile(w, r, Path)
}

func GzipFileServe(w http.ResponseWriter, r *http.Request, Path string, DistPath string) {
	//中文解码
	UrlDecode, err := url.QueryUnescape(Path)
	if err != nil {
		http.Error(w, "404 not found", 404)
		return
	}
	path := DistPath + UrlDecode
	//判断是否支持压缩
	isGzip := strings.HasPrefix(r.Header.Get("Accept-Encoding"), "gzip")
	PathSuffix := GetFileSuffix(Path)
	if isGzip {
		//设置相应的请求头
		if PathSuffix == "css" {
			w.Header().Set("Content-Type", "text/css;charset=utf-8")
		} else {
			w.Header().Set("Content-Type", "text/html;charset=utf-8")
		}
		GzipFile(w, r, path)
		return
	}
	http.ServeFile(w, r, path)
}

func GzipFile(w http.ResponseWriter, r *http.Request, path string) {
	data, err := ioutil.ReadFile(path) //读取文件
	if err != nil {
		http.Error(w, "404 not found", 404)
		return
	}
	buf := &bytes.Buffer{}
	//新建一个gzip指针
	g, err := gzip.NewWriterLevel(buf, 9)
	if err != nil {
		http.Error(w, "404 not found", 404)
		return
	}
	//将读取到数据，进行压缩
	g.Write(data)
	g.Close()
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Vary", "Content-Encoding")
	w.Write(buf.Bytes())
}

func GetFileSuffix(filename string) string {
	FileNameSclie := strings.Split(filename, ".")
	FNSLen := len(FileNameSclie)
	return strings.ToLower(FileNameSclie[FNSLen-1])
}
