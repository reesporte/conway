# conway's game of life

this is [conway's game of life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) but implemented using a [toroidal array](https://en.wikipedia.org/wiki/Torus) so it looks nicer in a terminal display.

## usage 
```
$ go build .
$ ./life
```
![example output for size = 50](./img/50.png)  

## flags
the default matrix size is 50, and the default frame rate is about every 100 ms. to use different values, supply them to the `--size` or `--speed` flags like this: 

```
# use a matrix of size 60 by 60, with a refresh every 140 ms  
$ ./life --size=60 --speed=140
```

## installation on a *nix style os _(probably works)_
assuming you have go installed:
```
# build it
$ go build .
# copy the binary to a directory on your path
$ cp life /usr/local/bin
```
