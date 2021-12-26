package gmetrics

// GaugeOpts is an alias of Opts.
type GaugeOpts Opts

// IGauge gauge vec.
type IGauge interface {
	// Set sets the Gauge to an arbitrary value.
	Set(v float64, labels ...string)
	// Inc increments the Gauge by 1. Use Add to increment it by arbitrary
	// values.
	Inc(labels ...string)
	// Add adds the given value to the Gauge. (The value can be negative,
	// resulting in a decrease of the Gauge.)
	Add(v float64, labels ...string)
}
