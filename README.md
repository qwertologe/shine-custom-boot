# shine-custom-boot
custom boot.img fuer E-Book Reader erstellen

**Das Ausführen der Inhalte dieser Anleitung geschieht ausdrücklich auf eigene Verantwortung.
Ich gebe keine Garantie dafür, dass diese Anleitung zum Erfolg führt und bin nicht
verantwortlich für Defekte welche durch die Anwendung dieser Anleitung entstehen.**

## Custom boot.img erstellen

* Ein beliebiges (also zur eigenen Firmware passendes) update.zip für den E-Book Reader downloaden, in einen leeren Ordner kopieren und in diesen Ordner wechseln (Terminal)
* Das Script selbst benötigt zumindest das Kommandozeilen Utility "git" und eine Internetverbindung zum Download von https://github.com/CyanogenMod/android_system_core - alternativ kann man es natürlich manuell downloaden und in den erstellten Ordner entpacken
* Jetzt einfach das Script starten. Danach hat man ein neues boot.img im gleichen Ordner.

## Weitere Schritte

Bitte die Ausgaben genau ansehen und abwägen, ob alles gut lief.

Die restliche Installation, wieder wie immer: fastboot flash boot boot.img. Ich möchte hierbei nicht auf Details eingehen - dies wird schon an vielen anderen Stellen erläutert.

Bei mir ging danach mein 2HD (getestet mit 10.1.0) mit WLAN und zusätzlichen Apps :-).
