package string

import "testing"

func TestImplode(t *testing.T) {
	tests := []struct {
		args      []interface{}
		separator string
		want      string
	}{
		{args: []interface{}{}, separator: ",", want: ""},
		{args: []interface{}{11, 22}, separator: ",", want: "11,22"},
		{args: []interface{}{"A", "B"}, separator: ",", want: "A,B"},
		{args: []interface{}{true, false}, separator: ",", want: "true,false"},
		{args: []interface{}{22}, separator: ",", want: "22"},
		{args: []interface{}{"", 22}, separator: ",", want: ",22"},
		{args: []interface{}{1.1, 2.2}, separator: ",", want: "1.1,2.2"},
	}
	for _, tt := range tests {
		got := Implode(tt.args, tt.separator)
		if got != tt.want {
			t.Errorf("Implode() => '%v', want '%v'", got, tt.want)
		}
	}
}
