Method
===

## Method란?

method는 type에 속한 함수입니다. 함수는 독립적이지만 `method는 type에 종속되어 있습니다.`

또한 method는 함수로 변경하여 사용할 수 있습니다. (`기능적으로 동일하게 실행되지만 객체의 개념을 이용하기 위해 method 이용`)

<br>

## type에 종속되어 있다??

type에 종속되어 있다는 말은 즉 `상태와 기능이 합쳐있다는 것을 의미하며` `객체`의 개념으로 확장할 수 있습니다.

기존 go의 struct에는 `상태에 해당 하는 field 들이 정의되어 있었습니다.` 여기에 `method`라는 개념을 추가하여 `기능을 추가한 것입니다.`

타언어 java의 field와 method가 하나의 class에 들어있는 것처럼 `go언어에서는 type과 type에 종속된 method로 객체를 나타냅니다.`

<br>

### method는 객체와 객체의 관계를 나타냅니다.

```go
type Student struct{
    Name string
}

func (s *Student) EnrollClass (sub *Subject) {
    // 등록하는 logic
}
```

`student는 Subject를 이용하고 있습니다(상호작용이 생김)` 즉 이를 통해 Student의 객체와 Student의 객체의 관계가 형성되었습니다.

<br>

## Usage

method를 선언할 때 함수를 선언하는 것과 크게 달라지는 것은 없으나 `리시버`의 개념이 들어옵니다.

함수명 앞에 타입과 변수를 붙여 `리시버를 나타냅니다.` 즉 리시버가 붙은 함수는 `독릭적인 함수가 아닌 type에 종속된 method임을 나타냅니다.`

```go
type Person struct{
    Name string
    Age int
}

func (p Person) PrintName(){
    fmt.Println(p.Name)
}

func (p *Person) AddAge(){
    p.Age++
}

person := Person{
    Name : "foo",
    Age : 26
}

person.PrintName()
person.AddAge()
```

위의 code와 같이 함수 명 앞에 `타입과 변수`를 붙임으로 `PrintName`과 `AddAge`는 Person이라는 type에 `종속된 method임을 나타냅니다.`

<br>

## Value Type Method vs Pointer Type Method

리시버에는 두가지 종류가 존재합니다. 

<br>

### Case #1

첫번째 PrintName과 같은 `Value Type`의 리시버를 이용한 method입니다. Value Type으로 리시버를 이용할 경우 해당하는 method의 body에서는 해당하는 `person 객체를 복사하여 사용합니다.`

`즉 함수에서 parameter로 받은 변수를 내부에서 복사해서 사용하는 것과 동일합니다.`

그렇기 때문에 `생성된 person과 method 안에서의 p와는 다른 인스턴스 입니다.`

PrintName을 함수로 변경하면 다음과 같습니다.

```go
func PrintName (p Person){
    fmt.Println(p.Name)
}
```

<Br>

### Case #2

두번째 AddAge와 같은 `Pointer Type`의 리시버를 이용한 method입니다.
Pointer Type으로 리시버를 이용할 경우 해당하는 method의 body에서는 해당하는 `person 객체의 주소값을 사용하게 됩니다.`

`즉 method 내부에서 해당하는 type의 주소값에 접근하여 field를 handling 할 수 있다는 이야기 입니다.`

그렇기 때문에 AddAge `method안에서 Age값을 증가시키면 person의 field에 값이 증가됩니다.`

AddAge를 함수로 변경하면 다음과 같습니다.

```go
func AddAge(p *Person){
    p.Age++
}
```

<br>

### 그렇다면 언제 value Type? pointer Type?

method가 종속되어 있는 type의 성격에 따라 사용하면 됩니다. 

field의 값이 변경되었을 때 type이 새로운 인스턴스가 되어야 하면(근본적인 개념적 실체가 변경되면) `value Type`의 리시버를 사용,

그렇지 않아야 할 때는 `pointer type`의 리시버를 사용하면 됩니다.

time과 timer의 package로 예를 들어보겠습니다.

time은 `시각`을 나타냅니다 10시를 나타내는 time과 11시를 나타내는 time은 시간이라는 개념은 동일하지만 다른 시간이며 `서로 다른 실체입니다.`

하지만 timer를 20분으로 설정한 후 2분이 지난 timer와 10분이 지난 timer는 다른 timer가 아니라 처음 `20분을 설정할 때 사용 된 timer와 동일한 실체입니다.`

<br>

## go에서는 생성자 함수가 없다?

go에서는 생성자 함수를 지원하지 않습니다. 그렇기 때문에 생성자 함수를 개발자가 만들어 사용해야 합니다.

관례로는 함수 이름으로 `New + type 이름()`이 이용되며 생성한 type의 `pointer`을 return 해줍니다.

```go
func New<type 이름>(초기화에 필요한 값들...) *type{
    // 초기화 로직
    return &type{}
}

// ex
type Person struct{
    Name string
    Age int
    pID string
}

func NewPerson(name string, age int, pID string) *Person{
    return Person{
        Name : name,
        Age : age,
        pID : pID,
    }
}
```

생성자 함수를 만들어 `expose 되지 않는 field들을 초기화 한 후 해당 type에 종속된 method에서 사용할 수 있다.` 

<br>


## REFERENCE

https://www.youtube.com/watch?v=-ijeABV8vLU

https://jacking75.github.io/go_struct_pattern/
