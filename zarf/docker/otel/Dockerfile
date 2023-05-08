FROM otel/opentelemetry-collector-contrib:0.53.0

COPY ./collector.yaml .

ENTRYPOINT ["/otelcol-contrib", "--config=/collector.yaml"]