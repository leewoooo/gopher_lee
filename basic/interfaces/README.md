Interface
===

## 인터페이스란?

인터페이스는 객체가 구현할 수 있는 method들을 `function 시그니처의` 집합이라 할 수 있으며 구체화 된 객체(concret object)가 아닌 `추상화` 된 method를 이용해 객체 간의 `상호작용으로 관계를 표현합니다.`(go에서는 class가 없지만 method로 객체 간 관계를 표현)

인터페이스의 기본 역할은 메소드명, 매개변수 입력값, 반환값의 타입을 제공하는 일 입니다. 

<br>

### Syntax

interface 또한 type입니다. `즉 interface를 변수로 선언 할 수 있습니다.`

interface를 정의하는 방법은 `struct`와 비슷합니다. 하지만 field를 갖지 않고 `구현되지 않은 method들을` interface body안에 정의합니다.

```go
type interface명 interface{
    method명(parameter...) (return...)
    ...
}
```

<br>

### interface의 선언 규칙

1. method는 반드시 method명이 있어야 합니다.
2. parameter와 return type이 다르고 `동일한 method명을 가진 method를 선언할 수 없습니다.`
3. interface의 method는 `구현을 포함하지 않습니다.`

<br>

### interface를 왜 사용할까?

`code의 의존성을 줄이기 위해 사용합니다.` interface를 정의하고 구현한 구현체로 code를 작성하면 이 후 `initialize 부분만 변경하여 code의 유지보수 측면에서 이점을 가져갈 수 있습니다.`(Decoupling)

또한 `interface를 주입받는 객체를 test하는 code를 작성할 때` interface를 구현하는 `mock 객체를 만들어 test의 비용을 줄일 수 있습니다.`

결정적으로 `다형성을 이용하기 위해서 입니다.(Polymorphism)`
>다형성이란 프로그램 언어 각 요소들(상수, 변수, 식, 객체, 메소드 등)이 다양한 자료형(type)에 속하는 것이 허가되는 성질을 가리킨다. - 위키피디아 중 -

정리하자면 다형성이란 하나의 type에 여러 객체를 대입할 수 있는 성질로 이해하면 될 것입니다. 다형성을 활용하면 기능을 확정하거나, 객체를 변경해야 할 때 타입 변경 없이 객체의 주입만으로 수정이 일어나게 할 수 있습니다. 

[interface in go](https://medium.com/rungo/interfaces-in-go-ab1601159b3a) 에서는 interface는 go에서 다형성을 구현할 수 있는 유일한 방법이라 이야기 하고 있습니다.

<br>

### inteface의 구현

OOP언어에 익숙하다면 interface를 구현하는 과정에서 `implement`라는 keyword를 종종 보았을 것입니다.

Java와 같은 경우 class가 interface를 구현할 때 implement keyword를 이용해 `자기가 누구인지 밝히며 interface를 구현합니다.`

하지만!

go에서는 implement와 같이 `keyword를 명시적으로 사용하지 않습니다.` interface가 가지고 있는 method들을 `리시버를 이용해 객체에 종속시켜` 해당 interface를 구현할 수 있습니다. -> duck typing

<br>

## Duck Typing

>만약 어떤 새가 오리처럼 걷고, 오리처럼 헤엄치고, 오리처럼 꽥꽥거린다면 나는 그 새를 오리라고 부를 것이다. (덕 테스트)

go에서는 duck typing을 지원합니다 duck typing을 살펴보기 전에 정적, 동적 언어의 특징과 차이점을 간단하게 알아보기를 원합니다.

<br>

### Static, Dynamic Language

정적언어의 특징은 다음과 같습니다.

- 변수 또는 함수의 type을 미리 지정해야 합니다.
- 컴파일 시점에 type Check가 가능합니다.

동적 언어의 특징은 다음과 같습니다.

- 변수 또는 함수의 type을 지정하지 않고 사용할 수 있습니다.
- type Check는 runtime 시점에서만 알 수 있습니다.
- Type지정이 안되어 있는 소스 분석이 어렵습니다.

두 언어의 가장 큰 차이점은 `Type정보를 지정해야 하나? 아니냐?`로 볼 수 있습니다.

go언어는 `정적 type의 언어입니다.` 하지만 `동적 언어의 특성 또한 수용하였습니다.`

즉 컴파일러의 보장을 받으면서 동적 언어의 장점을 사용할 수 있다는 것 입니다.(동적 언어의 장점으로는 문법적인 간결함과 잦은 기능변경이 필요한 곳에서 변경이 용이하다는 점입니다.)

이를 가능하게 해주는 것이 바로 `duck typing` 입니다.(go는 정적 type의 언어이다 보니 해당 type이 interface를 구현했는지 complie time에 check를 하게 됩니다.)

<br>

### duck typing을 왜 사용할까?

결론부터 이야기 하자면 `사용자 중심의 code작성이 가능해지기 때문입니다.`

interface의 구현 여부를 type선언시 하지 않고 interface가 사용 될 때 결정하기 때문에 서비스 제공자는 `구체화 된 객체를 제공하고` 사용자가 `필요에 따라 interface를 정의하여 사용할 수 있기 때문입니다.` 


<br>

## 인터페이스의 내부적인 포인터 사용

인터페이스는 두가지 유형의 type을 내부적으로 갖습니다. 첫번째로는 interface를 구현한 `인스턴스의 memory주소의 값`, 두번째는 interface를 구현한 `인스턴스의 type 정보`를 가지고 있습니다.

즉 인터페이스는 `인터페이스를 구현한 타입의 인스턴스를 가리키고 있는 것 입니다.` 이로 인해 인터페이스를 구현한 인스턴스를 숨기고 `인스턴스를 가리키고 있는 인터페이스를 노출 시킬 수 있습니다.`

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r *Rect) Perimeter() float64 {
	return 2 * (r.width * r.height)
}
```

Shape라는 interface가 존재하고 해당 인터페이스의 method를 구현하여 Rect라는 구조체의 `리시버를 이용하여 종속시켰습니다. (Rect는 Shape를 구현하였다.)`

이 후 인터페이스를 선언하고 Rect 구조체를 선언 및 초기화 하여 해당 인스턴스를 인터페이스 변수에 할당해보겠습니다.

```go
var shape Shape

