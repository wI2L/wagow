# wagow

A simple HTTP service that sends Wake-on-Wan magic packets.

### How-to

The service expose a single route at `POST /` that sends a wake-on-wan magic packet to the destination machine. The parameters can be sent using either `application/json` or `application/x-form-www-urlencoded` content types.

   - `address`: IP address or FQDN of the machine in the format _host:port_
   - `target`: MAC-48 hardware address of the machine
   - `password`: the SecureOn password (_optional_)

If the destination address does not include a port, a default one (9) will be added automatically.

## License

Copyright (c) 2017, William Poussier (MIT Licence)
