package old

import (
	"crypto/md5"
	"encoding/hex"
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

//32位md5加密
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//验证名称是否符合规范
func VerifyNameFormat(name string) error {
	if name == "" {
		return errors.New("名称不能为空！")
	}

	if strings.Index(name, " ") != -1 {
		return errors.New("名称不能有空格！")
	}

	if strings.Index(name, "<") != -1 {
		return errors.New("名称不能有“<”符号!！")
	}
	if strings.Index(name, ">") != -1 {
		return errors.New("名称不能有“>”符号!！")
	}
	if strings.Index(name, "//") != -1 {
		return errors.New("名称不能有“//”符号!！")
	}
	if strings.Index(name, "%") != -1 {
		return errors.New("名称不能有“%”符号!！")
	}
	if strings.Index(name, "*") != -1 {
		return errors.New("名称不能有“*”符号!！")
	}
	if strings.Index(name, ";") != -1 {
		return errors.New("名称不能有“;”符号!！")
	}
	if strings.Index(name, ",") != -1 {
		return errors.New("名称不能有“,”符号!！")
	}
	if strings.Index(name, "\"") != -1 {
		return errors.New("名称不能有“\"”符号!！")
	}
	if strings.Index(name, "'") != -1 {
		return errors.New("名称不能有“'”符号!！")
	}
	return nil
}

//检查是否是验证码
func VerifyCodeFormat(code string) bool {
	pattern := `^[0-9A-Z]{5}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(code)
}

//是否是md5
func VerifyMd5Format(md5 string) bool {
	pattern := `^[0-9a-z]{32}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(md5)
}

//是否为邮箱地址
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//是否为手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//获取字段名称
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)

	}
	return result
}

//f1 完整的struct  f2 一段不完整的struct，  例如：用户登录时只使用到了struct里的name和pass字段，但完整的用户struct的字段要多得多，那f1就是完整的用户struct，用户登录提交的struct就是f2。
//fu 一段字段验证方法，需要自定义 示例：
/*
func Attribute_format(name string, v interface{}) error {
	switch name {
	case "Attribute_name":
		v := v.(string)
		err := tools.VerifyNameFormat(v)
		if err != nil {
			return err;
		}
		break

	}
	return nil
}
*/
func Return_validate(f1 interface{}, f2 []interface{}, fu func(s string, v interface{}) error) (interface{}, error) {

	var f_list []string
	var m_d interface{}

	if f2 != nil && len(f2) > 0 {
		f_list = GetFieldName(f2[0])
		m_d = f2[0]
	} else {
		f_list = GetFieldName(f1)
		m_d = f1
	}
	t := reflect.ValueOf(m_d)
	t2 := reflect.ValueOf(f1)
	f1_list := GetFieldName(f1)
	for i, value1 := range f_list {
		v := t.Field(i)

		isFind := false
		for i2, value2 := range f1_list {
			if value2 == value1 {
				isFind = true
				ty2 := t2.Field(i2).Type()
				if ty2 != v.Type() {
					return nil, errors.New(value1 + "数据类型应该为：" + ty2.String())
				}
			}
		}
		if isFind == false {
			//return nil,errors.New(value1+"没有找到这个字段")
			continue
		}

		var err error
		switch v.Type().String() {
		case "string":
			err = fu(value1, v.Interface().(string))
			break
		case "int64":
			err = fu(value1, v.Interface().(int64))
			break
		case "time.Time":
			err = fu(value1, v.Interface().(time.Time))
			break
		case "bool":
			err = fu(value1, v.Interface().(bool))
			break
		case "int":
			err = fu(value1, v.Interface().(int))
			break
		case "int8":
			err = fu(value1, v.Interface().(int8))
			break
		default:
			fmt.Println("发现新类型，请添加" + v.Type().String())
			os.Exit(0)
			break
		}

		if err != nil {
			return nil, err
		}
	}
	return m_d, nil
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
