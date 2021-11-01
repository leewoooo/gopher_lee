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

golang standard library에서는 3개의 자료구조를 기본적으로 제공해준다. 

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

#### 맨 앞의 요소를 추가하기 

배열의 맨 앞의 요소를 추가하는 경우에는 맨 처음부터 끝까지 요소를 뒤로 한칸씩 밀어 자리를 만들고 맨 앞의 자리에 새로 추가 할 요소를 저장하게 됩니다.

**하지만!**

list의 경우는 `pointer`를 연결해 놓은 구조이기 때문에 맨 앞의 요소의 `prev를 새로 추가되는 요소의 memory와 연결만 시켜주면 됩니다.`

맨 앞의 요소 추가가 아니더라도 중간 요소를 추가하는 경우 배열은 해당 위치 뒤에 있는 요소들을 전부 이동해야 하지만 list의 경우 해당 위치 `앞, 뒤에 요소의 next와 prev만 추가 될 요소의 memory로 이어주기만 하면 됩니다.`

그렇기에 `요소를 맨 뒤가 아닌 어떠한 위칭에 추가할 때는 list가 우세합니다.`

<br>

#### 특정 요소에 접근하기 (random access)

list는 특정 요소에 접근을 할 때 `list의 처음부터 해당 요소까지 next를 조사하여 요소를 찾아야 합니다. 혹은 맨 뒤에서 부터 prev를 조사하여 찾아야 합니다.`

list는 연속된 memory가 아니기 때문에 다음 요소에 접근하려면 `그 주소 값을 직접 찾아가야 합니다.`

**하지만!**

배열의 경우에는 `연속된 memory`이며 index를 사용할 수 있기 때문에 특정 요소에 손쉽게 접근을 할 수 있습니다.

`배열의 시작주소 + (index * type 크기)` 의 연산을 진행하면 특정 요송에 접근을 빠르게 할 수 있습니다.

그렇기에 `특정 요소에 접근할 때는 list보다 배열이 우세합니다.`

<br>

#### 데이터의 지역성 

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



## REFERENCE

https://andrew0409.tistory.com/148