package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/equinix-labs/otel-init-go/otelinit"
	"github.com/wendall-robinson/tracing/pkg/loggers"
	"github.com/wendall-robinson/tracing/pkg/tracer"
	"go.uber.org/zap"
)

const (
	URL string = "https://api.publicapis.org/entries"
)

var (
	appName = "tracing-example"
	logger  *zap.SugaredLogger
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx, otelShutdown := otelinit.InitOpenTelemetry(ctx, appName)
	defer otelShutdown(ctx)

	GetSomeURL(ctx)
}

func GetSomeURL(ctx context.Context) {

	// start a span
	params := tracer.BuildSpanParams("url", URL)
	span := tracer.TraceWithAttributes(ctx, loggers.Zap(logger), appName, "GetSomeURL", params)
	defer span.EndSpan()

	// Create a new HTTP client
	client := &http.Client{}

	// Send HTTP GET request
	response, err := client.Get(URL)
	if err != nil {
		fmt.Printf("Error sending request: %s", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
