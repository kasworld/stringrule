// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// 문자열 규칙 검사기
package stringrule

import (
	"testing"
	"unicode"
)

func TestStringRule_DoCheck(t *testing.T) {
	latinnum := StringRule{
		RuleName: "latinnum",
		Min:      4,
		Max:      13,
		RuleFns: []RuleFn{
			CheckSpaceRule,
			myStructRuleFn,
		},
		UnicodeRanges: []*unicode.RangeTable{
			unicode.Latin,
			unicode.Number,
		},
	}
	s := "hello33"
	ms := &myStruct{}
	r := latinnum.DoCheck(ms, s)
	t.Logf("str check %v %v", s, r)
}

type myStruct struct {
}

func (ms *myStruct) check(name string) error {
	return nil
}

func myStructRuleFn(robj interface{}, name string) error {
	return robj.(*myStruct).check(name)
}
