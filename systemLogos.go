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

const NIXOS = `
          ::::.    ':::::     ::::'
          ':::::    ':::::.  ::::'
            :::::     '::::.:::::
      .......:::::..... ::::::::
     ::::::::::::::::::. ::::::    ::::.
    ::::::::::::::::::::: :::::.  .::::'
           .....           ::::' :::::'
          :::::            '::' :::::'
 ........:::::               ' :::::::::::.
:::::::::::::                 :::::::::::::
 ::::::::::: ..              :::::
     .::::: .:::            :::::
    .:::::  :::::          '''''    .....
    :::::   ':::::.  ......:::::::::::::'
     :::     ::::::. ':::::::::::::::::'
            .:::::::: '::::::::::
           .::::''::::.     '::::.
          .::::'   ::::.     '::::.
         .::::      ::::      '::::.`

const MX = `
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNMMMMMMMMM
MMMMMMMMMMNs..yMMMMMMMMMMMMMm: +NMMMMMMM
MMMMMMMMMN+    :mMMMMMMMMMNo' -dMMMMMMMM
MMMMMMMMMMMs.   'oNMMMMMMh- 'sNMMMMMMMMM
MMMMMMMMMMMMN/    -hMMMN+  :dMMMMMMMMMMM
MMMMMMMMMMMMMMh-    +ms. .sMMMMMMMMMMMMM
MMMMMMMMMMMMMMMN+'   '  +NMMMMMMMMMMMMMM
MMMMMMMMMMMMMMNMMd:    .dMMMMMMMMMMMMMMM
MMMMMMMMMMMMm/-hMd-     'sNMMMMMMMMMMMMM
MMMMMMMMMMNo   -' :h/    -dMMMMMMMMMMMM
MMMMMMMMMd:       /NMMh-   '+NMMMMMMMMMM
MMMMMMMNo':mMMN+'   '-hMMMMMMMM
MMMMMMh.            'oNMMd:    '/mMMMMMM
MMMMm/                -hMd-      'sNMMMM
MMNs'                   -          :dMMM
Mm:                                 'oMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
`

const ELEMENTARY = `
         eeeeeeeeeeeeeeeee
      eeeeeeeeeeeeeeeeeeeeeee
    eeeee  eeeeeeeeeeee   eeeee
  eeee   eeeee       eee     eeee
 eeee   eeee          eee     eeee
eee    eee            eee       eee
eee   eee            eee        eee
ee    eee           eeee       eeee
ee    eee         eeeee      eeeeee
ee    eee       eeeee      eeeee ee
eee   eeee   eeeeee      eeeee  eee
eee    eeeeeeeeee     eeeeee    eee
 eeeeeeeeeeeeeeeeeeeeeeee    eeeee
  eeeeeeee eeeeeeeeeeee      eeee
    eeeee                 eeeee
      eeeeeee         eeeeeee
         eeeeeeeeeeeeeeeee
`

const UBUNTU = `
            .-/+oossssoo+/-.
        ':+ssssssssssssssssss+:'
      -+ssssssssssssssssssyyssss+-
    .ossssssssssssssssssdMMMNysssso.
   /ssssssssssshdmmNNmmyNMMMMhssssss/
  +ssssssssshmydMMMMMMMNddddyssssssss+
 /sssssssshNMMMyhhyyyyhmNMMMNhssssssss/
.ssssssssdMMMNhsssssssssshNMMMdssssssss.
+sssshhhyNMMNyssssssssssssyNMMMysssssss+
ossyNMMMNyMMhsssssssssssssshmmmhssssssso
ossyNMMMNyMMhsssssssssssssshmmmhssssssso
+sssshhhyNMMNyssssssssssssyNMMMysssssss+
.ssssssssdMMMNhsssssssssshNMMMdssssssss.
 /sssssssshNMMMyhhyyyyhdNMMMNhssssssss/
  +sssssssssdmydMMMMMMMMddddyssssssss+
   /ssssssssssshdmNNNNmyNMMMMhssssss/
    .ossssssssssssssssssdMMMNysssso.
      -+sssssssssssssssssyyyssss+-
        ':+ssssssssssssssssss+:'
            .-/+oossssoo+/-.`

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
	} else if strings.Contains(os, "nixos") {
		return NIXOS
	} else if strings.Contains(os, "mx") {
		return MX
	} else if strings.Contains(os, "elementary") {
		return ELEMENTARY
	} else if strings.Contains(os, "ubuntu") {
		return UBUNTU
	} else {
		return `¯\_(ツ)_/¯`
	}
}