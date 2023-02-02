.PHONY: update_tag

update_tag:
	CURRENT_TAG=$$(git describe --abbrev=0 --tags); \
	VERSION=$$(echo $${CURRENT_TAG} | grep -E -o '[0-9]+$$'); \
	NEW_VERSION=$$(expr $${VERSION} + 1); \
	NEW_TAG=v0.0.$${NEW_VERSION}; \
	git tag -a $${NEW_TAG} -m "Nova tag: $${NEW_TAG}"; \
	git push origin $${NEW_TAG}



