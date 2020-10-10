package main

import (
	"fmt"
	_ "github.com/dodoao/gt"
	"github.com/dodoao/gt/gt_http"
)

func main() {

	//创建日志示例
	//gt_log.Year("以年为单位创建日志文件")
	//gt_log.Month("以月为单位创建日志文件")
	//gt_log.Day("以日为单位创建日志文件")
	//gt_log.Hour("以小时为单位创建日志文件")

	//===========================

	//配置示例
	//type Conf_struct struct {
	//	Id                int64
	//	Id2               int
	//	Fold_ float32
	//	Fold2 float64
	//	FileName          string
	//	ServerName        string
	//	MongodbURi        string
	//	MongodbDatabase   string
	//	MongodbCollection string
	//	MongodbObjectId   primitive.ObjectID
	//}
	//var conf Conf_struct
	//conf.FileName = "kk.ini"
	//conf.Id = 1
	//conf.Id2 = 64
	//conf.Fold_=1.1
	//conf.Fold2=2.2
	//conf.ServerName = "0.0.0.0:8080"
	//conf.MongodbURi = "mongodb://localhost:27017"
	//conf.MongodbDatabase = "myNewDatabase"
	//conf.MongodbCollection = "myCollection"
	//conf.MongodbObjectId, _ = primitive.ObjectIDFromHex("5f713db3f499d0c4c2dc73d4")
	//gt_config.Set("test.conf", &conf)
	//
	//var conf2 Conf_struct
	//gt_config.Read("test.conf",&conf2)
	//fmt.Println(conf2)

	//===========================

	//类型的值转string
	//var int_ int
	//int_ = 1
	//stringInt := gt.Type_to_string(&int_)
	//fmt.Println(stringInt)
	//
	//var int64_ int64
	//int64_ = 1
	//stringInt64 := gt.Type_to_string(&int64_)
	//fmt.Println(stringInt64)
	//
	//var ObjectID_ primitive.ObjectID
	//ObjectID_ = primitive.NewObjectID()
	//stringObjectId := gt.Type_to_string(&ObjectID_)
	//fmt.Println(stringObjectId)
	//
	//////string转类型
	//var int_2 int
	//stringInt2 := "1"
	//bool_ := gt.String_to_Type(stringInt2, &int_2)
	//if bool_ {
	//	fmt.Println(int_2)
	//}
	//
	//var int64_2 int64
	//stringInt642 := "1"
	//bool_2 := gt.String_to_Type(stringInt642, &int64_2)
	//if bool_2 {
	//	fmt.Println(int64_2)
	//}
	//
	//var ObjectID_2 primitive.ObjectID
	//stringObjectId2 := "5f7fd350ea61c224b008d98b"
	//bool_3 := gt.String_to_Type(stringObjectId2, &ObjectID_2)
	//if bool_3 {
	//	fmt.Println(ObjectID_2)
	//}

	//===========================

	//Http 示例

	url := "https://www.baidu.com/"
	type Ex_Data struct {
		Name string
		Test string
	}
	var Ex_val Ex_Data
	Ex_val.Name = "测试名称"
	Ex_val.Test = "测试内容"

	result, bool_ := gt_http.Http_get(url)
	if bool_ {
		fmt.Println(result)
	}

	result, bool_ = gt_http.Http_get_for_data(url, &Ex_val)
	if bool_ {
		fmt.Println(result)
	}

	result, bool_ = gt_http.Http_post(url, &Ex_val)
	if bool_ {
		fmt.Println(result)
	}

	//===========================

}
