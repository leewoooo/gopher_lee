Function 고급편
===

## 가변인자함수

함수의 parameter에서 같은 type의 값들을 slice로 받을 때 사용합니다. 


가변 인자를 사용할 경우 가변인자의 parameter 위치는 가장 뒤에 있어야 합니다.
```go
// can only use ... with final parameter in list
func example(a int, b int, c ...int) { // error 발생
    // ex
    for i , v := range c {
        //...
    }
}

// can only use ... with final parameter in list
func example2(a int, b ...int, c int) { // error 발생

}
```

<br>

## 지연 실행 keyword defer

함수 안에서 `defer` keyword를 이용하여 함수가 종료되는 시점 바로 전에 실행되는 codee를 정의할 수 있습니다. (java의 try-catch-finally중 finally와 비슷한 역할) 

### 언제 사용할까?

`OS`의 자원을 반납할 때 주로 사용하게 됩니다. 예를 들어 file을 Handling하는 프로그램을 작성할 때 process는 다음과 같습니다.

1. 파일 작업시 파일 프로그램은 파일 핸들러를 OS에 요청합니다.
2. OS는 파일 핸들러를 제공합니다.
3. 프로그램은 파일 작업이 완료 후 파일 핸들러를 OS에게 반납합니다. (go에서는 defer를 이용)

프로그램이 종료되면 자원을 자동으로 다 반납하지만 `프로그램이 계속 실행중이라면 사용한 자원은 반드시 반납되어야 합니다.`

<br>

## ex

defer의 사용은 아래의 code와 같습니다.

```go
func CreateFile(name string) {
    openedFile,err := os.Crete(name)
    if err != nil{
        log.Fatal(err)
    }
    defer openFile.Close()
}
```

defer를 이용해 실행할 code를 익명함수로 작성할 수도 있습니다. 만약 defer로 실행해야 하는 code에서 return 값으로 error를 뱉는다면 아래와 같이 익명함수를 만들어 error handling이 가능합니다.

```go
func CreateFile(name string){
    openedFile, err := os.Create(name)
    if err != nil{
        log.Fatal(err)
    }
    defer func(){
        err := opendFile.Close()
        // error handling
    }()
}
```
<br>

### defer의 call stack

하나의 함수의 defer가 여러개 일 경우는 어떤 순서대로 실행이 될까요?

```go
func TestDeferCallStack(t *testing.T) {
	defer log.Println("1")
	defer log.Println("2")
	defer log.Println("3")
}
```

위에서 부터 순서대로 1, 2, 3을 loggig하는 code를 지연시켰을 때 call stack에는 다음과 같이 들어갑니다.

1 -> 2 -> 3

stack은 `FILO`의 성격을 가진 자료구조입니다. 그렇기 때문에 호출되는 순서는 다음과 같습니다.

3 -> 2 -> 1

```go
// output
2021/10/27 22:34:01 3
2021/10/27 22:34:01 2
2021/10/27 22:34:01 1
```

<br>

## 함수 type의 변수

함수 type변수란 함수를 값으로 갖는 변수를 이야기합니다. 

### 그렇다면 어떻게 함수를 값으로 가질수 있을까요?

예전 함수 호출과정을 살펴보았을 때 `IP(instruction pointer)`를 옮겨 가며 함수를 실행한다는 것을 알게 되었습니다. (함수가 호출될 때 return할 주소를 가지고 함수가 있는 곳으로 IP포인터를 이동하여(JMP) 함수를 실행 후 return할 주소로 다시 돌아온다.)

컴퓨터는 숫자로 된 data를 다룹니다. 즉 함수의 시작 위치도 `숫자값으로 표현`이 가능하다는 것 입니다. 여기서 숫자 값은 `함수의 주소를 나타내며 변수는 함수의 주소를 참조하고 있는 것 입니다.`

### syntax

함수 시그니처로 표현을 합니다. 함수 시그니처라 하면 아래와 같습니다.

```go
// 함수
func add(a, b int) int {
    return a + b
}

// 함수 시그니처
func (int, int) int
```

주로 `별칭타입을 만들어서` 사용을 하며 별칭타입 선언은 아래와 같습니다.

```go
type Add func(a int, b int) int

func makeAdd() Add {
	return func(a, b int) int {
		return a + b
	}
}

func TestAdd(t *testing.T) {
	funcAdd := makeAdd()

	result := funcAdd(1, 2)

	if result != 3 {
		t.Fatal("result should be 3")
	}
}
```

<Br>

## 함수 리터럴 (람다)

소스 코드에 고정된 값을 `리터럴 이라합니다.`

```go
var a int = 3 // 여기서 3이 리터럴에 해당함.

// 여기서 익명 함수 즉 f에 할당되는 function이 함수 리터럴이다.
f := func (a, b int) int {
    return a + b
}
```

go에서 함수 리터럴은 익명 함수를 나타냅니다. 함수명을 갖지 않은 함수입니다. (Anonymous Function)이라고 부릅니다.

보통 익명 함수는 `함수 전체를 변수에 할당하여 다른 함수의 인자로 전달하거나(고계 함수) 반복사용 혹은 즉시 실행함수를 작성할 때 사용합니다.`

익명함수의 다른말로는 `람다 함수`가 있습니다.

<br>

### 함수 리터럴은 내부 상태를 가질 수 있다? (함수 리터럴은 외부의 변수를 capture해서 사용할 수 있다.)

부재를 길게 적어놨지만 다른언어에서의 `closure` 개념을 사용할 수 있다는 의미입니다.

예제 code를 살펴본 후 이야기 해보자면

```go
func TestClosure(t *testing.T) {
	i := 0

	t.Log(unsafe.Pointer(&i))

	AddTen := func() {
		i += 10
		t.Log(unsafe.Pointer(&i))
	}
	AddTen()

	if i != 10 {
		t.Fatal("i should be 10")
	}
}

//output
2021/10/27 23:03:07 0x140001161c0
2021/10/27 23:03:07 0x140001161c0
```

golang은 기본적으로 call by value입니다. 

하지만!!

위의 경우 결과를 보면 `i의 값은 10으로 변경`되어 있고 `outer의 i의 주소와 AddTen의 i는 같은 주소를 가지고 있는 것을 확인 할 수 있습니다.`

즉 함수 리터럴안에서 외부의 지역변수를 사용할 때는 `값이 복사가 되는 것이 아니라 reference가 복사되기 때문입니다. (go에서는 reference라는 개념보다는 pointer)`