rect := Rect{
    width:  3.0,
    height: 4.0,
}

shape = &rect
```

Rect는 Shape라는 인터페이스를 구현하였기 때문에 해당 code가 정상적으로 동작합니다. 이 과정에서 memory는 어떻게 구성되는지 살펴보겠습니다.

<img src = https://user-images.githubusercontent.com/74294325/138889263-1dfafa37-368d-4cfe-815f-0baeabbe9308.png>

간단하게 나타낸 것이지만 `shape는 내부적으로 Rect의 주소와 type 정보를 가지고 있습니다.` 이와 같이 인터페이스는 `인터페이스를 구현한 타입의 포인터를 동적으로 가지고 있다고 볼 수 있습니다.`

<br>

## 인터페이스의 기능 더 알아보기

### 포함된 인터페이스

인터페이스가 인터페이스를 포함할 수 있습니다. io package에서 자주 찾아볼 수 있는데 아래와 같은 code의 모양을 하게 됩니다.

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}
```

ReadCloser는 `Reader와 Closer`를 포함하고 있습니다. 만약 ReadCloser를 구현하려 한다면 Reader의 `Read`, Closer의 `Close`를 구현하면 해당 type은 `ReadCloser를 구현했다고 할 수 있습니다.`

<br>

### 빈 인터페이스

아무런 method도 가지고 있지 않은 인터페이스를 `빈 인터페이스`라고 합니다. 빈 인터페이스는 `interface{}`로 표현을 하는데 빈 인터페이스는 아무런 method도 선언되지 않았기 때문에 `모든 type이 빈 인스턴스를 구현했다고 할 수 있습니다.`

fmt package의 `Println` method를 보면 `빈 인터페이스의 사용을 이해하기 쉬울 것 입니다.`

```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```
Println의 parameter로 빈 인터페이스를 가변인자로 받습니다. Println의 parameter로 모든 type이 가능했던 이유이기도 합니다.

<br>

### 인터페이스의 타입 Assertion

type assertion은 인터페이스가 가지고 있는 실제 값(concrete value)에 접근을 할 수 있게 해줍니다. 즉 인터페이스 아래에 숨겨진 동적 값을 찾아낼 수 있다는 이야기 입니다. 
>A type assertion provides access to an interface value's underlying concrete value.(https://tour.golang.org/methods/15)

type assertion은 인터페이스 변수 뒤에 `.(interface를 구현한 type)`을 붙여 작성하면 됩니다. 그럼 두 개의 return 값을 받을 수 있습니다.

```go
// syntax
value , ok := i.(Type)

// ex
r, ok := shape.(*Rect)
if !ok {
    log.Fatal("interface를 구현한 type은 Rect가 아닙니다.")
}

t.Log("width :", r.width)
t.Log("height :", r.height)
```

1. interface를 구현한 type의 pointer (Type이 interface를 구현하지 않았다면 nil)
2. Type이 i라는 인터페이스를 구현하였는지에 대한 boolean값 (true, false)

type assetion은 `인터페이스 타입을 띄는 구조체의 한 필드에 접근할 때 type assertion을 사용하는 것이 좋다.` 

<br>

## REFERENCE
https://www.popit.kr/golang%EC%9C%BC%EB%A1%9C-%EB%A7%8C%EB%82%98%EB%B3%B4%EB%8A%94-duck-typing/

https://hoonyland.medium.com/%EB%B2%88%EC%97%AD-interfaces-in-go-d5ebece9a7ea

https://2kindsofcs.tistory.com/13