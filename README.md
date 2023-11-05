# Go BRL rate get
 
This is a small project to fetch data from BCB (Brazilian Central Bank) to get
all exchanges rates between a currency and BRL

## To Install

```bash
go install github.com/uelei/brl-rate-get@latest

# OR

go install github.com/uelei/brl-rate-get@v0.0.20
```

## Help
 
```bash 
brl-rate-get -help
Usage of brl-rate-get:
  -get
        To get yesterday closing currency
  -range
        To get a range of closing currency rates
```

## How to use the -get command

To get last rate of a currency

```bash 
brl-rate-get -get usd
2023/11/05 09:17:38 Found 5 records of price..

____________________________________
|#######====================#######|
|#  BRL          EUR              #|
|#         ====                   #|
|#               Buy:   5.2474    #|
|#  1,00         Sell:  5.2500    #|
|#                                #|
|##==============================##|
|    Rate Date   2023-11-03        |
 ----------------------------------
```

## How to use the -range command

To get a range of data from April first until April 10 use

```bash 
brl-rate-get -range usd 2022-04-01 2022-04-10

2023/11/04 21:28:58 Found 30 records of price..
2023/11/04 21:28:58 Found 6 records of fechamento price..
2023/11/04 21:28:58 Writing a file.
File saved  usd_brl_2022-04-01-2022-04-10.csv

```
