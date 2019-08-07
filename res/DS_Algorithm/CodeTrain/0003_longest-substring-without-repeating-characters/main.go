package main

import "fmt"

//给定一个字符串，找出不含有重复字符的最长子串的长度。
//
//示例：
//
//给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
//
//给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
/
//给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。


// 滑动窗口算法
func lengthOfLongestSubstring(s string) int {
    // 初始化
    MaxLen := 0
    StringStart := 0 // 重复字符上次出现的索引，上一个重复字符和当前重复字符的之间的子串即有效子串
    StringMap := make(map[int32]int, len(s)) // 减少内存分配

    for StringEnd, char := range s {
        // 字符重复，从重复字符下一位开始重新算子串
        if val, ok := StringMap[char]; ok{
            if val > StringStart {
                StringStart = val
            }
        }
        // 存入Hash中
        temp := StringEnd - StringStart + 1
        if temp > MaxLen {
            MaxLen = temp
        }
        StringMap[char] = StringEnd + 1
    }
    return MaxLen
}


func main() {
	//len := lengthOfLongestSubstring(" ")
	//len := lengthOfLongestSubstring("dvdf")
	len := lengthOfLongestSubstring("abcabcbb")
	//len := lengthOfLongestSubstring("abba")
	fmt.Println("len=", len)
}
