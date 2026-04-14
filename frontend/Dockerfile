FROM oven/bun AS build
WORKDIR /app
COPY package.json bun.lock ./
COPY scripts ./scripts
COPY packages ./packages
COPY apps/web ./apps/web
RUN bun install
WORKDIR /app/apps/web
ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL
RUN bun run build

FROM nginx:alpine
COPY --from=build /app/apps/web/dist /usr/share/nginx/html
COPY apps/web/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
