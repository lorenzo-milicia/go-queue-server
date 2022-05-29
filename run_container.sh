#!/bin/bash
exec docker run -it --mount type=bind,src=$PWD,target=/app/src/ --net=host queue-server:v0.0.2 
