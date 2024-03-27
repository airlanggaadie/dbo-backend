package postgresql

const (
	queryGetUsersPaginate          = `SELECT id, username, name, created_at, updated_at FROM users OFFSET $1 LIMIT $2;`
	queryGetUsersByNamePaginate    = `SELECT id, username, name, created_at, updated_at FROM users WHERE name ILIKE $1 OFFSET $2 LIMIT $3;`
	queryGetUsersCount             = `SELECT count(*) FROM users;`
	queryGetUsersCountByName       = `SELECT count(*) FROM users WHERE name ILIKE $1;`
	queryGetUserById               = `SELECT id, username, name, created_at, updated_at FROM users WHERE id = $1;`
	queryGetUserByUsername         = `SELECT id, username, name, created_at, updated_at FROM users WHERE username = $1;`
	queryGetSimpleUsersByName      = `SELECT id, name FROM users WHERE name ILIKE $1 LIMIT 10;`
	queryInsertUser                = `INSERT INTO users (id, username, name) VALUES ($1, $2, $3);`
	queryInsertUserPassword        = `INSERT INTO user_password (user_id, password) VALUES ($1, $2);`
	queryInsertUserLoginHistory    = `INSERT INTO user_login_history (user_id) VALUES ($1);`
	queryUpdateUser                = `UPDATE users SET username = $2, name = $3, updated_at = NOW() WHERE id = $1 RETURNING id, username, name, created_at, updated_at;`
	queryDeleteUser                = `DELETE FROM users WHERE id = $1;`
	queryDeleteUserPassword        = `DELETE FROM user_password WHERE user_id = $1;`
	queryGetPasswordByUserId       = `SELECT password FROM user_password WHERE user_id = $1;`
	queryGetUserLoginHistoryDetail = `SELECT h.user_id, u.username, h.created_at FROM user_login_history h INNER JOIN users u ON u.id = h.user_id OFFSET $1 LIMIT $2;`
	queryGetUserLoginHistoryCount  = `SELECT count(*) FROM user_login_history;`

	queryGetOrdersPaginate        = `SELECT id, code, buyer_name, item_quantity, total_price, created_at, updated_at FROM orders OFFSET $1 LIMIT $2;`
	queryGetOrdersByNamePaginate  = `SELECT id, code, buyer_name, item_quantity, total_price, created_at, updated_at FROM orders WHERE code ILIKE $1 OFFSET $2 LIMIT $3;`
	queryGetOrdersCount           = `SELECT count(*) FROM orders;`
	queryGetOrdersCountByCode     = `SELECT count(*) FROM orders WHERE code ILIKE $1;`
	queryGetOrderById             = `SELECT id, code, buyer_name, item_quantity, total_price, created_at, updated_at FROM orders WHERE id = $1;`
	queryGetOrderItemByOrderId    = `SELECT id, order_id, code, name, quantity, unit_price, total_price, created_at, updated_at FROM orders WHERE order_id = $1;`
	queryDeleteOrder              = `DELETE FROM orders WHERE id = $1;`
	queryDeleteOrderItemByOrderId = `DELETE FROM order_item WHERE order_id = $1;`
	queryGetSimpleOrdersByName    = `SELECT id, code FROM orders WHERE code ILIKE $1 LIMIT 10`
	queryInsertOrder              = `INSERT INTO orders (id, code, buyer_name, item_quantity, total_price) VALUES ($1, $2, $3, $4, $5);`
	queryInsertOrderItem          = `INSERT INTO order_item (id, order_id, code, name, quantity, unit_price, total_price) VALUES `
	queryUpdateOrder              = `UPDATE orders SET buyer_name = $2, updated_at = NOW() WHERE id = $1;`
)
