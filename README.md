# sleepuntil
<img src="https://travis-ci.org/jnidzwetzki/sleepuntil.svg?branch=master">

This program acts as the famous ``sleep`` unix command. The main difference is,  that ``sleepuntil`` sleeps until a certain time is reached. ``sleep`` in contrast sleeps a fixed amount of seconds.

``sleepuntil`` was written to take over simple jobs in maintenance work (e.g., shutdown the mysql server at 06:00 am).

## Examples
```bash
# Sleep until 04:00 and shutdown the MySQL server afterwards
sleepuntil 04:00; service mysql stop

# Sleep until 2019-01-02 15:04:05, show the remaining time and restart the apcahe webserver afterwards
sleepuntil 2019-01-02 15:04:05 -animate; service apache2 restart
```