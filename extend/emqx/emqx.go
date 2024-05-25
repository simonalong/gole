package emqx

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/logger"
	t0 "time"
)

func init() {
	config.LoadConfig()

	if config.ExistConfigFile() && config.GetValueBoolDefault("gole.emqx.enable", false) {
		err := config.GetValueObject("gole.emqx", &config.EmqxCfg)
		if err != nil {
			logger.Warn("读取emqx配置异常, %v", err.Error())
			return
		}

		mqtt.DEBUG = emqxLogger{"DEBUG"}
		mqtt.WARN = emqxLogger{"WARN"}
		mqtt.CRITICAL = emqxLogger{"CRITICAL"}
		mqtt.ERROR = emqxLogger{"ERROR"}
	}
}

type emqxLogger struct {
	Level string
}

func (log emqxLogger) Println(v ...interface{}) {
	switch log.Level {
	case "DEBUG":
		logger.DebugDirect(v...)
	case "WARN":
		logger.WarnDirect(v...)
	case "CRITICAL":
		logger.ErrorDirect(v...)
	case "ERROR":
		logger.ErrorDirect(v...)
	}
}
func (log emqxLogger) Printf(format string, v ...interface{}) {
	switch log.Level {
	case "DEBUG":
		logger.Debug(format, v...)
	case "WARN":
		logger.Warn(format, v...)
	case "CRITICAL":
		logger.Error(format, v...)
	case "ERROR":
		logger.Error(format, v...)
	}
}

func NewEmqxClient() (mqtt.Client, error) {
	if !config.GetValueBoolDefault("gole.emqx.enable", false) {
		logger.Error("emqx没有配置，请先配置")
		return nil, errors.New("emqx没有配置，请先配置")
	}

	_emqxClient := mqtt.NewClient(localEmqxOptions())
	if token := _emqxClient.Connect(); token.Wait() && token.Error() != nil {
		logger.Error("链接emqx client失败, %v", token.Error().Error())
		return nil, token.Error()
	}
	return _emqxClient, nil
}

func localEmqxOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	for _, server := range config.EmqxCfg.Servers {
		opts.AddBroker(server)
	}

	if config.EmqxCfg.ClientId != "" {
		opts.SetClientID(config.EmqxCfg.ClientId)
	}

	if config.EmqxCfg.Username != "" {
		opts.SetUsername(config.EmqxCfg.Username)
	}

	if config.EmqxCfg.Password != "" {
		opts.SetPassword(config.EmqxCfg.Password)
	}

	if config.EmqxCfg.CleanSession != true && config.GetValueString("gole.emqx.clean-session") != "" {
		opts.SetCleanSession(config.EmqxCfg.CleanSession)
	}

	if config.EmqxCfg.Order != true && config.GetValueString("gole.emqx.order") != "" {
		opts.SetOrderMatters(config.EmqxCfg.Order)
	}

	if config.EmqxCfg.WillEnabled != false && config.GetValueString("gole.emqx.will-enabled") != "" {
		opts.WillEnabled = config.EmqxCfg.WillEnabled
	}

	if config.EmqxCfg.WillTopic != "" {
		opts.WillTopic = config.EmqxCfg.WillTopic
	}

	if config.EmqxCfg.WillQos != 0 {
		opts.WillQos = config.EmqxCfg.WillQos
	}

	if config.EmqxCfg.WillRetained != false && config.GetValueString("gole.emqx.will-retained") != "" {
		opts.WillRetained = config.EmqxCfg.WillRetained
	}

	if config.EmqxCfg.ProtocolVersion != 0 {
		opts.ProtocolVersion = config.EmqxCfg.ProtocolVersion
	}

	if config.EmqxCfg.KeepAlive != 30 && config.GetValueString("gole.emqx.keep-alive") != "" {
		opts.KeepAlive = config.EmqxCfg.KeepAlive
	}

	if config.EmqxCfg.PingTimeout != "10s" && config.GetValueString("gole.emqx.ping-timeout") != "" {
		t, err := t0.ParseDuration(config.EmqxCfg.PingTimeout)
		if err != nil {
			logger.Warn("读取配置【gole.emqx.ping-timeout】异常：%v", err.Error())
		} else {
			opts.PingTimeout = t
		}
	}

	if config.EmqxCfg.ConnectTimeout != "30s" && config.GetValueString("gole.emqx.connect-timeout") != "" {
		t, err := t0.ParseDuration(config.EmqxCfg.ConnectTimeout)
		if err != nil {
			logger.Warn("读取配置【gole.emqx.connect-timeout】异常：%v", err.Error())
		} else {
			opts.PingTimeout = t
		}
	}

	if config.EmqxCfg.MaxReconnectInterval != "10m" && config.GetValueString("gole.emqx.max-reconnect-interval") != "" {
		t, err := t0.ParseDuration(config.EmqxCfg.MaxReconnectInterval)
		if err != nil {
			logger.Warn("读取配置【gole.emqx.max-reconnect-interval】异常：%v", err.Error())
		} else {
			opts.MaxReconnectInterval = t
		}
	}

	if config.EmqxCfg.AutoReconnect != true && config.GetValueString("gole.emqx.auto-reconnect") != "" {
		opts.AutoReconnect = config.EmqxCfg.AutoReconnect
	}

	if config.EmqxCfg.ConnectRetryInterval != "30s" && config.GetValueString("gole.emqx.connect-retry-interval") != "" {
		t, err := t0.ParseDuration(config.EmqxCfg.ConnectRetryInterval)
		if err != nil {
			logger.Warn("读取配置【gole.emqx.connect-retry-interval】异常：%v", err.Error())
		} else {
			opts.ConnectRetryInterval = t
		}
	}

	if config.EmqxCfg.ConnectRetry != false && config.GetValueString("gole.emqx.connect-retry") != "" {
		opts.ConnectRetry = config.EmqxCfg.ConnectRetry
	}

	if config.EmqxCfg.WriteTimeout != "0" && config.GetValueString("gole.emqx.write-timeout") != "" {
		t, err := t0.ParseDuration(config.EmqxCfg.WriteTimeout)
		if err != nil {
			logger.Warn("读取配置【gole.emqx.write-timeout】异常：%v", err.Error())
		} else {
			opts.WriteTimeout = t
		}
	}

	if config.EmqxCfg.ResumeSubs != false && config.GetValueString("gole.emqx.resume-subs") != "" {
		opts.ResumeSubs = config.EmqxCfg.ResumeSubs
	}

	if config.EmqxCfg.MaxResumePubInFlight != 0 {
		opts.MaxResumePubInFlight = config.EmqxCfg.MaxResumePubInFlight
	}

	if config.EmqxCfg.AutoAckDisabled != false && config.GetValueString("gole.emqx.auto-ack-disabled") != "" {
		opts.AutoAckDisabled = config.EmqxCfg.AutoAckDisabled
	}

	return opts
}
