package basic_date

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	time := time.Time{}
	if now.After(time) {
		fmt.Println("yes time after")
	}
}

func TestTime2(t *testing.T) {
	d1, _ := time.ParseDuration("-11h")
	t1 := time.Now()
	t2 := time.Now().Add(d1)
	time.Sleep(1 * time.Microsecond)

	fmt.Println(t1.Unix())
	fmt.Println(t2.Unix())
}

func TestFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

func TestParse(t *testing.T) {
	str := "2018-06-22 11:10:17"
	time, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time)
	fmt.Printf("%T\n", time)
}

// 计算两个时间差
func TestTimeNow(tt *testing.T) {
	time1 := time.Now()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(time1))
}

func TestGetTime(t *testing.T) {
	nt := NewtimeSlide()
	fmt.Println(nt.maxT)
	fmt.Println(nt.minT)
	time.Sleep(time.Second * 5)
	fmt.Println(nt.maxT)
	fmt.Println(nt.minT)
	fmt.Println(nt.maxT.String())
	fmt.Println(nt.minT.String())
}

type timeSlide struct {
	maxT time.Duration
	minT time.Duration
}

func NewtimeSlide() *timeSlide {
	return &timeSlide{time.Second * 10, time.Second * 5}
}
