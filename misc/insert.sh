#!/bin/bash

bq --location=US load --source_format=NEWLINE_DELIMITED_JSON \
--schema=log_sink.stderr_20210607.schema.json \
log_sink.stderr sample_data.json   