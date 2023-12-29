import { resolve } from "path";
import { defineConfig } from "vite";

export default defineConfig({
    build: {
        lib: {
            entry: [resolve(__dirname, "src/htmx.min.js")],
            formats: ["es"],
            name: ["name"],
            fileName: ["name"],
        },
        outDir: "web/js",
        emptyOutDir: false
    }
});