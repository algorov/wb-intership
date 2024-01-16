package jsonutil

import (
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
	"time"
)

const schemaString = `{
  "type": "object",
  "properties": {
    "order_uid": {"type": "string"},
    "track_number": {"type": "string"},
    "entry": {"type": "string"},
    "delivery": {
      "type": "object",
      "properties": {
        "name": {"type": "string"},
        "phone": {"type": "string"},
        "zip": {"type": "string"},
        "city": {"type": "string"},
        "address": {"type": "string"},
        "region": {"type": "string"},
        "email": {"type": "string"}
      },
      "required": ["name", "phone", "zip", "city", "address", "region", "email"]
    },
    "payment": {
      "type": "object",
      "properties": {
        "transaction": {"type": "string"},
        "request_id": {"type": "string"},
        "currency": {"type": "string"},
        "provider": {"type": "string"},
        "amount": {"type": "integer"},
        "payment_dt": {"type": "integer"},
        "bank": {"type": "string"},
        "delivery_cost": {"type": "integer"},
        "goods_total": {"type": "integer"},
        "custom_fee": {"type": "integer"}
      },
      "required": ["transaction", "currency", "provider", "amount", "payment_dt", "bank", "delivery_cost", "goods_total", "custom_fee"]
    },
    "items": {"type": "array"},
    "locale": {"type": "string"},
    "internal_signature": {"type": "string"},
    "customer_id": {"type": "string"},
    "delivery_service": {"type": "string"},
    "shardkey": {"type": "string"},
    "sm_id": {"type": "integer"},
    "date_created": {"type": "string"},
    "oof_shard": {"type": "string"}
  },
  "required": ["order_uid", "track_number", "entry", "delivery", "payment", "items", "locale", "internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard"]
}
`

type DeliveryInfo struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type PaymentInfo struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type Order struct {
	OrderUID        string       `json:"order_uid"`
	TrackNumber     string       `json:"track_number"`
	Entry           string       `json:"entry"`
	Delivery        DeliveryInfo `json:"delivery"`
	Payment         PaymentInfo  `json:"payment"`
	Items           []Item       `json:"items"`
	Locale          string       `json:"locale"`
	InternalSig     string       `json:"internal_signature"`
	CustomerID      string       `json:"customer_id"`
	DeliveryService string       `json:"delivery_service"`
	Shardkey        string       `json:"shardkey"`
	SmID            int          `json:"sm_id"`
	DateCreated     time.Time    `json:"date_created"`
	OofShard        string       `json:"oof_shard"`
}

// ValidateJsonData ...
func ValidateJsonData(data string) bool {
	schemaLoader := gojsonschema.NewStringLoader(schemaString)
	documentLoader := gojsonschema.NewStringLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false
	}

	if result.Valid() {
		return true
	} else {
		return false
	}
}

// GetUnmarshallingJsonData ...
func GetUnmarshallingJsonData(data string) (order *Order, err error) {
	if err := json.Unmarshal([]byte(data), &order); err != nil {
		return nil, err
	}

	return order, nil
}
