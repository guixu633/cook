import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'https://cook-api.guixuu.com', // Updated to remote server IP
        changeOrigin: true,
      },
    },
  },
})
