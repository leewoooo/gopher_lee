valiable
===

프로그램이란 Data를 연산/조작 및 조작을 하는 것 입니다. Data를 조작할 때 `변수(valiable)`를 통해 진행합니다.

변수는 선언(Defination)을 통하여 사용을 할 수 있습니다.

## 선언 과정

- code에서 computer에게 명령을 한다 (내가 변수를 사용 할 거야) -> computer는 변수를 저장할 공간(memory)을 마련합니다. -> code에서 적은 변수명을 저장할 공간과 연결합니다. -> 이 후 변수명으로 memory공간에 접근하여 저장되어 있는 값을 사용합니다.

<br>

## 선언 방법

go언어의 변수 선언은 `var` keyword를 이용하여 선언을 합니다. 

```go
var <변수명> <자료형> = <값>
```

1. 명명 규칙은 Camel Case를 지향합니다.
2. go lint를 사용 시 '_'에 대한 warning을 나타냅니다.

<br>

## 대입연산자 '='

대입 연산자를 기준으로 좌측과 우측의 역할이 나뉘어 집니다.

```go
a = 3
```

위의 코드를 볼 때 대입연산자를 기준으로 다음과 같습니다.
- 좌측 = 메모리 공간
- 우측 = 값

<br>

## 변수의 4가지 속성

1. 이름(변수명) : 컴퓨터 입장에서는 컴파일 될 때 변수명을 사용하는 것이 아니라 memory 주소 값(16진수)을 이용합니다. 그럼 왜 변수 명을 사용할까? `개발자가 변수명을 통해 memory 주소 값을 구분하기 위해서 입니다.`

2. 값 : memory 주소에 저장되어 있는 값

3. 주소 : 할당된 memory의 시작 주소를 가지고 있다.

4. type(자료형) : type별 size를 통해 memory의 시작주소 부터 값을 읽어오기 위함

<br>

## 자료형에는 해당하는 Size정보를 가지고 있다.

```go
var a int8 = 3
```

다음과 같은 변수 선언이 있을 때 a라는 변수에 해당하는 memory 공간을 마련한 후, 마련한 공간에 3이라는 값을 복사하여 넣습니다.

그 이후 a라는 변수명으로 memory 공간에 접근을 하게 되는데 이 때 마련된 memory공간의 시작부터 자료형의 크기만큼 접근하게 됩니다.

int8형은 8bit의 크기이기 때문에 memory의 시작주소부터 8bit를 읽어 값을 가져오게 됩니다.

<br>

## 숫자 타입

### unsigned (양수) -> 정수 타입보다 약 2배 큰 값까지 나타낼 수 있다. 

| type명 | 크기 | 값의 범위 |
| :-- | :-- | :-- |
uint | 32bit or 64bit | uint32범위 or uint64범위
uint8 | 8bit(1byte) | 0 ~ 255
uint16 | 16bit(2byte) | 0 ~ 65535 
uint32 | 32bit(4byte) | 0 ~ 4294967295(약 42억)
uint64 | 64bit(8byte) | 0 ~ 18446744073709551615

### singed (정수)

| type명 | 크기 | 값의 범위 |
| :-- | :-- | :-- |
int | 32bit or 64bit | int32범위 or int64범위
int8 | 8bit(1byte) | -128 ~ 127
int16 | 16bit(2byte) | -32768 ~ 32767
int32 | 32bit(4byte) | -2147483648 ~ 2147483647
int64 | 64bit(8byte) | -9223372036854775808 ~ 9223372036854775807

int type은 컴퓨터에 따라 다르게 작동한다. 32bit 컴퓨터에서는 int32로 64bit 컴퓨터에서는 64bit로 적용됩니다.

여기서  32bit, 64bit는 hardware의 레지스터크기를 의미합니다.(한번 연산하는데 올라가는 크기)

### 실수

| type명 | 크기 | 값의 범위 |
| :-- | :-- | :-- |
float32 | 32bit(4byte) | IEEE-754 32비트 단정밀도 부동소수점, 7자리 정밀도 보장
float64 | 64bit(8byte) | IEEE-754 64비트 배정밀도 부동소수점, 15자리 정밀도 보장

### byte

| type명 | 크기 | 값의 범위 |
| :-- | :-- | :-- |
byte | 8bit | uint의 별칭(Alias)

byte는 보통 16진수, 문자 값으로 저장을 합니다. 주로 바이너리 파일에서 데이터를 읽거나 쓸 때, 데이터를 암호화나 복호화할 때 주로 사용합니다.

### rune

| type명 | 크기 | 값의 범위 |
| :-- | :-- | :-- |
rune | 32bit(4byte) | uint32의 별칭(Alias)

rune는 유니코드(UTF-8) 문자 코드를 저장할 때 사용합니다.

''로 묵어주어야 하며 문자 그대로 저장하거나 유니코드 문자 코드로 저장을 할 수 있습니다.

Go는 문자 한글자를 1~3 byte로 나타낸다. 최소 3byte의 공간을 가지고 있어야 하는데 가장 가까운 변수 타입이 uint32입니다.

### bool

| type명 | 크기 | 값 |
| :-- | :-- | :-- |
bool | 8bit(1byte) | true or false

<br>

## Overflow && Underflow

### Overflow

각 자료형에 저장할 수 있는 최대 크기를 넘으면 오버플로우가 발생합니다.
```go
var a uint8 = 255 + 1 //compile error
// cannot use 255 + 1 (untyped int constant 256) as uint8 value in variable declaration (overflows)
```

uint8 최대 값을 할당 후 1을 더하면 다시 0으로 돌아갑니다. (uint는 양의 정수이기 때문에 -로 내려가지 않음.)
int8 최대 값을 할당 후 1을 더하면 bit가 죄측으로 하나씩 밀려 부호 bit가 채워져 -128이 됩니다.
```go
var b uint8 = math.MaxUint8
b = b + 1 // b = 0

var c int8 = math.MaxInt8
c = c + 1 // c = -128
```

### Underflow

각 자료형에 저장할 수 있는 최소 크기보다 더 작아지면 언더플로우가 발생합니다.

```go
var a uint8 = 0 - 1 // complie error
// cannot use 0 - 1 (untyped int constant -1) as uint8 value in variable declaration (overflows)
```
uint8 최솟 값을 할당 후 -1을 하면 최댓값인 255가 됩니다.

```go
var b uint8 = 0
b = b - 1 // 255
```

<br>

## type의 크기 구하기

각 변수의 자료형에 해당하는 size를 구하려면 unsafe 패키지의 sizeof 함수를 사용합니다. 여기서 크기는 byte단위입니다.

```go
var a int8 = 123
var b int16 = 32000

unsafe.Sizeof(a) //1
unsafe.Sizeof(b) //2
```






