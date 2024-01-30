import * as path from 'path'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import codegen from 'vite-plugin-graphql-codegen'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 3000,
  },
  plugins: [react(), { ...codegen(), apply: 'serve' }],
  resolve: {
    alias: [
      {
        find: '@',
        replacement: path.resolve(__dirname, 'src'),
      },
    ],
  },
})
