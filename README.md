kgen
==========

Meme generator written in Go

# Requirements

* ImageMagick
    $ sudo apt-get install imagemagick

# Usage

Build and start. Supply images in img/. Configure /pub for public access on $host/pub. You probably want to use a reverse-proxy for the application server.

# License

GPL3+

# Backlog

* Softcode strings / localization
* Make port configurable
* Compositing should be possible in Javascript frontend code instead of ImageMagick
