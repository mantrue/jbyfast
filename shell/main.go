package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"shell/parameter"
	"time"
)

func main() {
	//增加企业基本信息
	/*data := parameter.ImportInfo()

	body := bytes.TrimPrefix([]byte(data), []byte("\xef\xbb\xbf"))

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
		idmap, ok := dat["data"].(map[string]interface{})
		if ok == true {
			cid := idmap["id"]
			//关联公司关系cid,A1,京版云
			ua := parameter.UpdateAppInfo(cid.(string))
			fmt.Println(ua)

			//更新APP关联类型信息
			uac := parameter.UpdateAppClass(cid.(string))
			fmt.Println(uac)

			//增加行政许可
			l := parameter.LicenseInfo(cid.(string))
			fmt.Println(l)
			//失信被执行人
			ud := parameter.UpdateDishonest(cid.(string))
			fmt.Println(ud)

			//经营异常
			uab := parameter.UpdateAbnormal(cid.(string))
			fmt.Println(uab)

			//法律诉讼
			uj := parameter.UpdateJudicial(cid.(string))
			fmt.Println(uj)
		}

	} else {
		fmt.Println(err.Error())
	}*/

	//产品内容监管  需要三个接口（更新安全扫描日志记录信息,增加用户反馈信息,增加用户反馈信息）
	/*us := parameter.UpdateSafeScanInfo() //数量一致就是绿的能通过（1淫秽，2反党反共）

	postValue := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"}, //1—投诉,2—举报
		"supportType":  {"2"}, //1—不具备视力保护功能并可支持家长监管 2--学习无关的网络游戏等内容及链接 3--企业泄露隐私 4－存在侵权内容 5－集体入驻型平台向个人收费
		"description":  {"暴力"},
	}
	uf := parameter.UserFeed(postValue)

	fmt.Println(us)
	fmt.Println(uf)*/

	//产品经营监管 需要三个接口(增加课程基本信息,增加用户订单,增加任教教师基本信息)
	/*aci := parameter.AddCourseInfo()
	aoi := parameter.AddOrderInfo()
	fmt.Println(aci)
	fmt.Println(aoi)

	ati := parameter.AddTeacherInfo()
	fmt.Println(ati)

	postValue1 := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"},
		"supportType":  {"3"},
		"description":  {"隐私泄露"},
	}
	us1 := parameter.UserFeed(postValue1)
	fmt.Println(us1)*/

	//舆情监管 需要两个接口(增加APP评论信息,增加用户订单)
	/*ui := parameter.AppCommentInfo()
	fmt.Println(ui)
	uaa := parameter.UpdateAppInfoCopy()
	fmt.Println(uaa)*/

	//产品功能监管 需要1个接口(增加用户反馈信息)
	/*postValue := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"},
		"supportType":  {"1"},
		"description":  {"产品功能监管护眼和家长监督"},
	}
	us := parameter.UserFeed(postValue)

	postValue1 := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"},
		"supportType":  {"3"},
		"description":  {"信息安全泄露"},
	}
	us1 := parameter.UserFeed(postValue1)
	fmt.Println(us)
	fmt.Println(us1)*/

	//企业资质管理 需要1个接口(企业信用监管需要有各种url，增加用户反馈信息)
	/*ui := parameter.UpdateInfo()
	fmt.Println(ui)*/

	//投诉解决数据 需要2个接口(增加用户反馈信息,更新用户反馈) 投诉并设置解决
	dates := []string{"2019-11-25 00:00:00", "2019-10-25 00:00:00", "2019-09-25 00:00:00", "2019-08-25 00:00:00", "2019-07-25 00:00:00"}

	for _, dt := range dates {
		ui := parameter.UserFeedTouSu(dt)
		var r Res
		body := bytes.TrimPrefix([]byte(ui), []byte("\xef\xbb\xbf"))
		err := json.Unmarshal([]byte(body), &r)

		if err != nil {
			fmt.Println(err)
			continue
		}

		uuf := parameter.UpdateUserFeedbackInfo(r.Data.Id)
		fmt.Println(uuf)
		fmt.Println(ui)
	}

	//产品投诉比例

	postValue := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"},
		"supportType":  {"1"},
		"description":  {"产品功能监管"},
		"create_at":    {"2019-10-25 00:00:00"},
	}
	us := parameter.UserFeed(postValue)
	postValue1 := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"},
		"supportType":  {"3"},
		"description":  {"信息安全泄露"},
		"create_at":    {"2019-09-25 00:00:00"},
	}
	us1 := parameter.UserFeed(postValue1)
	fmt.Println(us)
	fmt.Println(us1)

	//产品营收 接口(新增订单)
	aoi := parameter.AddOrderInfo()
	fmt.Println(aoi)

	//用户类型分类
	ati := parameter.AddTeacherInfo() //增加用户教师
	fmt.Println(ati)

	ast := parameter.AddStudentInfo() //增加学生
	fmt.Println(ast)

	api := parameter.AddParentInfo() //增加家长
	fmt.Println(api)

	aai := parameter.AddAuthInfo() //增加教育部门基本信息
	fmt.Println(aai)

	//增加用户登录信息
	aul := parameter.AddUserLoginInfo()
	fmt.Println(aul)

	//增加用户时长信息
	alt := parameter.AddOnlineTimeInfo()
	fmt.Println(alt)

	//触发定时任务接口
	parameter.DoHttp()
	time.Sleep(time.Second * 2)
	fmt.Println("=================do end=================")
}

type Res struct {
	Status  int
	Message string
	Data    DataInfo
}

type DataInfo struct {
	Id string
}
