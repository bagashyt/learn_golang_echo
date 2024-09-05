# Insert 1 milion data from CSV file to Database Server

Insert 1 milion data from CSV file to Database Server using Worker Pool, Database Connection Pool and Failover.

## DB Preparation

Download csv data

http://downloads.majestic.com/majestic_million.csv

Make a Mysql DB

```sh
CREATE DATABASE IF NOT EXISTS test;
USE test;
CREATE TABLE IF NOT EXISTS domain (
    GlobalRank int,
    TldRank int,
    Domain varchar(255),
    TLD varchar(255),
    RefSubNets int,
    RefIPs int,
    IDN_Domain varchar(255),
    IDN_TLD varchar(255),
    PrevGlobalRank int,
    PrevTldRank int,
    PrevRefSubNets int,
    PrevRefIPs int
);
```
## How to run

run program on directory
```sh
go run main.go
```
