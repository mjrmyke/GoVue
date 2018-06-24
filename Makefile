PROJECT = "GoVue Server"


prod:
	cd VueApp && npm install
	cd VueApp && npm run build
	go build
	./GoVue

make dev:
	make devfront devback -j2
	
devfront:
	cd VueApp && npm install
	cd VueApp && npm run dev

devback:
	go build
	./GoVue -dev