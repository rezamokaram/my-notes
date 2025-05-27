# io-max-nr in Linux & ScyllaDB

## What it is

io-max-nr is a Linux kernel setting (/proc/sys/fs/aio-max-nr) that defines the maximum number of allowed concurrent AIO (Asynchronous I/O) requests system-wide.  
more on [doc](https://www.kernel.org/doc/Documentation/sysctl/fs.txt)

## In ScyllaDB

ScyllaDB uses Linux AIO for high-performance disk I/O. If io-max-nr is too low, ScyllaDB may hit a limit and throw errors or throttle performance.

## Default

Often set to 65536 by default — usually too low for ScyllaDB on modern hardware.

## Recommended

ScyllaDB recommends setting it to a much higher value, like 1048576, during tuning.

## To View or Set It

```bash
# View current value
cat /proc/sys/fs/aio-max-nr

# Set it temporarily
echo 1048576 > /proc/sys/fs/aio-max-nr

# Set it permanently (add to /etc/sysctl.conf)
fs.aio-max-nr = 1048576
sysctl -p /etc/sysctl.conf
```
