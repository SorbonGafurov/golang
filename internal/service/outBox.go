package service

import (
	"IbtService/internal/config"
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type OutBoxMessage struct {
	Message string
}

type OutBox struct {
	db *sql.DB
}

func OutBoxOpen(cfg *config.Config) (*OutBox, error) {
	db, err := sql.Open("postgres", cfg.PostreSqlConnStr)
	if err != nil {
		return nil, err
	}

	oBox := OutBox{db: db}

	return &oBox, nil
}

func (o *OutBox) Close() {
	o.db.Close()
}

func (o *OutBox) InsertOutBox(message string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "insert into outboxmessages (message, status) values ($1, $2);"
	result, err := o.db.ExecContext(ctx, query, message, "new")
	if err != nil {
		return 0, err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func (o *OutBox) SelectOutBox() (*OutBoxMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	om := &OutBoxMessage{}

	query := `
		UPDATE outboxmessages
		SET status = 'processing'
		WHERE id = (
			SELECT id FROM outboxmessages
			WHERE status = 'new'
			LIMIT 1
			FOR UPDATE SKIP LOCKED
		)
		RETURNING message;
	`
	err := o.db.QueryRowContext(ctx, query).Scan(&om.Message)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return om, nil
}
