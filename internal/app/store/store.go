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

// AddOrder ...
func (s *Store) AddOrder(order *jsonutil.Order) (bool, error) {
	if _, err := s.addOrder(order); err != nil {
		return false, err
	}

	if _, err := s.addDeliveryInfo(order.OrderUID, &order.Delivery); err != nil {
		return false, err
	}

	if _, err := s.addPaymentInfo(order.OrderUID, &order.Payment); err != nil {
		return false, err
	}

	if _, err := s.addItems(order.OrderUID, &order.Items); err != nil {
		return false, err
	}

	return true, nil
}

// GetOrders ...
func (s *Store) GetOrders() ([]jsonutil.Order, error) {
	orders := []jsonutil.Order{}

	rows, err_order := s.db.Query(
		"SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, shardkey, sm_id, date_created, oof_shard FROM orders")
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
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard); err != nil {
			return nil, err
		}

		deliveryInfo, err_delivery := s.getDeliveryInfo(order.OrderUID)
		if err_delivery != nil {
			return nil, err_delivery
		}

		order.Delivery = deliveryInfo

		paymentInfo, err_payment := s.getPaymentInfo(order.OrderUID)
		if err_payment != nil {
			return nil, err_payment
		}

		order.Payment = paymentInfo

		items, err_items := s.getItems(order.OrderUID)
		if err_items != nil {
			return nil, err_items
		}

		order.Items = items

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Store) addOrder(order *jsonutil.Order) (int, error) {
	result, err := s.db.Exec(
		"INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSig,
		order.CustomerID,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(affected), nil
}

func (s *Store) addDeliveryInfo(orderUID string, deliveryInfo *jsonutil.DeliveryInfo) (int, error) {
	result, err := s.db.Exec(
		"INSERT INTO delivery_info (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		orderUID,
		deliveryInfo.Name,
		deliveryInfo.Phone,
		deliveryInfo.Zip,
		deliveryInfo.City,
		deliveryInfo.Address,
		deliveryInfo.Region,
		deliveryInfo.Email)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(affected), nil
}

func (s *Store) getDeliveryInfo(orderUID string) (jsonutil.DeliveryInfo, error) {
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
		return deliveryInfo, err
	}

	return deliveryInfo, nil
}

func (s *Store) addPaymentInfo(orderUID string, paymentInfo *jsonutil.PaymentInfo) (int, error) {
	result, err := s.db.Exec(
		"INSERT INTO payment_info (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		orderUID,
		paymentInfo.Transaction,
		paymentInfo.RequestID,
		paymentInfo.Currency,
		paymentInfo.Provider,
		paymentInfo.Amount,
		paymentInfo.PaymentDt,
		paymentInfo.Bank,
		paymentInfo.DeliveryCost,
		paymentInfo.GoodsTotal,
		paymentInfo.CustomFee)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(affected), nil
}

func (s *Store) getPaymentInfo(orderUID string) (jsonutil.PaymentInfo, error) {
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
		return payInfo, err
	}

	return payInfo, nil
}

func (s *Store) addItems(orderUID string, items *[]jsonutil.Item) (int, error) {
	var totalAffected int = 0

	for _, item := range *items {
		result, err := s.db.Exec(
			"INSERT INTO item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
			orderUID,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.RID,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status)
		if err != nil {
			return totalAffected, err
		}

		affected, err := result.RowsAffected()
		if err != nil {
			return totalAffected, err
		}

		totalAffected += int(affected)
	}

	return totalAffected, nil
}

func (s *Store) getItems(orderUID string) ([]jsonutil.Item, error) {
	rows, err := s.db.Query(
		"SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM item WHERE order_uid = $1", orderUID)
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

	return items, nil
}
