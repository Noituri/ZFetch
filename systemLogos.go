package main

import "strings"

const BEDROCK = `
--------------------------------------
--------------------------------------
--------------------------------------  
---\\\\\\\\\\\\-----------------------   
----\\\      \\\----------------------    
-----\\\      \\\---------------------   
------\\\      \\\\\\\\\\\\\\\\\------    
-------\\\                    \\\-----    
--------\\\                    \\\----   
---------\\\        ______      \\\---    
----------\\\                   ///---    
-----------\\\                 ///----    
------------\\\               ///-----    
-------------\\\////////////////------    
--------------------------------------   
--------------------------------------    
--------------------------------------`

//TODO colorize
const GENTOO = `
         -/oyddmdhs+:. 
     -odNMMMMMMMMNNmhy+-
-yNMMMMMMMMMMMNNNmmdhy+-
omMMMMMMMMMMMMNmdmmmmddhhy/
omMMMMMMMMMMMNhhyyyohmdddhhhdo    
.ydMMMMMMMMMMdhs++so/smdddhhhhdm+ 
oyhdmNMMMMMMMNdyooydmddddhhhhyhNd.
:oyhhdNNMMMMMMMNNNmmdddhhhhhyymMh   
.:+sydNMMMMMNNNmmmdddhhhhhhmMmy   
/mMMMMMMNNNmmmdddhhhhhmMNhs:  
oNMMMMMMMNNNmmmddddhhdmMNhs+  
sNMMMMMMMMNNNmmmdddddmNMmhs/. 
 /NMMMMMMMMNNNNmmmdddmNMNdso: 
+MMMMMMMNNNNNmmmmdmNMNdso/-   
yMMNNNNNNNmmmmmNNMmhs+/-      
/hMMNNNNNNNNMNdhs++/-         
/ohdmmddhys+++/:.             
-//////:--.  `

const VOID = `
                __.;=====;.__
            _.=+==++=++=+=+===;.
             -=+++=+===+=+=+++++=_
         .     -=:;=''' '--==+=++==.
        _vi,    '           --+=++++:
      .uvnvi.       _._        -==+==+.
     .vvnvnI'    .;==|==;.      :|=||=|.
+QmQQmpvvnv; _yYsyQQWUUQQQm #QmQ#:QQQWUV$QQmL
-QQWQWpvvowZ?.wQQQE==<QWWQ/QWQW.QQWW(: jQWQE
-$QQQQmmU'  jQQQ@+=<QWQQ)mQQQ.mQQQC+;jWQQ@'
   -$WQ8YnI:   QWQQwgQQWV'mWQQ.jQWQQgyyWW@!
     -1vvnvv.     '~+++'        ++|+++
      +vnvnnv,                 '-|===
      +vnvnvns.           .      :=-
       -Invnvvnsi..___..=sv=.     '
          +Invnvnvnnnnnnnnvvnn;.
           ~|Invnvnvvnvvvnnv}+'
                 -~|{*l}*|~`

const ARCH = `
                   -'
                  .o+'
                 'ooo/
                '+oooo:
               '+oooooo:
               -+oooooo+:
             '/:-:++oooo+:
            '/++++/+++++++:
           '/++++++++++++++:
          '/+++ooooooooooooo/'
         ./ooosssso++osssssso+'
        .oossssso-''''/ossssss+'
       -osssssso.      :ssssssso.
      :osssssss/        osssso+++.
     /ossssssss/        +ssssooo/-
   '/ossssso+/:-        -:/+osssso+-
   '+sso+:-'               '.-/+oso:
   '++:.                        '-/+/
   .'                               '/`

const DEBIAN = `
       _,met$$$$$gg.
    ,g$$$$$$$$$$$$$$$P.
  ,g$$P"     """Y$$.".
 ,$$P'              '$$$.
',$$P       ,ggs.     '$$b:
'd$$'     ,$P"'   .    $$$
$$P      d$'     ,    $$P
$$:      $$.   -    ,d$$'
 $$;      Y$b._   _,d$P'
 Y$$.    '.'"Y$$$$P"'
 '$$b      "-.__
  'Y$$
   'Y$$.
    '$$b.
     'Y$$b.
       '"Y$b._
          '"""`

func GetASCII(os string) string {
	os = strings.TrimSpace(strings.ToLower(os))

	if strings.Contains(os, "bedrock") {
		return BEDROCK
	} else if strings.Contains(os, "gentoo") {
		return GENTOO
	} else if strings.Contains(os, "void") {
		return VOID
	} else if strings.Contains(os, "arch") {
		return ARCH
	} else if strings.Contains(os, "debian") {
		return DEBIAN
	} else {
		return `¯\_(ツ)_/¯`
	}
}