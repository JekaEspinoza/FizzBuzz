package main

import (
"fmt"
)

func main() {
for i := 1; i <= 100; i++ {
if test := FizzBuzz(i); test != "" {
fmt.Println(i, test)
continue
}
fmt.Println(i)
}
}

func FizzBuzz(val int) (res string) {
if val % 3 == 0 {
res += "Fizz"
}
if val % 5 == 0 {
res += "Buzz"
}

return res

}
