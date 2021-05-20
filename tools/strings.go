package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//获取指定个数的随机字符串
//          size                     指定的字符串个数
//            kind                     0，纯数字；1，小写字母；2，大写字母；3，数字+大小写字母
//        string                      返回生成的随机字符串
func RandStr(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

//将interface{}转为字符串，适合bool，数字等
//@v            interface{}         需要转化为字符串的值
func Interface2String(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

//将字符串切片转成interface切片
func StringSliceToInterfaceSlice(slice []string) (ret []interface{}) {
	for _, v := range slice {
		ret = append(ret, v)
	}
	return ret
}

//将字符串切片数组转成map
func StringSliceToMap(slice []string) (maps map[string]string) {
	maps = make(map[string]string)
	for _, v := range slice {
		maps[v] = v
	}
	return maps
}

//字符串截取
func SubStr(str interface{}, start, length int) string {
	v := fmt.Sprintf("%v", str)
	if start < 0 {
		start = 0
	}
	slice := strings.Split(v, "")
	l := len(slice)
	if l == 0 || start > l {
		return ""
	}
	if start+length+1 > l {
		return strings.Join(slice[start:], "")
	}
	return strings.Join(slice[start:length], "")
}

//首字母大写
func UpperFirst(str string) string {
	if len(str) > 0 {
		strings.Replace(str, str[0:1], strings.ToUpper(str[0:1]), 1)
	}
	return str
}

func LtrimString(source string, pattern string) string {
	return strings.TrimLeft(source, pattern)
}

// string转float64
func StringToFloat64(str string) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {

	}
	return i
}
