# Logp hook for logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" /> [![Build Status](https://travis-ci.org/andrewkroh/logrus-logp-hook.svg?branch=master)](https://travis-ci.org/andrewkroh/logrus-logp-hook)

Use this hook to send the logs the [logp logger](https://github.com/elastic/beats/tree/master/libbeat/logp) that used in Elastic Beats.

## Usage

```go
package main

import (
        "io/ioutil"

        "github.com/Sirupsen/logrus"
        "github.com/andrewkroh/logrus-logp-hook"
)

func main() {
        logrus.SetOutput(ioutil.Discard)
        logrus.SetLevel(logrus.DebugLevel)
        logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true, DisableTimestamp: true})
        logpHook, _ := logp.NewHook()
        logrus.AddHook(logpHook)

        ctx := logrus.WithFields(logrus.Fields{
                "package": "mypackage", // This is your logp debug selector.
        })

        ctx.Debug("hello")
}
```

This is how it will look like:

```
2016/09/20 17:09:19.507085 logp.go:37: DBG  level=debug msg=hello package=mypackage
```
