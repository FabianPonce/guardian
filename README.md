# guardian
Guardian is a pluggable tool for using machine learning to identify whether objects are in view of a camera and then take some action based on that.

## Requirements

* Darwin (x86_64) or Linux (arm64). Linux x86_64 probably works but is untested.
* An AWS Account, if using the Rekognition classification driver
  
### Darwin
* OpenCV

### Linux
* alsa-lib (libasound2-dev on Debian)

## Usage
First, modify the `guardian.yml` file to your desired settings. Camera indices start from 0 and increase from there. 

Guardian will always look for the file in the current working directory.

`make install` will install it system-wide, and `make` will by default just build it in this directory.