package sensor

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"time"
)

type MqttSensor struct {
	Sensor
}

func NewMqttSensor(options *Options) *MqttSensor {
	s := &MqttSensor{}
	s.options = *options
	return s
}

func (s *MqttSensor) Update(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	s.UpdatedTime = time.Now()
	var value float64
	var err error
	value, err = strconv.ParseFloat(string(msg.Payload()), 8)

	if err == nil {
		s.SetValue(value)
	}
}
