package main

import "unicode"
import "strings"
// import "fmt"

func RemoveEven(a []int) []int {
    n := len(a)
    b := []int{}
    for i := 0; i < n; i++ {
        if a[i] % 2 != 0 {
            b = append(b, a[i])
        }
    }
    return b
}

func PowerGenerator(x int) func() int {
    y := 1
    return func() int {
        y *= x
        return y
    }
}

func DifferentWordsCount(s string) int {
    prev := '.'
    ans := 0
    has := map[string]bool{}
    cur := ""

    s += "."

    for _, c := range s {
        if unicode.IsLetter(prev) && !unicode.IsLetter(c) {
            cur = strings.ToLower(cur)
            if _, ok := has[cur]; !ok {
                ans++
                has[cur] = true
            }
            cur = ""
        }
        if unicode.IsLetter(c) {
            cur += string(c)
        }
        prev = c
    }
    return ans
}

// func main() {
//     // input := []int{0, 3, 2, 5}
//     // result := RemoveEven(input)
//     // fmt.Println(result)

//     // gen := PowerGenerator(2)
//     // fmt.Println(gen()) // Должно напечатать 2
//     // fmt.Println(gen()) // Должно напечатать 4
//     // fmt.Println(gen()) // Должно напечатать 8

//     fmt.Println(DifferentWordsCount("Hello, w_orld!HELLO  wOrlD...12keke"))
// }
