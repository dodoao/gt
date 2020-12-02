package gt

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/***
 * 获取程序运行路径你
 */
func Get_current_directory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//int 转 string
func Int_to_string(int_ int) string {
	return strconv.Itoa(int_)
}

func Int_to_int64(int_ int) int64 {
	int64_ := int64(int_)
	return int64_
}

func Int64_to_int(int64_ int64) (int, bool) {
	strInt64 := strconv.FormatInt(int64_, 10)
	int_, bool_ := String_to_int(strInt64)
	if bool_ == false {
		return 0, false
	} else {
		return int_, true
	}
}

func String_to_int(str string) (int, bool) {
	int_, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	} else {
		return int_, true
	}
}

func Int64_to_string(int64 int64) string {
	return strconv.FormatInt(int64, 10)
}

func String_to_int64(string string) (int64, bool) {
	int64, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		return 0, false
	} else {
		return int64, true
	}
}

func Float32_to_string(float32 float32) string {
	return strings.Trim(strconv.FormatFloat(float64(float32), 'f', 6, 64), "0")
}

func Float64_to_string(float64 float64) string {
	return strings.Trim(strconv.FormatFloat(float64, 'f', 6, 64), "0")
}

func String_to_float32(str string) (float32, bool) {
	float32_, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, false
	} else {
		return float32(float32_), true
	}
}

func String_to_float64(str string) (float64, bool) {
	float64_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, false
	} else {
		return float64_, true
	}
}

func PrimitiveObjectID_to_string(id primitive.ObjectID) string {
	objId := id.String()
	lastIndex := strings.LastIndex(objId, "\")")
	return objId[10:lastIndex]
}

func String_to_PrimitiveObjectID(strKey string) (primitive.ObjectID, bool) {
	MongodbObjectId, err := primitive.ObjectIDFromHex(strKey)
	if err != nil {
		return MongodbObjectId, false
	} else {
		return MongodbObjectId, true
	}
}

//
func String_to_Time(str string) (time.Time, bool) {

	var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
	st, err := time.Parse(timeLayoutStr, str)
	if err != nil {
		return st, false
	}
	return st, true
}

func Time_to_String(st time.Time) string {
	return st.String()
}

//string转类型传递的数据结构
type GT_Type_SendData_struct struct {
	Field    *reflect.Value
	TypeName string
}

//类型的值转string
func Type_to_string(Type_ interface{}) string {
	var field reflect.Value
	tmpTest := reflect.ValueOf(Type_)
	field = tmpTest.Elem()
	typeOfTypeName := field.Type().Name()
	if typeOfTypeName == "GT_Type_SendData_struct" {
		NewSturct := field.Interface().(GT_Type_SendData_struct)
		field = *NewSturct.Field
		typeOfTypeName = NewSturct.TypeName
	}

	if typeOfTypeName == "primitive.ObjectID" {
		typeOfTypeName = "ObjectID"
	}

	switch typeOfTypeName {
	case "string":
		return field.String()
		break
	case "ObjectID":
		objId := field.Interface().(primitive.ObjectID)
		return PrimitiveObjectID_to_string(objId)
		break
	case "int":
		tmpInt := field.Int()
		return Int64_to_string(tmpInt)
		break
	case "int64":
		return Int64_to_string(field.Int())
		break
	case "float32":
		return Float64_to_string(field.Float())
		break
	case "float64":
		return Float64_to_string(field.Float())
		break
	}
	return ""
}

//string转类型
func String_to_Type(itemValue string, Type_ interface{}) bool {

	var field reflect.Value
	tmpTest := reflect.ValueOf(Type_)
	field = tmpTest.Elem()
	typeOfTypeName := field.Type().Name()
	if typeOfTypeName == "GT_Type_SendData_struct" {
		NewSturct := field.Interface().(GT_Type_SendData_struct)
		field = *NewSturct.Field
		typeOfTypeName = NewSturct.TypeName

	}
	if typeOfTypeName == "primitive.ObjectID" {
		typeOfTypeName = "ObjectID"
	}
	switch typeOfTypeName {
	case "string":
		field.SetString(itemValue)
		return true
		break
	case "ObjectID":
		tmpId, bool_ := String_to_PrimitiveObjectID(itemValue)
		if bool_ == true {
			field.Set(reflect.ValueOf(tmpId))
			return true
		}
		return false
		break
	case "int":
		tmpInt, bool_ := String_to_int64(itemValue)
		if bool_ == true {
			field.SetInt(tmpInt)
			return true
		}
		return false
		break
	case "int64":
		TmpInt, bool_ := String_to_int64(itemValue)
		if bool_ == true {
			field.SetInt(TmpInt)
			return true
		}
		return false
		break
	case "float32":
		tmpInt, bool_ := String_to_float32(itemValue)
		if bool_ == true {
			field.Set(reflect.ValueOf(tmpInt))
			return true
		}
		return false
		break
	case "float64":
		tmpInt, bool_ := String_to_float64(itemValue)
		if bool_ == true {
			field.SetFloat(tmpInt)
			return true
		}
		return false
		break

	}
	return false
}
