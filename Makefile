.PHONY: dev
dev:
	@echo "Starting..."
	cd backend && air

.PHONY: build_front
build_front:
	@echo "Forcing frontend build"
	$(eval RELATIVE := $(filter true, $(relative)))
	@if [ "$(RELATIVE)" = "true" ]; then \
		cd ../frontend && npm run build; \
	else \
		cd frontend && npm run build; \
	fi
