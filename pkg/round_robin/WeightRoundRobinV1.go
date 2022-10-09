package main

import (
	"errors"
	"fmt"
	"time"
)

//[go实现路由加权算法]https://baijiahao.baidu.com/s?id=1737148450169550635&wfr=spider&for=pc
//https://www.jb51.net/article/231401.htm

type RoundRobinBalance struct {
	currentIdx int
	ass        []string
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) <= 0 {
		return errors.New("params len error")
	}
	for _, adds := range params {
		r.ass = append(r.ass, adds)
	}
	return nil
}

func (r *RoundRobinBalance) Next() string {
	slen := len(r.ass)
	if slen <= 0 {
		return ""
	}
	if r.currentIdx >= slen {
		r.currentIdx = 0
	}
	currentAdds := r.ass[r.currentIdx]
	r.currentIdx = (r.currentIdx + 1) % slen
	return currentAdds
}

func main() {
	rs := new(RoundRobinBalance)
	var err error
	err = rs.Add("127.0.0.1:80", "127.0.0.1:81", "127.0.0.1:82", "127.0.0.1:83", "127.0.0.1:84")
	err = rs.Add()
	defer func() {
		if errs := recover(); errs != nil {
			errObj := fmt.Errorf("panic error: %v", errs)
			fmt.Println(errObj.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
	for {
		adds := rs.Next()
		fmt.Println(adds)
		time.Sleep(time.Second * 1)
	}
}
