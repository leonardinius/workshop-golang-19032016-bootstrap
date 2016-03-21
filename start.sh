#/bin/bash

#sudo docker run --rm -it --user=1000:1000 -v $(pwd)/w:/go golang:1.6 bash
sudo docker run \
  --rm \
  -it \
  -p 8081:8081 \
  -v $(pwd)/w:/go \
  golang:1.6 \
  bash
