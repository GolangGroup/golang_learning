package iniconfig

import (
	"io/ioutil"
	"fmt"
	"strings"
	"reflect"
	"strconv"
)

func parseSection(line string, result interface{})(resultSectionName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		return	
	}
	
	if line[0] == '[' && line[len(line)-1] != ']' {
		return
	}
	
	sectionName := line[1:len(line)-1]
	if len(sectionName) == 0 {
		return
	}
	
	resultType := reflect.TypeOf(result)
	if resultType.Kind() != reflect.Ptr {
		err = fmt.Errorf("result is not ptr")
		return
	}

	resultConfigType := resultType.Elem()
	
	if resultConfigType.Kind() != reflect.Struct {
		err = fmt.Errorf("resulttype is not struct")
		return
	}

	for i := 0; i < resultConfigType.NumField(); i++ {
		if sectionName == resultConfigType.Field(i).Tag.Get("ini") {
			resultSectionName = resultConfigType.Field(i).Name
		} 
	}

	return
}

func parseItem(resultSectionName string, line string, result interface{})(err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("line sytnax err %s", line)
		return
	}
	
	itemName := line[:index]
	itemValue := line[index+1:]
	
	itemName = strings.TrimSpace(itemName)
	itemValue = strings.TrimSpace(itemValue)
	
	if len(itemName) == 0 {
		err = fmt.Errorf("line sytnax err %s", line)
		return
	}
	
	resultValue := reflect.ValueOf(result).Elem()
	resultType := resultValue.Type()
	
	var resultSectionValue reflect.Value
	var resultSectionType reflect.Type
	
	for i := 0; i < resultValue.NumField(); i++ {
		if resultSectionName == resultType.Field(i).Name {
			resultSectionValue = resultValue.Field(i)
			resultSectionType = resultSectionValue.Type()
			break
		}	
	}
	
	if resultSectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field %s must be struct", resultSectionName)
		return
	}
	var resultItemValue reflect.Value
	var resultItemName string = ""
	for j := 0; j < resultSectionValue.NumField(); j++ {
		if itemName == resultSectionType.Field(j).Tag.Get("ini") {
			resultItemValue = resultSectionValue.Field(j)
			resultItemName = resultSectionType.Field(j).Name
		}
	}
	
	if len(resultItemName) == 0 {
		return
	}
	
	switch resultItemValue.Type().Kind() {
	case reflect.String:
		resultItemValue.SetString(itemValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(itemValue, 10, 64)
		if err != nil {
			return err
		}
		resultItemValue.SetInt(value)
	default:
		err = fmt.Errorf("unsupport type %d", resultItemValue.Type().Kind())
	}


	return


}

func UnMarshalFileConfig(filename string, result interface{})(err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ioutil.ReadFile err=%v", err)
		return err
	}
	
	lineArray := strings.Split(string(data), "\r\n")
	var resultSectionName string
	for index, line := range lineArray {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		
		if line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			resultSectionName, err = parseSection(line, result)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
				return err
			}
			continue
		}
		if len(resultSectionName) == 0 {
			continue
		}

		err = parseItem(resultSectionName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)
			return err
		}
	}
	
	
	return
}


func MarshalFileConfig(filename string, conf interface{})(err error) {
	confValue := reflect.ValueOf(conf)
	confType := confValue.Type()
	
	if confType.Kind() != reflect.Struct {
		err = fmt.Errorf("conf must be struct")
		return
	}
	
	var data []byte

	for i := 0; i < confValue.NumField(); i++ {
		confSectionType := confType.Field(i)
		confSectionValue := confValue.Field(i)
		if confSectionType.Type.Kind() != reflect.Struct {
			continue
		}
		str := fmt.Sprintf("[%s]\r\n", confSectionType.Tag.Get("ini"))
		data = append(data, []byte(str)...)
		for j := 0; j < confSectionType.Type.NumField(); j++ {
			str := fmt.Sprintf("%s=%v\r\n", confSectionType.Type.Field(j).Tag.Get("ini"), confSectionValue.Field(j).Interface())
			data = append(data, []byte(str)...)
		}
	}
	
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		return
	}
	return
}














