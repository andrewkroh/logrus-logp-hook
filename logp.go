package logp

import (
	"github.com/Sirupsen/logrus"
	"github.com/elastic/beats/libbeat/logp"
)

// Hook is a Logrus hook to send logs to Elastic Beats logp logger.
type Hook struct{}

// NewHook creates and returns a new logrus Hook to send logs to the logp logger.
func NewHook() (*Hook, error) {
	return &Hook{}, nil
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		logp.Err("failed to convert logrus entry to string. entry=%+v, err=%v", *entry, err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		logp.WTF(line)
	case logrus.FatalLevel:
		logp.Critical(line)
	case logrus.ErrorLevel:
		logp.Err(line)
	case logrus.WarnLevel:
		logp.Warn(line)
	case logrus.InfoLevel:
		logp.Info(line)
	case logrus.DebugLevel:
		if pkg, found := entry.Data["package"]; found {
			pkgStr, ok := pkg.(string)
			if ok {
				logp.Debug(pkgStr, line)
				break
			}
		}
		logp.Debug("logrus", line)
	}

	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
