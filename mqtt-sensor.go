package sensor

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"time"
)

type MqttOptions struct {
	Options
	ReportToAddress  string
	ReportToQos      byte
	ReportToRetained bool
}

type MqttSensor struct {
	Sensor
	options        MqttOptions
	ReportToClient mqtt.Client
}

func NewMqttSensor(options *MqttOptions) *MqttSensor {
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

func (s *MqttSensor) ReportTo() {
	s.ReportToClient.Publish(s.options.ReportToAddress, s.options.ReportToQos, s.options.ReportToRetained, s.value)
}

func (s *MqttSensor) SetValue(value float64) {
	s.UpdatedTime = time.Now()
	s.State = Active
	s.value = value

	if s.Propagate != nil {
		s.Propagate(s)
	}

	if s.ReportToClient != nil {
		s.ReportTo()
	}

}
