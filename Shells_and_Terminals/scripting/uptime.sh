#!/bin/bash

showuptime() {
  up=$(uptime -p | cut -c4-)
  since=$(uptime -s)
  bat << EOF
-----
This machine has been up for ${up}
It ha sbeen runnig sice ${since}
-----
EOF
}
showuptime
