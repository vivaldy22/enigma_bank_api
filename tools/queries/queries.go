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
)
