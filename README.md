# Directory-io-with-balance
```bash
--Fonctionne pour les adresses compressées. 
--Adresse type : Pubkeyhash, Scripthash, Bech32 pubkeyhash
--Mode aléatoire.
--Mode afficher la balance (vraiment fonctionnel).
```
## Possibilitée de changer le code pour afficher les adresses non compressées et même supprimer le mode random.
```bash
Install Go 1.10

Installer les dependences:

$ go get github.com/btcsuite/btcd
```
si cela ne fonctionne pas de vous conseil de télécharger :
```bash
pour btcd : https://github.com/btcsuite/btcd/archive/refs/tags/v0.21.0-beta.zip
pour btcutil : https://github.com/btcsuite/btcutil/archive/refs/tags/psbt/v1.0.2.zip
```
```bash
$ go run directory.go
```
```bash
Par default sur le port port 80. http://localhost:8080
 ```
