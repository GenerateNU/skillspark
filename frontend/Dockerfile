FROM oven/bun AS build
WORKDIR /app
COPY package.json bun.lock ./
COPY packages ./packages
COPY apps/web ./apps/web
RUN bun install
WORKDIR /app/apps/web
RUN bun run build

FROM nginx:alpine
COPY --from=build /app/apps/web/dist /usr/share/nginx/html
COPY apps/web/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]