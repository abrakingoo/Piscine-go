#!/bin/bash

# Check if HERO_ID is set in the environment
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID environment variable is not set."
  exit 1
fi

# Use curl to fetch the superhero data and jq to process it
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq -r --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id | tonumber)) | .connections.relatives | gsub("\"";"") | gsub("\n"; " ")'

