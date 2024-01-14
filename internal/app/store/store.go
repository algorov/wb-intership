package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"l0Service/internal/util/jsonutil"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// GetOrders ...
func (s *Store) GetOrders() (*[]jsonutil.Order, error) {
	orders := []jsonutil.Order{}

	rows, err_order := s.db.Query(
		"SELECT order_uid, track_number, entry, locale, internal_sig, custom_id, shard_key, sm_id, date_creates, oof_shard FROM orders")
	if err_order != nil {
		return nil, err_order
	}

	for rows.Next() {
		order := jsonutil.Order{}

		if err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSig,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard); err != nil {
			return nil, err
		}

		deliveryInfo, err_delivery := s.GetDeliveryInfo(order.OrderUID)
		if err_delivery != nil {
			return nil, err_delivery
		}

		order.Delivery = deliveryInfo

		paymentInfo, err_payment := s.GetPaymentInfo(order.OrderUID)
		if err_payment != nil {
			return nil, err_payment
		}

		order.Payment = paymentInfo

		items, err_items := s.GetItems(order.OrderUID)
		if err_items != nil {
			return nil, err_items
		}

		order.Items = items

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &orders, nil
}

// GetDeliveryInfo ...
func (s *Store) GetDeliveryInfo(orderUID string) (*jsonutil.DeliveryInfo, error) {
	deliveryInfo := jsonutil.DeliveryInfo{}
	if err := s.db.QueryRow(
		"SELECT name, phone, zip, city, address, region, email FROM delivery_info WHERE order_uid = $1", orderUID,
	).Scan(
		&deliveryInfo.Name,
		&deliveryInfo.Phone,
		&deliveryInfo.Zip,
		&deliveryInfo.City,
		&deliveryInfo.Address,
		&deliveryInfo.Region,
		&deliveryInfo.Email); err != nil {
		return nil, err
	}

	return &deliveryInfo, nil
}

// GetPaymentInfo ...
func (s *Store) GetPaymentInfo(orderUID string) (*jsonutil.PaymentInfo, error) {
	payInfo := jsonutil.PaymentInfo{}
	if err := s.db.QueryRow(
		"SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment_info WHERE order_uid = $1", orderUID,
	).Scan(
		&payInfo.Transaction,
		&payInfo.RequestID,
		&payInfo.Currency,
		&payInfo.Provider,
		&payInfo.Amount,
		&payInfo.PaymentDt,
		&payInfo.Bank,
		&payInfo.DeliveryCost,
		&payInfo.GoodsTotal,
		&payInfo.CustomFee); err != nil {
		return nil, err
	}

	return &payInfo, nil
}

// GetItems ...
func (s *Store) GetItems(orderUID string) (*[]jsonutil.Item, error) {
	rows, err := s.db.Query(
		"SELECT chrt_id, track_number, price, rid, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1", orderUID)
	if err != nil {
		return nil, err
	}

	items := []jsonutil.Item{}
	for rows.Next() {
		item := jsonutil.Item{}
		if err := rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &items, nil
}
