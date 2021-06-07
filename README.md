# sinkronizer
Utility to replay Cloud Logging into BigQuery

Goals:
- Provide retroactive catch-up for Log Sink
- Replicate native Log Sink format
- Catch up to when a Log Sink was enabled (if it has been)
- Idenpotency-ish: tool can be run at any time and won't create duplicate entries

Implementation:
- Go binary in Cloud Run using GCP SDK
- UI for launching execution
