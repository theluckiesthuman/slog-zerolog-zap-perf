package main

import (
	"bytes"
	"testing"

	"log/slog"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// BenchmarkZerolog measures the performance of zerolog
func BenchmarkZerolog(b *testing.B) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf).With().Timestamp().Logger()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Int("iteration", i).
			Str("service", "benchmark").
			Str("environment", "production").
			Int("user_id", 123456).
			Str("request_id", "abc-def-ghi").
			Float64("response_time", 0.234).
			Bool("cache_hit", true).
			Str("method", "POST").
			Str("path", "/api/v1/users").
			Int("status_code", 200).
			Int64("content_length", 1234567).
			Str("user_agent", "Mozilla/5.0").
			Str("client_ip", "192.168.1.1").
			Str("host", "api.example.com").
			Str("protocol", "HTTP/2.0").
			Msg("test message")
	}
}

// BenchmarkZap measures the performance of zap
func BenchmarkZap(b *testing.B) {
	var buf bytes.Buffer
	enc := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enc),
		zapcore.AddSync(&buf),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test message",
			zap.Int("iteration", i),
			zap.String("service", "benchmark"),
			zap.String("environment", "production"),
			zap.Int("user_id", 123456),
			zap.String("request_id", "abc-def-ghi"),
			zap.Float64("response_time", 0.234),
			zap.Bool("cache_hit", true),
			zap.String("method", "POST"),
			zap.String("path", "/api/v1/users"),
			zap.Int("status_code", 200),
			zap.Int64("content_length", 1234567),
			zap.String("user_agent", "Mozilla/5.0"),
			zap.String("client_ip", "192.168.1.1"),
			zap.String("host", "api.example.com"),
			zap.String("protocol", "HTTP/2.0"),
		)
	}
}

// BenchmarkSlog measures the performance of slog
func BenchmarkSlog(b *testing.B) {
	var buf bytes.Buffer
	opts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewTextHandler(&buf, &opts))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test message",
			"iteration", i,
			"service", "benchmark",
			"environment", "production",
			"user_id", 123456,
			"request_id", "abc-def-ghi",
			"response_time", 0.234,
			"cache_hit", true,
			"method", "POST",
			"path", "/api/v1/users",
			"status_code", 200,
			"content_length", 1234567,
			"user_agent", "Mozilla/5.0",
			"client_ip", "192.168.1.1",
			"host", "api.example.com",
			"protocol", "HTTP/2.0",
		)
	}
}
