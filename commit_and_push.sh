echo "First arg: $1"
if [ $1 == '' ]
then
	echo "please provide the commit name"
else
	git add .
	git commit -m $1
	git push origin master
	filename='../gittok.txt'
	pass=""
	while read line; do
		# reading each line
		pass=$line
		n=$((n+1))
	done < $filename
	sleep 1
	echo   "m.dowlatabadi.ce@gmail.com"
	read -n1 KEY
	sleep 1
	echo $pass
	read -n1 KEY
	echo "Press enter if commit is ready to be pushed!"
fi

