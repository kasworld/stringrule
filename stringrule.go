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
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"
)

func GetFunctionName(i interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	ss := strings.Split(s, ".")
	return ss[len(ss)-1]
}

type RuleFn func(robj interface{}, s string) error

func (msfn RuleFn) String() string {
	return GetFunctionName(msfn)
}

type StringRule struct {
	RuleName      string
	Min, Max      int // include max
	RuleFns       []RuleFn
	UnicodeRanges []*unicode.RangeTable
}

func IsIn(x int, r1, r2 int) bool {
	return x >= r1 && x < r2
}

func (sr StringRule) NewAdd(arg StringRule) StringRule {
	rtn := StringRule{
		sr.RuleName,
		sr.Min, sr.Max,
		append(sr.RuleFns, arg.RuleFns...),
		append(sr.UnicodeRanges, arg.UnicodeRanges...),
	}
	return rtn
}

func (sr StringRule) DoCheck(robj interface{}, name string) error {
	if !utf8.ValidString(name) {
		return fmt.Errorf("Invalid String %v", name)
	}

	namelen := utf8.RuneCountInString(name)
	if !IsIn(namelen, sr.Min, sr.Max+1) {
		return fmt.Errorf("%v Invalid string len %v,  %v", sr.RuleName, namelen, name)
	}
	for _, fn := range sr.RuleFns {
		if err := fn(robj, name); err != nil {
			return fmt.Errorf("%v fail %v", sr.RuleName, err)
		}
	}
	if len(sr.UnicodeRanges) == 0 {
		return nil
	}
	for _, r := range name {
		if !unicode.In(r, sr.UnicodeRanges...) {
			return fmt.Errorf("%v can't use charactor '%v'", sr.RuleName, r)
		}
	}
	return nil
}

func CheckSpaceRule(robj interface{}, name string) error {
	if name != strings.TrimSpace(name) {
		return fmt.Errorf("can't start,end whitespace")
	}
	// noPreSuf := " "
	// if strings.HasPrefix(name, noPreSuf) || strings.HasSuffix(name, noPreSuf) {
	// 	return fmt.Errorf( "can't start,end charactor %v", noPreSuf)
	// }
	if strings.Contains(name, "  ") {
		return fmt.Errorf("no double space in name")
	}
	return nil
}
