package sensor

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

type MqttSensor struct {
	Sensor
}

func NewMqttSensor(options *Options) *MqttSensor {
	s := &MqttSensor{}
	s.State = Inactive
	s.options = *options
	return s
}

//goland:noinspection GoUnusedParameter
func (s *MqttSensor) Update(client mqtt.Client, msg mqtt.Message) {

	value, err := strconv.ParseFloat(string(msg.Payload()), 8)

	if err == nil {
		s.SetValue(value)
	}
}
