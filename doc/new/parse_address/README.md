mail package의 ParseAddress
===
 
Email형식의 String type을 validation하기 위해 reference를 찾다가 golang에서 제공해주는 mail package를 찾게 되었습니다.

mail package에서 지원하는 API를 이용하여 validation을 할 수 있을까 했지만 충분히 원하는 결과는 얻지 못하였습니다. 하지만 mail Pacakge의 API를 사용하면서 알게 된 점을 공유하기 원합니다.

## mail.ParseAddress(address string)

mail Package에서 제공해주는 API중 `ParseAddress(address string)`이라는 API가 있습니다.

내부 code는 다음과 같이 작성되어 있습니다.
```go
// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
func ParseAddress(address string) (*Address, error) {
	return (&addrParser{s: address}).parseSingleAddress()
}
```
<br>

address에 들어갈 수 있는 형식은 RFC 5322를 따르고 있습니다.
```
(
    address         =   mailbox / group
    mailbox         =   name-addr / addr-spec
    name-addr       =   [display-name] angle-addr
    angle-addr      =   [CFWS] "<" addr-spec ">" [CFWS] /
                        obs-angle-addr
    group           =   display-name ":" [group-list] ";" [CFWS]
    display-name    =   phrase
    mailbox-list    =   (mailbox *("," mailbox)) / obs-mbox-list
    address-list    =   (address *("," address)) / obs-addr-list
    group-list      =   mailbox-list / CFWS / obs-group-list
)

(
    3.4-1
    address-spec    =   local-part "@" domain
    domain          =   dot-atom / domain-literal / obs-domain

    (The local-part portion is a domain-dependent string.)
)

(
    <domain>        ::= <subdomain> | " "

    <subdomain>     ::= <label> | <subdomain> "." <label>

    <label>         ::= <letter> [ [ <ldh-str> ] <let-dig> ]
)
```
참조 

RFC-5322 3.4(https://datatracker.ietf.org/doc/html/rfc5322#section-3.4)

RFC-5322 3.4.1(https://datatracker.ietf.org/doc/html/rfc5322#section-3.4-1)

RFC-1034 3.5(https://datatracker.ietf.org/doc/html/rfc1034#section-3.5)

<Br>

내부적으로 `parseSingleAddress()`가 호출되는데 내부에는 인터프리터로 address값을 분석하여 결과를 retrun해줍니다.

참조 

mail_package parseSingleAddress(https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/net/mail/message.go;drc=3ef8562c9c2c7f6897572b05b70ac936a99fd043;l=345)


<br>

## 사용방법

golang의 공식 doc에서 제공해주는 exmaple처럼 사용하면 됩니다. 

```go
func main() {
	e, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.Name, e.Address)

}

//output
Alice alice@example.com
```

<br>

**하지만!!**

정규식을 사용했을 때 보다는 더 정교하게 validation하기는 어렵습니다. RFC-5322, RFC-1034 에 따르면 다음과 같은 문자열도 정상적으로 작동합니다.

```
alice@example
```

왜냐면 RFC-5322를 따랐을 때 domain에 `.com`과 같이 붙지 않아도 domain으로 취급되기 때문입니다. 

<br>

## Test

다음과 같은 mockData를 이용하여 Test를 진행하였습니다.
```go
"Alice <alice@example.com>",
"<alice@example.com>",
"alice@example.com",
"alice@example",
"bad-example",
"",
"@",
```

<br>

결과는 아래와 같습니다.
```go
2021/11/30 00:29:24 name: Alice, address: alice@example.com
2021/11/30 00:29:24 name: , address: alice@example.com
2021/11/30 00:29:24 name: , address: alice@example.com
2021/11/30 00:29:24 name: , address: alice@example
2021/11/30 00:29:24 mail: missing '@' or angle-addr
2021/11/30 00:29:24 mail: no address
2021/11/30 00:29:24 mail: no angle-addr
```

`@`가 존재하지 않는 Address나 empty string 혹은 `angle-addr`에 맞지 않은 경우 `ParseAddress(address string)`에서 error를 return을 확인할 수 있었습니다.

<br>

## 결국은 정규식??

위의 mail package를 이용하여 email 형식의 문자열을 validation 할 수도 있지만!

**로그인/가입 양식**유효성 검사에서 이를 사용하면 많은 사람들이 부분적 또는 잘못된 이메일 주소를 입력하여 프로덕션 에 잘못된 레코드 가 잔뜩 생성 될 수도 있습니다.

그렇기 때문에 결국 해당 로직을 구현할 때 정규식을 사용하였습니다. go에서 정규식을 사용할 때는 `regexp` 패키지를 사용합니다. 예제코드는 다음과 같습니다. 

```go
emailreg := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

ok, err := regexp.MatchString(emailreg, <정규식에 대입하여 검증할 문자열>)

//ex
given := []string{
    "Alice <alice@example.com>",
    "<alice@example.com>",
    "alice@example.com",
    "alice@example",
    "bad-example",
    "",
    "@",
}

for _, v := range given {
    ok, err := regexp.MatchString(emailreg, v)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%s Match result: %v", v, ok)
}

// output
2021/11/30 21:11:26 Alice <alice@example.com> Match result: false
2021/11/30 21:11:26 <alice@example.com> Match result: false
2021/11/30 21:11:26 alice@example.com Match result: true
2021/11/30 21:11:26 alice@example Match result: false
2021/11/30 21:11:26 bad-example Match result: false
2021/11/30 21:11:26  Match result: false
2021/11/30 21:11:26 @ Match result: false
```
