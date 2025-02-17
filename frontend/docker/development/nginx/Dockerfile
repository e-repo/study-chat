FROM nginx:1.25-alpine

COPY conf.d /etc/nginx/conf.d

WORKDIR /app

HEALTHCHECK --interval=30s --timeout=5s --start-period=1s CMD curl --fail http://127.0.0.1/health || exit 1