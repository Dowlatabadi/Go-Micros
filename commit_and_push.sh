echo "First arg: $1"
if [ $1 == '' ]
then
	echo "please provide the commit name"
else
	git add .
	git commit -m $1
	git config user.name "m.dowlatabadi.ce@gmail.com"

	filename='../gittok.txt'
	pass=""
	while read line; do
		# reading each line
		pass=$line
		n=$((n+1))
	done < $filename
	echo "m.dowlatabadi.ce@gamil.com"
	echo $pass
	git push origin master 
fi

