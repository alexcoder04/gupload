
# gupload

This is a simple HTTP server which allows you to transfer files to and from a machine on the local network.

It's a rewrite of the [old Python version](https://github.com/ac04-dump/dump/tree/main/uploader) in combination with some features of [sharebox](https://github.com/ac04-dump/sharebox) and even some additions.

I use it to transfer some files quickly to/from devices on my local network, mostly Android/iOS/Windows ones.

**Warning:**
The transfer happens over plain HTTP, so it's *unencrypted*.

## Build Instructions

```sh
go build .
```
