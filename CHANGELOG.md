# 0.0.1
<p> Initial commit <p>

[Add]
- ./apis
  - idp.go: has data -> json data.
    - func ServeLanding
    - func ServeAssetWhole
    - func ServeDebtWhole
    - func ServeMacroWhole
- ./asset
  - idpChecklist.json
  - idpDept.json
  - idpMacro.json
- ./orm
  - (placeholder for database connection)
  - dbaQuery.go: makes database queries
  - dbaStruct.go: makes necessary database struct, types


[Change]

[Fix]

[Remove]


# 0.1.0
<p> Major Change in API 1) changed file tree structrue - divide file by query 2) pre-load data 3) pre-load files when initiating the server 4) add year and some classifiers to the api (as an example) <p>


[Add]
- ./apis
  - idpStrc.go
    - struct IDPDataSet 
    - struct IDPAsset
    - struct assets
    - struct IDPDebt
    - struct debts
    - struct IDPMacro
    - struct macros
    - struct macroRow
  - idpData.go
    - func MntData
    - func mntDebt
    - func mntMacro
    - func mntAsset
    
[Change]
- ./apis
  - QryAsset.go
    - func ServeAssetWhole: add parameter (strategy)
  - QryDebt.go
    - func ServeDebtWhole: add parameter (year)
  - QryMacro.go
    - func ServeMacroWhole: add parameter (year, commodity)

[Fix]

[Remove]


# 0.1.1

<p>

Changes to functions, simplify functions by splitting them.

</p>

[Add]
- annotations to all structures
- ./apis
  - meta.go: provide meta for swagger config
  - idpReqStrc.go: records all structures that gathers params inside request body
    - struct ReqRowCount
    - struct ReqIDPAsset
    - struct ReqIDPDebt
    - struct ReqIDPMacro

[Change]
- ./apis
  - QryDebt.go
    - func ServeDebtWhole
      - separate it with procDebtQry | serveDebtWhole for better maintainence
  - QryAsset.go
    - func ServeAssetWhole
      - separate it with procAssetQry | serveAssetWhole
  - QryRowCount.go
    - rename func debtRowCount -> func procRowCountQry
    - change procRowCountQry parameters to take ReqRowCount struct
  - QryMacro.go
    - func ServeMacroWhole
      - separate it with procMacroQry | ServeMacroWhole
      - simplify switch-case grammar with hashmap


[Fix]

[Remove]


# 0.2.0

<p>
insert proc*Param function. edit QryDebt.go, QryRowCount.go to fit new design.
* Delete QryRowCount.go - integrated into QryDebt.go
* Delete ReqIDPDebt Structure.
* Enable Optional Queries to QryDebt - other Macro and Asset should follow
</p>

[Add]
- ./apis
  - idp.go
    - func IsWithInSlider
    - func IsWithInChoice - enables multi conditions linked by "-"


[Change]
- ./apis
  - QryAsset.go, QryMacro.go
    - func procAssetParam - separated from serveAssetWhole
    - func procMacroParam - separated from serveMacroWhole
  - Output of QryDebt changed from []debts -> IDPDebt struct

[Fix]
- ./apis
  - QryDebt.go
    - func ServeDebtWhole -> ServeDebt
    - func procDebtQry - now supports optional queries.

[Remove]
- ./apis
  - QryDebt.go
    - func procDebtParam - integrated into procDebtQry to support optional queries
  - QryRowCoung.go - the whole thing
  - idpRespStrc.go
    - struct IDPRowCount
  - idpReqStrc.go
    - struct ReqRowCount
    - struct ReqIDPDebt - useless in processing optional queries


# 0.2.1

<p>
add endpoint for graph 1 and graph 2.
</p>

[Add]

[Change]
- ./apis
  - QryDebt.go
    - func procDebtQry - add integer forGraph -> point out whether it's for datatable, graphleft, graphright.
  - idpRespStrc.go
    - type IDPDebt
      - add DataGraphLeft, DataGraphRight Key
    - type debtsGraphLeft
    - type debtsGraphRight

[Fix]

[Remove]


# 0.2.2

<p>
add endpoint & functions for single fund search
</p>

[Add]
- ./apis
  - QrySingle.go
    - func ServeSingle
    - func procSingleQry
  - idpRespStrc.go
    - struct IDPSingle
  - idp.go
    - const TEST_URL_SINGLE, MSG_SINGLE

