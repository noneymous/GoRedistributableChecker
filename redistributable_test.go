package redistributable

import (
	"fmt"
	"testing"
)

func TestIsInstalled(t *testing.T) {
	type args struct {
		redistributableVersion int
	}
	tests := []struct {
		name string
		args args
	}{
		{"VC2005x86", args{VC2005x86}},
		{"VC2005x64", args{VC2005x64}},
		{"VC2008x86", args{VC2008x86}},
		{"VC2008x64", args{VC2008x64}},
		{"VC2010x86", args{VC2010x86}},
		{"VC2010x64", args{VC2010x64}},
		{"VC2012x86", args{VC2012x86}},
		{"VC2012x64", args{VC2012x64}},
		{"VC2013x86", args{VC2013x86}},
		{"VC2013x64", args{VC2013x64}},
		{"VC2015x86", args{VC2015x86}},
		{"VC2015x64", args{VC2015x64}},
		{"VC2017x86", args{VC2017x86}},
		{"VC2017x64", args{VC2017x64}},
		{"VC2015to2019x86", args{VC2015to2019x86}},
		{"VC2015to2019x64", args{VC2015to2019x64}},
	}
	for _, tt := range tests {
		installed := IsInstalled(tt.args.redistributableVersion)
		if installed {
			fmt.Printf("%s: INSTALLED\n", tt.name)
		} else {
			fmt.Printf("%s: missing\n", tt.name)
		}
	}
}
