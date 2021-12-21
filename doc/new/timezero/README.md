time.Time의 Zero Value에 대해서
===

코드를 작성할 때 time.Time을 많이 사용하기는 했지만 `time.Time`의 **Zero Value**에 대해서는 생각해보지 않았다.

진행 흐름은 다음과 같았다. 

Request에서 만약 `time.Time` 필드에 Null이 들어왔을 때 `sql.NullTime`을 이용하여 DB에 null값을 저장한다.

<br>

나의 고민은 여기서 부터 시작이였다. **time.Time**을 어떻게 **zero Value** 체크를 할 수 있을까? 

**아래와 같이 nil을 이용하여 체크할 수 있을까?**
```go
testTime := time.Time{}

// cannot convert nil (untyped nil value) to struct{}
if testTime == nil{ // 불가능
    // ... 
}
```
역시 정답은 공식 문서에 있었다... (https://pkg.go.dev/time#Time.IsZero)

<br>

`time.Time` 패키지에서 `IsZero()` API를 지원하고 있다.

```go
// IsZero reports whether t represents the zero time instant, January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool {
    return t.sec() == 0 && t.nsec() == 0
}
```
<br>

내부적으로 `sec`와 `nsec`가 **0과 같은지** 판단 하여 `boolean`값을 return해준다.

<br>

`time.Time`의 구조체를 생성 후 test를 진행해보면 다음과 같다.
```go
//given
testTime := time.Time{} // 1

//when
isZero := testTime.IsZero() // 2

//then
if !isZero {
    t.Fatal("testTime should be zero")
}

t.Log(testTime)
t.Log(testTime.Second())
t.Log(testTime.Nanosecond())

//output
/Users/leewoooo/dev/go/src/gopher_lee/doc/new/timezero/timezero_test.go:20: 0001-01-01 00:00:00 +0000 UTC
/Users/leewoooo/dev/go/src/gopher_lee/doc/new/timezero/timezero_test.go:21: 0
/Users/leewoooo/dev/go/src/gopher_lee/doc/new/timezero/timezero_test.go:22: 0
```

1. `time.Time`의 구조체를 생성한다. (zero Value)

2. `IsZero()`를 이용하여 해당 time변수가 zero Value인지 확인한다.
