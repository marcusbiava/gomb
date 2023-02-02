.PHONY: update-tag

update-tag:
	@echo "Atualizando tag do Git"
	@last_tag=$$(git ls-remote --tags origin | awk -F"/" '{print $$3}' | sort -V | tail -1); \
  	next_tag=$$(echo "$$last_tag" | awk -F"." '{print $$1"."$$2"."$$3+1}'); \
	git tag -a "$$next_tag" -m "Nova tag: $$next_tag"; \
	git push origin "$$next_tag"
