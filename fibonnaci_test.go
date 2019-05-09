package fibonacci

import "testing"

func Test_fibR(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fib(0) = 0", args{0}, 0},
		{"fib(1) = 1", args{1}, 1},
		{"fib(2) = 1", args{2}, 1},
		{"fib(3) = 2", args{3}, 2},
		{"fib(20) = 2", args{20}, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibR(tt.args.n); got != tt.want {
				t.Errorf("FibR() = %v, want %v", got, tt.want)
			}
		})
	}
}

// http://www.suguru.jp/Fibonacci/Fib100.html
func Benchmark_fibR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibR(40)
	}
}

func Test_fibRM(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fib(0) = 0", args{0}, 0},
		{"fib(1) = 1", args{1}, 1},
		{"fib(2) = 1", args{2}, 1},
		{"fib(3) = 2", args{3}, 2},
		{"fib(20) = 2", args{20}, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibRM(tt.args.n); got != tt.want {
				t.Errorf("FibRM() = %v, want %v", got, tt.want)
			}
		})
	}
}

// http://www.suguru.jp/Fibonacci/Fib100.html
func Benchmark_fibRM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRM(40)
	}
}

func Test_fibL(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fib(0) = 0", args{0}, 0},
		{"fib(1) = 1", args{1}, 1},
		{"fib(2) = 1", args{2}, 1},
		{"fib(3) = 2", args{3}, 2},
		{"fib(20) = 2", args{20}, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibL(tt.args.n); got != tt.want {
				t.Errorf("FibL() = %v, want %v", got, tt.want)
			}
		})
	}
}

// http://www.suguru.jp/Fibonacci/Fib100.html
func Benchmark_fibL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibL(40)
	}
}
