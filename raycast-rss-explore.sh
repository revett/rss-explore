#!/bin/bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title rss-explore - Convert YouTube URL → Channel RSS Feed
# @raycast.mode compact

# Optional parameters:
# @raycast.icon 🔗
# @raycast.argument1 { "type": "text", "placeholder": "e.g. https://youtube.com/watch?v=7LICrnxWd38" }
# @raycast.packageName @revcd

# Documentation:
# @raycast.author Charlie Revett (@revcd)
# @raycast.authorURL https://revcd.com

DATA="{\"url\":\"$1\"}"
RESPONSE=$(curl -l -s -r POST 'https://rss-explore.revcd.com/youtube/convert' --header 'Content-Type: application/json' --data-raw $DATA)
RSS_FEED=$(echo $RESPONSE | jq -r '.url')

if [ $RSS_FEED = "null" ]; then
  echo "Error: $RESPONSE"
  exit 1
fi

echo $RSS_FEED | pbcopy
echo $RSS_FEED
