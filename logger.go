package logger

import (
	"fmt"

	"github.com/sasha-s/go-deadlock"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// DefaultFieldsHook implements the logrus Hook interface
// We use it to add a set of default fields to all our logs
type DefaultFieldsHook struct {
	defaults log.Fields
	mu       *deadlock.Mutex
}

// Levels implements the Levels methods for the logrus Hook interface
func (h *DefaultFieldsHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire implements the Levels methods for the logrus Hook interface
// This is where the defaults are added to the logrus Entry.Data map
func (h *DefaultFieldsHook) Fire(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	for k, v := range h.defaults {
		e.Data[k] = v
	}
	return nil
}

// FieldHelper encompasses all field manipulation methods.
type FieldHelper interface {
	Fields() log.Fields

	AddField(key string, value interface{})
	AddFields(fields log.Fields)

	RemoveFieldsByKey(key ...string)
	RemoveFields(fields log.Fields)
	ResetFields()
}

// DefaultHelper encompasses all default manipulation methods.
type DefaultHelper interface {
	Defaults() log.Fields
	SetDefaults(defaults log.Fields)
}

// ErrorHelper encompasses all error helper methods.
type ErrorHelper interface {
	WrapAndPrintWithError(err error, format string, args ...interface{}) error
	PrintWithError(err error, format string, args ...interface{})
}

// Helper interface is a set of wrappers that the package
// provides over logrus.
type Helper interface {
	FieldHelper
	DefaultHelper
	ErrorHelper
}

// Cloner interface provides a shallow copy method
type Cloner interface {
	Clone() Logger
}

// Logger interface wraps the logrus interface and some additional helper functions
type Logger interface {
	Logrus
	Helper
	Cloner
}

type logger struct {
	*log.Entry
	fields   log.Fields
	defaults log.Fields

	mu *deadlock.RWMutex
}

// New returns a new instance of the logger struct
func New() *logger {
	return &logger{
		log.NewEntry(log.New()),
		make(log.Fields),
		make(log.Fields),
		&deadlock.RWMutex{},
	}
}

func (l *logger) updateFields() {
	l.mu.Lock()
	for k, v := range l.fields {
		l.Entry.Data[k] = v
	}
	l.mu.Unlock()
}

// Defaults returns the set default fields map
func (l *logger) Defaults() log.Fields {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.defaults
}

// Fields returns the variable fields map
func (l *logger) Fields() log.Fields {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.fields
}

// SetDefaults sets the default fields
func (l *logger) SetDefaults(defaults log.Fields) {
	l.defaults = defaults
	l.Logger.AddHook(&DefaultFieldsHook{l.defaults, &deadlock.Mutex{}})
}

// AddField adds a single key value pair to the fields map
func (l *logger) AddField(key string, value interface{}) {
	l.mu.Lock()
	l.fields[key] = value
	l.Data[key] = value
	l.mu.Unlock()
}

// AddFields adds one or more kv pairs to the fields map
func (l *logger) AddFields(fields log.Fields) {
	defer l.updateFields()
	l.mu.Lock()
	for k, v := range fields {
		l.fields[k] = v
		l.Data[k] = v
	}
	l.mu.Unlock()
}

// RemoveFieldsByKey removes one or more entries from the fields map by their key
func (l *logger) RemoveFieldsByKey(keys ...string) {
	defer l.updateFields()
	l.mu.Lock()
	for _, k := range keys {
		delete(l.fields, k)
		delete(l.Data, k)
	}
	l.mu.Unlock()
}

// RemoveFieldsByKey removes one or more entries from the fields map using the same object used to set the kv pairs
func (l *logger) RemoveFields(fields log.Fields) {
	defer l.updateFields()
	l.mu.Lock()
	for k, _ := range fields {
		l.RemoveFieldsByKey(k)
	}
	l.mu.Unlock()
}

// ResetFields resets all the fields by clearing the fields map
func (l *logger) ResetFields() {
	defer l.updateFields()
	l.mu.Lock()
	l.Data = make(map[string]interface{})
	l.fields = make(map[string]interface{})
	l.defaults = make(map[string]interface{})
	l.SetDefaults(l.defaults)
	l.mu.Unlock()
}

// WrapAndPrintWithError is a combination function that helps us log the error as well as wrap it with the custom error
// message
func (l *logger) WrapAndPrintWithError(err error, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	err = errors.Wrap(err, msg)
	l.WithError(err).Error(msg)
	return err
}

// PrintWithError is a wrapper that helps reduce the verbosity of the WithError logrus method
func (l *logger) PrintWithError(err error, format string, args ...interface{}) {
	l.WithError(err).Error(fmt.Sprintf(format, args...))
}
