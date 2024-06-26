package test

import (
	"errors"
	"fmt"
	"github.com/liuzhaomax/maxblog-user/internal/core"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFormatInfo(t *testing.T) {
	cases := []struct {
		name string
		have string
		want string
	}{
		{
			name: "成功例子",
			have: "配置文件加载成功",
			want: "OK: 配置文件加载成功",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := core.FormatInfo(tc.have)
			assert.Equal(t, got, tc.want, fmt.Sprintf("\n*** Expected: \n %#v \n*** Got: \n %#v", tc.want, got))
		})
	}
}

func TestFormatError(t *testing.T) {
	cases := []struct {
		name string
		have core.Error
		want string
	}{
		{
			name: "没有error payload",
			have: core.Error{
				Code: core.Unknown,
				Desc: "配置文件加载失败",
				Err:  nil,
			},
			want: "Unknown: 配置文件加载失败",
		},
		{
			name: "有error payload",
			have: core.Error{
				Code: core.Unknown,
				Desc: "配置文件加载失败",
				Err:  errors.New("未知错误"),
			},
			want: "Unknown: 配置文件加载失败: 未知错误",
		},
		{
			name: "未知错误码",
			have: core.Error{
				Code: 999999,
				Desc: "配置文件加载失败",
				Err:  nil,
			},
			want: "Code(999999): 配置文件加载失败",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := core.FormatError(tc.have.Code, tc.have.Desc, tc.have.Err)
			assert.Equal(t, got, tc.want, fmt.Sprintf("\n*** Expected: \n %#v \n*** Got: \n %#v", tc.want, got))
		})
	}
}
