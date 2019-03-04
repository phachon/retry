package retry

import (
	"errors"
	"fmt"
	"time"
)

var (
	// version
	Version = "v0.0.1"
)

var (

	// retry max limit
	errMaxRetryLimit = errors.New("retry count limit max")

	// retry timeout
	errRetryTimeout = errors.New("retry timeout")
)

func Retry(call func() error, options ...Option) error {

	ops := &Options{}
	for _, option := range options {
		option.Set(ops)
	}

	if ops.RetryTimeout == 0 {
		return simpleRetry(call, ops)
	}else {
		return timeoutRetry(call, ops)
	}
}

// simple retry
func simpleRetry(call func() error, options *Options) (err error) {

	var tryCount int
	var tryMaxCount = options.RetryCount + 1
	for tryCount = 1; tryCount <= tryMaxCount; tryCount++ {
		err = call()
		if err == nil {
			return
		}
		if options.IntervalTime > 0 {
			time.Sleep(options.IntervalTime)
		}
	}
	return
}

// retry check timeout
func timeoutRetry(call func() error, options *Options) (err error) {

	var tryCount int
	var tryMaxCount = options.RetryCount + 1
	var resError error
	for tryCount = 1; tryCount <= tryMaxCount; tryCount++ {
		errChan := make(chan error)
		go func() {
			defer func() {
				e := recover()
				if e != nil {
					errChan <- errors.New(fmt.Sprintf("%v", e))
				}
			}()
			errChan <- call()
		}()

		// retry timeout
		select {
		case resError = <-errChan:
		case <-time.After(options.RetryTimeout):
			resError = errRetryTimeout
		}
		if resError == nil {
			return nil
		}

		if options.IntervalTime > 0 {
			time.Sleep(options.IntervalTime)
		}
	}

	return resError
}