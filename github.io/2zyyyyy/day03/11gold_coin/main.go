package main

import "fmt"

/* 练习题：分金币
你有五十枚金币，需要分给以下几个人：
1.Matthew 2.Sarah 3.Augustus 4.Heidi 5.Emilie 6.Peter 7.Giana 8.Adriano 9.Aaron 10.Elizabeth
分配规则如下：
a.名字中每包含一个'e'或'E'分1枚金币
b.名字中每包含一个'i'或'I'分2枚金币
c.名字中每包含一个'o'或'O'分2枚金币
d.名字中每包含一个'u'或'U'分2枚金币
用golang编写程序，计算每个用户分到多少枚金币，以及最后剩余多少金币？
程序结构如下，请实现'dispatchCoin' */
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	/* 	思路：
	   	1.依次拿到每个人的名字，遍历users
	   	2.获取到一个名字根据规则分配金币
	   	  2.1.没人分得金币数应该保存到distribution中
		  2.2.计算剩余金币数量 */
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				// 满足该条件给给用户分配金币
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	// 第二种计算剩余金币方法 计算出map中所有value的和 然后 用金币总数减去
	count := 0
	for n := range distribution {
		count += distribution[n]
	}
	fmt.Println(count)
	fmt.Println("剩下：", left, (50 - count))
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}
