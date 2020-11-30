package gt_file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/***
 * 创建和覆盖，并写入文本
 */
func Create_and_cover(fileName string, str string) bool {
	os.Remove(fileName)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("不能创建文件", fileName)
		return false
	}
	defer f.Close()
	_, _ = f.WriteString(str)
	return true
}

/***
 * 创建或者追加
 */
func Create_or_append(fileName string, str string) bool {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("不能创建文件", fileName)
		return false
	}
	defer f.Close()
	_, _ = f.WriteString(str)
	return true
}

/***
 * 文件读取
 */
func Read_file(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ""
	}
	confContext := string(f)
	return confContext
}

//文件夹或文件是否存在
func Path_exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//目录搜索，参数：目录、遇到文件的函数、遇到目录的函数
func Path_search(path string, FileFunc func(Path, FileName string, FileType string), FolderFunc func(Path, FileName string)) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			FolderFunc(path, file.Name())
		} else {
			i := strings.LastIndex(file.Name(), ".")
			strType := ""
			if i != -1 {
				strType = file.Name()[:i]
			}
			FileFunc(path, file.Name(), strType)
		}
	}
}
