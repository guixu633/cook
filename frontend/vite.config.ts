import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://43.135.156.166:8080', // Updated to remote server IP
        changeOrigin: true,
      },
    },
  },
})
