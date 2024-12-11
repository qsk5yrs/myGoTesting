package util

import "strings"

// 用于混淆隐藏敏感信息的工具函数

func MaskLoginName(loginName string) string {
	// 判断loginName是邮箱还是手机号
	if strings.LastIndex(loginName, "@") != -1 {
		return MaskEmail(loginName)
	}

	return MaskPhone(loginName)
}

// MaskPhone 隐藏手机号中间4位地区码，如155****8888
func MaskPhone(phone string) string {
	if n := len(phone); n >= 8 {
		return phone[:n-8] + "****" + phone[n-4:]
	}

	return phone
}

// MaskEmail 隐藏邮箱ID的中间部分 zhang@go-mall.com ---> z***g@go-mall.com
func MaskEmail(address string) string {
	index := strings.LastIndex(address, "@")
	id := address[0:index]
	domain := address[index:]
	if len(id) <= 1 {
		return address
	}
	switch len(id) {
	case 2:
		id = id[0:1] + "*"
	case 3:
		id = id[0:1] + "*" + id[2:]
	case 4:
		id = id[0:1] + "**" + id[3:]
	case 5:
		id = id[0:1] + "***" + id[4:]
	default:
		masks := strings.Repeat("*", len(id)-4)
		id = id[0:2] + masks + id[len(id)-2:]
	}

	return id + domain
}

// MaskRealName 保留姓名首末位 如：张三--->张* 赵丽颖--->赵*颖 欧阳娜娜--->欧**娜
func MaskRealName(realName string) string {
	runeRealName := []rune(realName)
	if n := len(runeRealName); n >= 2 {
		if n == 2 {
			return string(append(runeRealName[0:1], rune('*')))
		} else {
			count := n - 2
			newRealName := runeRealName[0:1]
			for temp := 1; temp <= count; temp++ {
				newRealName = append(newRealName, rune('*'))
			}
			return string(append(newRealName, runeRealName[n-1]))
		}
	}

	return realName
}
