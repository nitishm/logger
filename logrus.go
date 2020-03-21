package logger

import log "github.com/sirupsen/logrus"

// Logrus interface is an alias for the upstream logrus interface
// This makes logrus methods concurrency safe.
type Logrus log.Ext1FieldLogger

// Add a single field to the Entry
func (l *logger) WithField(key string, value interface{}) *log.Entry {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Entry.WithField(key, value)
}

// Add a map of fields to the Entry.
func (l *logger) WithFields(fields log.Fields) *log.Entry {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Entry.WithFields(fields)
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (l *logger) WithError(err error) *log.Entry {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Entry.WithError(err)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Debugf(format, args...)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Infof(format, args...)
}
func (l *logger) Printf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Printf(format, args...)
}
func (l *logger) Warnf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Warnf(format, args...)
}
func (l *logger) Warningf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Warningf(format, args...)
}
func (l *logger) Errorf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Errorf(format, args...)
}
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Fatalf(format, args...)
}
func (l *logger) Panicf(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Panicf(format, args...)
}
func (l *logger) Debug(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Debug(args...)
}
func (l *logger) Info(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Info(args...)
}
func (l *logger) Print(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Print(args...)
}
func (l *logger) Warn(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Warn(args...)
}
func (l *logger) Warning(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Warning(args...)
}
func (l *logger) Error(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Error(args...)
}
func (l *logger) Fatal(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Fatal(args...)
}
func (l *logger) Panic(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Panic(args...)
}
func (l *logger) Debugln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Debugln(args...)
}
func (l *logger) Infoln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Infoln(args...)
}
func (l *logger) Println(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Println(args...)
}
func (l *logger) Warnln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Warnln(args...)
}
func (l *logger) Warningln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Warningln(args...)
}
func (l *logger) Errorln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Errorln(args...)
}
func (l *logger) Fatalln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Fatalln(args...)
}
func (l *logger) Panicln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Panicln(args...)
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Tracef(format, args...)
}
func (l *logger) Trace(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Trace(args...)
}
func (l *logger) Traceln(args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.Entry.Traceln(args...)
}

// Clone creates a shallow copy of the logger
// and returns a new instance.
// NOTE: That the logrus Entry ptr is not passed
// along from the original Logger.
func (l *logger) Clone() Logger {
	var clone *logger = New()
	clone.SetDefaults(l.Defaults())
	clone.AddFields(l.Fields())
	return clone
}
