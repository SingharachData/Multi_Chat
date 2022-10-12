package chat

const createChatTable = `CREATE TABLE IF NOT EXISTS chat (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	sender VARCHAR,
	name TEXT,
	sent_time DATETIME,
	text TEXT
);`

const insertMessageIntoChat = `INSERT INTO chat(sender, name, sent_time, text)
	VALUES(?, ?, ?, ?);`

const updateMessageInChat = `UPDATE chat
	SET text = ?
	WHERE id = ?;`

const deleteMessageFromChat = `DELETE from chat
	WHERE id = ?;`

const getAllMessagesFromChat = `SELECT * FROM chat;`

const getMessageFromChat = `SELECT * FROM chat WHERE id = ?;`
