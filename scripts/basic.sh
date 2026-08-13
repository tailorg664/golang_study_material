#!/bin/bash

echo "Hello User"
# What git does : 1) push code, 2) pull request, 3) add origin
# Repo builder and automater for github
# Function 1 : to link github repo with local git;
# Function 2 : store all the user logs (demo)

#------Functions------

#---------------------

echo "What do you want to do $USER?"
echo "a) Create Repo link"
echo "b) etc."

read IN

if  [ "$IN" == "b" ]; then
	 echo "Working on system"
elif [ "$IN" == "a" ]; then
	# Read the repo link of the user and store it in REPO variable
	read -p "Paste Repo Link: " REPO
	# Github link verification command
	# go run scripts/main.go $REPO
	sleep 1s
	go run main.go $REPO
	sleep 2s
	echo "you can work on your repo now"
	# use the REPO link for large structured work
	# git init
	# git add .
	# git commit -m "Initial commit"
	# git branch -M main
	# git remote add origin $REPO
	# git remote -v
	# git push -u origin main
	# git pull origin main

fi
