# Go Get Dollar
 
This is a small project to fetch data from BCB (Brazilian Central Bank) to get all exchanges rates between USD and BRL 
 
 
## Help  
 
```bash 
Usage: go-dollar-get <command>
 
Flags: 
  -h, --help     Show context-sensitive help.
      --debug    Enable debug mode.
 
Commands: 
  range <start_date> <end_date>
    Range of dates to query
 
Run "go-dollar-get <command> --help" for more information on a command.
  
``` 

## To use

To get a range of data from April first until April 10 use

```bash 
go-dollar-get range 2022-04-01 2022-04-10 

```
