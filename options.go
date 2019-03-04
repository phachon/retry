package retry

import "time"

// set option
type Option struct {
	Set func(*Options)
}

// retry option
type Options struct {

	// retry count
	RetryCount int

	// retry interval time
	IntervalTime time.Duration

	// retry timeout
	RetryTimeout time.Duration
}

// retry count >= 0
func SetRetryCount(retryCount int) Option {
	return Option{
		Set: func(options *Options) {
			options.RetryCount = retryCount
		},
	}
}

// interval time >= 0
func SetIntervalTime(intervalTime time.Duration) Option {
	return Option{
		Set: func(options *Options) {
			options.IntervalTime = intervalTime
		},
	}
}

// interval time >= 0, interval time = 0, do not check timeout
func SetRetryTimeout(retryTimeout time.Duration) Option {
	return Option{
		Set: func(options *Options) {
			options.RetryTimeout = retryTimeout
		},
	}
}

// set retry count and retry timeout
func SetRetryCountAndTimeout(retryCount int, retryTimeout time.Duration) Option {
	return Option{
		Set: func(options *Options) {
			options.RetryCount = retryCount
			options.RetryTimeout = retryTimeout
		},
	}
}

// set retry count and interval time
func SetRetryCountAndIntervalTime(retryCount int, intervalTime time.Duration) Option {
	return Option{
		Set: func(options *Options) {
			options.RetryCount = retryCount
			options.IntervalTime = intervalTime
		},
	}
}
