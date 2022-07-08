# IGIS WebApp BackEnd Service

```
go mod init IGISBackEnd
go get -u "github.com/gorilla/mux"
go get -u "github.com/go-sql-driver/mysql"
```

## API version: Current: version 0
<p>

배포 전 0.x.x 버전 - CHANGELOG.md 참고

Main router:
  1. api/v1: single, asset, macro 지원

Subrouter포인트: 
  1. api/v1/debt: dataTable, graphLeft, graphRight 등 지원
  2. api/v1/model: coef, info 등 모델 정보 지원

특징:
- 데이터베이스가 형성되지 않았기 때문에, csv를 그대로 json으로 쏴줄 것임. 
  - 서버 시작시 json을 미리 로드 하고, 요청시마다 가공해서 쏴줌
  - version 1에서는 달라질 것.
- Checklist ::, Macro ::, Debt ::, Single ::, Model :: 에서 요청사항처리 및 에러 확인 가능


### API 정보
<p>
http://localhost:8080에서 실행. GitBook에서 API 요청 확인 가능.
</p>


</p>

실행 후 Test 주소 모음

|Subroute| Content | Testing URL |
|-|--|--|
|main| Single Fund Query | http://localhost:8080/api/v1/single?fc=112001&idx=1 |
|main| Asset Query | http://localhost:8080/api/v1/asset?strat=Core |
|macro| Macro Query | http://localhost:8080/api/v1/macro/dataTable?dateFrom=20220101&dateUntil=20220801 |
|debt| Debt Query (Table) | http://localhost:8080/api/v1/debt/dataTable?at=%EC%98%A4%ED%94%BC%EC%8A%A4-%ED%98%B8%ED%85%94&seniorstr=%EC%84%A0&loancls=%EB%B8%8C%EB%A6%BF%EC%A7%80&debtFrom=1&debtUntil=1e13&pageCount=1  |
|debt| Debt Query (graph1) | http://localhost:8080/api/v1/debt/graphLeft?at=%EC%98%A4%ED%94%BC%EC%8A%A4-%ED%98%B8%ED%85%94&seniorstr=%EC%84%A0&loancls=%EB%B8%8C%EB%A6%BF%EC%A7%80&debtFrom=1&debtUntil=1e13&pageCount=1  |
|debt| Debt Query (graph2) | http://localhost:8080/api/v1/debt/graphRight?at=%EC%98%A4%ED%94%BC%EC%8A%A4-%ED%98%B8%ED%85%94&seniorstr=%EC%84%A0&loancls=%EB%B8%8C%EB%A6%BF%EC%A7%80&debtFrom=1&debtUntil=1e13&pageCount=1  |
|model| Model Info | http://localhost:8080/api/v1/model/info |
|model| Model Coefficient | http://localhost:8080/api/v1/model/coef |
|model| Model Prediction | http://localhost:8080/api/v1/model/pred?seniorstr=%EC%84%A0&loancls=%EB%B8%8C%EB%A6%BF%EC%A7%80 |