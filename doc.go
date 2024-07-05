/*
Copyright (c) 2018-2019 The SadNetwork developers
Copyright (c) 2013-2018 The btcsuite developers
Copyright (c) 2015-2016 The Decred developers
Copyright (c) 2013-2014 Conformal Systems LLC.
Use of this source code is governed by an ISC
license that can be found in the LICENSE file.

Sad is a full-node sad implementation written in Go.

The default options are sane for most users. This means sad will work 'out of
the box' for most users. However, there are also a wide variety of flags that
can be used to control it.

Usage:

	sad [OPTIONS]

For an up-to-date help message:

	sad --help

The long form of all option flags (except -C) can be specified in a configuration
file that is automatically parsed when sad starts up. By default, the
configuration file is located at ~/.sadnetwork/sad.conf on POSIX-style operating
systems and %LOCALAPPDATA%\sadnetwork\sad.conf on Windows. The -C (--configfile)
flag can be used to override this location.
*/
package main
