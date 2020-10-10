package gt_config

import (
	"github.com/dodoao/gt"
	"github.com/dodoao/gt/gt_file"
	"os"
	"reflect"
	"strings"
)

//检查conf文件夹在不在
func check_dir() {
	path := gt.Get_current_directory() + "/conf"
	if !gt_file.Path_exists(path) {
		os.Mkdir(path, os.ModePerm)
	}
}

//func Set(fileName string, confStruct interface{}) bool {
//	check_dir()
//	object := reflect.ValueOf(confStruct)
//	myref := object.Elem()
//	typeOfType := myref.Type()
//	fileName = gt.Get_current_directory() + "/conf/" + fileName
//	allConfStr := ""
//
//	for i := 0; i < myref.NumField(); i++ {
//		field := myref.Field(i)
//		allConfStr += typeOfType.Field(i).Name + "="
//		switch typeOfType.Field(i).Type.String() {
//		case "string":
//			allConfStr += field.String() + "\n"
//			break
//		case "primitive.ObjectID":
//			objId := field.Interface().(primitive.ObjectID)
//			allConfStr += gt.PrimitiveObjectID_to_string(objId) + "\n"
//			break
//		case "int":
//			tmpInt := field.Int()
//			allConfStr += gt.Int64_to_string(tmpInt) + "\n"
//			break
//		case "int64":
//			allConfStr += gt.Int64_to_string(field.Int()) + "\n"
//			break
//		case "float32":
//			allConfStr += gt.Float64_to_string(field.Float()) + "\n"
//			break
//		case "float64":
//			allConfStr += gt.Float64_to_string(field.Float()) + "\n"
//			break
//
//		}
//
//	}
//	gt_file.Create_and_cover(fileName, allConfStr)
//
//	return true
//}

func Set(fileName string, confStruct interface{}) bool {
	check_dir()
	object := reflect.ValueOf(confStruct)
	myref := object.Elem()
	typeOfType := myref.Type()
	fileName = gt.Get_current_directory() + "/conf/" + fileName
	allConfStr := ""

	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		allConfStr += typeOfType.Field(i).Name + "="

		var SendData gt.GT_Type_SendData_struct
		SendData.Field = &field
		SendData.TypeName = typeOfType.Field(i).Type.String()

		tmpString := gt.Type_to_string(&SendData)
		if tmpString != "" {
			allConfStr += tmpString + "\n"
		}

	}
	gt_file.Create_and_cover(fileName, allConfStr)

	return true
}

//func Read(fileName string, confStruct interface{}) {
//	fileName = gt.Get_current_directory() + "/conf/" + fileName
//	if !gt_file.Path_exists(fileName) {
//		return
//	}
//
//	object := reflect.ValueOf(confStruct)
//	myref := object.Elem()
//	typeOfType := myref.Type()
//
//	allConfStr := gt_file.Read_file(fileName)
//	if allConfStr == "" {
//		return
//	}
//
//	for i := 0; i < myref.NumField(); i++ {
//		field := myref.Field(i)
//		itemNameTmp := typeOfType.Field(i).Name + "="
//		itemValue := get_key_value(allConfStr, itemNameTmp)
//		switch typeOfType.Field(i).Type.String() {
//		case "string":
//
//			field.SetString(itemValue)
//			break
//		case "primitive.ObjectID":
//			tmpId, bool_ := gt.String_to_PrimitiveObjectID(itemValue)
//			if bool_ == true {
//				field.Set(reflect.ValueOf(tmpId))
//			}
//			break
//		case "int":
//			tmpInt,bool_ := gt.String_to_int64(itemValue)
//			if bool_ == true {
//				field.SetInt(tmpInt)
//			}
//			break
//		case "int64":
//			TmpInt,bool_ := gt.String_to_int64(itemValue)
//			if bool_ == true {
//				field.SetInt(TmpInt)
//			}
//			break
//		case "float32":
//			tmpInt,bool_ := gt.String_to_float32(itemValue)
//			if bool_ == true {
//				field.Set(reflect.ValueOf(tmpInt))
//			}
//			break
//		case "float64":
//			tmpInt,bool_ := gt.String_to_float64(itemValue)
//			if bool_ == true {
//				field.SetFloat(tmpInt)
//			}
//			break
//
//		}
//
//	}
//	gt_file.Create_and_cover(fileName, allConfStr)
//
//	return
//}

func Read(fileName string, confStruct interface{}) {
	fileName = gt.Get_current_directory() + "/conf/" + fileName
	if !gt_file.Path_exists(fileName) {
		return
	}

	object := reflect.ValueOf(confStruct)
	myref := object.Elem()
	typeOfType := myref.Type()

	allConfStr := gt_file.Read_file(fileName)
	if allConfStr == "" {
		return
	}

	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		itemNameTmp := typeOfType.Field(i).Name + "="
		itemValue := get_key_value(allConfStr, itemNameTmp)

		var SendData gt.GT_Type_SendData_struct
		SendData.Field = &field
		SendData.TypeName = typeOfType.Field(i).Type.String()

		gt.String_to_Type(itemValue, &SendData)

	}
	gt_file.Create_and_cover(fileName, allConfStr)

	return
}

func get_key_value(str, key string) string {
	lenKey := len(key)
	if len(str) <= lenKey {
		return ""
	}
	ArrItem := strings.Split(str, "\n")
	for _, v := range ArrItem {
		if len(v) <= lenKey {
			continue
		}
		if v[:lenKey] == key {
			return v[lenKey:]
		}
	}

	return ""
}
