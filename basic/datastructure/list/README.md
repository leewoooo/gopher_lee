Data Structure
===

## 자료구조란?

사전적인 의미로는 자료(Data)의 집합을 의미하며, 각 원소들이 논리적으로 정의된 규칙에 의해 나열되며 자로에 대한 처리를 효율적으로 수행할 수 있도록 자료를 구분하여 표현한 것 입니다.

즉 Data들을 어떠한 형태로 저장할까에 대한 이야기 입니다.

<br>

### 자료구조의 목적은?

자료를 `더 효율적으로 저장하고, 관리하기 위해 사용되며` 잘 선택된 자료구조는 실행 시간을 단축시켜 주거나 메모리 용량의 절약을 이끌어 낼 수 있습니다.

<br>

### 선택 기준

1. 자료의 처리 시간
2. 자료의 크기
3. 자료의 활용 빈도
4. 자료의 갱신 정도
5. 프로그램의 용이성

<br>

### golang에서 container package를 이용해보자.

golang standard library에서는 3개의 자료구조를 기본적으로 제공해줍니다. (이 markdown에서는 list만 다루겠습니다.)

1. list (linked list)
2. heap (min heap)
3. ring

<br>

## list (linked list)

배열과 함께 기본적인 `선형` 자료구조 중 하나입니다. [list api문서](https://pkg.go.dev/container/list@go1.17.2)에는 list를 이렇게 소개하고 있습니다.

>Package list implements a doubly linked list.

말 그대로 double linked를 구현해 놓은 api입니다.

<br>

### 배열 vs list

배열은 연속된 memory에 Data를 저장하는 자료구조 입니다. 그에 비해 list는 연속된 memory가 아닌 `각각 위치하는 memory를 pointer를 이용하여 연결해 놓은 자료구조 입니다.`

<Br>

### 맨 앞의 요소를 추가하기 

배열의 맨 앞의 요소를 추가하는 경우에는 맨 처음부터 끝까지 요소를 뒤로 한칸씩 밀어 자리를 만들고 맨 앞의 자리에 새로 추가 할 요소를 저장하게 됩니다.

**하지만!**

list의 경우는 `pointer`를 연결해 놓은 구조이기 때문에 맨 앞의 요소의 `prev를 새로 추가되는 요소의 memory와 연결만 시켜주면 됩니다.`

맨 앞의 요소 추가가 아니더라도 중간 요소를 추가하는 경우 배열은 해당 위치 뒤에 있는 요소들을 전부 이동해야 하지만 list의 경우 해당 위치 `앞, 뒤에 요소의 next와 prev만 추가 될 요소의 memory로 이어주기만 하면 됩니다.`

그렇기에 `요소를 맨 뒤가 아닌 어떠한 위칭에 추가할 때는 list가 우세합니다.`

<br>

### 특정 요소에 접근하기 (random access)

list는 특정 요소에 접근을 할 때 `list의 처음부터 해당 요소까지 next를 조사하여 요소를 찾아야 합니다. 혹은 맨 뒤에서 부터 prev를 조사하여 찾아야 합니다.`

list는 연속된 memory가 아니기 때문에 다음 요소에 접근하려면 `그 주소 값을 직접 찾아가야 합니다.`

**하지만!**

배열의 경우에는 `연속된 memory`이며 index를 사용할 수 있기 때문에 특정 요소에 손쉽게 접근을 할 수 있습니다.

`배열의 시작주소 + (index * type 크기)` 의 연산을 진행하면 특정 요송에 접근을 빠르게 할 수 있습니다.

그렇기에 `특정 요소에 접근할 때는 list보다 배열이 우세합니다.`

<br>

### 데이터의 지역성 

데이터가 인접해 있을수록 cache의 성공률은 올라갑니다. 즉 성능이 좋아집니다.

cpu가 연산을 할 때 `memory에서 필요한 data를 가지고 와 연산을 진행합니다.` 이 때 `필요한 data만 가져오는 것이 아니라 인접한 영역의 data를 cache로 가져옵니다.`

그 이유는 `현재 연산에 다음 연산에서 사용되는 data가 인접한 영역의 data를 이용할 확률이 높기 때문입니다.`

cache 성공이란 `현재 연산에 다음 연산이 cache에 있는 data를 이용`하는 것을 이야기 합니다. 또한 cache 실패는 `현재 연산에 다음 연산이 cache에 있는 data를 이용하지 않는` 것을 이야기 합니다.

cache 실패가 되면 `현재 cache에 있는 data를 지우고 다시 memory에서 data와 그 인접한 영역을 가져오게 됩니다.` 즉 cache를 지우고 다시 가져오는 연산이 추가됩니다.

list는 `연속된 memory가 아니기 때문에` 현재 연산에서 사용되는 data와 인접한 영역의 data를 가져와서 연산을 진행한다고 하여도 `현재 data의 인접한 영역에 있는 data들은 추 후 연산과 연관이 없을 경우가 높습니다.` 즉 cache 성공률이 낮습니다.

**하지만!**

배열의 경우 `연속된 memory`이기 때문에 현재 연산에서 사용되는 data와 인접한 영역의 data를 가져와서 연산을 할 경우 `배열안에 있는 data`들이기 때문에 추 후 연산과 연관이 높습니다. 즉 cache 성공률이 높습니다.

그렇기에 `데이터 지역성의 특징으로 보았을 때 list보다 배열이 우세합니다.`

<br>

### 그럼 언제 배열을 쓰고 언제 list를 써야하나?

위에서 정리한 배열 vs list를 바탕으로 이야기 하자면 다음과 같습니다.

- 특정 요소에 대한 접근이 많을 경우 list보다는 `배열을 이용합니다.`
- 요소에 대한 삽입과 삭제의 빈도수가 많을 경우 배열보다 `list를 이용합니다.`

>일반적으로 요소의 수가 적은 경우 리스트 보다 배열이 빠릅니다. <br>(https://www.youtube.com/watch?v=r_oVchBrQkc&t=2410s)

`요소의 수가 많이지고 기준이 애매하다면` 프로파일링을 이용하여 배열 혹은 list를 선택하면 됩니다.
>프로파일링:프로파일링 또는 성능 분석은 프로그램의 시간 복잡도 및 공간, 특정 명령어 이용, 함수 호출의 주기와 빈도 등을 측정하는 동적 프로그램 분석의 한 형태이다 <br>(https://ko.wikipedia.org/wiki/프로파일링_(컴퓨터_프로그래밍))

<br>

### double linked list? (single) linked list?

그럼 double linked list와 linked list의 차이는 무엇일까요? list의 element가 되는 구조체의 method를 보면 알 수 있습니다.

```go
type Element struct {
	Value interface{}
}

// Next
func (e *Element) Next() *Element

// Prev
func (e *Element) Prev() *Element
```

Next와 Prev가 구현되어 있습니다. 이와 같이 하나의 요소가 `자기 자신을 기준으로 이전과 이후에 대한 정보를 가지고 있으면 double, Next의 정보만 가지고 있으면 single` 이라 합니다.

<br>

## 배열과 list를 이용하여 Stack과 Queue을 구현해보자.

### stack

- 사진

	<img src = https://user-images.githubusercontent.com/74294325/119511133-17219200-bdad-11eb-9008-7a5eb1a16814.png>

위의 사진과 같이 비어있는 Stack에 값들이 추가되는 것을 Push라 하며 Stack의 가장 위에 있는 값을 얻는 것을 Pop라고 합니다.

Stack은 주로 `Momory 구조`나 `함수 call stack`에서 사용됩니다.

stack은 배열로 구현로 구현을 하는 경우가 많은데 그 이유는 배열과 같은 경우 `FILO(First in Last Out)`과 같은 구조를 가지고 있기 떄문입니다.

요소가 추가되는 것도 자료구조의 `제일 뒤`로 추가가 되고 요소를 꺼내어 사용할 때도 `제일 뒤`에서 사용하기 때문에 `요소 삽입에 대한 제약을 받지 않습니다.`

#### 구현

```go
var stack interface{}{}

// push
stack = append(stack, 추가할 값)

// pop
val, stack = stack[len(stack)-1], stack[:len(stack)-1]
```

<br>

### Queue

- 사진
s
    <img src = https://user-images.githubusercontent.com/74294325/119514493-06bee680-bdb0-11eb-87cd-e13f043f8895.png>	


Queue 또한 Stack과 구현하는 것이 크게 다를 것이 없습니다. (stack을 배열로 하였으니 queue는 list를 이용하겠습니다.) 

Queue는 Stack과 다르게 `FIFO(First in Frist Out)`과 같은 구조를 가지고 있습니다.

Queue 같은 경우에는 대기열에 많이 사용됩니다. 들어온 순서에 따라 로직을 처리해야 할 때 Queue를 이용하게 됩니다.(예를 들어 번호표를 뽑고 그 번호표 대로 실행 하고 싶을 때 사용)

#### 구현

```go
type Queue struct {
	l *list.List
}

func NewQueue(l *list.List) *Queue {
	return &Queue{l: l}
}

func (q *Queue) Push(val interface{}) {
	q.l.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	element := q.l.Front()

	if element != nil {
		return q.l.Remove(element)
	}

	return nil
}
```
<br>

## REFERENCE

https://www.youtube.com/watch?v=r_oVchBrQkc&t=2410s

https://andrew0409.tistory.com/148