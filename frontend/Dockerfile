FROM node:18-slim as builder
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable && corepack prepare pnpm@latest --activate
COPY . /app
WORKDIR /app
# Expose buildtime args
ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=${VITE_API_BASE_URL}
RUN pnpm install --frozen-lockfile
RUN pnpm run build

FROM caddy:2.7.4-alpine
# COPY --from=prod-deps /app/node_modules /app/node_modules
COPY --from=builder /app/dist /app/dist
EXPOSE 80 443