package service

import (
	"IbtService/internal/config"
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn    *amqp.Connection
	returns chan amqp.Return
	done    chan struct{}
}

func ConnRabbit(cfg *config.Config) (*Rabbit, error) {
	conn, err := amqp.Dial(cfg.UrlRabbit)
	if err != nil {
		return nil, err
	}

	r := &Rabbit{
		conn:    conn,
		returns: make(chan amqp.Return),
		done:    make(chan struct{}),
	}

	// Постоянный слушатель возвратов
	go func() {
		for {
			select {
			case ret, ok := <-r.returns:
				if !ok {
					log.Println("return channel closed — stopping listener")
					return
				}
				log.Printf("❌ Returned: %s (key=%s body=%s)", ret.ReplyText, ret.RoutingKey, string(ret.Body))

			case <-r.done:
				log.Println("return listener stopped by shutdown")
				return
			}
		}
	}()

	return r, nil
}

func (r *Rabbit) PublishToRabbit(data []byte) (bool, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return false, err
	}
	defer ch.Close()

	/*confirms := make(chan amqp.Confirmation)
	ch.NotifyPublish(confirms)
	err = ch.Confirm(false)
	failOnError(err, "Failed to confirm")*/

	//returns := make(chan amqp.Return)
	ch.NotifyReturn(r.returns)

	err = ch.ExchangeDeclare(
		"transfers", // name
		"direct",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"transfers",   // exchange
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

	select {
	case <-returns:
		return false, nil
	case <-time.After(2 * time.Second):
		return true, nil
	}
}

func (r *Rabbit) Close() {
	if r.conn != nil {
		_ = r.conn.Close()
	}
}
