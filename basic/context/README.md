Context
===

## Context란?

공식 golang document에는 다음과 같이 설명되어 있습니다.. 

>Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes

한글로 풀어 설명하면 "Context"란 패키지는 프로세스와 API간 전달되는 "Context"라는 type을 정의하는 패키지입니다.

이 type은 deadline(마감시간), cancellation(취소 시그널), request-scoped 값을 지닙니다.

"Context"는 하나의 `맥락`이고, 이 `맥락`을 유지하기 위해 다른 함수를 호출할 때 인자로 전달됩니다.

```go
//ex
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

여기서 주의해야 할 점은 Context는 struct로 가지고 있으면 안됩니다.

위 예제 코드처럼 항상 명시적으로 전달을 해야합니다. `nil`을 Context로 전달이 허용되도 절대 `nil`을 전달하면 안된다. 어떠한 Context를 전달해야 할지 모르겠다면 `context.TODO`를 전달하면 됩니다.
>Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.

<br>

## Context Type

context type은 interface이며 4개의 method로 구성되어 있습니다.
```go
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

<br>

### Done

Done method는 해당 Context가 cancel 혹은 타임아웃 되었을 때 `닫힌 channel을 리턴합니다.`

만약 cancel될 수 없는 context라면 Done method는 nil을 리턴할 수도 있습니다.
> Done may return nil if this context can never be canceled

cancel이 될 수 없는 context는 다음과 같습니다.

1. <code>context.Background()</code>
2. <code>context.TODO()</code>

<br>

### Err

만약 Done이 `닫혀있지 않다면 Err은 nil`을 return합니다. `Done이 닫혀있다면 Err은 non-nil error`를 리턴한다. error은 왜 context가 cancel되었는지 설명이 존재합니다.
>If Done is not yet closed, Err returns nil. If Done is closed, Err returns a non-nil error explaining why

<br>

### Deadline

Deadline method는 `마감기간이 존재할 때` 주어진 context의 마감기간을 리턴합니다. 만약 마감기간이 존재하지 않는 context라면 bool값으로 false를 리턴합니다.
>Deadline returns the time when work done on behalf of this context should be canceled. Deadline returns ok==false when no deadline is set

<br>

### Value

Value는 `context가 가지고 있는 값을 return`합니다. 인자로 key값을 넣으면 해당하는 값을 리턴합니다. 만약 key에 해당하는 값이 `존재하지 않는다면 nil을 return합니다.`
>Value returns the value associated with this context for key, or nil

<br>

## Empty Context

`context.Background()`, `context.TODO()`는 empty context입니다. 이 context들은 값을 지니지 않고, deadline도 없고 절대 cancel되지 않습니다.

Background()
>Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
보통 main함수, initalization, test, request 최상단에 선언되고 사용됩니다.

TODO()
>TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
어떠한 context를 사용해야 할지 불분명할 때 주로 사용됩니다.

empty context의 구현을 보면 아래와 같은데 context가 가지고 있는 method의 return값은 전부 `nil`이다.

<br>

```go
// An emptyCtx is never canceled, has no values, and has no deadline. It is not
// struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool){
    return
}

func (*emptyCtx) Done() <-chan struct{} {
    return nil
}

func (*emptyCtx) Err() error {
	  return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	  return nil
}

func (e *emptyCtx) String() string {
	  switch e {
	  case background:
		  return "context.Background"
	  case todo:
		  return "context.TODO"
	  }
	  return "unknown empty Context"
}

var (
  	background = new(emptyCtx)
	  todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```

<br>

## Cancel context

Context 사용에 있어서 이 cancel context 부분을 이해하는 것이 중요한데 Go를 동시성있게 고루틴을 이용한 코드를 작성할 수 있는 것이 cancel할 수 있는 context가 있기 때문입니다.

고루틴을 사용할 때 종료를 제대로 해주지 않게 되면 종료가 되지 않는 작업을 만들게 될 수도 있기 때문입니다.

그 때, 이 고루틴 종료를 관리해줄 수 있는 것이 바로 cancel 가능한 context입니다.

예를 들어보자면

유저의 `데이터를 가져오는 API와 이를 처리하는 Go 서버`가 있다고 생각해보았을 때 유저가 데이터를 달라고 `서버에 request를 보내면 서버에서 request를 처리하기 위해 필요한 정보들을 DB에서 가져오고`, 그 데이터를 적절하게 `가공하여 유저에게 다시 response를 보냅니다.`

<br>

<img src = https://user-images.githubusercontent.com/74294325/128026393-44d25ce4-b389-4c73-9ed9-615c67e46c39.png>

