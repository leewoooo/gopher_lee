Slice
===

## Slice란?

슬라이스는 Go에서 지원하는 동적인 배열 type입니다. 동일한 type의 데이터를 순차적으로 저장할 때 사용한다는 것은 배열과 동일하지만

`차이점은 배열은 고정길이지만 slice는 길이를 동적으로 다룰 수 있습니다.`

### 동적? 정적?

정적이라는 것은 compile time에 정해진다는 것 입니다. (Build time) 즉 한번 정해지면 고정되는 것을 이야기 합니다.

동적이라는 것은 runtime 도중 변경 될 수 있는것을 이야기 합니다.

<br>

## Slice도 내부적으로 Pointer를 이용합니다!

string에서 살펴보았듯 slice도 내부적으로 pointer를 사용합니다. string의 구조체와 동일하지만 다른점은 cap field가 추가되었다는 것 입니다.

```go
type SliceHeader struct{
    Data uintptr // 실제 배열이 저장되어 있는 memory의 주소값
    Len int // 요소의 갯수
    Cap int // slice의 최대 길이 (현재 slice가 할당 받은 memory에서 넣을 수 있는 최대 요소의 갯수)
}
```

즉 `slice는 Goㅇ에서 지원하는 배열을 가리키는 pointer type입니다.`

### Slice도 internig이 될까?

이전 String을 test했던 것 처럼 slice 또한 reflect package와 unsafe package를 이용해서 확인해 보면 다음과 같다.

```go
//given
arr1 := []int{1, 2, 3}

arr2 := []int{1, 2, 3}

//when
log.Println((*reflect.SliceHeader)(unsafe.Pointer(&arr1)))
log.Println((*reflect.SliceHeader)(unsafe.Pointer(&arr2)))

//output
2021/10/18 23:25:54 &{1374390878280 3 3}
2021/10/18 23:25:54 &{1374390878304 3 3}
```

`slice는 internig은 되지 않는다..!` slice를 생성할 때마다 새롱운 memory를 할당받아 만들고 저장한다.

### slice 변수의 크기는 24byte

내부적으로 위에서 작성한 구조체를 상용하다보니 slice 변수의 크기는 24byte이다. 

```go
// 64bit 기준
type SliceHeader struct{
    Data uintptr // 8byte
    Len int // 8byte
    Cap int // 8byte
}

//given
arr1 := []int{1, 2, 3}
arr2 := []int{1, 2, 3}

//when
log.Println(unsafe.Sizeof(arr1))
log.Println(unsafe.Sizeof(arr2))

//output
2021/10/18 23:34:53 24
2021/10/18 23:34:53 24
```

<br>

## usage

slice도 type이기 때문에 선언을 하여 memory를 할당받고 초기화를 하여 값을 채워 넣어야 합니다. 기본적으로 선언 및 초기화는 배열과 동일합니다.

하지만 배열과는 다를게 `길이를 지정하지 않습니다.`

```go
var slice []int = []int{value...}
```

