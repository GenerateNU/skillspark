FROM oven/bun AS build
WORKDIR /app
# Copy root workspace files
COPY frontend/package.json frontend/bun.lock ./
# Copy packages directory (includes api-client)
COPY frontend/packages ./packages
# Copy web app
COPY frontend/apps/web ./apps/web
# Copy scripts
COPY scripts ./scripts
# Install all workspace dependencies
RUN bun install
# Build the web app
WORKDIR /app/apps/web
RUN bun run build
# Production stage
FROM nginx:alpine
COPY --from=build /app/apps/web/dist /usr/share/nginx/html
COPY frontend/apps/web/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]