참조(https://devjin-blog.com/golang-context/)


그렇지만, 만약 유저가 response를 받기 전에 request를 cancel하게 되도, `서버는 그대로 request를 처리하기 위해 DB에서 데이터를 가져오고, 가공한 다음 response를 보내게 됩니다.`(물론 유저가 request를 cancel했기 때문에 response 전달도 실패한다.) 

이러한 경우 early cancel이 되어서 그 뒤에 있는 작업들은 request가 cancel되었다는 것을 모르고 `불필요한 작업을 계속 하게 되는 것`이다.

```이 때, cancel 가능한 context가 필요한 이유가 생기는 것입니다.```

<br>

<img src = https://user-images.githubusercontent.com/74294325/128026960-85353d28-82b2-4029-9b73-83dcde1b0563.png>

참조(https://devjin-blog.com/golang-context/)

<br>

cancel 가능한 context가 있다면, `유저가 response가 오기 전에 request를 cancel하게 되면 그 이상 작업을 진행하지 않고 작업을 종료할 수 있게 됩니다.`

이 때 최상단의 작업이 cancel됨을 밑에 진행되는 작업들이 스스로 cancel될 수 있어야 한다. Go의 cancel 가능한 context는 이런 부분들을 가능케 합니다.

이러한 context의 특징 덕분에, 고루틴을 사용해서 병렬적으로 작업을 진행하고 있더라도, `하나의 고루틴이 cancel되면 나머지 모든 고루틴의 작업들도 빠르게 cancel시킬 수 있게 되는 것 입니다.`

sql package의 내부 `Exec`관련 함수 내부적으로 다음과 같이 ctx가 done 되었는지 확인하는 code가 있는 것을 볼 수 있습니다.

```go
// ExecContext -> exec -> execDC -> ctxDriverExec
func ctxDriverExec(ctx context.Context, execerCtx driver.ExecerContext, execer driver.Execer, query string, nvdargs []driver.NamedValue) (driver.Result, error) {
	if execerCtx != nil {
		return execerCtx.ExecContext(ctx, query, nvdargs)
	}
	dargs, err := namedValueToValue(nvdargs)
	if err != nil {
		return nil, err
	}

	select {
	default:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return execer.Exec(query, dargs)
}
```

이처럼 ctx가 Done 되었다면 `return되어 Exec를 호출하지 않게 됩니다.` Done되지 않았다면 `Exec를 호출하여 query를 처리합니다.` 

<br>

## WithCancel

cancel할 수 있는 context를 만들어주는 WithCancel의 구현을 하나씩 살펴보기를 원합니다.

간략하게 WithCancel에 대해서 설명하자면 이 함수는 parent context의 copy본과 새로운 Done channel을 리턴합니다. 이 새로운 Done channel은 함수가 리턴하는 ```cancel 함수가 호출되던지, 혹은 parent context의 Done channel이 닫힐 때 같이 닫히게 됩니다.```
>WithCancel returns a copy of parent with a new Done channel. The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}
```

WithCancel 함수를 총 5가지 step으로 나눠서 보겠습니다.

### 1.Parameter

함수의 parameter값을 보면 `이 함수를 호출하기 전에 미리 parent context (context.Background() or context.TODO())를 만들고` 해당 context를 인자로 전달해야 함을 알 수 있습니다.
```go
ctx , cancel := context.WithCancel(context.Background())
```
<br>

### 2. Parent context Validation

위처럼 valid한 context를 WithCancel함수에 전달해야 하는 이유는 함수의 `첫줄에 nil을 validation하는 로직이 있기 때문입니다.` 이 때 만약 Context type의 값이 정상 값이 아니라면 panic을 일으킵니다.
```go
if parent == nil {
		panic("cannot create context from nil parent")
	}
```

<br>

### 3. 새로운 Cancel을 할 수 있는 context 생성

만약 parent context가 nil이 아니라면 새로운 cancel 가능한 context를 newCancelCtx라는 함수를 호출해서 만듭니다.
```go
c := newCancelCtx(parent)
```

newCancelCtx 함수는 cancelCtx라는 Struct를 리턴하는데, 이 struct의 Context type에 전달받은 `parent의 context를 initalize하고 return합니다.`
```go
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}

type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}
```

mutex는 다른 값들을 지켜주는 역할을하고, 이 메커니즘을 통해 context 패키지는 동시성이 가능해지는 것입니다.

cancelCtx는 cancel될 수 있습니다. 그리고 cancel될 때는 `이 canceler를 구현한 모든 children들이 함께 cancel된다.`

즉, newCancelCtx 함수에 parent context를 인자로 넣으면 새로은 context를 만들어서 리턴해주는 데, `이 새로운 context가 cancel되면 그 밑에 있는 모든 children들까지 cancel되는 것입니다.`

<br>

### 4. Propagate Cancel

WithCancel 함수에서 새로운 cancel 가능한 context가 만들어진 다음에 parent context와 방금 새로 만들어진 cancel가능한 context를 propagate라는 함수에 전달합니다.
```go
propagateCancel(parent, &c)
```

이 함수도 하나씩 분석해보자면 
```go
var goroutines int32

