package CTL

import (
	"../httpRouter"
	"GoLang-WEB/Demo/models"
	"fmt"
	"strconv"
)

var userFields map[string]FieldType = map[string]FieldType{
	"username": FieldType{
		"string",
		map[string]bool{
			"edit": false,
			"add":  true,
		},
	},
	"sex": FieldType{
		"int",
		map[string]bool{
			"edit": true,
			"add":  true,
		},
	},
	"password": FieldType{
		"string",
		map[string]bool{
			"edit": true,
			"add":  true,
		},
	},
}

func userFilterFields(body map[string]interface{}, opType string) (map[string]interface{}, error) {
	return FilterFields(body, opType, userFields)
}

func AddUser(c *odserver.Context) {
	resp := Resp{
		DEFAULT_RETCODE,
		"添加用户成功",
		"",
	}
	_resp := &resp
	defer RESP(c, _resp)
	body, err := userFilterFields(c.PostParams(), "add")
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     err.Error(),
		}
		return
	}
	result, err := Models.InsertUser(body)
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "修改信息失败",
			Data:    err.Error(),
		}
		return
	}
	_resp.Data = map[string]int64{
		"uid": result,
	}
}

func DeleteUser(c *odserver.Context) {
	resp := Resp{
		RetCode: DEFAULT_RETCODE,
		Msg:     "用户删除成功",
		Data:    "",
	}
	_resp := &resp
	defer RESP(c, _resp)
	uid := c.GetParams()["uid"]
	i, err := strconv.Atoi(uid) // string类型转换为Int
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "uid类型错误",
			Data:    err.Error(),
		}
		return
	}
	result, err := Models.DeleteUser(i)
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "修改信息失败",
			Data:    err.Error(),
		}
		return
	}
	if result == 0 {
		resp = Resp{
			RetCode: "0",
			Msg:     "删除失败",
			Data:    "",
		}
	}
}

func PutUser(c *odserver.Context) {
	resp := Resp{
		RetCode: DEFAULT_RETCODE,
		Msg:     "用户修改成功",
		Data:    "",
	}
	_resp := &resp
	defer RESP(c, _resp)
	uid := c.GetParams()["uid"]
	i, err := strconv.Atoi(uid) // string类型转换为Int
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "uid类型错误",
			Data:    err.Error(),
		}
		return
	}

	body := c.PostParams()
	body, err = userFilterFields(body, "edit")
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     err.Error(),
			Data:    "",
		}
		return
	}
	result, err := Models.UpdateUser(i, body)
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "修改信息失败",
			Data:    err.Error(),
		}
		return
	}
	if result == 0 {
		resp = Resp{
			RetCode: "0",
			Msg:     "修改信息失败",
			Data:    "",
		}
	}
}

func FindUser(c *odserver.Context) {
	resp := Resp{
		RetCode: DEFAULT_RETCODE,
		Msg:     "用户查找成功",
		Data:    "",
	}
	_resp := &resp
	defer RESP(c, _resp)
	uid := c.GetParams()["uid"]
	i, err := strconv.Atoi(uid) // string类型转换为Int
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "uid类型错误",
			Data:    err.Error(),
		}
		return
	}
	result, getid, err := Models.UserDetail(i)
	fmt.Println("result", result)
	fmt.Println("getid", getid)
	if err != nil {
		resp = Resp{
			RetCode: "0",
			Msg:     "查询信息失败",
			Data:    err.Error(),
		}
		return
	}
	if result == 0 {
		resp = Resp{
			RetCode: "0",
			Msg:     "查询失败",
			Data:    getid,
		}
	}
	if result == 1 {
		resp = Resp{
			RetCode: "1",
			Msg:     "用户查找成功",
			Data:    getid,
		}
	}
}
