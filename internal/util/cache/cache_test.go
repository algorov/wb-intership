package cache

import (
	"github.com/stretchr/testify/assert"
	"l0Service/internal/util/jsonutil"
	"testing"
)

func TestCache_GetOrder(t *testing.T) {
	cases := []struct {
		case_name string
		orders    []jsonutil.Order
		query     string
		expected  *jsonutil.Order
		exists    bool
	}{
		{
			case_name: "Order exist",
			orders: []jsonutil.Order{
				{OrderUID: "anything1"},
				{OrderUID: "anything2"},
			},
			query:    "anything2",
			expected: &jsonutil.Order{OrderUID: "anything2"},
			exists:   true,
		}, {
			case_name: "Order not exist",
			orders: []jsonutil.Order{
				{OrderUID: "anything1"},
				{OrderUID: "anything2"},
			},
			query:    "errrrrr",
			expected: nil,
			exists:   false,
		},
	}

	for _, test_case := range cases {
		t.Run(test_case.case_name, func(t *testing.T) {
			cache := New()

			for _, order := range test_case.orders {
				cache.AddOrder(order)
			}

			result, flag := cache.GetOrder(test_case.query)
			assert.Equal(t, test_case.exists, flag)
			assert.Equal(t, test_case.expected, result)
		})
	}
}
