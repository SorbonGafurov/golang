package service

import (
	"IbtService/internal/config"
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	returnCh chan amqp.Return
	confirms chan amqp.Confirmation
	close    chan *amqp.Error
}

func ConnRabbit(cfg *config.Config) (*Rabbit, error) {
	conn, err := amqp.Dial(cfg.UrlRabbit)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("ошибка открытия канала: %w", err)
	}

	r := &Rabbit{
		conn:     conn,
		ch:       ch,
		returnCh: make(chan amqp.Return, 10),
		confirms: make(chan amqp.Confirmation, 10),
		close:    make(chan *amqp.Error),
	}

	err = r.ch.Confirm(false)
	if err != nil {
		return nil, err
	}

	go func() {
		for c := range r.ch.NotifyPublish(r.confirms) {
			if !c.Ack {
				log.Println("Nacked")
			}
		}
	}()

	go func() {
		for err := range r.ch.NotifyClose(r.close) {
			if err != nil {
				log.Println("Close")
			}
		}
	}()

	go func() {
		for range r.ch.NotifyReturn(r.returnCh) {
			log.Println("Return")
		}
	}()

	return r, nil
}

func (r *Rabbit) PublishToRabbit(data []byte) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.ch.PublishWithContext(ctx,
		"transfer",    // exchange
		"transferKey", // routing key
		true,          // mandatory
		false,         // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Rabbit) Close() {
	if r.ch != nil {
		_ = r.ch.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
}
