# Files
- **build.sh** : script to build binaries
- **hammer.mongo.go** : main file for the tool

# Usage:
- Command line options:
   <pre>
Usage of ./bin/hammer.macos:
  -config="": To use config file
  -debug=false: debug flag (true|false)
  -initdb=false: Drop DB before start testing
  -max=false: To find out Max
  -monitor=1: Monitor interval
  -profile="": to specify a traffic profile, all UPPERCASE
  -quiet=false: To silent monitor output
  -rps=500: Set Request Per Second, 0 to find Max possible
  -run="": To specify run id, used for archive test report
  -server="localhost:27017": Define server to be tested, default to localhost:27017
  -thread=0: Number of system thread to be used
  -total=0: Total request to be sent, default to unlimited (0)
  -totaltime=0: To set how long (seconds) to run the test
  -warmup=0: To set how long (seconds) for warmup DB
  -worker=10: Number of workers, every worker will have two connections to mongodb
</pre>

a sample run:
<pre>
GOPATH=\`pwd\`:$GOPATH go run hammer.mongo.go -monitor 1 -max -worker 9 -server localhost:27017 -max -initdb=true -profile=insert
</pre>

use the binary 
<pre>
./hammer.linux -monitor 1 -server ec2-107-21-153-123.compute-1.amazonaws.com:27017 -thread 4 -max -initdb=false -max=true -profile=insert -worker 32 -total=100000
./hammer.linux -monitor 1 -server ec2-107-21-153-123.compute-1.amazonaws.com:27017 -thread 4 -max -initdb=false -max=true -profile=singleQuery -worker 32 -total=200000
./hammer.linux -monitor 1 -server ec2-107-21-153-123.compute-1.amazonaws.com:27017 -thread 4 -max -initdb=false -max=true -profile=inplaceupdate -worker 32 -total=100000
./hammer.linux -monitor 1 -server ec2-107-21-153-123.compute-1.amazonaws.com:27017 -thread 4 -max -initdb=false -max=true -profile=extendedupdate -worker 32 -total=100000
./hammer.linux -monitor 1 -server ec2-107-21-153-123.compute-1.amazonaws.com:27017 -thread 4 -max -initdb=false -max=true -profile=insertsmall -worker 8 -total=95000
</pre>

# Quick Install Using Released Binaries:

To use pre-build binaries, just do this:

<pre>
wget --no-check-certificate https://raw.githubusercontent.com/rzh/hammer.mongo/master/scripts/bootstrap.sh -O - | bash
</pre>

It will download the binaries, and necessary scripts. No need to install Go. Linux/64 and Darwin/64 are support at this moment. Simple run script to run the workload

<pre>
./simple_insert.sh
</pre>

You can also use go get to install hammer.mongo, run following command 
<pre>
  go install github.com/rzh/hammer.mongo
</pre>

hammer.mongo shall be installed into your <i><b>$GOPATH/bin</b></i>, make sure your PATH is properly configured, you can simply issue hammer.mongo to run this. 

# Releases:
Release are made and binaries for Linux/amd64 and Darwin/amd64 will be part of the binary package, you can get it from
<pre>
https://github.com/rzh/hammer.mongo/releases/latest
</pre>

# How to Build or Run from the source:
The build.sh will build hammer.mongo for Mac OS and Linux 64, to do it, you have to:
- To build the tool, please make sure your Go is installed and setup properly. Please refer to this doc http://golang.org/doc/install. For Mac OS, the easier way it to use homebrew to install with --cross-compile-common option:
<pre>
  brew install go --cross-compile-common
</pre>
- For cross platform build, make sure you compile Go runtime properly, please refer to this blog for details: http://dave.cheney.net/2012/09/08/an-introduction-to-cross-compilation-with-go. You will not need this if you install with homebrew according to the above step.
- You will need install proper Go package, run following command after your GOPATH is properly set
<pre>
  GOPATH=\`pwd\`:$GOPATH go get -d
</pre>
- Run build.sh, which will build binaries under ./bin folder,
<pre>
$ ls -1 ./bin
hammer.linux
hammer.macos
</pre>
- To run from source, simple run <b><i>go run hammer.mongo.go ...</i></b> with all the options.