[Change]
- main.go
  - add names, and specify methods for /dataTable, /graphRight, /graphLeft
  - add /single endpoint

[Fix]

[Remove]

# 0.3.2

<p>
add index for IDPDebt query json to sort out single fund
</p>

[Add]
- ./asset
  - idpDebt2.json

[Change]
- ./apis
  - idpData.go
    - func mntDebt - read from idpDebt2.json
  - idpRespStrc.go
    - struct debts - add field UniqueIndex
  - QrySingle.go
    - func procSingleQry
      - var uniqIdx - gets from idx
[Fix]

[Remove]

# 0.3.3
<p>
add PageCount for IDPDebt query
</p>

[Add]
- ./apis
  - idp.go
    - SINGLE_PAGE_INFO
  - QryDebt.go
    - func divDebtArray

[Change]
- ./apis
  - idp.go - change TEST URL for Debt
  - QryDebt.go
    -  func procDebtQry - add new var pgn. return divided Debt Array
      - var pgn takes value from 1 ~ ...

[Fix]

[Remove]

# 0.3.4
<p>
add path to image
</p>

[Add]

[Change]
- ./apis
  - idpData.go
    - func mntDebt - use idpDebt3.json
  - idpRespStrc.go
    - struct debts - add Image field

[Fix]

[Remove]
  

# 0.3.5
<p>
add path to file location
</p>

[Add]

[Change]
- ./apis
  - idpRespStrc.go
    - struct debts - add file field

[Fix]

[Remove]

# 0.3.6
<p>
add endpoint for api parameters. prepare logging middleware
</p>

[Add]
- ./apis
  - QryModel.go
    - func ServeModelInfo
    - func ServeModelCoef
  - idp.go
    - const DATA_PANIC_MODEL, DATA_ERR_MODEL, MSG_MODEL
  - idpData.go
    - func mntModelInfo
    - func mntModelCoef
- main.go - add endpoint
  - /api/v1/model/info
  - /api/v1/model/coef

[Change]
- ./apis
  - QryDebt.go
    - map[string]string prepare for middleware

[Fix]

[Remove]

# 0.3.7
<p>
fix readme - version info goes only on CHANGELOG.md. Other information such as example page will be shown on README. fix row count. add asset name to graphleft. 
</p>

[Add]
- main.go
  - func init - initial ascii art

[Change]
- ./apis
  - QryDebt.go
    - func procDebtQry - returns fixed row counts per query.

[Fix]

[Remove]

# 0.3.8
<p>
add info to graphLeft
</p>

[Add]

[Change]

[Fix]
- ./apis
  - idpRespStrc.go
    - struct debtsGraphLeft - add AssetName
  - QryDeb.go
    - func procDebtQry - adjust code to fit new field AssetName
    

[Remove]

# 0.3.9
<p>
add sorting mechanism. change model api
</p>

[Add]
- ./apis
  - QryDebt.go
    - func divDebtArray - sortByKey inserted
  - QryDebtSort.go
    - script for custom sorting functions
      - custom types
      - custom Len, Less, Swap for QuickSort Prep
    - func sortByKey
  - QryModel.go
    - add http headers to API

[Change]
- ./apis
  - QryDebt.go
    - func procDebtQry - shorten var declaration, change cndAssetType.. declaration to ":="

[Fix]

[Remove]

# 0.3.10
<p>
fix sortByKey. Problem::if sortKey is not present, api panics
</p>

[Add]

[Change]

[Fix]
- ./apis
  - QryDebtSort.go
    - func sortByKey - add return d to default.

[Remove]

# 0.3.11
<p>
model prediction calculation will be served from the server side. add servemodelcalc function and endpoint /pred
</p>

[Add]
- ./apis
  - QryModel.go
    - func findRecentMacro
    - func findDataPointMap
    - func genParameterMap
    - func procModelQuery
    - func calcInterest
    - func ServeModelCalc
  - idpRespStrc.go
    - func IDPModelCoef
    - func coefficient
    - func ModelPrediction

[Change]
- ./apis
  - QryModel.go 
    - func ServeModelCoef - throws JSON -> []byte encoded file.
  - idpData.go
    - func mntModelCoef - unmarshal JSON file with struct 

[Fix]

[Remove]

# 0.3.12
<p>
change model prediction calculation output data type from single point to array of point.
</p>

[Add]

