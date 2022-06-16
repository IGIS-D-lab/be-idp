# IGIS WebApp BackEnd Service

```
go mod init IGISBackEnd
go get -u "github.com/gorilla/mux"
go get -u "github.com/go-sql-driver/mysql"
	

```

## v1 api
<p>

API 버전 1

Subrouter포인트: api/v1
특징:
- 데이터베이스가 형성되지 않았기 때문에, csv를 그대로 json으로 쏴줄 것임. 
- 3개의 엔드포인트 제공 (localhost:8080/api/v1/...)
  1. /asset
  2. /debt
  3. /macro

- 쿼리문 받아서 쓰는 기능 추가 필요해보임 
  - /debt 엔드포인트의 큰 데이터 크기
- Checklist ::, Macro ::, Debt :: 에서 요청사항처리 및 에러 확인 가능

실행 후 python으로 확인작업

</p>

```python
import requests

asset_test = "http://localhost:8080/api/v1/asset"
debt_test = "http://localhost:8080/api/v1/debt"
macro_test = "http://localhost:8080/api/v1/macro"


# asset
r0 = requests.get(macro_test)
print(r0.status_code)
d = r0.json()
print(d.keys())

# debt
r1 = requests.get(debt_test)
print(r1.status_code)
d1 = r1.json()
print(d1.keys())

# macro
r2 = requests.get(macro_test)
print(r2.status_code)
d2 = r2.json()
print(d2.keys())
```

### json map

0. 공통

| jsonCode | Description | Type |
|----------|-------------|------|
|    fsht  | from sheet  | string|
|   desc   | data description | string |
| data| data rows | 엔드포인트마다 다름 |

```json
{
    "fsht": "checklist", 
    "desc": "IDP Excel. Checklist Sheet", 
    "last": "20220615T17:42:37", 
    "data": [
        {
            "univ": 1, 
            "fc": "112001", 
            "fn": "\uc774\uc9c0\uc2a4\uc804\ubb38\ud22c\uc790\ud615\uc0ac\ubaa8\ubd80\ub3d9\uc0b0\ud22c\uc790\uc2e0\ud0c11\ud638", 
            "an": "2001\uc544\uc6b8\ub81b(\uc548\uc591), NC\ubc31\ud654\uc810(\uac15\uc11c)", 
            "domfor": "\uad6d\ub0b4", 
            "at": "\ub9ac\ud14c\uc77c", 
            "ft": "Fund",
            "it": "\uc2e4\ubb3c", 
            "ismom": "\uc77c\ubc18", 
            "strat": "Core", 
            "iscmplt": "\uc644\ub8cc", 
            "univexcl": NaN, 
            "fsource": "ECM", 
            "fpath": "\ub0b4 PC\\IDCS_\ubb38\uc11c\uc911\uc559\uc11c\ubc84 (R:)\\\uc804\uc0ac \ud3f4\ub354\\\uc774\uc9c0\uc2a4\uc790\uc0b0\uc6b4\uc6a9\\02.\uc9d1\ud569\ud22c\uc790\uae30\uad6c\\\uc81c01\ud638(\uac15\uc11cNC)-AMC\ubcc0\uacbd\\\uacc4\uc57d\uc11c\\\uc548\uc591\uc810\ucd94\uac00\ub9e4\uc785\uacc4\uc57d\uc11c\\\ub300\ucd9c\uacc4\uc57d\uc11c", 
            "fname": "\ub300\ucd9c\uc57d\uc815\uc11c_20120228", 
            "misc": "NC\ubc31\ud654\uc810(\uac15\uc11c)"
        }, ...
}
```


####  Data 안 jsonID

1. /asset

| jsonCode | Description        | Type |
|----------|--------------------|------|
| univ     | Universe           |      |
| fc       | 펀드 코드          |      |
| fn       | 펀드명             |      |
| an       | 자산명             |      |
| domfor   | 국내 해외          |      |
| at       | 투자 자산 유형     |      |
| ft       | 펀드 구분          |      |
| it       | 투자 유형          |      |
| ismom    | 모자 구분          |      |
| strat    | 투자 전략          |      |
| iscmplt  | 완료 여부          |      |
| univexcl | Universe 제외 이유 |      |
| fsource  | 데이터 소스 (ECM)  |      |
| fpath    | 파일 경로          |      |
| fname    | 파일 명            |      |
| misc     | 비고               |      |


