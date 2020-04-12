# gopher


[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/prodhe/gopher) 

A server hosting files, dirs and links according to the gopher protocol.

Because gopherspace deserves more servers.

Work in progress as a playground for basic networking in go. And readers,
writers, buffers, []byte and whatnot...

## Gopher protocol

[ietf rfc](https://tools.ietf.org/html/rfc1436)

## License

MIT. See LICENSE.


## Docker hub

[docker hub](https://hub.docker.com/r/prodhe/gopher/)

<<<<<<< HEAD
## Docker compose

```yml
version: '3.1'

services:
  gopher:
    image: prodhe/gopher:latest
    ports:
      - 70:70
    environment:
      - GOPHER_ADDRESS=localhost
    volumes:
      - /opt/gopher/public:/public
      
```


=======
>>>>>>> 504b2c99d23e10de9bc1d1226c762bdead7f596d
