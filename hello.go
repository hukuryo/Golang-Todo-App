package main

import "fmt"

// 変数
var lang2 = "Golang" //初期値を渡す。

// 定数
const Num = 2

func main() {
   fmt.Println(lang2)
   fmt.Println(Num)                    // => 2
   fmt.Println(hello("Hello World"))   //=> Hello World
   fmt.Println(hello2("Hello", "World")) //=> Hello World
   fmt.Println(multipleArgs("Hello", "World")) //=> World Hello
   forRoup()
   fmt.Println(condition("Swift"))
}

//ex1:引数の戻り値の型を書く必要がある
func hello(arg string) string {
   return arg
}

//ex2:関数が２つ以上の同種類の引数を伴う際、型の省略する事が可能
func hello2(arg1, arg2 string) string {
   return arg1 + " " + arg2
}

//ex3:関数は複数の値を返すことができる
func multipleArgs(arg1, arg2 string) (string, string) {
   return arg2, arg1
}

func forRoup() {
   sum := 0
   for i := 0; i < 5; i++ {
   	sum += i
   }
   fmt.Println(sum) //=> 10
}

func condition(arg string)string{
	if v := "GO"; arg == v {
		return "This is Golang"
	}else{
		return "This is not Golang"
	}
}