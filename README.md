# go-mastodon-matrix-webhooks

Allows you to post signup/report info into a Matrix channel.    
    

# Install
Install golang. Install [matrix-webhook](https://github.com/nim65s/matrix-webhook).    
    
Copy .env.example to .env and fill in the blanks


```
MATRIX_WEBHOOK_URL="http://127.0.0.1:4785"
MATRIX_WEBHOOK_API_KEY=keySetupInMatrixWebhooks
MATRIX_ACCOUNT_CHANNEL="!channelID:matrix.org"
MATRIX_REPORT_CHANNEL="!channelID:matrix.org"
PORT=8081
IP2LOCATION_FILE="full path to .BIN file"
MASTODON_INSTANCE="mastodon.test"
```