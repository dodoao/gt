package gt_log

import (
	"os"
	"strconv"
	"time"
)
import "github.com/dodoao/gt"
import "github.com/dodoao/gt/gt_file"

//检查log文件夹在不在
func check_dir(folderName string) string {
	path := gt.Get_current_directory() + "/log"
	if folderName != "" {
		path += "/" + folderName
	}
	if !gt_file.Path_exists(path) {
		os.MkdirAll(path, os.ModePerm)
	}
	return path + "/"
}

//以年为单位创建日志
func Year(str string, folderName string) bool {
	str = time.Now().String() + ":" + str + "\n"
	path := check_dir(folderName)
	t1 := strconv.Itoa(time.Now().Year()) //年
	path += t1 + ".log"
	return gt_file.Create_or_append(path, str)
}

//以月为单位创建日志
func Month(str string, folderName string) bool {
	str = time.Now().String() + ":" + str + "\n"
	path := check_dir(folderName)
	t1 := strconv.Itoa(time.Now().Year()) //年
	t2 := time.Now().Month().String()     //月
	path += t1 + "_" + t2 + ".log"
	return gt_file.Create_or_append(path, str)
}

//以日为单位创建日志
func Day(str string, folderName string) bool {
	str = time.Now().String() + ":" + str + "\n"
	path := check_dir(folderName)
	t1 := strconv.Itoa(time.Now().Year()) //年
	t2 := time.Now().Month().String()     //月
	t3 := strconv.Itoa(time.Now().Day())  //日
	path += t1 + "_" + t2 + "_" + t3 + ".log"
	return gt_file.Create_or_append(path, str)
}

//以小时为单位创建日志
func Hour(str string, folderName string) bool {
	str = time.Now().String() + ":" + str + "\n"
	path := check_dir(folderName)
	t1 := strconv.Itoa(time.Now().Year()) //年
	t2 := time.Now().Month().String()     //月
	t3 := strconv.Itoa(time.Now().Day())  //日
	t4 := strconv.Itoa(time.Now().Hour()) //日
	path += t1 + "_" + t2 + "_" + t3 + "_" + t4 + ".log"
	return gt_file.Create_or_append(path, str)
}
