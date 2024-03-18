#get track from command 
track=$1
exercise=$2

if [ ! $track ] || [ ! $exercise ]; then
  echo "Please provide a track and an exercise"
  exit 1
fi

command="/usr/local/bin/exercism submit /root/exercism/$track/$exercise/*"

docker run --rm -it -v ./:/root/exercism md91/exercism-cli /bin/bash -c "$command"
