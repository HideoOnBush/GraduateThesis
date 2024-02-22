package line

import (
	model "GraduateThesis/biz/model/line"
	"GraduateThesis/conf"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var (
	exchangeName = "user.register.topic"
	queueName    = "user.register.queue"
	keyName      = "user.register.event"
)

func (l *Line) InitializeConsumers() {
	num, _ := strconv.Atoi(conf.GConfig.RabbitConCurrencyNum)
	limitC := make(chan struct{}, num)
	defer close(limitC)
	for {
		limitC <- struct{}{}
		go func() {
			defer func() {
				<-limitC
			}()
			if err := l.RABBITMQ.NewConsumer(queueName, func(body []byte) error {
				msg := model.Line{}
				err := json.Unmarshal(body, &msg)
				if err != nil {
					return errors.Wrap(err, "unmarshal in initializeConsumers failed")
				}
				err = l.Bulk(context.Background(), []*model.Line{&msg})
				log.Printf("consume msg from %s to %s", msg.Source, msg.Target)
				return nil
			}); err != nil {
				log.Fatalf("consume err: %v", err)
			}
		}()
	}
}
