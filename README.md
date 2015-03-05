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

GoRequest is MIT License.

