package round_robin

import (
	"errors"
	"fmt"
	"go_project/pkg/request"
	"syscall"
)

//[go实现路由加权算法]https://baijiahao.baidu.com/s?id=1737148450169550635&wfr=spider&for=pc
//https://www.jb51.net/article/231401.htm

type RoundRobinBalance struct {
	adds       []string
	currentIdx int
}

var r RoundRobinBalance

func init() {
	//todo conf
	appEnv, isSet := syscall.Getenv("APP_ENV")
	if appEnv == "" || !isSet {
		appEnv = "dev"
		syscall.Setenv("APP_ENV", appEnv)
	}

	if appEnv == "dev" {
		r.Add([]string{"http://101.35.50.230:30012"})
	} else {
		r.Add([]string{"http://php-project-service.ns-test.svc.cluster.local:8100"})
	}
	fmt.Println("init: adds config")
}

func (r *RoundRobinBalance) Add(addss []string) (*RoundRobinBalance, error) {
	if len(addss) < 0 {
		return nil, errors.New("adds is empty")
	}
	for _, adds := range addss {
		//	todo check adds
		r.adds = append(r.adds, adds)
	}
	return r, nil
}

func (r *RoundRobinBalance) Next() string {
	addlen := len(r.adds)
	if addlen < 0 {
		return ""
	}
	if r.currentIdx > addlen {
		r.currentIdx = 0
	}
	res := r.adds[r.currentIdx]
	r.currentIdx = (r.currentIdx + 1) % addlen
	return res
}

func CallPHP() (string, error) {

	url := r.Next()
	fmt.Println("url:", url)
	return request.Q(url)
}

//func main() {
//	for {
//		adds := r.Next()
//		fmt.Println(adds)
//		time.Sleep(time.Second)
//	}
//}
