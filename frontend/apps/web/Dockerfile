FROM oven/bun AS build

WORKDIR /app

# Copy root workspace files
COPY package.json bun.lock ./

# Copy packages directory (includes api-client)
COPY packages ./packages

# Copy web app
COPY apps/web ./apps/web

# Install all workspace dependencies
RUN bun install

# Build the web app
WORKDIR /app/apps/web
RUN bun run build

# Production stage
FROM nginx:alpine

COPY --from=build /app/apps/web/dist /usr/share/nginx/html
COPY apps/web/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]