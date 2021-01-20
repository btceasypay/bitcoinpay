# Ledger

### How to generate ledger

* You can use this command to generating ledger for the next bitcoinpay version.
```
~ cd ./tools/payledger
~ go build
~ ./payledger -h
```
* If you want to use all UTXOs from `srcdatadir`:
```
~ ./payledger --srcdatadir=[YourBitcoinpayDataPath] --endpoint=*
```

* If you want to use specific UTXOs from `srcdatadir`:
```
~ ./payledger --srcdatadir=[YourBitcoinpayDataPath] --endpoint=000005fd233345570677bc257e7c35e300dfe9b6d384bd8a0659c6619ff7ab30
```

* Then, you can build the next bitcoinpay version.
```
~ cd ./../../
~ go build
```

### How to show last result of generated ledger
```
~ ./payledger --last
```

### How to find end point
* You can print a list of recommendations when using `showendpoints` to set the specific number.
```
~ ./payledger --showendpoints=100
```
* Or skip some blocks
```
~ ./payledger --showendpoints=100 --endpointskips=200
```
* You can check if the end point works
```
~ ./payledger --checkendpoint=000005fd233345570677bc257e7c35e300dfe9b6d384bd8a0659c6619ff7ab30
```

### How to debug address
```
~ ./payledger --debugaddress=[Bitcoinpay Address]
or
~ ./payledger --debugaddress=[Bitcoinpay Address] --debugaddrutxo
or
~ ./payledger --debugaddress=[Bitcoinpay Address] --debugaddrvalid
```


### How to show blocks info
```
~ ./payledger --blocksinfo
```