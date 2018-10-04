package CTL

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const DEFAULT_RETCODE = "1"

type FieldType struct {
	fType         string
	FieldsOperate map[string]bool
}

// 对收到参数的进行过滤，只提取需要的参数
func FilterFields(fields map[string]interface{}, ope string, originFields map[string]FieldType) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range fields {
		if _, ok := originFields[k]; ok {
			if originFields[k].FieldsOperate[ope] {
				result[k] = v
			}
		}
	}
	if len(result) == 0 {
		return result, errors.New("没有任何参数")
	}
	err := CheckFields(result)
	if err == nil {
		if ope == "edit" {
			result = AddUpdatedTime(result)
		}
		if ope == "add" {
			result = AddCreatedTime(result)
		}
	}
	return result, err
}

// 增加updated参数
func AddUpdatedTime(fields map[string]interface{}) map[string]interface{} {
	fields["updated"] = time.Now().Unix()
	return fields
}

// 增加created参数
func AddCreatedTime(fields map[string]interface{}) map[string]interface{} {
	fields["created"] = time.Now().Unix()
	return fields
}

// 校验性别
func CheckSex(sex interface{}) error {
	var getSex int
	var err error
	if vs, ok := sex.(string); ok {
		getSex, err = strconv.Atoi(vs)
		fmt.Println("getSex", vs)
	} else if vm, ok := sex.(float64); ok {
		fmt.Println("getSex2", vm)
		getSex = int(vm)
	}
	if err != nil {
		return errors.New("性别参数错误")
	}
	if getSex > 3 || getSex < 1 {
		return errors.New("性别范围错误")
	}
	return nil
}

// 对常用参数进行格式校验
func CheckFields(body map[string]interface{}) error {
	if _, ok := body["sex"]; ok {
		if err := CheckSex(body["sex"]); err != nil {
			return err
		}
	}
	return nil
}

// 判断性别参数是否正确
func IsSex(sex int) bool {
	if sex <= 3 || sex >= 1 {
		return true
	}
	return false
}
