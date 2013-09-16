go-stringit
===========

A go library for formatting strings for functions similar to python's .format

## Purpose
I got tired of using `+` to concatenate strings in Go. So I decided that I would just
recreate python's .format function for strings in golang.

## Usage
Usage is limited to using only `{}` as a replacement tool. Types are also limited to 
`string`, `int`, `uint64`, and `float64`. We can use the function like:

        stringit.Format("The {} says {}", "cow", "MOO!")

Yields:

        "The cow says MOO!"

And:

        stringit.Format("I have {} bananas", 871)

Yields:

        "I have 871 bananas"
