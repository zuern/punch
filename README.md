# Punch - The Simple CLI timekeeper
__A dead simple punch card program to keep track of your time. Implemented in [Go](https://golang.org).__

## Usage
```
Usage:
  punch (-h | -v)
  punch (in|out|log) [<project_name>]
  punch list
Examples:
  punch in
  punch in "ACME Website Development"
  punch out "ACME Website Development"
  punch log "ACME Website Development"
Options:
  -h, --help
  -v, --version
```

## Installation
Installing punch can be done in a few ways. Probably the simplest way is as follows (for Unix-like systems):

```
git clone https://github.com/propheis/punch.git
cd punch
go build
sudo cp punch /usr/local/bin
cd ../
rm -r punch
```

If you already have Go installed and set up [on your path](https://golang.org/doc/code.html#GOPATH) it is as 
simple as running

```
go get github.com/propheis/punch
```

## How it Works
When you first run Punch it will create a folder at `~/.punch` where `~` evaluates to your `$HOME` directory on 
Unix-like systems and usually `C:\Users\YourName\` on Windows. Anytime you `punch in` or `punch out` it will 
write to `~/.punch/default.punch`. If you do something like `punch in "test"` it will write to 
`~/.punch/test.punch`. This is how punch keeps track of your various projects. These files are referred to as 
punchCards.

The punchCards contain the logs for each of your projects in an easy to parse format. Rows are separated by a 
newline character (`\n`) and columns by a tab character (`\t`). This makes it easy to process the data in other
programs or generate reports.

If you are using a Unix-like OS (Linux, BSD, MacOS, etc.) you can do cool stuff as well like piping the output
of Punch to other programs. For example, to find the list of all the times you punched out of a project called
"test" and save it to a file called `times.txt`, you could run the following command:

```
punch log test | grep "OUT" | awk '{print $3, $4, $5, $6, $8}' > times.txt
```

## Licensing Information

```
   Copyright © 2017 Kevin Zuern <kevin@kevinzuern.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
```