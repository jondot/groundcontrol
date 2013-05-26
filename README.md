# Ground Control

Ground Control is a Go based daemon that runs on your Pi and lets you
manage and monitor it with ease.

See a screenshot of [the management UI](https://raw.github.com/jondot/groundcontrol/master/ui-screenshot.png) and [my Pi's temperature on Librato](https://raw.github.com/jondot/groundcontrol/master/pi-librato.png).



## Usage

Download the Ground Control binary, or build from source (see
`Development`).



Then, transfer it to your Pi.


```
$ scp groundcontrol-v0.0.1.tar.gz pi-user@PI_HOST:
```

On the Pi, extract and set up.

```
$ tar zxvf groundcontrol-v0.0.1.tar.gz
$ cd groundcontrol-v0.0.1/
```

run with a simple command.

```
$ ./groundcontrol -config myconfig.json
```

You can access the UI from your browser on port `4571`.

```
http://PI_HOST:4571/
```

For configuration, use `groundcontrol.json.sample` as a basis and make
sure to customize these fields:


* `librato` - You can make a free account at
[Librato](http://librato.com) and then drop the key and user there.

* `tempodb` - You can make a free account at
[TempoDB](http://tempo-db.com) and then drop the key and user there.

Make sure to go over the plans (paid and free) and see what fits you
best. 

Both Librato and TempoDB were included because they have different
retention and resolution and features for the free plans.


Next up, set up your "controls". This is where you input a label, and an
"on", "off", "once" commands to automatically build a GUI around it.


Here is an example of having an `xbmc` control. It allows for shutting
down and turning on XBMC.

```javascript
  "controls" : {
    "xbmc": {
      "on" : "/etc/init.d/xbmc start",
      "off" : "/etc/init.d/xbmc stop"
    }
  }
```

### init.d

You might want to use an `init.d` or `upstart` script to keep
`groundcontrol` up at all times.

An default init.d script is included here `support/init.d/groundcontrol`.

Place your `groundcontrol` folder at:

```
/opt/groundcontrol/
```

And configuration should be at:

```
/etc/groundcontrol.json
```

You can edit your `support/init.d/groundcontrol` if you want to modify these paths.

After you've verified `/etc/init.d/groundcontrol start` to be working, to set up the default run order you can use:

```
$ update-rc.d groundcontrol defaults
```





## More Details

There's plenty more to Ground Control under the hood, let's list out a
few things.


### Temperature Monitoring

Ground control will read a file containing a temperature reading to
update its own health records. 

This is made so the mechanism is
flexible -- you can use Ground Control on any machine (not just a Pi) as
long as it will expose a temperature in a file-like device.


### Controls

You can add and remove controls (commands that your Ground Control can
run) by editing or specifying them with your configuration file..


Here's a full description of the format:


```
{
  "control_name": {
    "on" : "cmd for on, normally a 'start' for a service",
    "off" : "cmd for off, normally a 'stop' for a service",
    "status" : "A command that returns the status of the process",
    "once" : "a one time command, a cleanup, a shutdown etc."
  }
}
```

By convention `control_name` is snake_case and we turn it into "Control Name" on the UI.

The name should be nice for using in a REST API, in this case the commands turn into:


```
POST controls/control_name/on
POST controls/control_name/off
POST controls/control_name/once
GET controls/control_name/status
```

The `on`, `off`, and `once` commands are async, and we return a 202 OK for success.
The `status` command are async, and we return a 200 OK for success.


And you can easily build an app (mobile?) yourself that makes use of those.






## Development

Here's a short guide if you want to experiment with Ground Control yourself.

In each case, start of by taking a Ground Control repo clone.

### Compiling

Set up dependencies and build:

```
$ go get github.com/jondot/gosigar
& go build
```



### Cross Compiling

You probably want to build on your own (much more powerful) system rather than on the Pi itself to save time.

In my case I'll be compiling a Go binary on a Mac (OSX, x64), for a Raspberry PI (Linux, ARMv5/6).

Here's how to do it on a Mac and `brew`:

```
$ brew install go --devel --cross-compile-all    # I usually take --devel with Go, drop if you dont.
```

If you've got ZShell, or a nice alias-supporting shell, this is a nice alias:

```
alias go-pi='GOARCH=arm GOARM=5 GOOS=linux go'
```

and then I just

```
$ go-pi build
````


If you don't want to use an alias then this is the command to cross-compile for the Pi:

```
$ GOARCH=arm GOARM=5 GOOS=linux go build
```

### Implementation Details

Ground control surrounds around several concepts:

* Reporter - an entity that takes a Health, and reports it to somewhere.
* Health - the entity that's responsible to gather all of the important health metrics.
* Control - a switchboard-like entity that runs commands on request.

They are sorted by the level of fun/hackability you can get from it, but YMMV :)


Note, that [go-metrics](https://github.com/rcrowley/go-metrics) for example, could have replaced the entire reporter stack here using its various pluggable 
reporters, however, it only supports integers out of the box.

At the worst case, `go-metrics` itself can be implemented as a reporter (in fact it will be a aggregate reporter of reporters :).









# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2013 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



