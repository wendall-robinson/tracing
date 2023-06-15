# Tracing

This is a Go package designed to provide easy-to-use helpers for OpenTelemetry. The Tracing package focuses on creating, starting, and stopping SPANs with minimal setup, especially useful for functions or methods that require instrumentation, but don't have an existing instrumentation package.

The primary component of the package is the Tracer object that acts as a wrapper around the OpenTelemetry Tracer.
Features

* Create, start, and stop SPANs with minimal setup.
* Add arguments to the created SPANs.
* Supports custom loggers through a generic Logger interface. Currently, zap, logrus, standard logger, and glog are supported.

## Usage

First, you need to create a Tracer instance:
```golang
tracer := tracer.New("your-tracer-name", yourLogger)
```

Then, you can use Tracer to start a new SPAN:
```golang
ctx = tracer.StartSpan(ctx, "your-operation")
```

You can add parameters to your SPAN:
```golang
tracer.AddParams(yourParams)
```

And when you're done, you can end the SPAN:
```golang 
tracer.EndSpan()
```

## TraceWithAttributes

If you want to do these operations in one line, you can use TraceWithAttributes:
```golang
tracer := tracer.TraceWithAttributes(ctx, yourLogger, "your-tracer-name", "your-operation", yourParams)
```

## Loggers

Tracing package uses an agnostic logger so that the user can pass in any supported logger via an adapter method. Here's an example using zap logger:
```golang
zapLogger, _ := zap.NewProduction()
defer zapLogger.Sync()

logger := loggers.Zap(zapLogger.Sugar())
tracer := tracer.New("your-tracer-name", logger)
```

This package supports zap, logrus, standard logger, and glog.

## BuildSpanParams

For creating parameters for SPANs, you can use the BuildSpanParams method. This method accepts pairs of key and value arguments and builds a map out of them:
```golang
params := tracer.BuildSpanParams("key1", "value1", "key2", "value2")
tracer.AddParams(params)
```
License

This project is licensed under the MIT License. See the LICENSE file for details.
