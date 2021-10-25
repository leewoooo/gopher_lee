Interface
===

## 인터페이스란?

인터페이스는 객체가 구현할 수 있는 method들을 `function 시그니처의` 집합이라 할 수 있으며 구체화 된 객체(concret object)가 아닌 `추상화` 된 method를 이용해 객체 간의 `상호작용으로 관계를 표현합니다.`(go에서는 class가 없지만 method로 객체 간 관계를 표현)

인터페이스의 기본 역할은 메소드명, 매개변수 입력값, 반환값의 타입을 제공하는 일 입니다. 

<br>

## Syntax

interface 또한 type입니다. `즉 interface 또한 변수로 선언 할 수 있습니다.`

interface를 정의하는 방법은 `struct`와 비슷합니다. 하지만 field를 갖지 않고 `구현되지 않은 method들을` interface body안에 정의합니다.

```go
type interface명 interface{
    method명(parameter...) (return...)
    ...
}
```

### interface의 선언 규칙

1. method는 반드시 method명이 있어야 합니다.
2. parameter와 return type이 다르고 `동일한 method명을 가진 method를 선언할 수 없습니다.`
3. interface의 method는 `구현을 포함하지 않습니다.`

<br>

## interface를 왜 사용할까?

`code의 의존성을 줄이기 위해 사용합니다.` interface를 정의하고 구현한 구현체로 code를 작성하면 이 후 `initialize 부분만 변경하여 code의 유지보수 측면에서 이점을 가져갈 수 있습니다.`(Decoupling)

또한 `interface를 주입받는 객체를 test하는 code를 작성할 때` interface를 구현하는 `mock 객체를 만들어 test의 비용을 줄일 수 있습니다.`

<br>

## inteface의 구현

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



## 인터페이스의 내부적인 포인터 사용




## REFERENCE
https://www.popit.kr/golang%EC%9C%BC%EB%A1%9C-%EB%A7%8C%EB%82%98%EB%B3%B4%EB%8A%94-duck-typing/

https://hoonyland.medium.com/%EB%B2%88%EC%97%AD-interfaces-in-go-d5ebece9a7ea