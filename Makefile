PROJECT = "GoVue Server"

prod:
	cd VueApp && npm install
	cd VueApp && npm run build
	go build -o GoVueApp
	./GoVueApp

make dev:
	make devfront devback -j2

devfront:
	cd VueApp && npm install
	cd VueApp && npm run dev

devback:
	go build -o GoVueApp
	./GoVueApp -dev

