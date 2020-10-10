package gt_http

import (
	"encoding/json"
	"github.com/dodoao/gt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//返回http信息
func Http_get(url string) (string, bool) {
	resp, err := http.Get(url)

	if err != nil {
		return err.Error(), false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 && err == nil {
		return string(body), true
	} else {
		return string(body) + "\n" + err.Error(), false
	}
}

//Http get 并且根据struct创建参数
func Http_get_for_data(url string, data interface{}) (string, bool) {

	newUrl := url
	questionMarkPosition := strings.LastIndex(url, "?")

	if len(newUrl) < 5 {
		return "", false
	}
	lastWords := newUrl[len(newUrl)-1:]

	clearLastValue := []string{"/", "\\", "?", "&"}

	for _, v := range clearLastValue {
		if lastWords == v {
			newUrl = newUrl[:len(newUrl)-1]
		}
	}

	object := reflect.ValueOf(data)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {

		if questionMarkPosition == -1 && i == 0 {
			newUrl = newUrl + "?"
		} else {
			newUrl = newUrl + "&"
		}

		field := myref.Field(i)
		newUrl += typeOfType.Field(i).Name + "="

		switch typeOfType.Field(i).Type.String() {
		case "string":
			newUrl += field.String()
			break
		case "ObjectID":
			objId := field.Interface().(primitive.ObjectID)
			newUrl += gt.PrimitiveObjectID_to_string(objId)
			break
		case "int":
			tmpInt := field.Int()
			newUrl += gt.Int64_to_string(tmpInt)
			break
		case "int64":
			newUrl += gt.Int64_to_string(field.Int())
			break
		case "float32":
			newUrl += gt.Float64_to_string(field.Float())
			break
		case "float64":
			newUrl += gt.Float64_to_string(field.Float())
			break
		}

	}

	resp, err := http.Get(newUrl)
	if err != nil {
		return err.Error(), false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 && err == nil {
		return string(body), true
	} else {
		return string(body) + "\n" + err.Error(), false
	}
}

func Http_post(url string, data interface{}) (string, bool) {

	jsons, _ := json.Marshal(data)
	result := string(jsons)
	jsoninfo := strings.NewReader(result)
	req, _ := http.NewRequest("POST", url, jsoninfo)
	req.Header.Add("appCode", "winner")
	//req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", false
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", false
	}

	//fmt.Println(res)
	return string(body), true
}