[Change]
- ./apis
  - idpRespStrc.go
    - struct ModelPrediction - change field's value's data type from float64 to []float64
  - QryModel.go
    - func ServeModelCalc - change accordingly
  
[Fix]

[Remove]


# 0.3.13
<p>
Fix debt query search. write annotation for function. 
</p>

[Add]
- (annotations)

[Change]

[Fix]
- ./apis
  - QryDebt.go
    - now support investType query (it)

[Remove]

# 0.3.14
<p>
change model prediction from point to band. prep for docker. todo compile with external file
</p>

[Add]
- ./apis
  - idpRespStrc.go
    - struct IDPModelInfo
    - struct modelmeta
- ./Dockerfile  - prep for docker deploy

[Change]  
- ./apis
  - idpData.go
    - func mntModelInfo - unmarshal it into struct
    - func mntModelCoef - change filesource from idpCoef.json to idpCoef2.json
  - idpV1Strc.go
    - struct IDPDataSet - change accordingly
  - QryModel.go
    - func ServeModelCalc - add band to sendpacket point predictions

[Fix]

[Remove]


# 0.4.1
<p>
Fix band. Add Macro. Query macro by dateFrom and dateUntil
</p>

[Add]
- ./apis
  - QryModel.go
    - func parseModelInfo
  - QryMacro.go
    - func ServeMacro
    - func procMacroQuery
    - func procByAsset
  - idp.go
    - func IsWithInDate
- main.go
  - func routeDebt
  - func routeModel
  - func routeMacro
  - func routeLanding

[Change]
- main.go
  - func main - separate it
    - routeLanding
    - routeDebt
    - routeModel
    - routeMacro
    - routes moved to routing.go (see [Add])

[Fix]
- ./apis
  - QryModel.go
    - func ServeModelCalc - fix band
    - func findRecentMacro - BUG FIX - loop now updates maxDate also.

[Remove]
- main.go
  - remove /api/v1/macro endpoint
- ./apis
  - QryMacro.go
    - func serveMacroQry
    - func serveMacroWhole

# 0.4.2
<p>
deployable main file. Macro New Data Post system
</p>

[Add]
- README.md
  - information regarding building docker
- ./apis
  - QryMacro.go
    - func UpdateMacro
    - func procMacroUpdate
  - idpRespStrc.go
    - struct newMacroPost
- main.go
  - func routeMacro - new endpoint /update

[Change]
- ./apis
  - func ServeMacro - now read file on request

[Fix]
- main.go
  - Hosting addr: 127.0.0.1:8080 -> 0.0.0.0:8080. Deployed version.
- Dockerfile
  - RUN go build
  - CMD ["/app/IGISBackEnd"]

[Remove]


# 0.4.4
<p>
Skip 0.4.3 to match docker deploy version. Model pred endpoint returns (low, high, base)
</p>

[Add]

[Change]
- ./apis
  - QryModel.go
    - func ServeModelCalc - gives []float64{bandLow, bandHigh, bandMidpoint(base)}

[Fix]

[Remove]

# 0.5.0 - Branch origin/0.5
<p>
Move macro data into redis-cloud instance
</p>

[Add]
- ./orm
  - dbaStruct.go
    - func processKey
    - func CreateDatabaseObject
    - func RedisJSONGet
    - func RedisJSONSet

[Change]
- ./apis
  - QryMacro.go
    - func ServeMacro - now gets data from `func mntMacroRedis` which queries from redis cloud
    - func UpdateMacro - now supports updating database cloud using `func RedisJSONSet`
  - idpData.go
    - func parseResult - parse raw string result by Redis
    - func mntMacroRedis - queries and gets data from RedisCloud
  - main.go - change accordingly.

[Fix]


[Remove]
- ./orm
  - delete dbaQuery.go

# 0.5.1 - Branch origin/0.5
<p>
Use go generics to simplify codes
</p>

[Add]
- ./orm
  - ormConn.go
    - func processKey
    - func Conn
    - func DataBaseConfig
  - ormReJson.go
    - use go generics to get arbitrary types of container and parse it into Database. 
    - func JSONGet[T any]
    - func JSONSet[T any]

[Change]
- ./apis
  - idpData.go
    - func mntMacroRedis - use orm.JSONGet to simplify things.

[Fix]

[Remove]
- ./orm
  - dbaStruct.go - removed due to go generics function