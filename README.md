# GoQuotes
A REST API made in Go that handles the "/quote [text]" command. Also monitors
the #general chat for starAddedEvents using the real time slack api. These
starred messages are stored the same way as the /quote command would.

## Dependencies
* go get github.com/nlopes/slack
* go get github.com/nlopes/slack
* go get code.google.com/p/gcfg
* go get gopkg.in/dancannon/gorethink.v1
* go get github.com/gorilla/schema

## Wishlist
* Daily activity api
* Api Authentication
* Comments on quotes 
