# Mech

> Listen, um, I don’t really know what’s going on here, but I’m leaving. I
> don’t know where exactly, but I’m gonna start over.
>
> Come with me.
>
> [Paint it Black (2016)][1]

Download media or send API requests

Some users might want to make anonymous requests, because of privacy or any
number of other reasons. This module allows people to do that. Most API these
days only offically support authenticated access. This is useful for the
company providing the API, as they can use the data for their own purposes
(analytics etc). However authentication really doesnt do anything for the end
user. Its just a pointless burden to getting the information you need for a
program you may be writing. Consider that in many cases, the same information
is available via HTML on the primary website, usually without being logged in.
So why can you do that with HTML, but not with the API? Well you can, using this
module.

## How to build?

Check here first:

https://github.com/89z/mech/releases

I dont do a build for every tag, but some tags will have builds available. If
you need a newer build, and cant build yourself for some reason, comment on this
thread:

https://github.com/89z/mech/issues/7

To build yourself, download Go from here:

https://golang.org/dl

and extract archive. Then download Mech:

https://github.com/89z/mech/archive/refs/heads/master.zip

and extract archive. Then navigate to `mech-master/cmd/youtube`, and enter:

~~~
go build
~~~

## Apps used

Installs  | ID
----------|-----------------------------
10.830 B  | `com.google.android.youtube`
3.877 B   | `com.instagram.android`
1.894 B   | `com.zhiliaoapp.musically`
331.715 M | `com.soundcloud.android`
83.771 M  | `com.reddit.frontpage`
30.338 M  | `com.vimeo.android.videoapp`
29.692 M  | `bbc.mobile.news.ww`
2.525 M   | `com.bandcamp.android`
2.086 M   | `com.pbs.video`

## Docs

- https://github.com/89z/mech/tree/master/docs
- https://godocs.io/github.com/89z/mech

[1]://f002.backblazeb2.com/file/ql8mlh/Paint.It.Black.2016.mp4
