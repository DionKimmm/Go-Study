package main

import (
  "fmt"
  "main/0_MainPackage/something"
)

func main() {
  fmt.Println("Hello, World!")
  something.SayHello()
}

// # 1.0

// Go 언어의 컴파일 실행 파일명은 반드시 main.go이어야 한다.
// 파일명(main.go의 main)과 package명(package main의 main)이 같아야 한다.
// 컴파일 하기 위해서 들어가야 하는 함수는 func main() {} 이다.

// go run main.go : 컴파일 커맨드
// 실행하면 자동으로 컴파일러는 main 패키지와 그 안에 있는 main 함수를 먼저 찾고 실행시킨다.

// go의 컨셉 두가지 
// 1. package : 파일 맨 위에 니가 작성할 패키지의 이름을 작성하는 곳
// 2. func : JS나 C랑 비슷

// # 1.1

// 어떤 함수가 export 되기 위해선 그 함수의 첫글자가 대문자여야 한다.
// something.sayHello 는 되지 않지만, something.SayHello 는 실행됨.