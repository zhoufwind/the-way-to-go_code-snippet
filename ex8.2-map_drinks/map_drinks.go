// Exercise 8.2
// Construct a collection which maps English names of drinks to the French (or your native language)
// translations; print first only the drinks available, and then print both (the name and the translation).
// Then produce the same output, but this time the English names of the drinks must be sorted.

// 构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。

package main

import "fmt"
import "sort"

var (
    drinkMap = map[string]string{"cola": "可乐", "sprite": "雪碧", "mirinda": "美年达", "seven-up": "七喜", "coffee": "咖啡", "tea": "茶"}
)

func main() {
    var drinkArr = make([]string, len(drinkMap))
    i := 0
    for k, _ := range drinkMap {
        drinkArr[i] = k
        i++
    }
    fmt.Println(drinkArr)

    for k, v := range drinkMap {
        fmt.Printf("English: %v, Chinese: %v\n", k, v)
    }

    sort.Strings(drinkArr)
    fmt.Println("sorted:")
    fmt.Println(drinkArr)
}

/*
[cola sprite mirinda seven-up coffee tea]
English: coffee, Chinese: 咖啡
English: tea, Chinese: 茶
English: cola, Chinese: 可乐
English: sprite, Chinese: 雪碧
English: mirinda, Chinese: 美年达
English: seven-up, Chinese: 七喜
sorted:
[coffee cola mirinda seven-up sprite tea]
*/
