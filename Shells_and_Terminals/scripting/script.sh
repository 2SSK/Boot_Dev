#!/bin/bash

# VARIABLES
 #name='saurav'
 #
 ## PRINTING VARIABLES
 #echo "Hello $name, what is your age?"
 #read age # Take input from user
 #
 ## CONDITIONAL STATEMENTS
 #if [[ $age -lt 18 ]]; then
   #echo "\nYou are a minor"
 #elif [[ $age -lt 60 ]]; then
   #echo "\nYou are an adult"
 #else  
   #echo "\nYou are a senior citizen"
 #fi
 #
 ## POSITIONAL ARGUMENTS
 #echo "\nname: $name\tage: $age"

#echo "\nEnter command : "
#read input

 #if [ ${1,,} = saurav ]; then
   #echo -e "\nOh, you're the boss here. Welcome!"
 #elif [ ${1,,} = help ]; then
   #echo -e "\nJust enter your username, duh!"
 #else
   #echo -e "\nI don't know who you are. But you're not hte boss of me!"
 #fi

# CASE STATEMENTS
 #case ${1,,} in 
   #saurav | adminstrator)
     #echo -e "\nHello, you're the boss here!"
     #;;
   #help)
     #echo -e "\nJust enter your username!"
     #;;
   #*)
     #echo "Hello there. You're not the boss of me. Enter a valid username!"
   #;;
 #esac

# LOOPS
MY_FIRST_LIST=(one  two three four five)

for i in ${MY_FIRST_LIST[@]}; do
  echo $i
done

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
