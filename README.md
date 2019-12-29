# janumg20cm-modbus-tcp


Read modbus arguments from 1000 to 1416 address, and returns a json object.

Programm flags:

-ip - janitza ip address (defaut value "localhost");

-port - janitza modbus tcp port (defaut value 502);

-id - janitza modbus slave ID (defaut value 1);

-type -  from 0 to 5 (defaut value 0).

Example:

`janumg20cm_modbus_tcp -ip=192.168.10.10 -port=502 -id=2 -type=1`
