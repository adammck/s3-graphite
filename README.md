# S3 to Graphite

This is a tiny Go app which polls an S3 bucket+prefix for the number of objects
it contains, and writes it to Graphite. This is handy when one is using S3 as a
ghetto job queue, and possibly other reasons.


## Usage

Fetch the code and build:

	go get github.com/adammck/s3-graphite
	cd $GOPATH/adammck/s3-graphite
	go build

Set the following environment vars:

	# AWS keys with read access to the bucket
	export AWS_ACCESS_KEY_ID=xxxxxxxxxx
	export AWS_SECRET_ACCESS_KEY=yyyyyyyyyy
	export AWS_REGION=us-east-1

	# the bucket to watch
	export S3_BUCKET=my-bucket
	export S3_PREFIX=dir/subdir

	# the server to send metrics to
	export GRAPHITE_ADDRESS=metrics.example.com
	export GRAPHITE_PREFIX=s3-count.my-bucket.dir.subdir

Run it:

	./s3-graphite


## License

MIT.
