package main

type ShapeHolder interface {
	getShape() *Shape
	ContentsOf(*Cell) string
}
