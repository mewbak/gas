gas [![License](http://img.shields.io/:license-gpl3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0.html) [![GoDoc](http://godoc.org/github.com/opennota/gas?status.svg)](http://godoc.org/github.com/opennota/gas) [![Build Status](https://travis-ci.org/opennota/gas.png?branch=master)](https://travis-ci.org/opennota/gas)
===

A tool for extracting functions from object files and transforming them into Go assembly.

## Installation

    go get github.com/opennota/gas

## Usage

```
gas object some_function
```

where object is an executable, an object file (`*.o`), or a static or shared library.
