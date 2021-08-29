# Google Authorizer CLI

So simple, so cute UwU<br>
This is just a simple switch to *`console mode`*<br>

If you wondering how it works, here is your answer:
> It uses HMAC SHA-256 to generate 6 digit codes
It is something that google authorizer uses, but on console.

You can use it for setting every 2fa where Google Authorizer is mentioned<br>
And.. ah the syntax is `gauth --key hereYourKey`<br>
In place of `hereYourKey` place your manual 2fa secret.

## Building
Here is note for someone who may want to build this CLI and use it.<br>
To build use `go` cli, go into our project directory and type: `go build ./src/main.go -o gauth`<br>
Call it on unix systems by doing `./gauth`, if you use windows, type `gauth.exe` and change `-o` param value into `gauth.exe`