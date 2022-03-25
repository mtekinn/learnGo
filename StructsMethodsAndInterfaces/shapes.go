package main

import "math"

type Shape interface {
    Area() float64
}

// Dikdörtgen
type Rectangle struct {
    Width float64
    Height float64
}

// Dikdörtgenin Alanı
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Dikdörtgenin Çevresi
func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.Width + rectangle.Height)
}

// Daire
type Circle struct {
    Radius float64
}

// Dairenin Alanı
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Üçgen
type Triangle struct {
    Base float64
    Height float64
}

// Üçgenin Alanı
func (c Triangle) Area() float64 {
    return (c.Base * c.Height) * 0.5
}