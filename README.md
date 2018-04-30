# RPN Calculator

A calculator for reverse polish notation.

For more information see http://en.wikipedia.org/wiki/Reverse_Polish_notation.

## Usage

Get the program:

	$ go get github.com/sauerbraten/rpncalc

This will put the `rpncalc` executable into `$GOPATH/bin`.

### Syntax:

	rpncalc <input>

`<input>` consists of float64 literals and operators; divided by spaces. Operators currently supported are:

- `+`
- `-`
- `·`
- `/`
- `%`
- `^`

All operators need 2 arguments. If you can't type the `·`, see here: http://en.wikipedia.org/wiki/Interpunct#Keyboard_input. The asterisk (`*`) doesn't work since it is reserved in bash and expands to all folder contents in the current directory.

### Examples:

	$ rpncalc 2 3 +
	5

 `2` and `3` are the arguments for the `+` operator.

	$ rpncalc 2 3 + 4 5 + ·
	45

	$ rpncalc 2 3 + 4 5 + · 20 - 10 / 2 ^ 3 %
	1

- `2 3 +` → `5`
- `4 5 +` → `9`
- `5 9 ·` → `45`
- `45 20 -` → `25`
- `20 10 /` → `2`
- `2 2 ^` → `4`
- `4 3 %` → `1`


	$ rpncalc 3 4 + 5 2 ^ · 4 6 · - 8.75 -
	142.25

## License

This code is licensed under a BSD License:

Copyright (c) 2013-2016 Alexander Willing. All rights reserved.

- Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
