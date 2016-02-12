# Slackify

Only on mac, allows you to add a slackbot and ask him to choose and next your songs in the office.

You need to run it locally on the office mac running spotify

## Within slack
```
!play <your spotify uri>
!next
```

## Compile it (on mac only)
```
go build -o slackify_companion 
```


## Running it

SLACK_TOKEN=xxxxx ./slackify_companion