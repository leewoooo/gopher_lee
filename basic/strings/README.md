String
===

## 문자열이란?

문자열은 간단히 말하면 문자들이 모인 것을 의미합니다.(문자열은 문자들이 배열이라고 생각하는 것이 편할 수도 있습니다.)

go에서는 char type은 존재하지 않습니다. rune(int32)를 이용하여 문자의 코드 값을 표현할 수 있습니다.(`4byte를 이용하여 문자 1자를 표현`)

go는 기본적으로 utf-8을 이용합니다.

<br>

## String은 내부적으로 Pointer를 이용한다?

go에서는 같은 type끼리 복사가 가능합니다. type을 알면 size가 정해지기 때문입니다 그렇다면 문자열은 어떻게 복사가 이뤄질까요?

먼저 아래의 코드를 확인해봅시다.

```go
str1 := "Hello String"
str2 := ""
// something...
str2 = str1
```

<br>

str1과 str2를 선언하면서 초기화를 하였습니다. str1에는 "Hello String"이라는 값을 str2에는 ""를 할당하였습니다.

그렇다면 str1의 크기는 12byte이고 str2의 크기는 0byte일까요? 확인하면 결과는 아래와 같습니다.
```go
str1 := "Hello String"
str2 := ""

log.Println(unsafe.Sizeof(str1)) // 16
log.Println(unsafe.Sizeof(str2)) // 16
```

<Br>

둘 다 16byte입니다. 어떻게 둘의 size가 같은 것일까요? go에서는 string을 내부적으로 아래와 같은 struct를 이용합니다.
```go
type StringHeader struct {
	Data uintptr 
	Len  int
}
```

<br>

즉 go에서의 string은 포인터와 길이를 포함하는 구조체입니다. 해당 구조체의 field를 보면 `Data(uintPtr)과 Len인 int를 포함하고 있습니다.` 그렇기 때문에 해당 구조체의 크기는 `16byte(uintptr + int)`입니다.

Data field는 문자열의 주소값을 값으로 같은 포인터변수이며 Len field는 해당 문자열의 길이를 나타내는 field입니다. 즉 str1을 str2에 복사를 하게 되면 `16byte의 memory가 통체로 복사가 되어 str2변수에 할당된 memory에 덮어 쓰여지게 됩니다.`

이로 인해 str1과 str2는 같은 문자열의 `주소값을` 바라보고 있습니다.

<br>

## string은 불변(immutable)입니다.

문자열은 대입할 때는 먼저 만들어져 있는 String을 공유하지만 String을 변경할 때는 새로 만들게 됩니다. 한번 생성된 string은 `읽어 올 수만 있고 변경 할 수는 없습니다.` -> 같은 String 객체에 대해서는 별도의 공간을 지정하지 않아도 됩니다.

<br>

```go
str1 := "Hello String"
log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

// 부분 수정은 불가능 하다
// str1[2] = 'e'  error -> cannot assign to str1[2] (value of type byte)compiler (UnassignableOperand)

str1 = "Hello Golang"
log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

// output (서로 다른 주소값을 나타낸다.) -> 즉 string을 새로 만든다.
2021/10/16 21:53:05 &{4300555103 12}
2021/10/16 21:53:05 &{4300555091 12}
```

<br>

string의 경우 부분 수정을 하게 되면 assign error가 발생한다. 선언되어 있는 string 변수에 `새로운 string값을 할당하게 된다면 새로 만들거나 이미 memory상에 존재하는 string이라면 그것을 찾아오게 된다.`

<br>

## go 문자열 interning

문자열 인턴은 메모리에 각 고유 문자열의 복사본을 하나만 저장하는 기술입니다. 즉 위에서 설명한대로 `이미 memory상에 존재하는 string이라면 그것을 찾아오게 된다.` 가 이에 해당하는 것 입니다.
> internig : 이미 생성되어 있는 객체를 재사용 하는것을 말합니다. 보통 immutable 객체에 때해 internig을 사용합니다. (string은 immutable)

`문자열은 한번만 저장이 되고 한번 생성 된 문자열은 변경할 수 없으며 저장된 문자열을 공유합니다.` string이 `pointer`를 사용하기 때문에 만들어진 `문자열의 주소값을 Data field의 값으로만 가지고 있으면 되기 때문에 하나의 문자열을 여러 string이 공유 할 수 있습니다.`

<br>

컴파일 time에 정해지는 string은 internig이 적용 됩니다.
```go
str1 := "12"
log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

str2 := "1" + "2"
log.Println((*reflect.StringHeader)(unsafe.Pointer(&str2)))

// output
2021/10/16 22:32:04 &{4301270589 2}
2021/10/16 22:32:04 &{4301270589 2}
```

<br>

하지만 runtime시 생성되는 string에는 interning가 적용되지 않습니다.

```go
str1 := "12"
log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

str2 := strconv.Itoa(12)
log.Println((*reflect.StringHeader)(unsafe.Pointer(&str2)))

// output
2021/10/16 22:27:57 &{4365217341 2}
2021/10/16 22:27:57 &{4365257214 2}
```

<Br>

## ASCII, UTF-8

### ASCII

하나의 문자를 나타내는데 `1byte`를 이용합니다. (0 ~ 255) 하지만 ASCII로 표현하기에는 힘든 문자들이 등장하면서 UTF-8이 등장하게 됩니다.

