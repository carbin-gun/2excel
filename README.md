# 2excel


### core elements

 - `source` ,which file you want to convert from .
 - `target` ,the location you want the generated excel to be put into. do not provide filename entrance.the excel name will be the same with source.but with `xlsx` extension.
 - `delimiter`,delimiter for row fields.


### how to use

```
 2excel to_convert_file
```
the above command will automatically convert file in which the row fields delimiter is comma or tab ,then generate a excel file in current directory.

or 
```
2excel to_convert_file target_dir
```
you can specify the source file and the target dir where the generated excel will be put into.


if your file row fields' delimiter is not comma or tab.you can specify it by -d option.like the following .

```
2excel -f to_convert_file  -d "user_specify_delimeter"  

```
for example,you have a file in which every row is delimitered by a vertical line .
you can do with a input like  ` 2excel -f test.txt -d "|" ` .


if you do not like the target excel file in the current directory .you can specify the directory with ` -t ` option . like the following

```
2excel to_convert_file -t target_dir
```

### help

you can just input ` 2excel help ` ,you can see the command-line help doc.

```

NAME:
   2excel - convert sql export or csv file to excel with one command!

USAGE:
   2excel [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR(S):
   carbin-gun <cilendeng@gmail.com>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -f 			the file which the program will convert from
   -d 			the delimeter of different fields of the same row.the tab or comma is support by default,you can point it out or not
   -t 			the target directory the excel will be generated at,the default is current dir.
   --help, -h		show help
   --version, -v	print the version
   
```