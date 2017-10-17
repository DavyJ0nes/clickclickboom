[![Go Report Card](https://goreportcard.com/badge/github.com/DavyJ0nes/clickclickboom)](https://goreportcard.com/report/github.com/DavyJ0nes/clickclickboom)
# clickclickboom

## Description
This is a pretty basic test of you're Docker Service's resiliance to single instance failures.

In a nut shell, the program picks a defined number of random instances from a defined Docker Swarm service and kills them.
The idea behind it was firstly to keep writing stuff in Go and also to play around with the Docker APIs.

## Usage
```
Usage of clickclickboom:
Version: 0.1 (d3c2c6) 2017-10-17_05:12:14PM
  -list
        Do you want to just list running containers?
  -name string
        the name of the container you want to kill
  -version
        Version Info
```

## TODO
- [ ] Refactor to use Services rather than just running containers
- [ ] Add flag to choose how many instances to blow up
- [ ] Add output option to output as JSON
