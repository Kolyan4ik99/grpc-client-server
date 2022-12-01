package config

import (
	"flag"
	"time"
)

type Client struct {
	UserName string
	Password string

	DialInterval time.Duration
	DialDeadline time.Time
	Buffer       *buffer
	ServerURL    string
}

func NewClientConfig() (*Client, error) {
	userName := flag.String("user_name", "root", "Credential name")
	password := flag.String("user_password", "admin123", "Credential password")

	dialIntervalStr := flag.String("dial_interval", "2s", "Communication interval")
	dialDeadlineStr := flag.String("dial_deadline", "47s", "Deadline timer limit")

	bufferSize := flag.Int64("buffer_size", 10, "Max buffer size")
	bufferThresholdStr := flag.String("buffer_threshold", "19s", "Max threshold before flush and clear buffer")

	flag.Parse()

	dialInterval, err := time.ParseDuration(*dialIntervalStr)
	if err != nil {
		return nil, err
	}
	dialDeadline, err := time.ParseDuration(*dialDeadlineStr)
	if err != nil {
		return nil, err
	}
	bufferThreshold, err := time.ParseDuration(*bufferThresholdStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		UserName: *userName,
		Password: *password,

		DialInterval: dialInterval,
		DialDeadline: time.Now().Add(dialDeadline),
		Buffer: &buffer{
			Size:      *bufferSize,
			Threshold: bufferThreshold,
		},
		ServerURL: "127.0.0.1:5053",
	}, nil
}

type buffer struct {
	Size      int64
	Threshold time.Duration
}
