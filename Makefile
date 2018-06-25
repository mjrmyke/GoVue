PROJECT = "GoVue Server"
current_dir = $(shell pwd)

prod:
	cd VueApp && npm install
	cd VueApp && npm run build
	go build -o GoVue
	./GoVue

make dev:
	make devfront devback -j2

devfront:
	cd VueApp && npm install
	cd VueApp && npm run dev

devback:
	go build -o GoVue
	./GoVue -dev