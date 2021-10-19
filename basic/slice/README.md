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

즉 `slice는 Go에서 지원하는 배열을 가리키는 pointer type입니다.`

그렇기 때문에 아래와 같은 code가 가능해 지는 것 입니다.

```go
func changeSlice(slice []int) {
	slice[1] = 10
}

func changeArray(arr [3]int) {
	arr[1] = 10
}

func TestChange(t *testing.T) {
	slice := []int{1, 2, 3}
	arr := [3]int{1, 2, 3}

	changeSlice(slice)
	changeArray(arr)

	log.Println("slice: ", slice)
	log.Println("arr: ", arr)
}
```

<br>

배열과 같은 경우에는 `changeArray`안으로 들어가면서 값이 복사되어 함수 안의 인스턴스와 main함수의 인스턴스가 `다른 인스턴스이기 때문에` 함수 안에서 값을 바꿔도 변함이 없습니다.

하지만!

slice와 같은 경우는 내부적으로 `pointer`를 사용하기 때문에 slice의 구조체가 함수안에 복사되어 `다른 인스턴스가 되어도 가리키는 배열은 동일하기 때문에` 함수 안에서 값을 변경하면 `pointer`가 가리키는 배열의 값이 바뀌기 때문에 `해당 배열을 바라보고 있는 slice 구조체들에게는 동일하게 적용됩니다.`

<br>

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

`slice는 internig은 되지 않는다..!` slice를 생성할 때마다 새로운 memory를 할당받아 만들고 저장한다.

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

```go
// case #1
var slice []int = []int{1, 2, 3}

// case #2
slice := make([]int, len, cap)
```

배열을 생성하는 것과 비슷하게 slice도 선언 및 초기화를 할 수 있습니다. case #1번을 보면 배열을 선언 및 초기화와 비슷하지만 `길이를 지정하지 않습니다.`

case #2와 같은 경우에는 `make`라는 내장함수를 이용하여 slice를 선언합니다. 

이 때 `len, cap` 을 함수의 argument로 넣어주게 되는데 `len의 값만큼 선언된 slice는 zero Value를 가지게 되며 cap는 선언된 slice memory에 담을 수 있는 최대 요소의 갯수를 나타냅니다` 

초기화는 하지 않았지만 len만큼 `index로 접근이 가능합니다.`

또한 cap의 크기보다 요소가 더 많이 들어가게 되면 내부적으로 새롭게 memory를 할당받아 slice에 요소를 추가합니다. 

### Nil Slice

```go
var a []int
var b []int = make([]int,0)
var c []int = []int{}

log.Println("a, b, c", a, b, c)

//output
2021/10/19 23:46:11 a, b, c [] [] []
```

위와 같이 slice를 선언을 하게되면 a, b, c 모두가 `cap, len이 0이게 됩니다.` 

하지만 a와 b,c는 다른점이 있습니다. 

a와 같은 경우는 변수를 생성은 하였지만 해당하는 변수에 아직 memory가 할당되어 있지 않아 `nil`인 상태입니다.

하지만 b,c와 같은 경우는 변수를 생성하고 선언한 상태에 해당되기 때문에 `nil이 아닌 빈 배열을 가리키고 있게 됩니다.`

<br>

## append

slice에 요소를 추가할 때 go 내장함수인 `append`를 사용할 수 있습니다. append는 argument로 받은 slice에 해당 type을 추가한 후 `추가 된 결과의 slice를 반환합니다.`

### append의 동작 원리 (append의 결과는 새로운 slice일까? 기존의 slice일까?)

`append로 return 받는 slice는 새로운 slice일 수도 있고 기존의 slice일 수도 있습니다.`

그럼 어떻게 이게 결정될까요?

`slice의 구조체 중 len과 cap에 관련이 있습니다.` 현재의 `Len값이 Cap보다 작다면` 기존의 배열에 추가할 요소를 반환합니다. 하지만 현재의 `Len의 값이 Cap와 같다면` 즉 현재 slice가 선언할 때 할당받은 Memory에 요소가 가득 차있다면 `새로운 Memory를 할당받아 slice에 요소를 추가 후` 새로 생성된 배열을 `pointer`로 가리키게 됩니다.

>The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements
(https://pkg.go.dev/builtin#append)
```go
slice := make([]int, 3, 5)
log.Println("origin: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

slice = append(slice, 4)
log.Println("append 1: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

slice = append(slice, 5)
log.Println("append 2: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

slice = append(slice, 6)
log.Println("append 3: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

//output
2021/10/19 23:21:29 origin:  &{1374389633600 3 5}
2021/10/19 23:21:29 append 1:  &{1374389633600 4 5}
2021/10/19 23:21:29 append 2:  &{1374389633600 5 5}
2021/10/19 23:21:29 append 3:  &{1374389649648 6 10}
```

<br>

위의 예제 code에서 확인할 수 있듯 len과 cap이 동일해지기 전까지 (`할당받은 memory에 공간이 남아있다면`) append시 `동일한 slice를 가리키고 있는 것을 볼 수 있습니다.`

하지만!

len과 cap이 동일한 상태에서 append를 할 경우(`할당받은 memory에 공간이 남아있지 않다면`) append시 `새로운 배열을 만들어 요소를 추가 후` 생성된 memory를 가리킵니다. 이 때 cap또한 같이 증가하게 됩니다.

<br>

## REFRENCE

https://www.youtube.com/watch?v=z-_6o7WYkiE








