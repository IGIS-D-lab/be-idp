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

Create swagger API document
- swagger logs instance by analyzing the formatted annotation
- add annotation to each handleFuncs

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