package old

import (
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

//检查是否是验证码
func VerifyCodeFormat(code string) bool {
	pattern := `^[0-9A-Z]{5}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(code)
}

func Clear_data_models(a interface{}, b interface{}) interface{} {
	a_list := GetFieldName(a)
	b_list := GetFieldName(b)
	//a_t := reflect.ValueOf(a)
	b_t := reflect.ValueOf(b)

	m_st := reflect.TypeOf(a)
	m_vl := reflect.New(m_st)
	m_vl = m_vl.Elem()

	for a_i, a_v := range a_list {
		for b_i, b_v := range b_list {
			if a_v == b_v {
				m_vl.Field(a_i).Set(b_t.Field(b_i))
			}
		}
	}
	return m_vl.Interface()
}

//返回登录信息
//func Return_login_info(ctx iris.Context) interface{} {
//
//}

// 用b的所有字段覆盖a的
// 如果fields不为空, 表示用b的特定字段覆盖a的
// a应该为结构体指针
func CopyFields(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)

	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}

	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)

		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}

/*
获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//根据操作系统类型格式化文件夹路径
func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		fmt.Println("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}

//复制文件夹
func CopyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E", "/Y", "/C")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}
	_, e := cmd.Output()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
}

//转换编码 gbk to utf 8
func ConvertStrGbk2Utf8(str string) string {
	ret, _ := simplifiedchinese.GBK.NewDecoder().String(str)
	return ret
}

//转换编码utf-8 to gbk
func ConvertStrUtf82GBk(str string) string {
	enc := mahonia.NewEncoder("GBK")
	str = enc.ConvertString(str)
	return str
}

func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

//循环文件夹
func Ergodic_floader(path string, fu func(p string, f os.FileInfo)) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			Ergodic_floader(path+file.Name()+"/", fu)
		}

		fu(path, file)

	}
}

//文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//创建文件
func CreateFile(path string, str string) (error, bool) {
	f, err := os.Create(path)
	defer f.Close()
	if err == nil {
		_, err = f.Write([]byte(str))
		if err != nil {
			return err, false
		} else {
			return nil, true
		}
	} else {
		return err, false
	}
}

//获取LS
func Get_Floder_LS(path string) []string {
	fs, _ := ioutil.ReadDir(path)
	var ar []string
	for _, file := range fs {
		ar = append(ar, file.Name())
	}
	fmt.Println(ar)

	return ar

}
