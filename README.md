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