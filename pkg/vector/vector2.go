package vector

import (
	"fmt"
	"math"
)

type Numeric interface {
	int | int32 | int64 | float32 | float64
}

type Intager interface {
	int | int32 | int64
}

type Vec2[T Numeric] struct {
	X, Y T
}

func (v1 Vec2[T]) Sub(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vec2[T]) Add(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vec2[T]) Mul(arg interface{}) Vec2[T] {
	switch v2 := arg.(type) {
	case Vec2[T]:
		return Vec2[T]{
			X: v1.X * v2.X,
			Y: v1.Y * v2.Y,
		}
	case T:
		return Vec2[T]{
			X: v1.X * v2,
			Y: v1.Y * v2,
		}
	default:
		panic("unsupported type")
	}
}

func (v1 Vec2[T]) Div(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v1.X / v2.X,
		Y: v1.Y / v2.Y,
	}
}

func (v1 Vec2[T]) Modulo(val Vec2[T]) (Vec2[T], error) {
	switch any(v1.X).(type) {
	case int, int32, int64:
		return Vec2[T]{
			X: T(int(v1.X) % int(val.X)),
			Y: T(int(v1.Y) % int(val.Y)),
		}, nil
	default:
		return Vec2[T]{}, fmt.Errorf("unsupported type for modulo")
	}
}

func (v1 Vec2[T]) Magn() float64 {
	return math.Sqrt(float64(v1.Dot(v1)))
}

func (v1 Vec2[T]) DistanceTo(v2 Vec2[T]) float64 {
	dist := v2.Sub(v1)
	return math.Sqrt(float64(dist.Dot(dist)))
}

func (v1 Vec2[T]) Len(arg interface{}) float64 {
	switch v2 := arg.(type) {
	case nil:
		return math.Sqrt(float64(v1.Dot(v1)))
	case Vec2[T]:
		vector := v2.Sub(v1)
		return math.Sqrt(float64(vector.Dot(vector)))
	default:
		panic("unsupported type")
	}
}

func (v1 Vec2[T]) Dot(v2 Vec2[T]) T {
	return (v1.X * v2.X) + (v1.Y * v2.Y)
}

func (v1 Vec2[T]) Up() Vec2[T] {
	return Vec2[T]{Y: v1.Y - 1, X: v1.X }
}

func (v1 Vec2[T]) Down() Vec2[T] {
	return Vec2[T]{Y: v1.Y + 1, X: v1.X }
}

func (v1 Vec2[T]) Left() Vec2[T] {
	return Vec2[T]{Y: v1.Y, X: v1.X - 1 }
}

func (v1 Vec2[T]) Right() Vec2[T] {
	return Vec2[T]{Y: v1.Y, X: v1.X + 1}
}

func (v1 Vec2[T]) Direction(r rune) Vec2[T] {
	switch r {
	case '>':
		return v1.Right()
	case '<':
		return v1.Left()
	case '^':
		return v1.Up()
	case 'v':
		return v1.Down()
	default:
		return Vec2[T]{}
	}
}
