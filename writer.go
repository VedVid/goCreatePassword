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
	"bufio"
	"fmt"
	"os"
	"time"
)

const dateFormat = "20060102-150405" //YYYYMMDD-HHMMSS 

func SavePass(s string) {
	/* Function SavePass saves password to the text file.
	Name is current date in format specified as dateFormat constant.
	Function takes current os time and creates file with date as name.
	If there is no error, password is written as plain text, and file closes.
	It takes single argument - password string, and returns nothing. */
	t := time.Now()
	name := t.Format(dateFormat)
	w, err := os.Create(name + ".txt")
	if err != nil {
		fmt.Println(err)
	} else {
		writer := bufio.NewWriter(w)
		fmt.Fprintln(writer, s)
		writer.Flush()
		fmt.Println("Saving done!")
	}
	defer w.Close()  //it's a bit unidiomatic
}