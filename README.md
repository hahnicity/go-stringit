go-stringit
===========

A go library for formatting strings for functions similar to python's .format

## Purpose
I got tired of using `+` to concatenate strings in Go. So I decided that I would just
recreate python's .format function for strings in golang.

## Usage
Usage is currently limited to using only `{}` as a replacement tool. 

We can use the function like:

        stringit.Format("The {} says {}", "cow", "MOO!")

Yields:

        "The cow says MOO!"

And:

        stringit.Format("I have {} bananas", 871)

Yields:

        "I have 871 bananas"

## Supported Types
The following types are supported:

 * int
 * int8
 * int16
 * int32
 * int64
 * uint
 * uint8
 * uint16
 * uint32
 * uint64
 * float
 * float32
 * float64
 * bool
 * string (of course)
