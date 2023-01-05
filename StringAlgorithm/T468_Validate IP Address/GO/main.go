package main

import "fmt"

func main() {
	queryIP := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	fmt.Println("Output:", validIPAddress(queryIP))
}

func validIPAddress(queryIP string) string {
	ipv4SeparatorCount := 0
	ipv6SeparatorCount := 0
	for i := 0; i < len(queryIP); i++ {
		if queryIP[i] == '.' {
			ipv4SeparatorCount++
		} else if queryIP[i] == ':' {
			ipv6SeparatorCount++
		}
	}
	if ipv4SeparatorCount != 3 && ipv6SeparatorCount != 7 {
		return "Neither"
	} else if ipv4SeparatorCount == 3 {
		if isValidIPv4(queryIP) {
			return "IPv4"
		} else {
			return "Neither"
		}

	} else if ipv6SeparatorCount == 7 {
		if isValidIPv6(queryIP) {
			return "IPv6"
		} else {
			return "Neither"
		}
	} else {
		return "Neither"
	}

}

func split(s string, separator byte) []string {
	sections := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == separator {
			sections = append(sections, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		sections = append(sections, s[start:])
	}
	return sections
}

func isValidIPv4(queryIP string) bool {
	//分段
	sections := split(queryIP, '.')

	if len(sections) != 4 {
		return false
	}

	// 判断每一段是否合法
	for _, section := range sections {
		if len(section) == 0 {
			return false
		} else if len(section) > 1 && section[0] == '0' {
			return false
		} else {
			//转换成数字，再比较
			sectionInt := 0
			for i := 0; i < len(section); i++ {
				if section[i] < '0' || section[i] > '9' {
					return false
				}
				sectionInt = sectionInt*10 + int(section[i]-'0')
			}
			if sectionInt < 0 || sectionInt > 255 {
				return false
			}

		}
	}

	return true

}

func isValidIPv6(queryIP string) bool {
	//分段
	sections := split(queryIP, ':')
	if len(sections) != 8 {
		return false
	}
	for _, section := range sections {
		if len(section) < 1 || len(section) > 4 {
			return false
		}
		for i := 0; i < len(section); i++ {
			if !((section[i] >= '0' && section[i] <= '9') || (section[i] >= 'A' && section[i] <= 'F') || (section[i] >= 'a' && section[i] <= 'f')) {
				return false
			}
		}
	}
	return true
}
