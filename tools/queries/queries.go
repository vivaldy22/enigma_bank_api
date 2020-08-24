package queries

const (
	GET_ALL_LOGIN   = `SELECT * FROM login WHERE status_del = 0`
	GET_BY_ID_LOGIN = `SELECT * FROM login WHERE login_id = ? AND status_del = 0`
	CREATE_LOGIN    = `INSERT INTO login VALUES (NULL, ?, ?, 0)`
	UPDATE_LOGIN    = `UPDATE login 
						SET username = ?,
							password = ?
						WHERE login_id = ? AND
								status_del = 0`
	DELETE_LOGIN = `UPDATE login
						SET status_del = 1
						WHERE login_id = ?`

	GET_ALL_USER          = `SELECT * FROM user WHERE status_del = 0`
	GET_BY_ID_USER        = `SELECT * FROM user WHERE user_id = ? AND status_del = 0`
	GET_BY_ID_LOGIN_OWNER = `SELECT * FROM user WHERE login_owner_id = ? AND status_del = 0`
	CREATE_USER           = `INSERT INTO user VALUES (NULL, ?, ?, 0)`
	UPDATE_USER           = `UPDATE user 
						SET balance = ?
						WHERE login_owner_id = ? AND
								status_del = 0`
	DELETE_USER = `UPDATE user
						SET status_del = 1
						WHERE login_owner_id = ?`

	GET_ALL_TRANSACTION   = `SELECT * FROM transaction WHERE status_del = 0`
	GET_BY_ID_TRANSACTION = `SELECT * FROM transaction WHERE trans_id = ? AND status_del = 0`
	GET_BY_ID_USER_OWNER  = `SELECT * FROM transaction WHERE user_owner_id = ? AND status_del = 0`
	CREATE_TRANSACTION    = `INSERT INTO transaction VALUES (NULL, ?, ?, ?, ?, ?, 0)`
	UPDATE_TRANSACTION    = `UPDATE transaction 
							SET trans_date = ?,
								destination = ?,
								amount = ?,
								description = ?,
							WHERE trans_id = ? AND
									status_del = 0`
	DELETE_TRANSACTION = `UPDATE transaction
							SET status_del = 1
							WHERE trans_id = ?`
)
