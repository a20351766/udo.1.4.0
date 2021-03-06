/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flogging_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hyperledger/udo/common/flogging"
	"github.com/hyperledger/udo/common/flogging/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	logging, err := flogging.New(flogging.Config{})
	assert.NoError(t, err)
	assert.Equal(t, zapcore.InfoLevel, logging.DefaultLevel())

	_, err = flogging.New(flogging.Config{
		LogSpec: "::=borken=::",
	})
	assert.EqualError(t, err, "invalid logging specification '::=borken=::': bad segment '=borken='")
}

func TestNewWithEnvironment(t *testing.T) {
	oldSpec, set := os.LookupEnv("UDO_LOGGING_SPEC")
	if set {
		defer os.Setenv("UDO_LOGGING_SPEC", oldSpec)
	}

	os.Setenv("UDO_LOGGING_SPEC", "fatal")
	logging, err := flogging.New(flogging.Config{})
	assert.NoError(t, err)
	assert.Equal(t, zapcore.FatalLevel, logging.DefaultLevel())

	os.Unsetenv("UDO_LOGGING_SPEC")
	logging, err = flogging.New(flogging.Config{})
	assert.NoError(t, err)
	assert.Equal(t, zapcore.InfoLevel, logging.DefaultLevel())
}

//go:generate counterfeiter -o mock/write_syncer.go -fake-name WriteSyncer . writeSyncer
type writeSyncer interface {
	zapcore.WriteSyncer
}

func TestLoggingSetWriter(t *testing.T) {
	ws := &mock.WriteSyncer{}

	logging, err := flogging.New(flogging.Config{})
	assert.NoError(t, err)

	logging.SetWriter(ws)
	logging.Write([]byte("hello"))
	assert.Equal(t, 1, ws.WriteCallCount())
	assert.Equal(t, []byte("hello"), ws.WriteArgsForCall(0))

	err = logging.Sync()
	assert.NoError(t, err)

	ws.SyncReturns(errors.New("welp"))
	err = logging.Sync()
	assert.EqualError(t, err, "welp")
}

func TestNamedLogger(t *testing.T) {
	defer flogging.Reset()
	buf := &bytes.Buffer{}
	flogging.Global.SetWriter(buf)

	t.Run("logger and named (child) logger with different levels", func(t *testing.T) {
		defer buf.Reset()
		logger := flogging.MustGetLogger("eugene")
		logger2 := logger.Named("george")
		flogging.ActivateSpec("eugene=info:eugene.george=error")

		logger.Info("from eugene")
		logger2.Info("from george")
		assert.Contains(t, buf.String(), "from eugene")
		assert.NotContains(t, buf.String(), "from george")
	})

	t.Run("named logger where parent logger isn't enabled", func(t *testing.T) {
		logger := flogging.MustGetLogger("foo")
		logger2 := logger.Named("bar")
		flogging.ActivateSpec("foo=fatal:foo.bar=error")
		logger.Error("from foo")
		logger2.Error("from bar")
		assert.NotContains(t, buf.String(), "from foo")
		assert.Contains(t, buf.String(), "from bar")
	})
}

func TestInvalidLoggerName(t *testing.T) {
	names := []string{"test*", ".test", "test.", ".", ""}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			msg := fmt.Sprintf("invalid logger name: %s", name)
			assert.PanicsWithValue(t, msg, func() { flogging.MustGetLogger(name) })
		})
	}
}
