package gt_verify

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

//32位md5加密
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
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
