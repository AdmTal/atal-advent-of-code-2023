NAME ?= $(shell bash -c 'read -p "Enter a name for the folder >> " name; echo $$name')

folder:
	cp -r day_x_template/ $(NAME)

