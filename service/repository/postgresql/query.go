package postgresql

const (
	queryGetUsersPaginate       = `SELECT id, username, name, created_at, updated_at FROM users OFFSET $1 LIMIT $2;`
	queryGetUsersByNamePaginate = `SELECT id, username, name, created_at, updated_at FROM users WHERE name ILIKE $1 OFFSET $2 LIMIT $3;`
	queryGetUsersCount          = `SELECT count(*) FROM users;`
	queryGetUsersCountByName    = `SELECT count(*) FROM users WHERE name ILIKE $1;`
	queryGetUserById            = `SELECT id, username, name, created_at, updated_at FROM users WHERE id = $1;`
	queryGetSimpleUsersByName   = `SELECT id, name FROM users WHERE name ILIKE $1 LIMIT 10;`
	queryInsertUser             = `INSERT INTO users (id, username, name) VALUES ($1, $2, $3);`
	queryInsertUserPassword     = `INSERT INTO user_password (user_id, password) VALUES ($1, $2);`
	queryUpdateUser             = `UPDATE users SET username = $2, name = $3, updated_at = NOW() WHERE id = $1 RETURNING id, username, name, created_at, updated_at;`
	queryDeleteUser             = `DELETE FROM users WHERE id = $1;`
	queryDeleteUserPassword     = `DELETE FROM user_password WHERE user_id = $1;`
)
