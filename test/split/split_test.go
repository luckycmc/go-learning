package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v,got:%v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v,got:%v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	// 定义测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例的切皮娜
	tests := map[string]test{
		"simple": {
			input: "a:b:b",
			sep:   ":",
			want:  []string{"a", "b", "b"},
		},
		"wrong sep": {
			input: "a:b:c",
			sep:   ",",
			want:  []string{"a:b:c"},
		},
		"more sep": {
			input: "aabc",
			sep:   "ab",
			want:  []string{"a", "c"},
		},
		"leading sep": {
			input: "我是好人",
			sep:   "好",
			want:  []string{"我是", "人"},
		},
	}
	// 遍历切片，逐一执行
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s expected:%#v,got:%#v", name, tc.want, got)
		}
	}
}

func TestSplit3(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple": {
			input: "a:b:b",
			sep:   ":",
			want:  []string{"a", "b", "b"},
		},
		"wrong sep": {
			input: "a:b:c",
			sep:   ",",
			want:  []string{"a:b:c"},
		},
		"more sep": {
			input: "aabc",
			sep:   "ab",
			want:  []string{"a", "c"},
		},
		"leading sep": {
			input: "我是好人",
			sep:   "好",
			want:  []string{"我是", "人"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s expected:%#v,got:%#v", name, tc.want, got)
			}
		})

	}
}
