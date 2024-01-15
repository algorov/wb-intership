package store

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"l0Service/internal/util/jsonutil"
	"testing"
)

func TestStore_AddOrder(t *testing.T) {
	jsonData := "{\n  \"order_uid\": \"AHSDKASJDK\",\n  \"track_number\": \"WBILMTESTTRACK\",\n  \"entry\": \"WBIL\",\n  \"delivery\": {\n    \"name\": \"Test Testov\",\n    \"phone\": \"+9720000000\",\n    \"zip\": \"2639809\",\n    \"city\": \"Kiryat Mozkin\",\n    \"address\": \"Ploshad Mira 15\",\n    \"region\": \"Kraiot\",\n    \"email\": \"test@gmail.com\"\n  },\n  \"payment\": {\n    \"transaction\": \"b563feb7b2b84b6test\",\n    \"request_id\": \"\",\n    \"currency\": \"USD\",\n    \"provider\": \"wbpay\",\n    \"amount\": 1817,\n    \"payment_dt\": 1637907727,\n    \"bank\": \"alpha\",\n    \"delivery_cost\": 1500,\n    \"goods_total\": 317,\n    \"custom_fee\": 0\n  },\n  \"items\": [\n    {\n      \"chrt_id\": 9934930,\n      \"track_number\": \"WBILMTESTTRACK\",\n      \"price\": 453,\n      \"rid\": \"ab4219087a764ae0btest\",\n      \"name\": \"Mascaras\",\n      \"sale\": 30,\n      \"size\": \"0\",\n      \"total_price\": 317,\n      \"nm_id\": 2389212,\n      \"brand\": \"Vivienne Sabo\",\n      \"status\": 202\n    }\n  ],\n  \"locale\": \"en\",\n  \"internal_signature\": \"\",\n  \"customer_id\": \"test\",\n  \"delivery_service\": \"meest\",\n  \"shardkey\": \"9\",\n  \"sm_id\": 99,\n  \"date_created\": \"2021-11-26T06:22:19Z\",\n  \"oof_shard\": \"1\"\n}"
	config := Config{
		DatabaseURL: databaseURL,
	}

	s := New(&config)
	if err := s.Open(); err != nil {
		assert.Fail(t, err.Error())
	}

	order := jsonutil.Order{}

	if err := json.Unmarshal([]byte(jsonData), &order); err != nil {
		assert.Fail(t, err.Error())
	}

	result, err := s.AddOrder(&order)
	if err != nil {
		assert.Fail(t, err.Error())
		fmt.Printf("%+v\n", order)
	}

	assert.Equal(t, true, result)
}

func TestStore_GetOrders(t *testing.T) {
	config := Config{
		DatabaseURL: databaseURL,
	}

	s := New(&config)
	if err := s.Open(); err != nil {
		assert.Fail(t, err.Error())
	}

	orders, err := s.GetOrders()
	if err != nil {
		assert.Fail(t, err.Error())
		fmt.Printf("%+v\n", orders)
	}

	var isExist bool = false
	if orders != nil {
		isExist = true
	}

	assert.Equal(t, true, isExist)
}
