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
	git push https://m.dowlatabadi.ce@gmail.com:$pass@myrepository.biz/file.git --all 
fi

