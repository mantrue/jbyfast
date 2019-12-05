package parameter

import (
	"encoding/json"
	"net/url"
	"shell/config/setconst"
	"shell/config/setslice"
	"shell/httprun"
	"time"
)

//增加企业基本信息
func ImportInfo() string {
	//企业基本信息
	var data = map[string]string{
		"companyName": "O1公司", "companyLogo": "http://www.jingbanyun.com//Public/img/home/swing1.png", "category": "01",
		"registeredCapital": "100万", "foundedDate": time.Now().Format("2006-01-02 15:04:05"), "businessStartDate": time.Now().Format("2006-01-02 15:04:05"), "businessEndDate": time.Now().Format("2006-01-02 15:04:05"),
		"address": "北京市海淀区西四环中路", "legalRepresentative": "区块链", "businessScope": "在线教育", "companyTelephone": "0000-563245475", "copyRightCount": "100",
		"patentCount": "200", "staffCount": "20", "onlineEduStartTime": time.Now().Format("2006-01-02 15:04:05"), "schoolLicenseUrl": "", "businessLicenseUrl": "", "icpUrl": "",
		"broadcastLicenseUrl": "", "cultureLicenseUrl": "", "publishLicenseUrl": "",
	}

	j, _ := json.Marshal(data)
	return httprun.SendPost(setconst.URI+setslice.ImportInfo, j)
}

//关联公司关系
func UpdateAppInfo(id string) string {
	//行政许可
	postValue := url.Values{
		"company_id":   {id},
		"app_name":     {"O1公司"},
		"downloadNum":  {"1"},
		"chargingMode": {"定期收费"},
	}

	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.UpdateAppInfo, []byte(postString))
}

func UpdateAppClass(id string) string {
	//行政许可
	var data = []map[string]string{
		{
			"businessClass": "1",
		},
		{
			"businessClass": "2",
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.UpdateAppClass, j)
}

func UpdateAppInfoCopy() string {
	//行政许可
	postValue := url.Values{
		"downloadNum": {"1234"},
	}

	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.UpdateAppInfo, []byte(postString))
}

