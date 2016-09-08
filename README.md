# go-dojo-rectangles

This application loads rectangular polygons from a JSON file and prints the result of the following computations:
* Intersection
* Containment
* Adjacency

## Usage

This repository already provides executables for Linux, MacOS and Windows in the `binaries` folder. So,
if you're running Windows 64-bit, you should execute:
```
./binaries/rectangles_win64 -json testdata/valid.json
```

Based on `testdata/valid.json` (see the image below for a depiction of the JSON data), the output should look like:
 ```
 $ go run main.go -json testdata/valid.json
 => [A]
   => Intersects:  None
   => Contains:  None
   => Is adjacent to:  B
 => [B]
   => Intersects:  None
   => Contains:  None
   => Is adjacent to:  A, C
 => [C]
   => Intersects:  None
   => Contains:  None
   => Is adjacent to:  B
 => [D]
   => Intersects:  None
   => Contains:  None
   => Is adjacent to:  None
 => [E]
   => Intersects:  None
   => Contains:  None
   => Is adjacent to:  None
 => [F]
   => Intersects:  G (intersection points: [{X:-3 Y:-4} {X:-3 Y:-3} {X:-2 Y:-4} {X:-2 Y:-3}])
   => Contains:  None
   => Is adjacent to:  None
 => [G]
   => Intersects:  F (intersection points: [{X:-3 Y:-4} {X:-3 Y:-3} {X:-2 Y:-4} {X:-2 Y:-3}])
   => Contains:  None
   => Is adjacent to:  None
 => [H]
   => Intersects:  None
   => Contains:  I
   => Is adjacent to:  None
 => [I]
   => Intersects:  None
   => Contains:  None
   => Is adjacent to:  None
 => [J]
   => Intersects:  K (intersection points: [{X:-4 Y:5} {X:-3 Y:6}])
   => Contains:  A, B, C, D, E, F, G, H, I
   => Is adjacent to:  None
 => [K]
   => Intersects:  J (intersection points: [{X:-4 Y:5} {X:-3 Y:6}])
   => Contains:  None
   => Is adjacent to:  None
 ```

## Build

If you want to compile this application code, you'll need to [install Go](https://golang.org/doc/install).

Choose one from above and proceed to build:
```
GOOS=windows GOARCH=386 go build -o binaries/rectangles_win32.exe -ldflags '-w -extldflags=-static'
GOOS=windows GOARCH=amd64 go build -o binaries/rectangles_win64.exe -ldflags '-w -extldflags=-static'
GOOS=darwin GOARCH=amd64 go build -o binaries/rectangles_mac -ldflags '-w -extldflags=-static'
GOOS=linux GOARCH=386 go build -o binaries/rectangles_lin32 -ldflags '-w -extldflags=-static'
GOOS=linux GOARCH=amd64 go build -o binaries/rectangles_lin64 -ldflags '-w -extldflags=-static'
```

You can now jump to [Usage](#usage) or run it directly with Go:
```
go run main.go -json testdata/valid.json
```

## Tests

Package `geometry` tests feed from `testdata/valid.json` file, which contains rectangles as depicted below.
You should use the image for an easier understanding of said tests.

![alt Rectangles](rectangles.jpg)

The aforementioned package **has 100% test coverage**.
