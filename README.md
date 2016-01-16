# munafatask
task to query aws spot instances

### Installation instructions
* Install Golang version>=1.5
* run "git clone https://github.com/abhinavzspace/munafatask.git" inside $GOPATH/src
* run "go get -u github.com/aws/aws-sdk-go"
* run "go get github.com/abhinavzspace/munafatask/awsutil"
* run "go install"
* run "AWS_ACCESS_KEY_ID=thekey AWS_SECRET_ACCESS_KEY=thekey ./spotinstance -count=1 -instance="t1.micro" -price=0.10 -region="us-west-2" -zone="us-west-2a" " inside bin directory of $GOPATH
