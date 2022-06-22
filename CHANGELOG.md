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