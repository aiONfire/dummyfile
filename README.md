# dummyfile

This is an application to automatically generate datafiles without an extension for testing purposes. The application can be used for two different purposes.

## Generate some big files
The first operation mode for `big` files is optimized to create a few big files. The generation of the files is controlled by two different parameters. The `-fileCount <n>` to specify how many files to generate in total and the `-sizeStart <n>`parameter to specify the size of the first file in MB (Mega Byte). The second file has double the size of the first (initial file), the third file has double the file size of the provious file and so on.

> eg.: 10 MB -> 20 MB -> 40 MB -> 80 MB -> ....

dummyfile big -h

### Parameter
| Parameter | Default | Description | Example |
| --------- | ------- | ----------- | ------- |
| -fileCount | 8 | Count of files to generate | dummyfile big -n |
| -sizeStart | 10 |  1st file size in MB. further files will be doubled the size of the priviouse file |

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

## generate many tiny files

dummyfile tiny -h 



=================================================================
Generate big dummy files
=================================================================

dummyfile big -h

dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
Usage of big:
  -fileCount string
        Count of files to generate - file size is doubled every time (default "8")
  -sizeStart string
        1st file size in MB (default "10")

=================================================================
SAMPLES.....

dummyfile big -sizeStart 20 -fileCount 9
dummyfile big -sizeStart 5 -fileCount 1


=================================================================
Generate tiny dummy files
=================================================================

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


=================================================================
SAMPLES.....

dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 10
dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 100
dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
dummyfile tiny -fileCount 10
dummyfile tiny



=================================================================
FULL EXAMPLES .....


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


C:\Users\LaurenzWagner\go\dummyfile>dummyfile big -sizeStart 5 -fileCount 1
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
bigFileSizeStart: 5, bigFileCount: 1
5 MB written to bigfile1
---------------------------- DONE --------------------------------

C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -h
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
Usage of tiny:
  -fileCount string
        Count of files to generate (default "5000")
  -sizeMax string
        Maximum file size in kB (default "40")
  -sizeMin string
        dummyfile tiny <40> (Minimum file size in kB) (default "40")

C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 10
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2
0 MB written to tinyfile3
0 MB written to tinyfile4
0 MB written to tinyfile5
0 MB written to tinyfile6
0 MB written to tinyfile7
0 MB written to tinyfile8
0 MB written to tinyfile9
0 MB written to tinyfile10
---------------------------- DONE --------------------------------

C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 10
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2
0 MB written to tinyfile3
0 MB written to tinyfile4
0 MB written to tinyfile5
0 MB written to tinyfile6
0 MB written to tinyfile7
0 MB written to tinyfile8
0 MB written to tinyfile9
0 MB written to tinyfile10
---------------------------- DONE --------------------------------

C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -sizeMin 5 -sizeMax 5 -fileCount 100
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2

....

0 MB written to tinyfile99
0 MB written to tinyfile100
---------------------------- DONE --------------------------------



C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2
0 MB written to tinyfile3
0 MB written to tinyfile4
0 MB written to tinyfile5
0 MB written to tinyfile6
0 MB written to tinyfile7
0 MB written to tinyfile8
0 MB written to tinyfile9
0 MB written to tinyfile10
---------------------------- DONE --------------------------------



C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -sizeMin 40 -sizeMax 100 -fileCount 10
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2
0 MB written to tinyfile3
0 MB written to tinyfile4
0 MB written to tinyfile5
0 MB written to tinyfile6
0 MB written to tinyfile7
0 MB written to tinyfile8
0 MB written to tinyfile9
0 MB written to tinyfile10
---------------------------- DONE --------------------------------



C:\Users\LaurenzWagner\go\dummyfile>dummyfile tiny -fileCount 10
dummyfile - copyright (c) 2021 by Laurenz Wagner
------------------------------------------------------------------
0 MB written to tinyfile1
0 MB written to tinyfile2
0 MB written to tinyfile3
0 MB written to tinyfile4
0 MB written to tinyfile5
0 MB written to tinyfile6
0 MB written to tinyfile7
0 MB written to tinyfile8
0 MB written to tinyfile9
0 MB written to tinyfile10
---------------------------- DONE --------------------------------



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

