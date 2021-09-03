# Google Authorizer CLI
Simple CLI that allows to generate `6-digit` auth numbers for 2fa without mobile app

It uses `HMAC`, `SHA-256` and some other tools to generate auth code.
After that it is displayed on console and nothing else.

You can use it for setting every 2fa where Google Authorizer is mentioned<br>

## Syntax

The syntax is `gauth --key hereYourKey`<br>
In place of `hereYourKey` place your manual 2fa secret.

## Building

Here is note for someone who may want to build this CLI and use it.<br>
To build use `go` cli, go into our project directory and type: `go build ./src/main.go -o gauth`<br>
Call it on unix systems by doing `./gauth`, if you use windows, type `gauth.exe` and change `-o` param value into `gauth.exe`

**Note**
This project was originally written in JS but got rewrited into GO, as per of static typing and so on.