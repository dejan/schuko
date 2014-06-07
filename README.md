schuko
======

Subscribe a websocket client to a Redis PubSub channel. 


Server setup
------------

Here's a minimal server implementation using schuko:

```Go
package main

import (
        "net/http"
        "github.com/dejan/schuko"
)

func main() {

        // configure schuko if needed
        // schuko.RedisUrl = ":6379"

	err := http.ListenAndServe(":4000", schuko.NewHandler())
	if err != nil {
		panic(err.Error())
        }

}
```


Subscribe
---------

Subscribe with plain old WebSocket API.

```JavaScript
var ws = new WebSocket("ws://localhost:4000/coolchannel")
ws.onmessage = function(e) { console.log(e.data); }
```


Publish
-------

Publish messages through Redis.

```
$ redis-cli 
127.0.0.1:6379> publish /coolchannel "Hey Gringo!"
```

