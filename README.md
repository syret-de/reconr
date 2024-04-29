# Reconr


## Architecture

Task -> Une commande
Workflow -> Plein de commandes

Un container par task.

```mermaid
graph LR
A(subdomain) -->B(httpx)
    B --> C(scope)
    C -->D(katana)
    C -->E(nuclei)
    C -->F(fuff)
    D --> G(gf)
    E --> H(dashboard)
```
## TODO

- [x] Ajouter la possibilité de faire deux lignes de commandes
- [x] Fixer les logs
- [x] Faire un fichier target avec scope
- [x] Template workflow
- [x] Faire quelque chose pour le scope
- [ ] Droit sur les fichers out
- [ ] Procedure d'install
- [ ] Plus de config
- [ ] Faire fonctionner deux containers en parallèle
- [ ] Refactor le code
- [ ] Faire quelque chose avec dnsx, alterx, nabuu
https://github.com/1ndianl33t/Gf-Patterns
https://github.com/tomnomnom/waybackurls

## jq
```
 cat out2.txt | jq "[.url,.host]"
 cat nulcei.txt | jq '.info | select(.severity == "low")'
```

## Valid scope
```
grep -Ff scope.txt out2.txt | jq -r .url
```

## Nuclei
```
nuclei -u url.com -cloud-upload
```

## Scope
```
root@PCYann:/tmp# grepcidr -f scope.txt target.txt 
192.168.30.4

192.168.10.1/0 -> all ips
```
