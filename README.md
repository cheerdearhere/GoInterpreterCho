# I. 기본정보
- 사용언어: javascript
> go test src/main/lexer/lexer.go
# II. Token, Lexer, REPL(`Read Eval Print Loop`)
- 콘솔(console)/대화형(interactive) 모드
  - Read input
  - Eval by Interpreter
  - Print result
  - Loop
# III. Parser
- 입력값을 받아 자료구조를 만들어 프로그램이 처리할 수있도록 구조화하는 작업을 담당
- 구문이 올바른지 분석하거나 더 효율적으로 작동하도록 변환시키기도 함. 
- 보통은 lexical analyzer(어휘 분석기)를 따로 두기도 함
- Parser generator를 사용하기도하지만 지금은 배우는 시기이므로 그냥 만들어본다
## A. 전략
- 프로그래밍 언어를 파싱할때 크게 두가지 전략이 있다
  - Top-down(하향식)
    - recursive descent parsing(재귀적 하향 파싱): 여기서 만들어볼 파서(top down operator precedence parser)
    - earley parsing(얼리 파싱): 문맥무관 문법에 속한 문자열을 파싱하는 알고리즘
    - predictive parsing(예측적 파싱)
  - Bottom-up(상향식)
## B. 제작 순서
> let statement > return statement > expression statement > platte parser ...
## C. Let 문 파싱
- 변수 바인딩: `let` `<identifier>` = `<expression>`;
```javascript
// let 식별자 = 표현식;
    let x = 10;
    let y = 15;
    let add = function(a,b){
        return a + b;
    };
```
## D. Return 문 파싱
```javascript
return 5;
return 10;
return add(15);
```
## E. 표현식 파싱
### 1. 추가로 고민할 사항
  - 연산자 우선순위는 어떻게 구현할 것인가?
  - ()를 사용한 우선순위는?
  - 함수, 객체와 같은 포인터 이동은 어떻게 구현할 것인가
### 2. 프랫 파싱(vaughan pratt parsing)
- 이곳에서 구현할 표현식: 프랫 파싱(vaughan pratt parsing)
  - 하향식 연산자 우선 순위(Top Down Operator Precedence) 1973 출간
```
    (상략)...
    하향식 연산자 우선 순위는 쉽게 이해할 수 있을만큼 단순하며, 
    구현과 사용이 아주 쉽고, 
    이론적으로는 아닐지 몰라도 극도로 실용적이며, 
    합당한 구문적 요구 사항을 대다수 만족시키고도 남을 만큼 유연하다
    ...(하략)
```
- 자바스크립트를 구현한 생각
### 3. 용어정리
- 전위 연산자: prefix operator
  - 피연산자 앞에 붙는 연산자
  - 모든 연산자 중 가장 우선시한다
  ```javascript
  let a = 5; 
  console.log(--a);//감소연산자
  ``` 
- 후위 연산자: postfix operator
  - 피연산자 뒤에 붙는 연산자
  - 다른 연산이 종료된 후 처리된다
  ```javascript
  console.log(a--);//log(a)를 수행한 뒤 수행
  ```
- 중위 연산자: infix operator
  - 피연산자들 사이에 위치하는 연산자
  - 두개의 피연산자를 요구하므로 이항 표현식(binary expression)에서 사용된다
  ```javascript
  let a = 5 + 6 * 10;  
  ```
  - 연산자간 우선순위를 어떻게 판별할 것인가에대한 구분 필요
### 4. 표현식 AST
- 명령문 준비
```javascript
let x = 5;
x + 10;
```
### 5. 리터럴도 표현식이다
- 정수 literal
```javascript
let x = 5;
add(5, 10);
5 + 5 + 5;
```
- 정수 리터럴의 위치에 다른 표현식을 넣어도 이상이 없다
#### a. 전위연산자(prefix)
- 두개의 전위연산자
  - `-`: -5
  - `!`: !foobar
- 전위연산자의 구조: `<prefix operator><expression>`
  - 전위 연산자의 피연산자는 어떤 표현식이든 받을 수 있는 유연함이 있어야한다. 








