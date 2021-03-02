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

