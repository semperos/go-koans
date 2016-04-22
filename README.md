# Go Koans

I'm not sure how this ever got popular under my namespace, but it seems to have
enough interest that I feel compelled to maintain it now. The original author,
[Steven Degutis](https://github.com/sdegutis), committed the initial suite of
tests, 4c5e766, on Mar 10, 2012. I don't recall now how I discovered the
initial codebase beyond searching for "go koans" on GitHub. I do recall that
I was enlightened considerably by [Ruby Koans](http://rubykoans.com/) and
something similarly enlightening on my journey to the Nirvana of Go could be
a blessing for anyone.

Since my discovery of [The Go Programming Language](https://golang.org/), the
language and development environments have changed significantly. I will do my
best to balance current best practices and a low barrier of entry for newcomers
(whom I assume to be the vast majority of those with interest in this
repository). I will try to keep up to date with the latest stable releases. I
hope I can rely on this wonderful community to help me with this.

## Usage

Install Go. See the [Downloads page](https://golang.org/dl/).

Step into the directory for this project. You'll see there are a few different
folders to choose from. I'd suggest starting with the "easy" koans by running:

```
go test -v github.com/semperos/go-koans/easy
```

This will run just the tests for the koans in the "easy" folder. After you've
worked through those, you can move onto other folders.

### How to Solve

Running the tests will show you what's failing and thus what to work on next.
In the source code, you'll find identifiers prefixed with an underscore. You
need to either replace these with a legal value that makes the assertion true,
or in some cases you need to delete an assertion (it will say
`assert(_DeleteMe)`) so as to allow code execution to continue past that line.

Run the appropriate `go test` invocation again after you've made your changes to
see if you were right!

## Docker Usage

I have found that using [Docker](https://www.docker.com/) helps me keep my
development environment clean and portable. Here is an example of how I might
set up an environment dedicated to go through these koans.

Install/Setup:

```shell
luser@lolcathost:~ $ docker-machine create -d virtualbox golang
luser@lolcathost:~ $ eval $(docker-machine env golang)
luser@lolcathost:~ $ docker pull library/golang:1.6.0-alpine
luser@lolcathost:~ $ docker run --rm -ti -v "$PWD":/usr/src/koans -w /usr/src/koans golang:1.6.0-alpine /bin/sh
```

Now with an interactive shell inside of a minimal container you may itterate
through the same steps to enlightenment described above.

## Helpful References

Bookmark the [spec](http://golang.org/ref/spec) and the [packages listing](http://golang.org/pkg/).
You can also run the Go website locally with `godoc -http=:8080`.