<br>

### UTF-8

UTF-8은 유니코드를 위한 가변길이의 문자 인코딩 방식 중 하나입니다 UTF-8은 하나의 문자를 나타내는데 `1byte~3byte(4byte까지 사용합니다.)`를 이용하게 됩니다.

### 그렇다면 UTF-8은 어떠한 기준으로 1byte~4byte를 판단할까?

위키피디아의 설명은 먼저 다음과 같습니다.
>유니코드 코드 포인트를 나타내는 비트들은 여러 부분으로 나뉘어서, UTF-8로 표현된 바이트의 하위 비트들에 들어 간다. U+007F까지의 문자는 7비트 ASCII 문자와 동일한 방법으로 표시되며, 그 이후의 문자는 다음과 같은 4바이트까지의 비트 패턴으로 표시된다. 7비트 ASCII 문자와 혼동되지 않게 하기 위하여 모든 바이트들의 최상위 비트는 1이다. 
<img src = https://user-images.githubusercontent.com/74294325/137476332-665145f1-760a-47a6-a49d-64d92e3adf39.png>
(https://ko.wikipedia.org/wiki/UTF-8)



- 1byte: 첫 bit가 0으로 시작하게 되며 ASCII와 동일합니다.

- 2~4byte: 첫 bit가 1로 시작하며 그 뒤에 오는 1의 갯수를 가지고 2~4byte를 판단하게 됩니다.

### UTF-8의 장점

결론부터 이야기 하자면 memory를 효율적으로 사용하게 됩니다. UTF-16과 같은 경우는 하나의 문자를 `2byte`를 이용하여 나타내게 되는데 Web에서 사용하는 문자는 대부분 `1byte`로 표현되는 `영어나 숫자`가 대부분 입니다.

UTF-8과 같은 경우는 가변 길이의 문자 인코딩 방식이기 때문에 표현해야 하는 문자에 따라 `memory를 가변적으로 이용할 수 있기 때문입니다.`

또한 `영어나 숫자`와 같은 경우에는 ASCII와 호환이 가능합니다.

<br>

## Usage

문자열 또한 하나의 type이기 때문에 선언 및 초기화를 해주어야 합니다. 초기화 할 때 값을 정의하는 방법은 두 가지가 있습니다. 

1. ""(큰 따옴표)
2. `(백 스쿼트)

```go
var str string := "Hello String"
var str2 string := `Hello String2`
```

### Escape

문자열을 이용할 때 주로 사용하는 escape들입니다.
>Escape: 이스케이프 문자는 이스케이프 시퀀스를 따르는 문자들로서 다음 문자가 특수 문자임을 알리는 백슬래시(\\)를 사용한다.
https://ko.wikipedia.org/wiki/%EC%9D%B4%EC%8A%A4%EC%BC%80%EC%9D%B4%ED%94%84_%EB%AC%B8%EC%9E%90

```go
\\ -> \
\` -> `
\" -> "
\b -> 백 스페이스
\t -> 탭
\n -> 줄 바꿈
```
<br>

### 문자열 순회

문자열을 순회하는 방법은 3가지가 있습니다.

1. len()을 이용한 방법
2. str을 []rune로 변경하여 순회
3. range

```go
// given
	gString := "Hello 고랭"

	// #1
	for i := 0; i < len(gString); i++ {
		fmt.Printf("값: %d 문자:%c\n", gString[i], gString[i])
	}

	// #2
	arr := []rune(gString)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("값: %d 문자:%c\n", arr[i], arr[i])
	}

	// #3
	for _, v := range gString {
		fmt.Printf("값: %d 문자:%c\n", v, v)
	}
```

### #1

1번 방법에 대한 out put입니다.

```
값: 72 문자:H
값: 101 문자:e
값: 108 문자:l
값: 108 문자:l
값: 111 문자:o
값: 32 문자:
값: 234 문자:ê
값: 179 문자:³
값: 160 문자:
값: 235 문자:ë
값: 158 문자:
값: 173 문자:­
```

문자열에 len()을 이용했을 경우 return값으로 해당 문자열의 byte 크기를 return 받게 됩니다. 

해당 문자열은 총 12byte이기 때문에 12번의 반복문이 돌게 됩니다.

byte단위이다 보니 한글같은 경우는 3byte로 표현하다 보니 깨져서 출력되게 됩니다.

### #2

위와 같이 한글이 깨지는 경우를 방지하여 순회하기 위해 문자열을 `rune(int32) type`의 배열로 변경하여 순회를 하게되면 out put은 아래와 같습니다.

```
// arr = [72 101 108 108 111 32 44256 47021]
값: 72 문자:H
값: 101 문자:e
값: 108 문자:l
값: 108 문자:l
값: 111 문자:o
값: 32 문자:
값: 44256 문자:고
값: 47021 문자:랭
```

하나의 문자를 4byte단위로 처리하기 때문에 한글도 출력이 정상적으로 됩니다.

### #3

range를 이용한 결과는 #2와 동일합니다.

<Br>










## REFERENCE

https://stackoverflow.com/questions/52851788/string-memory-usage-in-golang

https://artem.krylysov.com/blog/2018/12/12/string-interning-in-go/