//添加行政许可
func LicenseInfo(id string) string {
	//行政许可
	var data = []map[string]string{
		{
			"licenseType": "1", "licenseNumber": "0032154524", "licenseName": "0078编码", "validStart": time.Now().Format("2006-01-02 15:04:05"),
			"validEnd": time.Now().Format("2006-01-02 15:04:05"), "adminOrgan": "工商局", "adminContent": "许可", "decisionDocumentNo": "004527548875465", "company_id": id,
		},
		{
			"licenseType": "1", "licenseNumber": "7854545", "licenseName": "0078编码", "validStart": time.Now().Format("2006-01-02 15:04:05"),
			"validEnd": time.Now().Format("2006-01-02 15:04:05"), "adminOrgan": "工商局", "adminContent": "许可", "decisionDocumentNo": "004527548875465", "company_id": id,
		},
		{
			"licenseType": "2", "licenseNumber": "124564546", "licenseName": "0078编码", "validStart": time.Now().Format("2006-01-02 15:04:05"),
			"validEnd": time.Now().Format("2006-01-02 15:04:05"), "adminOrgan": "工商局", "adminContent": "许可", "decisionDocumentNo": "004527548875465", "company_id": id,
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.LicenseInfo, j)
}

//失信被执行人
func UpdateDishonest(id string) string {
	//行政许可
	var data = []map[string]string{
		{
			"referNumber": "10121245421", "executionCourt": "北京市海淀区中级人民法院", "publishDate": time.Now().Format("2006-01-02 15:04:05"), "performance": "已履行",
			"person": "xxx", "registerDate": time.Now().Format("2006-01-02 15:04:05"), "documentNumber": "004527548875465", "company_id": id,
		},
		{
			"referNumber": "4545124548", "executionCourt": "北京市海淀区中级人民法院", "publishDate": time.Now().Format("2006-01-02 15:04:05"), "performance": "未履行",
			"person": "xxx", "registerDate": time.Now().Format("2006-01-02 15:04:05"), "documentNumber": "45412457454545", "company_id": id,
		},
		{
			"referNumber": "784551212415454", "executionCourt": "北京市海淀区中级人民法院", "publishDate": time.Now().Format("2006-01-02 15:04:05"), "performance": "已履行",
			"person": "xxx", "registerDate": time.Now().Format("2006-01-02 15:04:05"), "documentNumber": "7845121241548784", "company_id": id,
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.UpdateDishonest, j)
}

//经营异常
func UpdateAbnormal(id string) string {
	//行政许可
	var data = []map[string]string{
		{
			"inclusionDate": time.Now().Format("2006-01-02 15:04:05"), "inCause": "未审核通过", "decisionMakeAuth": "公安局", "removalDate": time.Now().Format("2006-01-02 15:04:05"),
			"outCause": "通过", "company_id": id,
		},
		{
			"inclusionDate": time.Now().Format("2006-01-02 15:04:05"), "inCause": "未审核通过", "decisionMakeAuth": "公安局", "removalDate": time.Now().Format("2006-01-02 15:04:05"),
			"outCause": "通过", "company_id": id,
		},
		{
			"inclusionDate": time.Now().Format("2006-01-02 15:04:05"), "inCause": "未审核通过", "decisionMakeAuth": "公安局", "removalDate": time.Now().Format("2006-01-02 15:04:05"),
			"outCause": "通过", "company_id": id,
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.UpdateAbnormal, j)
}

//法律诉讼
func UpdateJudicial(id string) string {
	//行政许可
	var data = []map[string]string{
		{
			"publishDate": time.Now().Format("2006-01-02 15:04:05"), "caseCause": "恶意伤人", "caseName": "闹事", "caseNumber": "1121545412215454211",
			"caseIdentity": "通过", "executionCourt": "北京市中级人民法院", "company_id": id,
		},
		{
			"publishDate": time.Now().Format("2006-01-02 15:04:05"), "caseCause": "打击", "caseName": "打击", "caseNumber": "9845454154557415245777",
			"caseIdentity": "通过", "executionCourt": "北京市中级人民法院", "company_id": id,
		},
		{
			"publishDate": time.Now().Format("2006-01-02 15:04:05"), "caseCause": "网络攻击", "caseName": "网络攻击", "caseNumber": "878454515154577878787",
			"caseIdentity": "通过", "executionCourt": "北京市中级人民法院", "company_id": id,
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.UpdateJudicial, j)
}

//调用产品内容监管
func UpdateSafeScanInfo() string {
	var data = []map[string]string{
		{
			"scanCount": "1000", "passCount": "1000", "scanDate": time.Now().Format("2006-01-02 15:04:05"), "scanType": "1",
		},
		{
			"scanCount": "2000", "passCount": "1700", "scanDate": time.Now().Format("2006-01-02 15:04:05"), "scanType": "2",
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.UpdateSafeScanInfo, j)
}

//调用产品功能监管
func UserFeed(postValue url.Values) string {
	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.UserFeed, []byte(postString))
}

//调用产品功能监管
func UserFeedTouSu(dt string) string {
	postValue := url.Values{
		"role":         {"1"},
		"userId":       {"10"},
		"feedBackType": {"1"},
		"supportType":  {"3"},
		"description":  {"反馈意见"},
		"create_at":    {dt},
	}

	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.UserFeed, []byte(postString))
}

//调用产品功能监管
func UpdateUserFeedbackInfo(Id string) string {
	postValue := url.Values{
		"id":             {Id},
		"feedBackStatus": {"2"},
	}

	postString := postValue.Encode()
	//fmt.Println(setconst.URI + setslice.UpdateUserFeedbackInfo + "=================")
	return httprun.SendFormUnicode(setconst.URI+setslice.UpdateUserFeedbackInfo, []byte(postString))
}

//调用产品功能监管
func UpdateInfo() string {
	postValue := url.Values{
		"update_at":          {time.Now().Format("2006-01-02 15:04:05")},
		"schoolLicenseUrl":   {"http://www.baidu.com"},
		"businessLicenseUrl": {"http://www.baidu.com"},
	}

	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.UpdateInfo, []byte(postString))
}

//舆情监管
func AppCommentInfo() string {
	postValue := url.Values{
		"text":     {"垃圾"},
		"scanDate": {"2019-12-03 00:00:00"},
	}

	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.AppCommentInfo, []byte(postString))
}

//http测试
func TestHttp() string {
	postValue := url.Values{}
	postValue.Set("a[0]", "b") //post数组数据
	postValue.Set("a[1]", "cb")

	postString := postValue.Encode()
	return httprun.SendFormUnicode(setconst.URI+setslice.TestHttp, []byte(postString))
}

func DoHttp() string {
	postValue := url.Values{}
	postString := postValue.Encode()
	return httprun.SendFormUnicode(setslice.DoHttp, []byte(postString))
}

//增加用户登录信息
func AddUserLoginInfo() string {
	var data = []map[string]string{
		{
			"role": "1", "userId": "1", "loginTime": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"role": "3", "userId": "2", "loginTime": "2019-10-25 00:00:00",
		},
		{
			"role": "3", "userId": "3", "loginTime": "2019-09-25 00:00:00",
		},
		{
			"role": "1", "userId": "1", "loginTime": "2019-08-25 00:00:00",
		},
		{
			"role": "2", "userId": "2", "loginTime": "2019-07-25 00:00:00",
		},
		{
			"role": "3", "userId": "3", "loginTime": "2019-06-25 00:00:00",
		},
		{
			"role": "4", "userId": "7", "loginTime": "2019-10-25 00:00:00",
		},
		{
			"role": "4", "userId": "8", "loginTime": "2019-09-25 00:00:00",
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddUserLoginInfo, j)
}

//增加用户时长信息
func AddOnlineTimeInfo() string {
	var data = []map[string]string{
		{
			"role": "1", "userId": "1", "startTime": time.Now().Format("2006-01-02 15:04:05"), "endTime": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"role": "1", "userId": "1", "startTime": "2019-10-25 00:00:00", "endTime": "2019-10-25 23:00:00",
		},
		{
			"role": "2", "userId": "2", "startTime": "2019-09-25 00:00:00", "endTime": "2019-09-25 23:00:00",
		},
		{
			"role": "2", "userId": "2", "startTime": "2019-08-25 00:00:00", "endTime": "2019-08-25 23:00:00",
		},
		{
			"role": "3", "userId": "3", "startTime": "2019-07-25 00:00:00", "endTime": "2019-07-25 23:00:00",
		},
		{
			"role": "3", "userId": "3", "startTime": "2019-06-25 00:00:00", "endTime": "2019-06-25 23:00:00",
		},
		{
			"role": "4", "userId": "4", "startTime": time.Now().Format("2006-01-02 15:04:05"), "endTime": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"role": "4", "userId": "4", "startTime": "2019-10-25 00:00:00", "endTime": "2019-10-25 23:00:00",
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddOnlineTimeInfo, j)
}

//产品经营监管
func AddCourseInfo() string {
	var data = []map[string]string{
		{
			"courseId": "1", "price": "120012.12", "subjectId": "1", "gradeId": "2", "versionId": "3", "courseName": "语文", "description": "课程大纲:\r\nLesson 1",
			"payType": "1", "courseDuration": "12", "courseCount": "10", "studyType": "1", "createTime": time.Now().Format("2006-01-02 15:04:05"), "startTime": time.Now().Format("2006-01-02 15:04:05"),
			"endTime": time.Now().Format("2006-01-02 15:04:05"), "teacherId": "1", "isForgien": "1", "isLive": "1", "regCount": "1", "finishRate": "60.0%",
		},
		{
			"courseId": "2", "price": "110000.12", "subjectId": "1", "gradeId": "2", "versionId": "4", "courseName": "数学", "description": "课程大纲:\r\n第二个知识点",
			"payType": "1", "courseDuration": "12", "courseCount": "10", "studyType": "1", "createTime": time.Now().Format("2006-01-02 15:04:05"), "startTime": time.Now().Format("2006-01-02 15:04:05"),
			"endTime": time.Now().Format("2006-01-02 15:04:05"), "teacherId": "1", "isForgien": "1", "isLive": "0", "regCount": "0", "finishRate": "60.0%",
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddCourseInfo, j)
}

//增加中小学
func AddTeacherInfo() string {
	var data = []map[string]string{
		{
			"name": "语文老师", "userId": "3", "idCardNumber": "001215454574", "certificationUrl": "http://www.baidu.com", "foreignCertUrl": "http://www.baidu.com", "createDate": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"name": "数学老师", "userId": "4", "idCardNumber": "001415452645", "certificationUrl": "http://www.baidu.com", "foreignCertUrl": "http://www.baidu.com", "createDate": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddTeacherInfo, j)
}

//增加学生 AddStudentInfo
func AddStudentInfo() string {
	var data = []map[string]string{
		{
			"name": "小张", "userId": "5", "districtId": "001215454574", "gradeId": "1", "registrationDate": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"name": "小李", "userId": "6", "districtId": "001415452645", "gradeId": "2", "registrationDate": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddStudentInfo, j)
}

//增加家长 AddParentInfo
func AddParentInfo() string {
	var data = []map[string]string{
		{
			"name": "小张家长", "userId": "7", "registrationDate": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"name": "小李家长", "userId": "8", "registrationDate": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddParentInfo, j)
}

//AddAuthInfo 增加教育部门

func AddAuthInfo() string {
	var data = []map[string]string{
		{
			"name": "小张教育部门", "userId": "9", "registrationDate": time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			"name": "小李教育部门", "userId": "10", "registrationDate": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddAuthInfo, j)
}

//产品经营监管
func AddOrderInfo() string {
	var data = []map[string]string{
		{
			"orderId": "5", "orderType": "1", "name": "一本书", "price": "15000.61", "courseId": "1", "orderState": "2", "createTime": "2019-11-25 00:00:00",
			"payTime": "2019-11-25 00:00:00", "orderDescription": "购买一本书", "goodsProperty": "attr", "userId": "1", "role": "2",
		},
		{
			"orderId": "6", "orderType": "2", "name": "一棵草", "price": "18000.61", "courseId": "2", "orderState": "2", "createTime": "2019-10-25 00:00:00",
			"payTime": "2019-10-25 00:00:00", "orderDescription": "购买一棵草", "goodsProperty": "attr", "userId": "2", "role": "3",
		},
		{
			"orderId": "7", "orderType": "3", "name": "衣服", "price": "10000", "courseId": "3", "orderState": "2", "createTime": "2019-09-25 00:00:00",
			"payTime": "2019-09-25 00:00:00", "orderDescription": "购买衣服", "goodsProperty": "attr", "userId": "100", "role": "3",
		},
		{
			"orderId": "8", "orderType": "1", "name": "玉米", "price": "18000.23", "courseId": "1", "orderState": "2", "createTime": "2019-08-25 00:00:00",
			"payTime": "2019-08-25 00:00:00", "orderDescription": "玉米", "goodsProperty": "attr", "userId": "1", "role": "2",
		},
		{
			"orderId": "9", "orderType": "1", "name": "大豆", "price": "100000", "courseId": "1", "orderState": "2", "createTime": "2019-07-25 00:00:00",
			"payTime": "2019-07-25 00:00:00", "orderDescription": "大豆", "goodsProperty": "attr", "userId": "1", "role": "2",
		},
	}

	j, _ := json.Marshal(data)
	return httprun.SendForm(setconst.URI+setslice.AddOrderInfo, j)
}
