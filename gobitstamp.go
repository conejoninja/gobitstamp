package gobitstamp

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type API struct {
	apiKey        string
	apiSecretKey  string
	apiClientId   string
	httpTransport *http.Transport
}

func NewAPI(apiKey, apiSecretKey, apiClientId string) *API {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &API{apiKey, apiSecretKey, apiClientId, tr}
}

func (this *API) Call(action, httpMethod string, params map[string]string) ([]byte, error) {

	var err error
	var res *http.Response

	if httpMethod=="POST" {
		nonce := strconv.Itoa(int(time.Now().UnixNano()))
		values := make(url.Values)
		values.Set("nonce", nonce)
		values.Set("key", this.apiKey)

		for key, val := range params {
			values.Set(key, val)
		}

		mac := hmac.New(sha256.New, []byte(this.apiSecretKey))
		mac.Write([]byte(nonce + this.apiClientId + this.apiKey))
		values.Set("signature", strings.ToUpper(hex.EncodeToString(mac.Sum(nil))))

		res, err = http.PostForm("https://www.bitstamp.net/api/"+action+"/", values)

	} else {
		res, err = http.Get("https://www.bitstamp.net/api/"+action+"/")
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body, err
}

func (this *API) Ticker() (interface{}, error) {
	dataStream, err := this.Call("ticker", "GET", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) OrderBook(unified bool) (interface{}, error) {
	unified_str := "0"
	if unified {
		unified_str = "1"
	}
	var params = map[string]string{
		"group": unified_str,
	}
	dataStream, err := this.Call("order_book", "POST", params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) Transactions(offset, limit int, sort string) (interface{}, error) {
	if limit==0 {
		limit = 100
	}
	if sort != "asc"{
		sort = "desc"
	}
	var params = map[string]string{
		"offset": strconv.Itoa(offset),
		"limit": strconv.Itoa(limit),
		"sort": sort,
	}
	dataStream, err := this.Call("transactions", "POST", params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) Rate() (interface{}, error) {
	dataStream, err := this.Call("eur_usd", "POST", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) Balance() (interface{}, error) {
	dataStream, err := this.Call("balance", "POST", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) UserTransactions(offset, limit int, sort string) (interface{}, error) {
	if limit==0 {
		limit = 100
	}
	if sort != "asc"{
		sort = "desc"
	}
	var params = map[string]string{
		"offset": strconv.Itoa(offset),
		"limit": strconv.Itoa(limit),
		"sort": sort,
	}
	dataStream, err := this.Call("user_transactions", "POST", params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) OpenOrders() (interface{}, error) {
	dataStream, err := this.Call("open_orders", "POST", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) CancelOrder(id int) (interface{}, error) {
	var params = map[string]string{
		"id": strconv.Itoa(id),
	}
	dataStream, err := this.Call("cancel_order", "POST", params)
	return string(dataStream), err
}

func (this *API) Buy(amount, price float64) (interface{}, error) {
	var params = map[string]string{
		"amount": strconv.FormatFloat(amount, 'e', 2, 32),
		"price": strconv.FormatFloat(price, 'e', 2, 32),
	}
	dataStream, err := this.Call("buy", "POST", params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) Sell(amount, price float64) (interface{}, error) {
	var params = map[string]string{
		"amount": strconv.FormatFloat(amount, 'e', 2, 32),
		"price": strconv.FormatFloat(price, 'e', 2, 32),
	}
	dataStream, err := this.Call("sell", "POST", params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) Withdrawal() (interface{}, error) {
	dataStream, err := this.Call("withdrawal_request", "POST", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) WithdrawalBitcoin(amount float64, address string) (interface{}, error) {
	var params = map[string]string{
		"amount": strconv.FormatFloat(amount, 'e', 2, 32),
		"address": address,
	}
	dataStream, err := this.Call("bitcoin_withdrawal", "POST", params)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) DepositBitcoin() (interface{}, error) {
	dataStream, err := this.Call("bitcoin_deposit_address", "POST", nil)
	return string(dataStream), err
}

func (this *API) UnconfirmedDeposits() (interface{}, error) {
	dataStream, err := this.Call("unconfirmed_btc", "POST", nil)
	data := map[string]interface{}{}
	json.Unmarshal(dataStream, &data)
	return data, err
}

func (this *API) WithdrawalRipple(amount float64, address, currency string) (interface{}, error) {
	var params = map[string]string{
		"amount": strconv.FormatFloat(amount, 'e', 2, 32),
		"address": address,
		"currency": currency,
	}
	dataStream, err := this.Call("ripple_withdrawal", "POST", params)
	return string(dataStream), err
}

func (this *API) DepositRipple() (interface{}, error) {
	dataStream, err := this.Call("ripple_address", "POST", nil)
	return string(dataStream), err
}