2. /debt

| jsonCode  | Description    | Type | jsonCode        | Description                        | Type |
|-----------|----------------|------|-----------------|------------------------------------|------|
| fc        | 펀드코드       |      | rate            | 금리 종류                          |      |
| fn        | 펀드명         |      | loanintcls      | 대출 이자율 분류                   |      |
| sdate     | 설정일         |      | loanint         | 대출 이자율                        |      |
| mdate     | 만기일         |      | loanintfloat    | 변동이자율기준                     |      |
| an        | 자산명         |      | sdaterate       | 체결일 이자율                      |      |
| am        | 자산수         |      | spread          | 스프레드                           |      |
| domfor    | 국내해외       |      | loanpremium     | 대출취급수수료                     |      |
| at        | 투자 자산 유형 |      | intdur          | 이자기간                           |      |
| ft        | 펀드 구분      |      | laterate        | 연체이자율                         |      |
| it        | 투자 유형      |      | lateratecls     | 연체이자율 분류                    |      |
| start     | 투자 전략      |      | sdatelaterate   | 체결일연체이자율                   |      |
| area      | 연면적 (평)    |      | earlypremium    | 조기상황수수료                     |      |
| equity    | EQUITY 총액    |      | earlypremiumcls | 조기상환수수료분류                 |      |
| loan      | LOAN 총액      |      | guranteelimit   | 담보한도율                         |      |
| aum       | AUM 총액       |      | guranteemax     | 담보채권최고액                     |      |
| ltv       | LTV            |      | dscr            | DSCR                               |      |
| loannum   | 대출계약서 수  |      | dscrval         | DSCR    rkqt                       |      |
| ro        | 롤오버         |      | intdeposit      | 이자유보금  interest deposit       |      |
| trnsdate  | 매매계약체결일 |      | default         | 채무불이행요건                     |      |
| loandate  | 대출계약체결일 |      | opinion         | 대주의의사결정                     |      |
| lpnum     | 대주수         |      | contact         | 통지처                             |      |
| lp        | 대주           |      | lender          | 차주                               |      |
| lpname    | 대주명         |      | trustee         | 신탁업자                           |      |
| lpcorp    | 대주명회사     |      | amc             | 집합투자업자                       |      |
| lpt       | 대주명분류     |      | financialinst   | 대리금융기관                       |      |
| seniorstr | 대출순위       |      | mm              | 자금관리자 Money Manager           |      |
| senior    | 우선순위       |      | cashsupp        | 자금보충인 cash deficiency support |      |
| loanuse   | 대출용도       |      | debtundwrt      | 채무인수인 debt underwrite         |      |
| sellto    | 매도인         |      | builder         | 시공사                             |      |
| loantype  | 대출 종류      |      | loanplan        | 대출실행예정일 loan plan date      |      |
| loancls   | 대출 분류      |      | loanexec        | 대출 실행일 loan execute date      |      |
| addr      | 토지 주소      |      | loanmatr        | 대출 만기일 loan maturity date     |      |
| loanratio | 대출 참가 비율 |      | duration        | 듀레이션 duration                  |      |
| loanamt   | 대출약정금     |      | loanmatrymd     | 대출만기일 YMD                     |      |
| loanrpy   | 대출금상환방식 |      |                 |                                    |      |

