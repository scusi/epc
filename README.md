# go-epc

A very simple implementation of the [EPC069-12 Standard](https://www.europeanpaymentscouncil.eu/sites/default/files/kb/file/2018-05/EPC069-12%20v2.1%20Quick%20Response%20Code%20-%20Guidelines%20to%20Enable%20the%20Data%20Capture%20for%20the%20Initiation%20of%20a%20SCT.pdf), also known as _Girocode_.

_Girocodes_ are basically SEPA Bank Transfers encoded in a standardized message format that is machine readable and can be encoded into QR-Codes.
This is very handy for users of mobile banking apps, who can just scan the QR-Code and directly have all transaction details without typing all the details.


## download lib

```
go get gitlab.scusi.io/flow/epc
```

### build lib

```
cd $GOSRC/gitlab.scusi.io/flow/epc
go build ./
```

## Examples

The lib contains 3 example programs that show how to use the lib in a program.
The following sections describe how to build and use these example programs.

### build example program

```
cd $GOSRC/gitlab.scusi.io/flow/epc
go build ./cmd/epc-simple
go build ./cmd/epc-parse
```

### Usage of epc-simple

`epc-simple` can output the EPC message as Text or as QR-Code as a PNG file.

### Text Format

With the following command you can create a new EPC message.

```
$ ./epc-simple -i "DE53200400600200400600" -n "Bündnis Entwicklung Hilft" \
	-a 5 -s "ARD/ Nothilfe Ukraine" -b "COBADEFFXXX" -format text 
BCD
002
1
SCT
COBADEFFXXX
Bündnis Entwicklung Hilft
DE53200400600200400600
EUR5


ARD/ Nothilfe Ukraine
```

Basically you can pipe the above command to `qrencode` to manually create a EPC-QR-Code from the message.

```
$ ./epc-simple -i "DE53200400600200400600" -n "Bündnis Entwicklung Hilft" \
	-a 5 -s "ARD/ Nothilfe Ukraine" -b "COBADEFFXXX" -format text\
	| qrencode -l H -t PNG -o images/test-qr.png 
```

The above command would write a new QR-Code into the file `images/test-qr.png`.

![test-qr.png](/images/test-qr.png)

### PNG Format

In order to create a new EPC-QR-Code you can do the follwoing:

```
$ ./epc-simple -i "DE53200400600200400600" -n "Bündnis Entwicklung Hilft" \
	-a 5 -s "ARD/ Nothilfe Ukraine" -b "COBADEFFXXX" -format png > images/test-qr2.png 
```

The above command would write a new QR-Code into the file `images/test-qr2.png`.

![test-qr2.png](/images/test-qr2.png)

### Parsing EPC Messages

You can also do the reverse and parse a text format EPC message into a EPC datastructure.

```
$ ./epc-simple -i "DE53200400600200400600" -n "Bündnis Entwicklung Hilft" \
        -a 5 -s "ARD/ Nothilfe Ukraine" -b "COBADEFFXXX" -format text > test.epc
$ ./epc-parse -f test.epc
BCD
002
1
SCT
COBADEFFXXX
Bündnis Entwicklung Hilft
DE53200400600200400600
EUR5


ARD/ Nothilfe Ukraine
```

The first command from the above example does create a new EPC Message, in Text format and pipes that to a file called `test.epc`
The second command reads `test.epc` and parses the content into a EPC datastruct (`epc.EPC`), before writeing it to STDOUT.


## Links

- GERMAN - Heise Artikel zum Thema: [Online-Banking: Rechnungen schneller mit QR-Codes überweisen](https://heise.de/-6543687)
- ENGLISH - EPC069-12 Standard: [EPC069-12](https://www.europeanpaymentscouncil.eu/sites/default/files/kb/file/2018-05/EPC069-12%20v2.1%20Quick%20Response%20Code%20-%20Guidelines%20to%20Enable%20the%20Data%20Capture%20for%20the%20Initiation%20of%20a%20SCT.pdf)

