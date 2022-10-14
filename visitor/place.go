package main

// Place provides an interface for place that the visitor should visit.
type Place interface {
	Accept(v Visitor) string
}
