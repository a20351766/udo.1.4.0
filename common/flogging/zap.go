/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flogging

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

// NewZapLogger creates a zap logger around a new zap.Core. The core will use
// the provided encoder and sinks and a level enabler that is associated with
// the provided logger name. The logger that is returned will be named the same
// as the logger.
func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
	return zap.New(
		core,
		append([]zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		}, options...)...,
	)
}

// NewGRPCLogger creates a grpc.Logger that delegates to a zap.Logger.
func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
	l = l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)
	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
}

// NewUDOLogger creates a logger that delegates to the zap.SugaredLogger.
func NewUDOLogger(l *zap.Logger, options ...zap.Option) *UDOLogger {
	return &UDOLogger{
		s: l.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar(),
	}
}

// A UDOLogger is an adapter around a zap.SugaredLogger that provides
// structured logging capabilities while preserving much of the legacy logging
// behavior.
//
// The most significant difference between the UDOLogger and the
// zap.SugaredLogger is that methods without a formatting suffix (f or w) build
// the log entry message with fmt.Sprintln instead of fmt.Sprint. Without this
// change, arguments are not separated by spaces.
type UDOLogger struct{ s *zap.SugaredLogger }

func (f *UDOLogger) DPanic(args ...interface{})                    { f.s.DPanicf(formatArgs(args)) }
func (f *UDOLogger) DPanicf(template string, args ...interface{})  { f.s.DPanicf(template, args...) }
func (f *UDOLogger) DPanicw(msg string, kvPairs ...interface{})    { f.s.DPanicw(msg, kvPairs...) }
func (f *UDOLogger) Debug(args ...interface{})                     { f.s.Debugf(formatArgs(args)) }
func (f *UDOLogger) Debugf(template string, args ...interface{})   { f.s.Debugf(template, args...) }
func (f *UDOLogger) Debugw(msg string, kvPairs ...interface{})     { f.s.Debugw(msg, kvPairs...) }
func (f *UDOLogger) Error(args ...interface{})                     { f.s.Errorf(formatArgs(args)) }
func (f *UDOLogger) Errorf(template string, args ...interface{})   { f.s.Errorf(template, args...) }
func (f *UDOLogger) Errorw(msg string, kvPairs ...interface{})     { f.s.Errorw(msg, kvPairs...) }
func (f *UDOLogger) Fatal(args ...interface{})                     { f.s.Fatalf(formatArgs(args)) }
func (f *UDOLogger) Fatalf(template string, args ...interface{})   { f.s.Fatalf(template, args...) }
func (f *UDOLogger) Fatalw(msg string, kvPairs ...interface{})     { f.s.Fatalw(msg, kvPairs...) }
func (f *UDOLogger) Info(args ...interface{})                      { f.s.Infof(formatArgs(args)) }
func (f *UDOLogger) Infof(template string, args ...interface{})    { f.s.Infof(template, args...) }
func (f *UDOLogger) Infow(msg string, kvPairs ...interface{})      { f.s.Infow(msg, kvPairs...) }
func (f *UDOLogger) Panic(args ...interface{})                     { f.s.Panicf(formatArgs(args)) }
func (f *UDOLogger) Panicf(template string, args ...interface{})   { f.s.Panicf(template, args...) }
func (f *UDOLogger) Panicw(msg string, kvPairs ...interface{})     { f.s.Panicw(msg, kvPairs...) }
func (f *UDOLogger) Warn(args ...interface{})                      { f.s.Warnf(formatArgs(args)) }
func (f *UDOLogger) Warnf(template string, args ...interface{})    { f.s.Warnf(template, args...) }
func (f *UDOLogger) Warnw(msg string, kvPairs ...interface{})      { f.s.Warnw(msg, kvPairs...) }
func (f *UDOLogger) Warning(args ...interface{})                   { f.s.Warnf(formatArgs(args)) }
func (f *UDOLogger) Warningf(template string, args ...interface{}) { f.s.Warnf(template, args...) }

// for backwards compatibility
func (f *UDOLogger) Critical(args ...interface{})                   { f.s.Errorf(formatArgs(args)) }
func (f *UDOLogger) Criticalf(template string, args ...interface{}) { f.s.Errorf(template, args...) }
func (f *UDOLogger) Notice(args ...interface{})                     { f.s.Infof(formatArgs(args)) }
func (f *UDOLogger) Noticef(template string, args ...interface{})   { f.s.Infof(template, args...) }

func (f *UDOLogger) Named(name string) *UDOLogger { return &UDOLogger{s: f.s.Named(name)} }
func (f *UDOLogger) Sync() error                     { return f.s.Sync() }
func (f *UDOLogger) Zap() *zap.Logger                { return f.s.Desugar() }

func (f *UDOLogger) IsEnabledFor(level zapcore.Level) bool {
	return f.s.Desugar().Core().Enabled(level)
}

func (f *UDOLogger) With(args ...interface{}) *UDOLogger {
	return &UDOLogger{s: f.s.With(args...)}
}

func (f *UDOLogger) WithOptions(opts ...zap.Option) *UDOLogger {
	l := f.s.Desugar().WithOptions(opts...)
	return &UDOLogger{s: l.Sugar()}
}

func formatArgs(args []interface{}) string { return strings.TrimSuffix(fmt.Sprintln(args...), "\n") }
