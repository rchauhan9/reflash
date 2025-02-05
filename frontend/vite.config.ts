import path from "path"
import react from "@vitejs/plugin-react"
import { defineConfig } from "vite"

const basenameProd = '/react-shadcn-starter'

export default defineConfig(({ command }) => {
  return {
    plugins: [react()],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
  }
})