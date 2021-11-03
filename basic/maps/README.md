Map 
===

## Map이란?

맵은 key와 value 형태로 data를 저장하는 자료구조 입니다. 언어에 따라 딕셔너리, 해쉬 테이블, 해쉬 맵 등으로 불립니다.

map의 종류로는 hash map과 sorted map이 있는데 go에서는 hash map을 이용합니다. 그렇기 때문에 삽입되는 요소의 `순서는 보장되지 않습니다.`

<br>

## syntax

맵을 선언 및 초기화를하는 방법은 두가지가 있습니다.

0. 선언

    ```go
    var m map[key type] value type
    ```

    선언만 한 경우는 map이 비어있습니다. 즉 비어있는 map입니다.(nil)

1. 내장함수 make를 이용

    ```go
    m := make(map[key type]value type)
    ```

2. 선언을 하면서 초기화

    ```go
    m := map[key type]value type {
        key : value
    }

    // ex
    m := map[string]string {
        "foo" : "bar",
    }
    ```
<br>

## map도 내부적으로 ref를 이용합니다.

map은 선언 및 초기화를 할 때 `memory에 해쉬 테이블을 만들고 해당 메모리의 주소를 가리키는 pointer를 return합니다.`

즉 map을 값으로 가지고 있는 변수의 크기는 다 동일하다는 것 입니다. -> `8byte`

`reflect package의 DeepEqual`의 Api를 이용하여 map을 만들어 비교해보겠습니다.
>Map values are deeply equal when all of the following are true: they are both nil or both non-nil, they have the same length, and either they are the same map object or their corresponding keys (matched using Go equality) map to deeply equal values. <br> map을 비교할 때는 둘다 nil인지 nil이 아닌지, 두 map이 동일한 길이와 같은 map 객체를 가지고 있는지, key와 value가 동일한지 비교합니다.

```go
func TestMapReference(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
	}

	returnM := returnMap(m)
	returnM["bar"] = "foo"

	t.Log(reflect.DeepEqual(m, returnM))

	t.Log(m)
	t.Log(returnM)
}

func returnMap(m map[string]string) map[string]string {
	return m
}

output
maps_test.go:32: true
maps_test.go:34: map[bar:foo foo:bar]
maps_test.go:36: map[bar:foo foo:bar]
```

test code를 이용하여 살펴본 결과 `returnMap`에서 내부적으로 복사한 map을 return하여 원본의 map과 비교하였을 때 결과는 `true`이며 `원본이 아닌 map에서 요소를 추가하였을 때 원본의 map에도 영향이 있는 것을 확인했습니다.`

<br>

## Map의 순회

내장 keyword인 range를 이용항여 map을 순회할 수 있는데 slice나 배열과 달리 map에 range를 이용하면 `key`와 `value`를 얻을 수 있다.

```go
for key, value := range m {
    //... 
}
```

<br>

## 요소 가져오기 및 존재 여부

map에서 요소를 가져올 때 아래와 같이 가져올 수 있습니다.

```go
m := map[string]string {
    "foo" : "bar"
}

val := m["foo"] // "bar"
val2 := m["noKey"] // ""
```
key를 이용하여 value에 접근을 합니다. key에 해당하는 `요소가 없을 경우에는 value type의 zero value`를 return 합니다.

**하지만!**

요소의 value가 type의 zero value와 같다면 어떻게 알 수 있을까요?

map에서 요소를 가져올 때 두번 째 값으로 요소가 존재하는지에 대한 여부를 확인할 수 있는 flag값을 제공합니다.

```go
m := map[string]string {
    "foo" : "bar"
}

val, ok := m["foo"] // "bar" , true
val2, ok2 := m["noKey"] // ""  , false
```

map안에 key와 value가 `존재한다면 value와 true를` return하고 존재하지 않는다면 `value type의 zero value와 false`를 return 합니다.






## REFERENCE

https://go.dev/blog/maps

https://velog.io/@zuyonze/%ED%95%B4%EC%8B%9C%ED%95%A8%EC%88%98%EC%97%90-%EB%8C%80%ED%95%9C-%EA%B0%9C%EB%85%90-%EC%A0%95%EB%A6%AC%ED%95%98%EA%B8%B0