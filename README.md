# RPN Calculator

A calculator for reversed polish notation, using function composition.

This was an exercise of concatenative programming; it only supports integer numbers. It is inspired by http://evincarofautumn.blogspot.mx/2012/02/why-concatenative-programming-matters.html.

For more information see http://en.wikipedia.org/wiki/Reverse_Polish_notation.

## Usage

Get the program:

	$ go get github.com/sauerbraten/rpncalc

Now, in your `$GOPATH/bin` there will be the `rpncalc` executable.

### Syntax:

	rpncalc <input>

`<input>` consists of int literals and operators; divided by spaces. Operators currently supported are:

- `+`
- `-`
- `·`
- `/`
- `%`
- `squared`

All operators need 2 arguments in front of it, except for `squared` which only needs one. The · one is `Alt Gr` + `,` on Linux.

### Examples:

	$ rpncalc 2 3 +
	5

	$ rpncalc 2 3 + 4 5 + ·
	45

	$ rpncalc 2 3 + 4 5 + · 20 - 10 / squared 3 %
	1

## License

This code is licensed under a BSD License:

Copyright (c) 2013 Alexander Willing. All rights reserved.

- Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.