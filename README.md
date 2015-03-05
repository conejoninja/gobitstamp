GoBitstamp
==========
Simple wrapper in Go for Bitstamp's API


## Installation


```bash
$ go get github.com/conejoninja/gobitstamp
```

## Documentation
See [Go Doc](http://godoc.org/github.com/conejoninja/gobitstamp) or [Go Walker](http://gowalker.org/github.com/conejoninja/gobitstamp) for usage and details.

## Example of use

```go
package main

import (
	"github.com/conejoninja/gobitstamp"
	"log"
)

func main() {
	var err error
	var msg interface{}
	var api *gobitstamp.API

	api = gobitstamp.NewAPI("YOUR_API_KEY", "YOUR_API_SECRET", "ACCOUNT_ID")

	msg, err = api.Ticker()
	log.Println(msg)
	log.Println(err)

	msg, err = api.DepositBitcoin()
	log.Println(msg)
	log.Println(err)

	msg, err = api.UnconfirmedDeposits()
	log.Println(msg)
	log.Println(err)

	msg, err = api.UserTransactions(0,100,"sort")
	log.Println(msg)
	log.Println(err)

	msg, err = api.OpenOrders()
	log.Println(msg)
	log.Println(err)

	msg, err = api.OrderBook(false)
	log.Println(msg)
	log.Println(err)

}
```

## Noted
This is my first Go project, I wouldn't use for anything serious or important.

## Contributing to GoBitstamp:

If you find any improvement or issue you want to fix, feel free to send me a pull request with testing.


## License

This is distributed under the Apache License v2.0

Copyright 2014 Daniel Esteban  -  conejo@conejo.me

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

