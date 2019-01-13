A simple Go script for sending SSH login notifications via Telegram.


Building:\
`go build -ldflags "-X main.telegramToken=XXXX:YYYYY -X main.userID=YOUR_TELEGRAM_USER_ID -X main.geoApiKey=ABCD"`

Execute the following commands on SSH login (put them in `/etc/ssh/sshrc`):

``ip=`echo $SSH_CONNECTION | cut -d " " -f 1``\
`/path/to/build SSH your_vps_name $ip`