// propogateChannel은 parent가 cancel될때 child도 cancel 될 수 있게 한다
func propagateCancel(parent Context, child canceler) {
	done := parent.Done()
	if done == nil {
		return // parent이고 절대 cancel 되지 않는다
	}

	select {
	case <-done:
		// parent가 이미 cancel되었다
		child.cancel(false, parent.Err())
		return
	default:
	}

	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent가 이미 cancel되었다
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		atomic.AddInt32(&goroutines, +1)
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}
```

가장 먼저 `parent의 Done()`을 확인합니다. parent가 context.Background() 이거나 context.TODO()이면 값이 nil일 것이기 때문에 바로 리턴될 것입니다. 

`즉 parent의 cancel 여부를 확인하는 로직입니다.` 만약 인자로 들어온 `parent context가 cancel가능한 context인 cancelCtx이면 Done을 호출했을 때 nil이 아닐 것이다.`
```go
done := parent.Done()
if done == nil {
	return // parent이고 절대 cancel 되지 않는다
}
```

그 다음 로직이 수행된다는 것은, parent context가 cancel이 가능한 cancelCtx type이라는 것입다.

이 때, `parent의 channel이 닫혀 있는지 확인`을 합니다. 만약 parent context가 닫혀있다면 child도 그냥 닫아립니다.

`즉 parent의 context가 먼저 cancel되면 그 밑에 있는 child context도 항상 cancel이 보장됨을 알 수 있습니다.`
```go
select {
	case <-done:
		// parent가 이미 cancel되었다
		child.cancel(false, parent.Err())
		return
	default:
	}
