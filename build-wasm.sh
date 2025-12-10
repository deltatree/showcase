#!/bin/bash
# Build script for WebAssembly version of Particle Symphony
# Uses Ebitengine for WASM-compatible rendering

set -e

echo "🚀 Building Particle Symphony WASM..."

# Create web directory if not exists
mkdir -p web

# Copy wasm_exec.js from Go installation (try different locations)
echo "📋 Copying wasm_exec.js..."
WASM_EXEC=""
for path in "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$(go env GOROOT)/lib/wasm/wasm_exec.js"; do
    if [ -f "$path" ]; then
        WASM_EXEC="$path"
        break
    fi
done

if [ -z "$WASM_EXEC" ]; then
    echo "❌ Error: wasm_exec.js not found"
    exit 1
fi

cp "$WASM_EXEC" web/

# Build WASM binary
echo "🔨 Compiling to WebAssembly..."
cd cmd/wasm
GOOS=js GOARCH=wasm go build -o ../../web/particle-symphony.wasm .
cd ../..

# Get WASM file size
WASM_SIZE=$(ls -lh web/particle-symphony.wasm | awk '{print $5}')
echo "✅ Build complete!"
echo "📦 WASM size: $WASM_SIZE"
echo ""
echo "🌐 To test locally:"
echo "   cd web && python3 -m http.server 8080"
echo "   Then open http://localhost:8080"
