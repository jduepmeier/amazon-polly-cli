# amazon-polly-cli
Minimal tool to use Amazon Polly

This tool uses Amazon Polly to synthesize speech.
It is written in [go](https://golang.org) and uses the aws go sdk [(aws-sdk-go)](https://github.com/aws/aws-sdk-go/)
to communicate with Amazon Polly.

## Installing

```
go get
go build
```

## Usage

You must setup the aws config and credentials first (see [aws documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)).

Here is an example use under linux:
```
./amazon-polly-cli -lang en -text <text> -out outfile.mp3
```
The tool can also read the text from stdin (```-text -```) or from file (```-input-file <file>```).

You can select the language with the ```-lang``` flag.
Currently supported language identifiers are
* de // german
* en // english us
* fr // french
* ca-fr // canadian-french
