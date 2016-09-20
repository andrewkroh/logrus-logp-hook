package logp_test

import (
	"io/ioutil"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/andrewkroh/logrus-logp-hook"
	beatslogp "github.com/elastic/beats/libbeat/logp"
)

func TestLogging(t *testing.T) {
	if testing.Verbose() {
		beatslogp.LogInit(beatslogp.LOG_DEBUG, "", false, true, []string{"mypackage"})
	}

	log := logrus.New()
	log.Out = ioutil.Discard
	log.Level = logrus.DebugLevel

	hook, err := logp.NewHook()
	if err != nil {
		t.Errorf("failed to create hook: %v", err)
	}
	log.Hooks.Add(hook)

	log.Info("Congratulations!")
	log.WithField("variable", 10.1).Warn("out of range")
	log.WithField("package", "mypackage").Debug("hello")
	log.Debug("not to be logged")
	beatslogp.Info("Done")
}

func ExampleAddHook() {
	beatslogp.LogInit(beatslogp.LOG_DEBUG, "", false, true, []string{"mypackage"})

	logrus.SetOutput(ioutil.Discard)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true, DisableTimestamp: true})
	logpHook, _ := logp.NewHook()
	logrus.AddHook(logpHook)

	ctx := logrus.WithFields(logrus.Fields{
		"package": "mypackage", // This is your logp debug selector.
	})

	ctx.Debug("hello")

	// 2016/09/20 17:09:19.507085 logp.go:37: DBG  level=debug msg=hello package=mypackage
}
