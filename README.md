# IGIS WebApp BackEnd Service

```
go mod init IGISBackEnd
go get -u "github.com/gorilla/mux"
go get -u "github.com/go-sql-driver/mysql"
```

## v0.1.0 api
<p>

API 버전 1

Subrouter포인트: api/v1
특징:
- 데이터베이스가 형성되지 않았기 때문에, csv를 그대로 json으로 쏴줄 것임. 
  - 서버 시작시 json을 미리 로드 하고, 요청시마다 가공해서 쏴줌
- 4개의 엔드포인트 제공 (localhost:8080/api/v1/...)
  1. /asset
  2. /debt
  3. /macro
  4. /debtRowCount

- Checklist ::, Macro ::, Debt :: 에서 요청사항처리 및 에러 확인 가능

실행 후 python으로 확인작업

</p>

실행 후 Test 주소 모음

* DebtRowCount: http://localhost:8080/api/v1/debtRowCount?yearFrom=2000&yearUntil=2021&aumFrom=1&aumUntil=100000000000&debtFrom=1&debtUntil=1000000000000
* Asset: http://localhost:8080/api/v1/asset?strat=Core
* Debt: http://localhost:8080/api/v1/debt?yearFrom=2010&yearUntil=2020
* Macro: http://localhost:8080/api/v1/macro?commodity=kr1y&yearFrom=2010&yearUntil=2020


### json map

0. /debtRowCount

* Request at 
  * http://localhost:8080/api/v1/debtRowCount?yearFrom=2000&yearUntil=2021&aumFrom=1&aumUntil=100000000000&debtFrom=1&debtUntil=1000000000000
  * yearFrom, yearUntil, aumFrom, aumUntil, debtFrom, debtUntil 전부 정수

| jsonCode | type | 
|----------|------|
| "rc"   |  int  |

```json
{"rc":402}
```


1. /asset

* Request at
  * http://localhost:8080/api/v1/asset?strat=Core
  * strat can be Core, Value-added, Oppotunistic 엑셀에 적힌대로

```json
[
    {"univ":1,
    "fc":"112001",
    "fn":"이지스전문투자형사모부동산투자신탁1호",
    "an":"2001아울렛(안양), NC백화점(강서)",
    "domfor":"국내",
    "at":"리테일",
    "ft":"Fund",
    "ismom":"일반",
    "strat":"Core",
    "iscmplt":"완료",
    "fsource":"ECM",
    "fpath":"내 PC\\IDCS_문서중앙서버 (R:)\\전사 폴더\\이지스자산운용\\02.집합투자기구\\제01호(강서NC)-AMC변경\\계약서\\안양점추가매입계약서\\대출계약서",
    "misc":""
    }
    , ... 
]
```

2. /debt

* Request at
  * http://localhost:8080/api/v1/debt?yearFrom=2010&yearUntil=2020
  * yearFrom, yearUntil 전부 정수

```json
[
    {
        "fc":"112001",
        "fn":"이지스전문투자형사모부동산투자신탁1호",
        "sdate":"2010-08-30",
        "mdate":"2050-03-11",
        "an":"2001아울렛(안양), NC백화점(강서)",
        "ac":"",
        "domfor":"국내",
        "at":"리테일",
        "ft":"Fund",
        "it":"실물",
        "strat":"",
        "area":"40971",
        "equity":"46,382,000,000",
        "loan":"176,000,000,000",
        "aum":"222,382,000,000",
        "ltv":"79.14%",
        "ro":"",
        "loandate":"2012-02-28",
        "lpnum":"3",
        "lp":"주식회사 우리은행",
        "lpcorp":"주식회사 우리은행",
        "lpt":"은행",
        "seniorstr":"선",
        "loanuse":"기존 대출계약(NC백화점 강서점을 취득하는데 필요한 자금을 조달하기 위하여 차주가 본건 투자신탁의 신탁업자의 자격으로 2010년 8월 30일 강서이랜드리테일제일차유한회사 등과 체결한 대출약정)에 따라 차주가 본건 투자신탁의 신탁업자의 자격으로 부담하고 있는 대출원리금의 채무의 상환(970 억원) 및 2001아울렛 안양점의 취득자금 및 운영자금의 조달(270 억원)",
        "loancls":"담보",
        "addr":"서울시 강서구 등촌동 689-2, 경기도 안양시 만안구 안양동 627-287",
        "loanamt":"54,000,000,000",
        "loanrpy":"대출만기일에 전액 상환",
        "rate":"고정",
        "loanintcls":"상수",
        "loanintfloat":"",
        "sdaterate":"5.39%",
        "spread":"1.86%",
        "loanpremium":"Null",
        "intdur":"1",
        "laterate":"19.00%",
        "lateratecls":"상수",
        "sdatelaterate":"",
        "earlypremium":"1%",
        "earlypremiumcls":"상수",
        "guranteelimit":"130%",
        "dscr":"검증조건 1.2",
        "intdeposit":"3",
        "default":"지급기일에 지급하지 아니한 경우",
        "opinion":"1/2 이상",
        "lender":"주식회사 국민은행",
        "trustee":"주식회사 국민은행",
        "amc":"피에스자산운용 주식회사",
        "financialinst":"Null",
        "mm":"주식회사 우리은행",
        "cashsupp":"Null",
        "debtundwrt":"주식회사 이랜드리테일",
        "builder":"Null",
        "duration":"36"
    }, 
    ...
]
```


3. /macro

* Request at
  * http://localhost:8080/api/v1/macro?commodity=kr1y&yearFrom=2010&yearUntil=2020
  * commodity 는 kr1y, kr3y, kr5y 엑셀에 적힌 대로(소문자로)
  * yearFrom, yearUntil 정수

```json
[
    {"date":"20100104","value":3.5},
    {"date":"20100105","value":3.43},
    {"date":"20100106","value":3.46},
    ...
]

```