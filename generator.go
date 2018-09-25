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
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func CreatePassword(l int, s []string, debugging bool) (string, error) {
	/* Function CreatePassword is password generator. It uses simple
	"math/rand" package because it is not supposed to serve cryptography.
	CreatePassword checks for some errors: empty password, and for empty 
	charSet slice.
	Finally, it creates password by choosing random chars from specified set.
	CreatePassword takes length of password, slice of character sets, and
	debugging switch as arguments, and returns password string and error. */
	var err error
	var pass string
	if l < 1 {
		err := errors.New("Password need to be at least 1 character long; " +
			"current length: " + strconv.Itoa(l) + ".")
		return pass, err
	}
	if len(s) < 1 {
		err := errors.New("You disallowed all character sets; it's impossible to create password now.")
		return pass, err
	}
	set := strings.Join(s, "")
	if debugging == true {
		fmt.Print("INFO: string of allowed characters:\n    ")
		fmt.Println(set)
	}
	pass = ""
	for i := 0; i < l; i++ {
		ch := string(set[rand.Intn(len(set))])
		pass = pass + ch
	}
	return pass, err
}