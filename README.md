# logger
A wrapper around https://github.com/sirupsen/logrus that allows setting default fields, and supports adding and removing configurable fields.

[![GoDoc](https://godoc.org/github.com/nitishm/logger?status.svg)](https://godoc.org/github.com/nitishm/logger)
# Usage
```go
import (
	"github.com/nitishm/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
  level := "info"
	l := logger.New()
	l.Logger.SetLevel(level)
	l.Logger.SetFormatter(&log.JSONFormatter{})
	l.Logger.SetReportCaller(true)
  ...
}

func() {
  ...
  defer l.RemoveFieldsByKey("label_selector")
  l.AddField("label_selectors", label.String())
  ...
  f := log.Fields{
      "num_bytes":   summary.NumBytes,
      "duration":    summary.Duration,
  }
  defer l.RemoveFields(f)
  l.AddFields(f)
  ...
  err := doSomething(context.TODO(), input)
  if err != nil {
    err = c.l.WrapAndPrintWithError(err, fmt.Sprintf("Failed to CreateOrUpdate Capture instance for ID %s", c.id))
    return err
  }
  ...
}
```

# Supported Methods
```go
// Encapsulates all supported logrus methods
log.Ext1FieldLogger
// Fields returns the set of configured non-default fields
Fields() log.Fields
// Defaults returns the set of default fields 
Defaults() log.Fields
// SetDefaults allows the user to set a default set of fields
// that are included in every log message
SetDefaults(defaults log.Fields)
// AddField allows the user to add a single key/value to the 
// non-default set of fields
AddField(key string, value interface{})
// AddFields allows the user to add multipme key/values to the 
// non-default set of fields
AddFields(fields log.Fields)
// RemoveFieldsByKey allows the user to remove multiple non-default
// fields using one or more key values
RemoveFieldsByKey(key ...string)
// RemoveFields allows the user to remove a set of fields
// This is useful when the user wishes to add a set of fields
// and defer their deletion subsequently after the function returns.
RemoveFields(fields log.Fields)
// ResetFields resets all the fields including the default and non-default
// fields.
ResetFields()
// WrapAndPrintWithError is a helper that allows the user to wrap an error
// using the errors package and log that to the output 
WrapAndPrintWithError(err error, format string, args ...interface{}) error
// PrintWithError allows a user to log the error output along with the 
// error value to the output
PrintWithError(err error, format string, args ...interface{})
```