```json
{
    "fsht": "macro", 
    "desc": "IDP Excel. Debt Information", 
    "last": "20220615T18:21:10", 
    "data": [
        {
            "fc": "112001", 
            "fn": "\uc774\uc9c0\uc2a4\uc804\ubb38\ud22c\uc790\ud615\uc0ac\ubaa8\ubd80\ub3d9\uc0b0\ud22c\uc790\uc2e0\ud0c11\ud638", 
            "sdate": "2010-08-30", 
            "mdate": "2050-03-11", 
            "an": "2001\uc544\uc6b8\ub81b(\uc548\uc591), NC\ubc31\ud654\uc810(\uac15\uc11c)", 
            "am": "2", 
            "domfor": "\uad6d\ub0b4", 
            "at": "\ub9ac\ud14c\uc77c", 
            "ft": "Fund", 
            "it": "\uc2e4\ubb3c", 
            "start": "Core", 
            "area": "40971", 
            "equity": "46,382,000,000", 
            "loan": "176,000,000,000", 
            "aum": "222,382,000,000", 
            "ltv": "79.14%", 
            "loannum": "2", 
            "ro": "1", 
            "trnsdate": "2012-03-20", 
            "loandate": "2012-02-28", 
            "lpnum": "3", 
            "lp": "\uc8fc\uc2dd\ud68c\uc0ac \uc6b0\ub9ac\uc740\ud589", 
            "lpname": "\ub300\uc8fc1", 
            "lpcorp": "\uc8fc\uc2dd\ud68c\uc0ac \uc6b0\ub9ac\uc740\ud589", 
            "lpt": "\uc740\ud589", 
            "seniorstr": "\uc120", 
            "senior": "1", 
            "loanuse": "\uae30\uc874 \ub300\ucd9c\uacc4\uc57d(NC\ubc31\ud654\uc810 \uac15\uc11c\uc810\uc744 \ucde8\ub4dd\ud558\ub294\ub370 \ud544\uc694\ud55c \uc790\uae08\uc744 \uc870\ub2ec\ud558\uae30 \uc704\ud558\uc5ec \ucc28\uc8fc\uac00 \ubcf8\uac74 \ud22c\uc790\uc2e0\ud0c1\uc758 \uc2e0\ud0c1\uc5c5\uc790\uc758 \uc790\uaca9\uc73c\ub85c 2010\ub144 8\uc6d4 30\uc77c \uac15\uc11c\uc774\ub79c\ub4dc\ub9ac\ud14c\uc77c\uc81c\uc77c\ucc28\uc720\ud55c\ud68c\uc0ac \ub4f1\uacfc \uccb4\uacb0\ud55c \ub300\ucd9c\uc57d\uc815)\uc5d0 \ub530\ub77c \ucc28\uc8fc\uac00 \ubcf8\uac74 \ud22c\uc790\uc2e0\ud0c1\uc758 \uc2e0\ud0c1\uc5c5\uc790\uc758 \uc790\uaca9\uc73c\ub85c \ubd80\ub2f4\ud558\uace0 \uc788\ub294 \ub300\ucd9c\uc6d0\ub9ac\uae08\uc758 \ucc44\ubb34\uc758 \uc0c1\ud658(970 \uc5b5\uc6d0) \ubc0f 2001\uc544\uc6b8\ub81b \uc548\uc591\uc810\uc758 \ucde8\ub4dd\uc790\uae08 \ubc0f \uc6b4\uc601\uc790\uae08\uc758 \uc870\ub2ec(270 \uc5b5\uc6d0)", 
            "sellto": "\uc8fc\uc2dd\ud68c\uc0ac \uc6b0\ub9ac\uc740\ud589", 
            "loantype": "\ub2f4\ubcf4\ub300\ucd9c", 
            "loancls": "\ub2f4\ubcf4", 
            "addr": "\uc11c\uc6b8\uc2dc \uac15\uc11c\uad6c \ub4f1\ucd0c\ub3d9 689-2, \uacbd\uae30\ub3c4 \uc548\uc591\uc2dc \ub9cc\uc548\uad6c \uc548\uc591\ub3d9 627-287", 
            "loanratio": "540/1240", 
            "loanamt": "54,000,000,000", 
            "loanrpy": "\ub300\ucd9c\ub9cc\uae30\uc77c\uc5d0 \uc804\uc561 \uc0c1\ud658", 
            "rate": "\uace0\uc815", 
            "loanintcls": "\uc0c1\uc218", 
            "loanint": "5.39%", 
            "loanintfloat": NaN, 
            "sdaterate": "5.39%", 
            "spread": "1.86%", 
            "loanpremium": "Null", 
            "intdur": "1", 
            "laterate": "19.00%", 
            "lateratecls": "\uc0c1\uc218", 
            "sdatelaterate": "19.00%", 
            "earlypremium": "1%", 
            "earlypremiumcls": "\uc0c1\uc218", 
            "guranteelimit": "130%", 
            "guranteemax": "70,200,000,000", 
            "dscr": "\uac80\uc99d\uc870\uac74 1.2", 
            "dscrval": "1.2", 
            "intdeposit": "3", 
            "default": "\uc9c0\uae09\uae30\uc77c\uc5d0 \uc9c0\uae09\ud558\uc9c0 \uc544\ub2c8\ud55c \uacbd\uc6b0", 
            "opinion": "1/2 \uc774\uc0c1", 
            "contact": "\ub300\uc8fc1_\uc8fc\uc2dd\ud68c\uc0ac \uc6b0\ub9ac\uc740\ud589_\uc724\uc0c1\uaddc_02-2002-5278 / \ub300\uc8fc2_\uc8fc\uc2dd\ud68c\uc0ac \ud55c\uad6d\uc678\ud658\uc740\ud589 \uc2e0\ucd0c\uc9c0\uc810_\uc591\uc6d0\ud6c8_02-718-0229(751) / \ub300\uc8fc3_\ud55c\uad6d\uc0b0\uc5c5\uc740\ud589 \ud3c9\ud0dd\uc9c0\uc810_\uc11c\uacbd\uc644_031-659-0524 / \ucc44\ubb34\uc778\uc218\uc778_\uc8fc\uc2dd\ud68c\uc0ac \uc774\ub79c\ub4dc\ub9ac\ud14c\uc77c_\ud64d\uacbd\uc77c_02-2029-3429 / \ucc28\uc8fc_\uc8fc\uc2dd\ud68c\uc0ac \uad6d\ubbfc\uc740\ud589_\uc2e0\ub3d9\ud76c_02-2073-5167 / \uc9d1\ud569\ud22c\uc790\uc5c5\uc790_\ud53c\uc5d0\uc2a4\uc790\uc0b0\uc6b4\uc6a9 \uc8fc\uc2dd\ud68c\uc0ac_\uad8c\uae30\ub9cc_02-2112-4009", 
            "lender": "\uc8fc\uc2dd\ud68c\uc0ac \uad6d\ubbfc\uc740\ud589", 
            "trustee": "\uc8fc\uc2dd\ud68c\uc0ac \uad6d\ubbfc\uc740\ud589", 
            "amc": "\ud53c\uc5d0\uc2a4\uc790\uc0b0\uc6b4\uc6a9 \uc8fc\uc2dd\ud68c\uc0ac", 
            "financialinst": "Null", 
            "mm": "\uc8fc\uc2dd\ud68c\uc0ac \uc6b0\ub9ac\uc740\ud589", 
            "cashsupp": "Null", 
            "debtundwrt": "\uc8fc\uc2dd\ud68c\uc0ac \uc774\ub79c\ub4dc\ub9ac\ud14c\uc77c", 
            "builder": "Null", 
            "loanplan": "2012-02-28", 
            "loanexec": "Null", 
            "loanmatr": "\ucd5c\ucd08\ub300\ucd9c\uc2e4\ud589\uc77c\ub85c\ubd80\ud130 3\ub144", 
            "duration": "36", 
            "loanmatrymd": NaN
            },
}
```

3. /macro

| jsonCode | Description | Type |
|----------|-------------|------|
| kr1y | 국고채 1년 (전체 시계열) | [{date, year}, ...]
| kr3y | 국고채 3년 (전체 시계열) |
| kr5y | 국고채 5년 (전체 시계열) |
| ifd1y |  ifd |
|cd91d | cd |
|cp91d | cp |
|koribor3m | koribor |

```json
{
    "fsht": "macro", 
    "desc": "IDP Excel. Macro Information", 
    "last": "20220615T17:42:37", 
    "data": {
        "kr1y": [
            {"date": "20100104", "value": 3.5}, 
            {"date": "20100105", "value": 3.43}, 
            ...
            ],
        "kr3y": ...
    }
}

```