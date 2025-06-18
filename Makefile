#Докер скрипт для запуска gigachat-adapter
gigachat:
	docker run -p 8000:8000 \
  -e GIGACHAT_CREDENTIALS=ZTIxNzM3NDktZjAxOC00YWQyLWJiOWEtNzc0Mjc4YmY1MzBlOmNiMzc5YTAxLTgwODMtNDAzNC04MWVkLWY2OGMwZmMxZDdlYg== \
  -e BEARER_TOKEN=secret123 \
  -e GIGACHAT_VERIFY_SSL_CERTS=False \
  antonk0/gigachat-adapter:latest