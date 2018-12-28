package basic_date

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"testing"
)

func TestBeegoHttp(t *testing.T) {
	req := httplib.Get("http://www.baidu.com")
	var result interface{}
	req.JSONBody(&result)
	fmt.Println(result)
}
