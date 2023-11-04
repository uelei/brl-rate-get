# Go BRL rate get
 
This is a small project to fetch data from BCB (Brazilian Central Bank) to get
all exchanges rates between a currency and BRL
 
## Help
 
```bash 
Usage: brl-rate-get <command>
```

## To use

To get yesterday close of a currency

```bash 
brl-rate-get -get usd

```

## To use

To get a range of data from April first until April 10 use

```bash 
brl-rate-get -range usd 2022-04-01 2022-04-10

```
