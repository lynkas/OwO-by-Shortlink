# OwO-by-Shortlink

a simple maybe buggy shortlink and text host server, write in golang with mongodb.

#### build
```
git clone https://github.com/lynkas/OwO-by-Shortlink.git
go build .
```

#### edit config.json

db => mongodb url, tokens are the map of ```token => client_name```, for discriminating different sources.
```
{
  "db": "mongodb://url",
  "token": {
    "strong_token_1": "client1",
    "strong_token_2": "client2"
  }
}
```
you can get free 512MB space mongodb service on atlas.


### use
#### new item

post data in following format to ```/new/```

```
{
    "format":"content type",
    "content":"content",
    "token":"token you leave in the config file",
    "identify":"user identity, such as cookie or waffle, not required"    
}   
```

as the repay, you will get

```
{
    "prefix":"the random chars",
    "text_serial":"a string for id",
    "serial": int
}
```

prefix+text_serial consists of the link.

for example, prefix is ```sdakj``` and serial_text is ```a24dd```, then the link token will be ```sdakja24dd```
 
serial (that number) is not really useful for the user. ;)

#### get link content

when you have a link with some confusing token, like ```sdakja24dd```, feed it to ```/query/{confusing_token}```.
Get the link and you will get

```
{
    "format":"content_type",
    "content":"content"
}
```

#### error

all errors have a 4xx code and with following format body:

```
{
    "message":"the error message"
}
```

That's all, thank you for your interest or using, and any contribution or suggestion are welcomed. 

Under MIT license. 