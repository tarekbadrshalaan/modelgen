# modelgen

[![CircleCI Status](https://img.shields.io/github/release/tarekbadrshalaan/modelgen.svg)](https://github.com/tarekbadrshalaan/modelgen/releases)
[![GoDoc](https://godoc.org/github.com/tarekbadrshalaan/modelgen?status.svg)](https://godoc.org/github.com/tarekbadrshalaan/modelgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/tarekbadrshalaan/modelgen)](https://goreportcard.com/report/github.com/tarekbadrshalaan/modelgen)
[![Build Status](https://travis-ci.org/tarekbadrshalaan/modelgen.svg?branch=master)](https://travis-ci.org/tarekbadrshalaan/modelgen)
[![CircleCI Status](https://circleci.com/gh/tarekbadrshalaan/modelgen.svg?style=shield)](https://circleci.com/gh/tarekbadrshalaan/modelgen)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)




Application to create (start app) go webservice with 3-Tier Architecture.

The Generated Application inculeds :- 

- Go mod
- Configeration file
- DAL,BLL,DTO,API and API_Tests for every Database Table 
- Compatalbe with `mysql` `postgres` `mssql` `sqlite` `oracle`
- Using [![GORM](https://github.com/jinzhu/gorm)](https://github.com/jinzhu/gorm) as ORM

## Installation

```
$ go get -u github.com/tarekbadrshalaan/modelgen
$ vi config.json
  
  {
    "Module"                :   "{{github.com/packagepath}}",
    "DBConnectionString"    :   "{{Database ConnectionString}}",
    "DBEngine"              :   "{{Database Engine}}",
    "WebAddress"            :   "0.0.0.0",
    "WebPort"               :   7070
  }

$ modelgen 
$ cd Application/
$ go build .
  go: finding github.com/tarekbadrshalaan/goStuff/configuration latest
  go: finding github.com/jinzhu/inflection latest
$ ./Application
```

## Example 
- postgres Database
http://www.postgresqltutorial.com/postgresql-sample-database/
to restor the database use : 
```
pg_restore -U postgres -d dvdrental ~/dvdrental.tar

$ vi config.json
  {
    "Module"                :   "github.com/Application",
    "DBConnectionString"    :   "{{Database ConnectionString}}",
    "DBEngine"              :   "postgres",
    "WebAddress"            :   "0.0.0.0",
    "WebPort"               :   7070
  }
  
$ modelgen 
$ cd Application/
  ├──> Application
    ├──> config.json
    ├──> go.mod 
    ├──> main.go
    db
    │	├──> database.go
    api
    │	├──> ActorAPI.go
    │ ...
    bll
    │	├──> ActorBLL.go
    │ ...
    dal
    │	├──> ActorDAL.go
    │ ...
    dto
    │	├──> ActorDTO.go
    │ ...
    test
    │	├──> Actor_test.go
    │ ...
    │	├──> test.json
    ─────────────────────────────

```


## standar
- in [![standar](https://github.com/tarekbadrshalaan/modelgen/tree/master/standard)](https://github.com/tarekbadrshalaan/modelgen/tree/master/standard) you can find the expected result of the generator. 

## Contributing

PRs accepted.


## License
[![License: MIT](https://img.shields.io/badge/License-MIT-ff69b4.svg)](https://opensource.org/licenses/MIT)
