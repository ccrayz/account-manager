# mnemonic to private key


Setup config.yaml
```
vi config.yaml
```

Edit the mnemonic and count under eth to match.
```
eth:
  mnemonic: your mnemonic
  count: number of ethereum accounts to look up
```

Run it like this
```
go run main.go
```
