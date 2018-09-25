/*
Copyright (c) 2018 Tomasz "VedVid" Nowakowski
This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.
Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:
1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package main

import(
	"flag"
	"fmt"
)

func ParseFlags() ([]string, int, bool, bool) {
	/* Function Parse Flags is supposed to be initial function of program.
	It parses flags that switches debugging mode, saving password to file,
	and adds character sets (specified in charSets.go) to allowed slice,
	that is used by generator.
	By default, sets A..Z, a..z and 0..9 are set to true.
	It takes no arguments, and returns slice of character sets, and
	bools for debugging and saving switches. */
	var allowed = []string{}
	debug := flag.Bool("debug", false, "shows debug information")
	d := flag.Bool("d", false, "short form of debug")
	length := flag.Int("length", 12, "sets length of password")
	l := flag.Int("l", 12, "short form of length")
	small := flag.Bool("smallLetters", true, "are small letters allowed?")
	sl := flag.Bool("sl", true, "short form of smallLetters")
	capital := flag.Bool("capitalLetters", true, "are capital letters allowed?")
	cl := flag.Bool("cl", true, "short form of capitalLetters")
	nums := flag.Bool("numbers", true, "are numbers allowed?")
	n := flag.Bool("n", true, "short form of numbers")
	syms := flag.Bool("symbols", false, "are symbols allowed? (default false)")
	sy := flag.Bool("sym", false, "short for of symbols")
	specs := flag.Bool("specials", false, "are special (unsafe) characters allowed? (default false)")
	sp := flag.Bool("spec", false, "short form of specials")
	save := flag.Bool("save", false, "save password to file? (default false)")
	flag.Parse()
	debugging := false
	if *debug == true || *d == true {
		debugging = true
	}
	passLen := 12
	if *length != 12 {
		passLen = *length
	} else if *l != 12 {
		passLen = *l
	}
	if debugging == true {
		fmt.Print("INFO: length of password:\n    ")
		fmt.Println(passLen)
	}
	// sets that are "true" by default
	if *small == true && *sl == true {
		allowed = append(allowed, SmallLetters)
	}
	if *capital == true && *cl == true {
		allowed = append(allowed, CapitalLetters)
	}
	if *nums == true && *n == true {
		allowed = append(allowed, Numbers)
	}
	// sets that are "false" by default
	if *syms == true || *sy == true {
		allowed = append(allowed, Symbols)
	}
	if *specs == true || *sp == true {
		allowed = append(allowed, Specials)
	}
	if debugging == true {
		fmt.Print("INFO: list of allowed characters:\n    ")
		fmt.Println(allowed)
	}
	return allowed, passLen, debugging, *save
}