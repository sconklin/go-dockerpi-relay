Golang library to interface to the DockerPi 4 channel relay I2C board
==============================================================

[![Build Status](https://travis-ci.org/sconklin/go-dockerpi-relay.svg?branch=master)](https://travis-ci.org/sconklin/go-dockerpi-relay)
[![Go Report Card](https://goreportcard.com/badge/github.com/sconklin/go-dockerpi-relay)](https://goreportcard.com/report/github.com/sconklin/go-dockerpi-relay)
[![GoDoc](https://godoc.org/github.com/sconklin/go-dockerpi-relay?status.svg)](https://godoc.org/github.com/sconklin/go-dockerpi-relay)
[![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

This library written in [Go programming language](https://golang.org/) to control a [DockerPi relay board](https://github.com/geeekpi/dockerpi).

![image](https://raw.github.com/sconklin/go-dockerpi-relay/master/docs/300px-Ddl-1.jpg)
![image](https://raw.github.com/sconklin/go-dockerpi-relay/master/docs/300px-Ddl-2.jpg)
![image](https://raw.github.com/sconklin/go-dockerpi-relay/master/docs/300px-Ddl-3.jpg)

Compatibility
-------------

Tested on Raspberry PI Zero W

Golang usage
------------

```go
func main() {
	i2c, err := i2c.NewI2C(0x20, 1)
	checkError(err)
	defer i2c.Close()
	rly, err := device.NewRelay(i2c)
	checkError(err)

	time.Sleep(1 * time.Second)
	fmt.Println("One ON")
	err = rly.On(1)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("One OFF")
	err = rly.Off(1)

	time.Sleep(1 * time.Second)
	fmt.Println("Two ON")
	err = rly.On(2)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("Two OFF")
	err = rly.Off(2)

	time.Sleep(1 * time.Second)
	fmt.Println("Three ON")
	err = rly.On(3)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("Three OFF")
	err = rly.Off(3)

	time.Sleep(1 * time.Second)
	fmt.Println("Four ON")
	err = rly.On(4)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("Four OFF")
	err = rly.Off(4)
}
```

Getting help
------------

GoDoc [documentation](http://godoc.org/github.com/sconklin/go-dockerpi-relay)

Installation
------------

```bash
$ go get -u github.com/sconklin/go-dockerpi-relay
```

Troubleshoting
--------------

- *How to obtain fresh Golang installation to RPi device (either any RPi clone):*
If your RaspberryPI golang installation taken by default from repository is outdated, you may consider
to install actual golang mannualy from official Golang [site](https://golang.org/dl/). Download
tar.gz file containing armv6l in the name. Follow installation instructions.

- *How to enable I2C bus on RPi device:*
If you employ RaspberryPI, use raspi-config utility to activate i2c-bus on the OS level.
Go to "Interfaceing Options" menu, to active I2C bus.
Probably you will need to reboot to load i2c kernel module.
Finally you should have device like /dev/i2c-1 present in the system.

- *How to find I2C bus allocation and device address:*
Use i2cdetect utility in format "i2cdetect -y X", where X may vary from 0 to 5 or more,
to discover address occupied by peripheral device. To install utility you should run
`apt install i2c-tools` on debian-kind system. `i2cdetect -y 1` sample output:
	```
	     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
	00:          -- -- -- -- -- -- -- -- -- -- -- -- --
	10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	70: -- -- -- -- -- -- 76 --    
	```

> NOTE 1: Library is not goroutine-safe, so use synchronization approach when multi-gorutine output expected to display in application.

> NOTE 2: If you experience issue with lcd-device stability play with strobe delays in routine `writeDataWithStrobe(data byte)`. Default settings: 200 ms (microseconds) for setting stober, and 30 ms for exposing it to zero. Try to increase them a little bit, if you expirience any malfunction.

Credits
-------

Contact
-------

Please use [Github issue tracker](https://github.com/sconklin/go-dockerpi-relay/issues) for filing bugs or feature requests.

License
-------

go-dockerpi-relay is licensed inder MIT License.
