package retry

import (
	"testing"
	"time"
)

func TestSetRetryCount(t *testing.T) {
	var options = &Options{}
	SetRetryCount(10).Set(options)
	if options.RetryCount != 10 {
		t.Fail()
	}
}

func TestSetIntervalTime(t *testing.T) {
	var options = &Options{}
	SetIntervalTime(time.Second * 10).Set(options)
	if options.IntervalTime != time.Second * 10 {
		t.Fail()
	}
}

func TestSetRetryTimeout(t *testing.T) {
	var options = &Options{}
	SetRetryTimeout(time.Second * 5).Set(options)
	if options.RetryTimeout != time.Second * 5 {
		t.Fail()
	}
}

func TestSetRetryCountAndIntervalTime(t *testing.T) {
	var options = &Options{}
	SetRetryCountAndIntervalTime(10, time.Second * 5).Set(options)
	if options.RetryCount != 10 {
		t.Fail()
	}
	if options.IntervalTime != time.Second * 5 {
		t.Fail()
	}
}

func TestSetRetryCountAndTimeout(t *testing.T) {
	var options = &Options{}
	SetRetryCountAndTimeout(10, time.Second * 10).Set(options)
	if options.RetryCount != 10 {
		t.Fail()
	}
	if options.RetryTimeout != time.Second * 10 {
		t.Fail()
	}
}