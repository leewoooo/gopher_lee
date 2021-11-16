json unmarshal with map[string]interface{}
===

## 구조체를 map[string]interface{}로 Unmarshal하면 int계열 자료형은 float64로?

code를 작성중에 구조체를 `map[string]interface{}`로 Unmarshal하는 과정에서 int 계열 자료형이 암시적으로 `float64`로 변환되는 것을 확인했습니다. code는 아래와 같습니다.

```go
type UserClaim struct {
	ID   int64   `json:"id"`
	Name string `json:"name"`
}

func (u *UserClaim) ConvertMap() (map[string]interface{}, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	var resultMap map[string]interface{}
	if err := json.Unmarshal(bytes, &resultMap); err != nil {
		return nil, err
	}

	return resultMap, nil
}

//testcode
func TestConvertMap_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	//given
	given := UserClaim{
		ID:   1,
		Name: "foo",
	}

	//when
	convertedMap, err := given.ConvertMap()

	//then
	assert.NoError(err)
	assert.NotEmpty(convertedMap)

	id, ok := convertedMap["id"].(int64) // int64 x float64 o
	// int64로 unmarshal되어 같을 것이라 예상되지만 int64 타입이 아니기에 zero value return
	assert.Equal(int64(1), id) 
	// int64가 아니기 때문에 false
	assert.True(ok) 
	t.Log(reflect.TypeOf(convertedMap["id"])) //float64

	name, ok := convertedMap["name"].(string)
	assert.Equal("foo", name)
	assert.True(ok)
}
```

<br>

## json package의 내부 code

unmarshal api를 쭉 타고 들어가다 보면 아래와 같은 code를 마주하게 됩니다. unmarshal의 대상이 되는 `리터럴이 interface{} 타입의 숫자형태면 아래와 같은 case를 지나게 됩니다.`

```go
//decode -> unmarshal -> value -> literalStore
default: // number
    if c != '-' && (c < '0' || c > '9') {
        // ...
        case reflect.Interface:
            n, err := d.convertNumber(s)
            if err != nil {
                d.saveError(err)
                break
            }
            if v.NumMethod() != 0 {
                d.saveError(&UnmarshalTypeError{Value: "number", Type: v.Type(), Offset: int64(d.readIndex())})
                break
            }
            v.Set(reflect.ValueOf(n))
		// ...

// convertNumber converts the number literal s to a float64 or a Number
// depending on the setting of d.useNumber.
func (d *decodeState) convertNumber(s string) (interface{}, error) {
	if d.useNumber {
		return Number(s), nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, &UnmarshalTypeError{Value: "number " + s, Type: reflect.TypeOf(0.0), Offset: int64(d.off)}
	}
	return f, nil
}
```

<br>

즉 convertNumber를 통하여 `float64가 되는 것을 확인할 수 있습니다.` 내부적으로 `ParseFloat()`를 호출하고 호출한 결과를 return하게 됩니다.

<br>

## 해결 방안

1. 강제 형 변환

float으로 암시적 형 변환이 된 값을 `float`으로 명시 후 강제적으로 형변환을 시킵니다.

```go
// ...
id := int64(convertedMap["id"].(float64)) 
assert.Equal(int64(1), id)
t.Log(reflect.TypeOf(id)) // uint
```

<br>

2. Decoder사용

decoder를 이용하여 `UseNumber`을 이용하여 decoder의 state값을 변경하여 사용합니다. `UseNmeber`를 호출하게 되면 내부적으로 decodeState 구조체의 필드 중 `useNumber`가 ture로 변경되게 됩니다.

그렇게 되면 `convertNumber` API 내부에서 조건문에 의해 `json.Number` 타입으로 리턴을 받게 됩니다.

```go
func (d *decodeState) convertNumber(s string) (interface{}, error) {
	if d.useNumber {
		return Number(s), nil
	}
	//...
```

이렇게 return받게 된 값을 사용하는 것은 아래의 예제코드와 같이 사용할 수 있습니다. 해당하는 값을 `json.Number`로 명시한 후 `Int64` API를 호출하여 `int64` 타입으로 사용할 수 있습니다.
```go
//...
marshaled, err := json.Marshal(given)
assert.NoError(err)

buf := bytes.NewBuffer(marshaled)

var target map[string]interface{}
d := json.NewDecoder(buf)
d.UseNumber()
d.Decode(&target)
assert.NoError(err)

//then
id, err := target["id"].(json.Number).Int64()
assert.Equal(int64(1), id)
assert.NoError(err)
```
