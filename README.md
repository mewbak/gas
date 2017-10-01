gas [![License](http://img.shields.io/:license-gpl3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0.html) [![GoDoc](http://godoc.org/github.com/opennota/gas?status.svg)](http://godoc.org/github.com/opennota/gas) [![Build Status](https://travis-ci.org/opennota/gas.png?branch=master)](https://travis-ci.org/opennota/gas)
===

A tool for extracting functions from object files and transforming them into Go assembly.

## Install

    go get -u github.com/opennota/gas

## Use

```
gas object some_function
```

where object is an executable, an object file (`*.o`), or a static or shared library.

## License

gas is released under the GNU General Public License version 3.0. As a special exception to the GPLv3, you may use the parts of gas output copied from gas source without restriction. Use of gas makes no requirements about the license of generated code.
