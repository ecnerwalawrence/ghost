# ghost

## How to build this?

1. Install go1.7.x
2. Install glide https://github.com/Masterminds/glide
e.g.
<pre>
brew install glide
</pre>
3. Create a workspace directory.
e.g. 
<pre>
mkdir -p /Users/ecnerwal/workspace/go
</pre>
4. Create code directories.
e.g. 
<pre>
mkdir -p /Users/ecnerwal/workspace/go/src/github.com/ecnerwalawrence/
</pre>
5. set your GOPATH environment variable.  Bash example below
<pre>
export GOPATH=/Users/ecnerwal/workspace/go
</pre>
6. clone this repo then cd into repo
e.g. 
<pre>
cd /Users/ecnerwal/workspace/go/src/github.com/ecnerwalawrence/
git clone git@github.com/ecnerwalawrence/ghost.git
cd ghost
</pre>
7. Build code 
<pre>
make setup
</pre>
8. Run test (optional)
<pre>
make test
</pre>
9. Run ghost
<pre>
./ghost
</pre>