```

channel이 닫혀있지 않다면 아직 parent가 cancel되지 않았다고 판단되서 그 밑의 로직들을 수행하게 됩니다.

밑에서는 parentCancelCtx에 parent context를 인자로 전달하고 리턴되는 bool값에 따라 if/else/ statement를 수행합니다.
```go
if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent가 이미 cancel되었다
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
} else {
		atomic.AddInt32(&goroutines, +1)
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
}
```

bool값을 return하는 parentCancelCtx함수는 아래의 코드와 같습니다.

이 함수의 목적을 설명하자면 `parent의 Value를 확인해서 가장 deep하게 있는 *cancelCtx를 찾고`, 이 `context의 Done()의 값과 parent의 Done()값이 같은지` 비교한다. 만약 값이 다르다면 *cancelCtx가 커스텀하게 한번 더 wrap되어서 사용되었기 때문에 `다른 done channel을 가지고 있음을 의미하게 된다.`
```go
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	done := parent.Done()
	if done == closedchan || done == nil {
		return nil, false
	}
	p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
	if !ok {
		return nil, false
	}
	p.mu.Lock()
	ok = p.done == done
	p.mu.Unlock()
	if !ok {
		return nil, false
	}
	return p, true
}
```

parentCancelCtx

1. true값이 리턴된다면 `현재 parent context에 cancel 가능한 cancelCtx type의 child를 덧붙힙니다.` parent에 `어떤 children들이 있는지 추적해야 하기 때문에 해당 정보를 map로 관리`한다. 만약 첫번째 child라면 map을 초기화 하고 그게 아니면 덧붙인다.

2. false값이 리턴된다면 `valid한 cancelCtx deep한 내부에 없다고 판단을 합니다.` 그러고는 parnet나 child의 Done() channel이 닫히는지 listen하는 고루틴을 활성화 합니다.

### context와 cancelfunc 리턴

위 과정이 다 마치면 WithCancel 함수는 새롭게 만들어진 cancel가능한 context와 그 context를 cancel할 수 있는 함수를 리턴하게 됩니다.
```go
return &c, func() { c.cancel(true, Canceled) }
```

<br>

## cancel()

WithCancel 함수를 통해 cancel 가능한 context를 만드는 것까지의 내용을 정리했다면 실제로 context가 어떤 방식으로 cancel되는지에 대한 이해가 남았습니다.
```go
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}
	for child := range c.children {
		// NOTE: acquiring the child's lock while holding parent's lock.
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

cancel가능한 context를 cancel하려고 cancel()함수를 호출하면 먼저 인자로 들어온 에러가 nil인지 확인합니다.

context의 에러를 확인하고 에러가 있다면 `이미 cancel된 것이기 때문에 early return을 합니다.`

아직 cancel이 되지 않았다면 `context의 channel을 닫아버립니다.` 자신의 channel도 닫아버리면서 `자신의 밑에 있는 모든 children context들도 순차적으로 cancel시킵니다.` 이렇게 되면 본인 및 본인 및에 있는 나머지 child context들은 다 cancel되는 것입니다.

마지막으로 남은 작업은 `본인과 parent context와의 관계를 끊는 것입니다.` removeChild를 호출하여 해당 작업을 수행합니다.
```go
func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}
```

<br>

## WithDeadline

WithDeadline은 WithCancel과 유사하게 cancel가능한 context와 cancel을 할 수 있는 함수를 리턴해줍니다. `차이점이 있다면 WithDeadline은 인자로 마감시간에 대한 정보를 같이받습니다.` 

또한 리턴되는 새로운 context는 특정 마감시간이 지나면 `자동으로 cancel됩니다.` (마감시간이 지나지 않더라도 cnacel함수를 통해 cnacel할 수 있다.)
```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		return WithCancel(parent)
	}
	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true, DeadlineExceeded) // 마감기간이 이미 지남
		return c, func() { c.cancel(false, Canceled) }
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded)
		})
	}
	return c, func() { c.cancel(true, Canceled) }
}
```

WithDeadline 또한 동일하게 가장먼저 parent context가 nil이 아닌지 확인합니다.
```go
if parent == nil {
	panic("cannot create context from nil parent")
}
```

그 다음에는 parent의 Deadline을 확인합니다.

이미 WithDeadline 함수로 전달된 `마감기간보다 이른 기한이라면 의미가없다고 판단`되서(parent가 cancel되면 child는 당연히 cancel되기 때문) WithCancel을 호출해서 마감기간은 없지만 cancel가능한 context를 리턴합니다.
```go
if cur, ok := parent.Deadline(); ok && cur.Before(d) {
	return WithCancel(parent)
}
```

위 조건을 패스한다면, `마감기간을 갖을수 있고 cacel이 가능한 context`라고 간주하고 timeCtx 타입의 context를 생성한다.
```go
c := &timerCtx{
	cancelCtx: newCancelCtx(parent),
	deadline:  d,
}

type timerCtx struct {
	cancelCtx
	timer *time.Timer // Under cancelCtx.mu.

	deadline time.Time
}
```

그 다음에는 WithCancel에서 그랫듯이 propagateCancel을 호출해서 새로운 context를 parent에 덧붙힙니다.
```go
propagateCancel(parent, c)
```

그 다음에는 context의 기간을 계산한다. 

만약 `마감시간이 이미 지났다면, cancel을 하고` context deadline exceeded 메세지를 보여주는 에러를 리턴한다. 마감기간이 `아직 지나지 않았다면 마감기간을 지났을 때 cancel되는 context를 리턴합니다.`
```go
dur := time.Until(d)
if dur <= 0 {
	c.cancel(true, DeadlineExceeded) // 마감기간이 이미 지남
	return c, func() { c.cancel(false, Canceled) }
}
c.mu.Lock()
defer c.mu.Unlock()
if c.err == nil {
	c.timer = time.AfterFunc(dur, func() {
		c.cancel(true, DeadlineExceeded)
	})
}
return c, func() { c.cancel(true, Canceled)
}
```

<br>

## WithTimeout

WithTimeout 함수에 `일정 시간을 지난 다음 context를 cancel하는` timeout시간을 전달하면 내부에서 현재시간 + 인자로 전달된 timeout시간을 더해 WithDeadline을 호출합니다.
```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}
```

<br>

## WithValue

WithValue는 cancel할 수 있는 contxt를 리턴하는 함수는 아니지만 인자로 `parent context와 key, value를 받아 인자로 받은 key와 value를 가지고 있는 context를 리턴한다.`
```go
func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}

// A valueCtx carries a key-value pair. It implements Value for that key and
// delegates all other calls to the embedded Context.
type valueCtx struct {
	Context
	key, val interface{}
}
```

context에서 값을 꺼내 사용할 때 주의해야 할 점이 있는데 context의 `value메서드의 리턴값은 interface{} 타입입니다.`

context의 값이 존재하지 않을 경우 `nil을 리턴합니다.` 그래서 context의 해당값이 `존재하는지` (value != nil), 그리고 `그 값이 원하는 타입이 맞는지` type assertion을 통해 확인(u, ok := v.(${type}))을 해야합니다.

<br>

## REFERENCE

https://devjin-blog.com/golang-context

https://jaehue.github.io/post/how-to-use-golang-context