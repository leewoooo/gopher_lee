Function
===

## 함수는 왜 사용하는 것일까?

궁극적으로 함수(function)을 사용하는 이유는 중복적인 작업을 줄이기 위해서 입니다.

<br>

## 함수의 발전을 통해 이해하는 함수 호출 구조

code로 작성된 프로그램을 사용할 때 IP라는 것이 존재합니다. IP란 `Instruction Point`의 약자로 현 작업 포인트를 이야기 합니다. 

`"기 승 전 결"`을 통해 함수의 발전 과정을 알아보려합니다.

<Br>

### 기

공통적으로 사용되는 code를 함수로 분리하여 작성을 합니다. 그 후 함수가 필요한 곳에서 함수를 호출하고 (`JMP 명령어를 이용하여 함수가 작성되어 있는 곳으로 IP를 옮깁니다.`) 그 후 다시 호출 한 곳으로 IP를 옮겨서 작업을 진행합니다.

`문제점`

함수를 호출하는 것으로 중복 code를 줄일 수 있게 되었지만 함수를 호출 후 항상 같은 IP로 옮겨지는 문제가 발생하였습니다.

### 승

공통적인 code를 함수로 작성하고 호출하지만 동일한 곳으로 돌아간다는 문제점으로 인해 함수를 호출할 때 `return point를 memory에 저장하는 방식으로 발전을 하게 됩니다.`

그로 인해 IP가 옮겨져서 함수를 호출한 후 memory에 저장되어 있는 `return point`를 읽어와서 해당하는 point로 IP를 옮깁니다.

`문제점`

완전이 동일하지는 않지만 비슷한 함수가 계속해서 반복되는 문제가 발생합니다. 예를 들어 같은 형태의 함수이지만 출력만 다른 경우가 생겼습니다.

### 전

return point를 memory에 저장할 때 함수에 필요한 data를 같이 저장해서 사용하는 방식으로 개선되었습니다. 함수에 필요한 data를 받아 저장 후 출력에 사용하여 같은 형태의 함수이지만 다른 출력이 가능해졌습니다.

`만약 함수에서 필요한 Data의 갯수가 여러개라면?`

### 결

필요한 Data가 여러개가 되면서 return point를 포함한 data들을 memory에 `stack` 자료구조를 이용하여 data를 저장하게 되었습니다.

stack는 FILO(First In Last Out)의 구조로 저장되는 순서는 다음과 같습니다.

```
---> push 순서
return point -> parameter1 -> parameter2 -> ...
```

또한 값을 가져오는 순서는 그 반대가 됩니다.

```
---> pop 순서
parameterX ... -> parameter2 -> parameter1 -> return point
```

마지막으로 memory에 저장되어 있는 data들은 `return`될 때 모두 소멸되게 됩니다.

<br>
<br>

## 함수 선언

```
func <함수 명>(parameter...)(return value...) {

}
```

### func

함수를 선언할 때 사용하는 golang keyword

<br>

### 함수명

숫자 혹은 특수 문자로 시작할 수 없습니다. ('_'는 가능) 
    
대문자로 함수명을 작성할 경우 expose가능하며 소문자로 작성할 경우 불가능합니다.

<br>

### parameter

함수에서 사용할 parameter를 정의할 수 있으며 정의하는 방법은 아래와 같습니다.

```
// parameter명 자료형
func example(a int)
```

여러개의 parameter를 사용할 수 있으며 가변 인자를 사용할 경우 가변인자의 parameter 위치는 가장 뒤에 있어야 합니다.

```go
// can only use ... with final parameter in list
func example2(a int, b ...int, c int) { // error 발생

}
```

<br>

### return value

함수를 실행 후 return 할 값의 자료형을 정의합니다. error를 return할 경우 error의 자리는 맨 뒤에 있는 것을 지향합니다. (go-lint)

```go
// error should be the last type when returning multiple items
func example3()(error, int){

}
```

또한 함수의 return도 동일하게 아래와 같이 작성할 수 있습니다.
```go
func example4(a int, err error){
    // 함수안에서 a와 err이라는 변수명을 가진 int type과 error type을 return
}
```
