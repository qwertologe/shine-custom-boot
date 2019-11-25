#!/bin/bash

if [ -r update.zip ]; then
 test -d android_system_core/ || git clone https://github.com/CyanogenMod/android_system_core.git
 rm -rf temp[1-9]
 mkdir temp1
 unzip -d temp1 update.zip boot.img
 mkdir temp2
 android_system_core/mkbootimg/unpackbootimg -i temp1/boot.img -o temp2
 mkdir temp3
 cd temp3
 gzip -dc ../temp2/boot.img-ramdisk.gz | cpio -imd
 sed -i -e 's/ro.secure=.*/ro.secure=0/' \
 -e 's/ro.debuggable=.*/ro.debuggable=1/' \
 -e 's/\(persist.sys.usb.config=.*\)/\1,adb/' default.prop
 find . ! -name . | LC_ALL=C sort | cpio -o -H newc -R root:root | gzip > ../temp2/boot.img-ramdisk.gz
 cd ../temp2
 ../android_system_core/mkbootimg/mkbootimg --kernel ./boot.img-zImage \
 --kernel_offset "$(echo $((16#$(< boot.img-kernel_offset))))" \
 --ramdisk ./boot.img-ramdisk.gz \
 --second ./boot.img-second \
 --second_offset "$(echo $((16#$(< boot.img-second_offset))))" \
 --cmdline "$(< boot.img-cmdline)" \
 --base "$(echo $((16#$(< boot.img-base))))" \
 --pagesize "$(< boot.img-pagesize)" \
 --dt boot.img-dt \
 --ramdisk_offset "$(echo $((16#$(< boot.img-ramdisk_offset))))" \
 --tags_offset "$(echo $((16#$(< boot.img-tags_offset))))" \
 --board "$(< boot.img-name)" \
 --output ../boot.img
 echo
 cd ..
 echo 'Genutzte Hilfsutilities sind noch im Ordner' a*
 cd temp1
 echo 'Hinweis: Bei mir hatten alle boot.img aus update.zip das Datum 24.7.2014'
 echo "Altes boot.img: $(md5sum boot.img|cut -f1 -d' ') (in update.zip)"
 ls -al boot.img
 cd ..
 echo 'Hinweis: Trotz Modifikation kann die Größe gleich bleiben'
 echo "Neues boot.img: $(md5sum boot.img|cut -f1 -d' ')"
 ls -al boot.img
 rm -rf temp[1-9]
 echo 'ACHTUNG: Ergebnis selber prüfen - keine Verantwortung von meiner Seite!'
else
 echo 'Benoetigt update.zip im (ansonsten leeren) Arbeitsordner - Abbruch'
fi
