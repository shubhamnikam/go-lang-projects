# go-slack-bulk-file-uploader using slack bot

#### slack setup
```bash
# goto
    - https://api.slack.com/apps
# create bots
# turn on socket mode
# goto OAuth & provide permissions in Bot Token Scopes
    channels:read
    chat:write
    files:read
    files:write
    im:read
    im:write
    mpim:history
    remote_files:read
    remote_files:share
    remote_files:write

# install app to workspace
# create & get Bot User OAuth Token and update into .env file
# get Slack Channels ID's and update into .env file
```

#### program setup
```bash
# get repo code
# install required packages => go get
# add bot into channel
# run go program
```