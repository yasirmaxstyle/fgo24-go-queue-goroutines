# ⏳ Go Concurrency Queue Simulation
This project is a simple Go-based simulation of concurrent queue processing using `goroutines`, `channels`, and `WaitGroups`. It demonstrates how multiple tasks (queues) can run in parallel and report their individual execution durations.

## Features
- Spawns a random number of "queues" (between 1–10).
- Each queue:
  - Sleeps for a random delay (1–5 seconds).
  - Sends the duration back through a `channel`.
- Results are printed in the order they complete (not necessarily the order they were started).
- Execution time is shown in seconds with nanosecond precision.

## Concepts Demonstrated
- Goroutines
- Channels
- WaitGroups
- Randomized execution with `math/rand`
- Time measurement with `time` package
- Nanosecond to second conversion using `1e9`

## How to run this program
### Prerequisite
- [Go Programming Language](https://go.dev/)

### Manual

1. Clone this repository
```bash
git clone https://github.com/yasirmaxstyle/fgo24-go-json-unmarshall.git
```
2. Get into the path
```bash
cd fgo24-go-json-unmarshall
```
3. Run the program
```bash
go run main.go
```

## Technologies
- [Go Programming Language](https://go.dev/)

## How to take part in this project
You are free to fork this project, make improvement and submit a pull request to improve this project. If you find this useful or if you have suggestion, you can start discussing through my social media below.
- [Instagram](https://www.instagram.com/yasirmaxstyle/)
- [LinkedIn](https://www.linkedin.com/in/muhamad-yasir-806230117/)

## License
This project is under MIT License
