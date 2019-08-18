# Strip Invaders Command Line Control Tool - sicc

This tiny tool is used to send control sequences to a _Strip Invaders_ device. (see [Strip Invaders](https://github.com/neophob/StripInvaders))

To test the tool, use the backend fake tool _fakend_:

```bash
# Run this command in the fakend folder
➜  fakend git:(master) ✗ go run .
2019/08/18 12:54:02 Listening osc server on: 127.0.0.1:8765
2019/08/18 12:54:02 Press Ctrl-C to stop server

# Run this command to in the sicc folder to see all available commands
➜  sicc git:(master) ✗ go run . -h
Usage:
  main [OPTIONS]

Application Options:
  -s, --server= Name or IP adress to connect to
  -p, --port=   Port on which the server is listening
  -c, --color=  Color to use (ex. #a2ff13)
  -m, --mode=   Mode to use (0-15) (default: -1)

Help Options:
  -h, --help    Show this help message
```

**Compile**
Run this command to generate the binary:

```bash
➜  sicc git:(master) ✗ go build
```
