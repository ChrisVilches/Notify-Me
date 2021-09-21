# Notify Me

A little command line tool that sends a push notification to your smartphone.

## How to use

This program is simply a wrapper for this app, which is free and you must install it on your phone in order to use this program:

https://play.google.com/store/apps/details?id=net.xdroid.pn

After installing it, take note of your API key.

NOTE: Link above is Android only. It's unknown if it's available for iOS.

Next, compile this program using:

```bash
make build # Or inspect Makefile to see the command, in case this doesn't work.
```

Then link the executable to your `PATH` so it can be run.

```bash
PATH=$PATH:~/Notify-Me/build # Should be done in ~/.bashrc so it's permanent (and more convenient).

notify_me -apiKey YOUR_API_KEY -title "notification title" -content "notification body, text here..."
```

If everything is OK, you should see this output, while getting a notification to your phone:

```bash
Sending to: YOUR_API_KEY
Title: notification title
Content: notification body, text here...
200 OK
```

You can also execute it by having the API key as an environment variable, instead of passing it as an option parameter (NOTE: Specifying `-apiKey YOUR_API_KEY` overrides the environment variable):

```bash
# Store it in ~/.bashrc
NOTIFY_ME_API_KEY=YOUR_API_KEY notify_me -title "notification title" -content "notification body, text here..."
```
