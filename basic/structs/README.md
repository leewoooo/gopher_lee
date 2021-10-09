Struct
===

## 구조체(struct)란?

구조체는 여러 `fields`를 묶어서 사용하는 type입니다. 

구조체 또한 type이기 때문에 선언이 필요하며 선언과 동시에 해당 구조체에 필요한 Memory를 할당받게 됩니다.

fields는 `변수들이다.` 각각의 이름과 type을 가지고 있으며 fields의 자료형은 `모두 같을 필요는 없습니다.`

<br>

## Usage

```go
type <구조체 명> struct {
    field명 type
    ...
}

//ex
type User struct {
    Name string
    Age int
}
```

위의 코드와 같이 `type`, `struct`의 keyword와 함께 struct를 정의합니다. field명을 정의할 때 `field명의 시작 첫 글자의 크기가 중요합니다. field명의 첫 글자로 인해 expose가 가능한지 아닌지의 여부가 결정되기 때문입니다.`

```go
// package domain
type User struct{
    Name string // expose
    pID int     // expose x
}

// another package or another directory
func test(){
    // package가 다르거나 디렉토리가 다르다면 expose가 되지 않은 field에 대한 초기화를 진행할 수 없음.
    user := domain.User{
        Name : "lee",
    }
}
```

만약 다른 package나 디렉토리에서 expose되지 않은 field를 접근하기 위해서는 구조체를 정의한 곳에서 생성자함수를 정의하여 초기화해준다.

```go
// package domain
func NewUser(name string, pid int) User{
    return User{
        Name : name,
        pID : pid
    }
}
```

<br>

## Embeded struct

struct의 field로 struct를 정의할 수 있습니다. 정의하는 방법은 2가지 입니다.

### case 1

```go
type People struct{
    Name string
    Age int
}

type Tenant struct {
    Member People
    TenantName string
}
```
Tenant라는 구조체가 People이라는 구조체를 field로 가지고 있습니다. Tenant를 선언할 때 People 구조체 또한 같이 초기화 하며 선언해줄 수있습니다.

```go
tenant := Tenamt{
    Member : People{
        Name : "lee",
        Age : 26
    },
    TenantName : "foo",
}
```

선언한 struct의 field를 접근할 때는 다음과 같이 접근이 가능합니다.

```go
tenant.Member.Name // People의 Name값에 접근 
tenant.Member.Age // People의 Age값에 접근 
```

<br>

### case 2

```go
type People struct{
    Name string
    Age int
}

type Tenant struct {
    People
    TenantName string
}
```

case 1과 크게 다를 것은 없지만 이번에는 Tenant라는 구조체가 field명 없이 People를 내장하고 있습니다. 선언 및 초기화는 case1과 같습니다. 하지만 field를 접근하는 방법이 조금 다릅니다. field명 없이 정의된 구조체는 아래와 같이 바로 값에 접근이 가능하게 됩니다.

```go
tenant.Name // Tenant에 내장된 people에 Name 값
tenant.Age // Tenant에 내장된 people에 Age값
```

만약 `embeded된 구조체와 부모 구조체의 field명이 중복이 발생한 경우에 우선순위는 부모 struct에게 있습니다.`

<br>

## 구조체의 크기

구조체 또한 type이기 때문에 선언할 때 memory가 할당됩니다. 구조체의 memory는 각 field의 type의 크기를 모두 합친 것과 동일합니다.

```go
type People struct{
    Age int32
    Score int
}
```

다음과 같은 People struct가 존재할 때 해당 struct의 memory를 예상해보자면 다음과 같습니다.

```go
// 64bit 기준
4byte(int32) + 8byte(int) = 12byte
```

하지만 결과는 다음과 같습니다. 

```go
p := People{
    Age:   1,
    Score: 2,
}

// 예상하기로는 12 하지만 16
log.Println(unsafe.Sizeof(p)) // output = 16
```

그 이유는 go compiler가 내부적으로 메모리를 정렬합니다. 

`64bit를 기준으로 score는 8byte이고 age는 4byte입니다. 이 때 compile는 성능을 높이기 위해 int 자료형에 맞게 age field에 해당하는 memory도 동일하게 8byte로 할당합니다.` 

register의 크기만큼 memory를 읽어와 연산을 진행할 텐데 `age를 4byte로 할당을 한 후 People 구조체의 memory의 시작부터 8byte를 읽어오게 된다면 처음 읽을 때 age의 값과 score의 memory중 절반만 읽힙니다.` 그 후 한번 더 8byte만큼 읽게 되면 People와 연관 없는 Memory의 값이 4byte만큼 읽힐 것 입니다.

그렇기 때문에 내부적으로 memory를 padding하여 `score의 크기를 기준으로 age의 memory 또한 8byte로 할당을 하게 됩니다.`

<br>

### 그렇다면 padding을 내부적으로 줄일 수 있을까?

field의 순서를 잘 정렬하면 내부적으로 padding을 줄일 수 있습니다.(padding을 줄인다라는 의미는 memory를 아낄수 있다는 것)

`즉 8byte보다 작은 field는 8byte 단위로 고려하여 정의하면 된니다.`





