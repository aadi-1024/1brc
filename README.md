## 1BRC - 1 Billion Row Challenge
I came across this challenge somewhere online: you had to read a file containing 1 Billion rows and aggregate the data, optimizing your approach as much as possible. This is my take on that challenge in Go, for ease I generated the data randomly by using a Python script (`scripts/gen_data.py`). Every row contains a randomly generated string (exactly 10k of those) along with an associated floating point value that was generated using a random gaussian distribution(random per "id", so one unique distribution per id)

### Procedure
- I generated two separate datasets, one containing 1 billion rows, for the actual challenge, and one containing 100M rows, to be able to quickly iterate on my solution.
- For a stable benchmark the process would run three times, and the time would be averaged out. In memory cache would be cleared beforehand.
- After every iteration, I look at the `pprof` generated profiles and figure out areas that can be optimized, trying to improve the benchmark every time.

### Benchmarks
All the benchmarks + development were performed on a Ryzen 5 5600U machine, with 24GB RAM, and Fedora 43 as the OS. Before running any benchmark, ensure that memory cache is cleared in order to prevent inconsistent results due to the dataset being cached. On Linux, run `sudo sh -c "echo 3 > /proc/sys/vm/drop_caches"`
| Version | Time for 100m | Time for 1B |
|---------|---------------|-------------|
|     v1    |    19.98s           |        221.88s     |
