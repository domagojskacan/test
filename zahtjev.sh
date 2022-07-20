#!/bin/bash

curl -i -X POST http://localhost:9091/podaci -H "Accept:application/json" -H "Content-Type:application/json" -d '{"ts": 12316549879, "metric_id": 5, "value": 105}'
