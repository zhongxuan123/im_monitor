package main

import (
	"net/http"
	"fmt"
)


func HeartDo(name string, path string) {
	fmt.Println("HeartDo - path:",path)
	resp, err := http.Get(path)
	if err != nil  {
		monitorHeart.Lock()
		times,ok := monitorHeart.heartMap[name]
		if !ok {
			times = 0
		}
		times = times+1
		monitorHeart.heartMap[name] = times
		monitorHeart.Unlock()
		return 
	}else if resp.StatusCode != 200{
		monitorHeart.Lock()
		times,ok := monitorHeart.heartMap[name]
		if !ok {
			times = 0
		}
		times = times+1
		monitorHeart.heartMap[name] = times
		monitorHeart.Unlock()
	}
	monitorHeart.Lock()
	monitorHeart.heartMap[name] = 0
	monitorHeart.Unlock()
	resp.Body.Close()
	fmt.Println(name, " HTTPDo success.")
}
