// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("go.opentelemetry.io/collector/scraper/scraperhelper")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("go.opentelemetry.io/collector/scraper/scraperhelper")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                      metric.Meter
	mu                         sync.Mutex
	registrations              []metric.Registration
	ScraperErroredLogRecords   metric.Int64Counter
	ScraperErroredMetricPoints metric.Int64Counter
	ScraperScrapedLogRecords   metric.Int64Counter
	ScraperScrapedMetricPoints metric.Int64Counter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// Shutdown unregister all registered callbacks for async instruments.
func (builder *TelemetryBuilder) Shutdown() {
	builder.mu.Lock()
	defer builder.mu.Unlock()
	for _, reg := range builder.registrations {
		reg.Unregister()
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.ScraperErroredLogRecords, err = builder.meter.Int64Counter(
		"otelcol_scraper_errored_log_records",
		metric.WithDescription("Number of log records that were unable to be scraped. [alpha]"),
		metric.WithUnit("{datapoints}"),
	)
	errs = errors.Join(errs, err)
	builder.ScraperErroredMetricPoints, err = builder.meter.Int64Counter(
		"otelcol_scraper_errored_metric_points",
		metric.WithDescription("Number of metric points that were unable to be scraped. [alpha]"),
		metric.WithUnit("{datapoints}"),
	)
	errs = errors.Join(errs, err)
	builder.ScraperScrapedLogRecords, err = builder.meter.Int64Counter(
		"otelcol_scraper_scraped_log_records",
		metric.WithDescription("Number of log records successfully scraped. [alpha]"),
		metric.WithUnit("{datapoints}"),
	)
	errs = errors.Join(errs, err)
	builder.ScraperScrapedMetricPoints, err = builder.meter.Int64Counter(
		"otelcol_scraper_scraped_metric_points",
		metric.WithDescription("Number of metric points successfully scraped. [alpha]"),
		metric.WithUnit("{datapoints}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
