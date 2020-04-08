# config
Package config contains settings of the service

## Configuration parameters

### HTTP API

ENV VARIABLE       	    	      | DESCRIPTION
--------------------------------|--------------------------------------------------------------------------------------------------------
ORDER_HTTP_ADDRESS              | address and port to be use for the API (Default: `:8080`)

### Order generator

ENV VARIABLE       	    	      | DESCRIPTION
--------------------------------|--------------------------------------------------------------------------------------------------------
ORDER_GENERATOR_KEYSIZE         | size of unique order identifier (Default: `2`)
ORDER_GENERATOR_KEYBYTES        | allowed symbols to be used for key generation (Default: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
ORDER_GENERATOR_POOLSIZE	      | length of order pool (Default: `50`)
ORDER_GENERATOR_REPLACEINTERVAL | an interval to replace some order in the pool with a new one (Default: `200ms`)
