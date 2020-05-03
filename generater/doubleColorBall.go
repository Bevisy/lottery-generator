package generater

import (
	"log"
	"math/rand"
	"sort"
	"strconv"
)

func DoubleColorBall() []string {
	var str []string
	red := randomBall("red")
	sort.Ints(red)
	for _, v := range red {
		str = append(str, func() string {
			if v <= 9 {
				return "0" + strconv.Itoa(v)
			} else {
				return strconv.Itoa(v)
			}
		}())
	}

	// 添加分隔符
	str = append(str, "|")

	blue := randomBall("blue")
	str = append(str, func(v int) string {
		if v <= 9 {
			return "0" + strconv.Itoa(v)
		} else {
			return strconv.Itoa(v)
		}
	}(blue[0]))

	return str
}

func randomBall(ball string) []int {
	//rand.Seed(time.Now().UnixNano()) // 禁止被调用函数内使用获取seed，并发或者高性能场景下，获取到的seed会是一致的，应该在全局定义一次seed
	var red []int  // 红球
	var blue []int // 蓝球
	if ball == "red" {
		for i := 0; i < 6; i++ {
			red = append(red, rand.Intn(33)+1)
		}
		return red
	} else if ball == "blue" {
		blue = append(blue, rand.Intn(16)+1)
		return blue
	} else {
		log.Fatal("The input ball color is INVALID.")
		return nil
	}
}
