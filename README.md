# dummyfile

This is an application to automatically generate datafiles without an extension for testing purposes. The application can be used for two different purposes.

## Generate some big files
The first operation mode for `big` files is optimized to create a few big files. The generation of the files is controlled by two different parameters. The `-fileCount <n>` to specify how many files to generate in total and the `-sizeStart <n>`parameter to specify the size of the first file in MB (Mega Byte). The second file has double the size of the first (initial file), the third file has double the file size of the provious file and so on.

> eg.: 10 MB -> 20 MB -> 40 MB -> 80 MB -> ....

### Parameter
| Parameter | Default | Description | Example |
| --------- | ------- | ----------- | ------- |
| -fileCount | 8 | Count of files to generate | dummyfile big -fileCount 5 -sizeStart 100 |
| -sizeStart | 10 | Size of firts file in MB. Further files will be doubled the size of the priviouse file | dummyfile big -sizeStart 1 |

### Example

#### Show the help
```bash

dummyfile big -h

dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
Usage of big:
  -fileCount string
        Count of files to generate - file size is doubled every time (default "8")
  -sizeStart string
        1st file size in MB (default "10")
```

#### Nine files starting with 20 MB (first file) 
```bash
dummyfile big -sizeStart 20 -fileCount 9
```
#### One file with 5 MB size
```bash
dummyfile big -sizeStart 5 -fileCount 1
```
#### Nine files starting with 20 MB (first file) - with full output
``` bash
C:\Users\LaurenzWagner\go\dummyfile>dummyfile big -sizeStart 20 -fileCount 9
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
bigFileSizeStart: 20, bigFileCount: 9
20 MB written to bigfile1
40 MB written to bigfile2
80 MB written to bigfile3
160 MB written to bigfile4
320 MB written to bigfile5
640 MB written to bigfile6
1280 MB written to bigfile7
2560 MB written to bigfile8
5120 MB written to bigfile9
---------------------------- DONE --------------------------------
```
#### further usage examples ...
```bash
dummyfile big -sizeStart 20 -fileCount 9
dummyfile big -sizeStart 5 -fileCount 1
dummyfile big
```

## Generate many tiny files
The second operation mode for `tiny` files is optimized to create a a lot of tiny files. The generation of the files is controlled by three different parameters. 
The `-fileCount <n>` to specify how many files to generate in total, the `-sizeMin <n>` to specify the minimum file size and `-sizeMax <n>`to specify the maximum file size. If the minimum file size `-sizeMin <n>` is identical with the maximum file size `-sizeMax <n>` then all files will have an identical file size. If the maximum file size `-sizeMax <n>` is bigger then the minmum file size `-sizeMin <n>` then all generated files will have a random file size between the value of `-sizeMin <n>` and `-sizeMax <n>`.


### Parameter
| Parameter | Default | Description | Example |
| --------- | ------- | ----------- | ------- |
| -fileCount | 5000 | Count of files to generate | dummyfile tiny -fileCount 500 -sizeMin 1 -sizeMax 4 |
| -sizeMin | 40 |  Minimum file size in kB. | dummyfile tiny -sizeStart 1 |
| -sizeMax | 40 |  Maximum file size in kB. | dummyfile tiny -sizeStart 1 |


### Example

#### Show the help
```bash

dummyfile tiny -h

dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
Usage of tiny:
  -fileCount string
        Count of files to generate (default "5000")
  -sizeMax string
        Maximum file size in kB (default "40")
  -sizeMin string
        dummyfile tiny <40> (Minimum file size in kB) (default "40")
```

#### 100 files with 5 kB size each 
```bash
dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 100
```

#### 10 files with as size between 40 kB and 100 kB 
```bash
dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
```

#### Using the deafault values to generate 5000 files with 40 kB each.
```bash
C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2
0 MB written to tinyfile3

....

0 MB written to tinyfile4998
0 MB written to tinyfile4999
0 MB written to tinyfile5000
---------------------------- DONE --------------------------------
```

#### further usage examples ...
```bash
dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 10
dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 100
dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
dummyfile tiny -fileCount 10
dummyfile tiny
```
