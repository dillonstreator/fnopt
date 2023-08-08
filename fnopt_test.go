package fnopt

import (
	"errors"
	"testing"
)

type TestConfig struct {
	count      int
	doThing    bool
	otherThing bool
}

func WithCount(count int) OptFn[TestConfig] {
	return func(cfg *TestConfig) {
		cfg.count = count
	}
}

func WithDoThing(doThing bool) OptFn[TestConfig] {
	return func(cfg *TestConfig) {
		cfg.doThing = doThing
	}
}

func WithOtherThing(otherThing bool) OptFn[TestConfig] {
	return func(cfg *TestConfig) {
		cfg.otherThing = otherThing
	}
}

type TestConfigE struct {
	count      int
	doThing    bool
	otherThing bool
}

func WithCountE(count int) OptFnE[TestConfigE] {
	return func(cfg *TestConfigE) error {
		cfg.count = count
		return nil
	}
}

func WithDoThingE(doThing bool) OptFnE[TestConfigE] {
	return func(cfg *TestConfigE) error {
		cfg.doThing = doThing
		return nil
	}
}

func WithOtherThingE(otherThing bool) OptFnE[TestConfigE] {
	return func(cfg *TestConfigE) error {
		cfg.otherThing = otherThing
		return nil
	}
}

var errWith = errors.New("error")

func WithError() OptFnE[TestConfigE] {
	return func(cfg *TestConfigE) error {
		return errWith
	}
}

func TestNew(t *testing.T) {
	cfg := New(WithCount(10), WithDoThing(true))

	if cfg.count != 10 {
		t.Errorf("expected count of 10 but got %d", cfg.count)
	}
	if !cfg.doThing {
		t.Error("expected doThing to be true but got false")
	}
	if cfg.otherThing {
		t.Error("expected otherThing to be false but got true")
	}
}

func TestFrom(t *testing.T) {
	cfg := &TestConfig{
		count:      5,
		doThing:    true,
		otherThing: true,
	}

	From(cfg, WithCount(10), WithDoThing(false))

	if cfg.count != 10 {
		t.Errorf("expected count of 10 but got %d", cfg.count)
	}
	if cfg.doThing {
		t.Error("expected doThing to be false but got true")
	}
	if !cfg.otherThing {
		t.Error("expected otherThing to be true but got false")
	}
}

func TestNewE(t *testing.T) {
	cfg, err := NewE(WithCountE(10), WithDoThingE(true))
	if err != nil {
		t.Errorf("unexpected non-nil err %v", err)
	}

	if cfg.count != 10 {
		t.Errorf("expected count of 10 but got %d", cfg.count)
	}
	if !cfg.doThing {
		t.Error("expected doThing to be true but got false")
	}
	if cfg.otherThing {
		t.Error("expected otherThing to be false but got true")
	}
}

func TestNewE_error(t *testing.T) {
	cfg, err := NewE(WithCountE(10), WithDoThingE(true), WithError())
	if !errors.Is(err, errWith) {
		t.Errorf("unexpected err value %v", err)
	}
	if cfg != nil {
		t.Errorf("unexpected non-nil cfg value %v", cfg)
	}
}

func TestFromE(t *testing.T) {
	cfg := &TestConfigE{
		count:      5,
		doThing:    true,
		otherThing: true,
	}

	err := FromE(cfg, WithCountE(10), WithDoThingE(false))
	if err != nil {
		t.Errorf("unexpected non-nil err %v", err)
	}

	if cfg.count != 10 {
		t.Errorf("expected count of 10 but got %d", cfg.count)
	}
	if cfg.doThing {
		t.Error("expected doThing to be false but got true")
	}
	if !cfg.otherThing {
		t.Error("expected otherThing to be true but got false")
	}
}

func TestFromE_error(t *testing.T) {
	cfg := &TestConfigE{
		count:      5,
		doThing:    true,
		otherThing: true,
	}

	err := FromE(cfg, WithCountE(10), WithDoThingE(true), WithError())
	if !errors.Is(err, errWith) {
		t.Errorf("unexpected err value %v", err)
	}
}
