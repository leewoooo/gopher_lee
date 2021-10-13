Pointer
===

## Pointer란

포인터는 `Memory의 주소를 값으로 갖는` type입니다. type이다 보니 당연히 선언이 필요하며 선언하는 과정에서 Memory를 할당받게 됩니다.

일반적인 변수는 Memory를 할당받아 그 Memory안에 값을 가지고 있게 되는데 pointer와 같은 경우는 Memory를 할당받아 어떠한 변수의 Memory값을 `할당받은 Memory에 가지고 있게 됩니다.`

golang에서 pointer를 정리하자면 다음과 같습니다.

1. 메모리 주소를 값으로 가진 변수를 pointer라 부릅니다.
2. 포인터는 메모리 주소를 값으로 가지고 있기 때문에 주소가 가지고 있는 값을 역참조가 가능합니다. (pointer가 가리키고 있는 변수의 memory에 있는 값)
3. c언어와 다르게 golang에서는 pointer의 연산을 허용하지 않습니다.

<br>

## usage

포인터의 선언은 다른 변수들의 선언과 다르지 않습니다. 변수를 선언할 때 `type앞에 "*" 를 붙임으로써 해당 변수가 pointer임을 명시합니다.`

변수의 앞에 `&를 붙여 해당 변수 type의 pointer를 pointer변수에 할당` 할 수 있습니다.

```go
// ex
var a int = 8
var b *int = &a

c := *b // 8

var d *int // nil
```
변수 a를 선언 및 초기화를 하여 8이라는 값을 할당 받은 memory에 저장합니다. 

변수 b도 동일하게 선언 및 초기화를 하게 되는데 type은 `int type pointer`이며 값은 `변수 a의 주소값`을 할당받은 memory에 저장합니다. 즉 변수 b는 `변수 a를 가리키는 pointer 변수`입니다. 해당 pointer를 이용하여 `a의 변수의 Memory에 값을 역으로 참조할 수 있습니다. (c 변수)`

변수 d와 같이 선언을 하고 초기화를 하지 않게 되면 pointer의 zero value인 `nil`가 됩니다. nil이란 다름 언어에서 `null`과 같습니다. `즉 아직 memory를 할당 받지 않은 상태 입니다.`

또한 `하나의 변수`를 여러 pointer가 참조할 수도 있습니다.

<br>

## 그럼 pointer를 왜 사용하는 것일까요?

함수를 호출하는 과정에서 해당 함수에 필요한 parameter를 함수 밖에서 함수 안으로 넣어주게 됩니다. 아래와 같이 말입니다.

```go
var a int = 1 //1
func Increase(a int){
    a++ //2
}
Increase(a) 
```

해당 함수를 실행하면 과연 a변수는 값이 2가 될까요? 결과부터 말씀드리면 a변수는 동일하게 1입니다.

`그 이유는 해당 함수의 parameter로 넣어준 값 a는 내부적으로 복사되어 사용되기 때문에 함수 밖에서의 a변수와(1) 함수 안에서의 a변수(2)는 서로 다른 Memory 공간을 갖게 됩니다.`

즉 함수 안에서 a라는 변수를 새로 만들어(2) parameter로 들어온 a변수(1)의 크기만큼 새로운 memory를 할당하여 `통째로 복사해서 사용하게 됩니다.`

이러한 문제를 해결하기 위해 pointer를 사용하게 됩니다.

함수 안에서 parameter로 받는 변수의 `주소에 있는 값을 handling하고 싶을 때` 사용하게 됩니다. 즉 `memory의 복사를 피하고 싶을 때 사용합니다.(memory를 아낄 수 있습니다.)`

그럼 위의 code를 pointer를 이용하여 변경하면 다음과 같습니다.
```go
var a int = 1
func Increase(a *int){
    *a++
}
Increase(&a)
```
<br>

## hip, stack(with pointer)

스택 메모리는 `함수 호출 스택을 저장하고` 로컬 변수, 인수, 반환 값도 여기에 저장하게 됩니다. 

힙 메모리는 함수 호출 스택과는 관계가 없고 `함수 범위에 얽매이지 않고 객체(instance)를 저장해 둡니다.` 힙 메모리에 있는 객체 중 참조되지 않는 객체가 존재한다면 다음 GC Time에 객체를 회수하게 됩니다. (회수할 때 처리 비용이 듭니다.)

golang은 Compiler가 해당 객체를 `stack에 둘 것인지 heap에 둘 것인지 결정을 하게 됩니다.`


```go
type User struct{
    Name string
    Age int
}

func NewUser(name string, age int) *User{
    var a string = "foo"
    
    u := &user{Name : name, Age: age}
    return u
}
```

예제 code를 c언어 입장에서 보자면 이해가 되지 않는 code일 것입니다. 그 이유는 `a변수와 u변수는 stack에 push되게 되고 함수 호출이 끝나게 되면 해당 변수들은 pop되어 소멸되기 때문입니다.` 

그럼 NewUser 함수를 호출한 곳에서 `return값으로 받은 User type의 pointer 변수를 참조하는 곳에서 dangling이 발생합니다.`
>포인터가 여전히 해제된 메모리 영역을 가리키고 있다면, 이러한 포인터를 댕글링 포인터(Dangling Pointer)라고 한다.

<br>

하지만 golang에서는 다르다!!

위에서 이야기 한 것처럼 `객체를 stack에 저장할 지 heap에 저장할지 compiler가  Escape Analysis 통해 결정한다.`
>Escape Analysis : 컴파일러 최적화에서 이스케이프 분석은 프로그램에서 포인터에 액세스 할 수있는 포인터의 동적 범위를 결정하는 방법입니다. 포인터 분석 및 모양 분석과 관련이 있습니다. 변수가 서브 루틴에 할당되면 변수에 대한 포인터가 다른 실행 스레드 또는 서브 루틴 호출로 벗어날 수 있습니다.

`컴파일러가 함수 내에서 확보한 값이 함수 밖에서도 필요하게 된다면 stack에서 heap으로 옮긴다.`

<Br>

컴파일러 플래그를 전달하여 testcode를 돌려보면 다음과 같이 출력된다.
```go
// bash
go test -gcflags -m   ./basic/pointer

// output
...
basic/pointer/pointer_test.go:67:17: moved to heap: u
...
```

<br>

## REFRENCE

https://jacking75.github.io/go_stackheap/

https://en.wikipedia.org/wiki/Escape_analysis