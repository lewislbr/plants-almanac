build-k8s:
	@docker build -t eu.gcr.io/plantdex-prod/plants --target prod plants && \
	docker build -t eu.gcr.io/plantdex-prod/users --target prod users && \
	docker build -t eu.gcr.io/plantdex-prod/web --target prod web

start:
	@docker compose -p plantdex up --build -d

start-%:
	@docker compose -p plantdex up --build $*

start-k8s: build-k8s
	@cd .kubernetes && \
	kubectl apply -f https://www.getambassador.io/yaml/aes-crds.yaml && \
	kubectl wait --for condition=established --timeout=90s crd -lproduct=aes && \
	kubectl apply -f https://www.getambassador.io/yaml/aes.yaml && \
	kubectl -n ambassador wait --for condition=available --timeout=90s deploy -lproduct=aes && \
	kubectl create namespace plantdex; \
	kubectl config set-context --current --namespace=plantdex && \
	kubectl apply -f users-secret.yaml && \
	kubectl apply -f plants-secret.yaml && \
	kubectl apply -f users-configmap.yaml && \
	kubectl apply -f plants-configmap.yaml && \
	kubectl apply -f web-configmap.yaml && \
	kubectl apply -f users.yaml && \
	kubectl apply -f plants.yaml && \
	kubectl apply -f web.yaml && \
	kubectl apply -f ingress.yaml && \
	cd ..

stop:
	@docker compose -p plantdex down

stop-k8s:
	@kubectl delete ns plantdex && \
	kubectl delete ns ambassador
