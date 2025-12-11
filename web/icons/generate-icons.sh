#!/bin/bash
# Generate PWA icons from SVG

# Create a simple SVG icon for Particle Symphony
cat > icon-src.svg << 'SVG'
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
  <defs>
    <linearGradient id="bg" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" style="stop-color:#667eea"/>
      <stop offset="50%" style="stop-color:#764ba2"/>
      <stop offset="100%" style="stop-color:#f093fb"/>
    </linearGradient>
    <radialGradient id="glow" cx="50%" cy="50%" r="50%">
      <stop offset="0%" style="stop-color:white;stop-opacity:0.8"/>
      <stop offset="100%" style="stop-color:white;stop-opacity:0"/>
    </radialGradient>
  </defs>
  <rect width="512" height="512" rx="80" fill="url(#bg)"/>
  <circle cx="256" cy="256" r="120" fill="url(#glow)" opacity="0.5"/>
  <!-- Particle dots -->
  <circle cx="256" cy="200" r="24" fill="white"/>
  <circle cx="200" cy="280" r="18" fill="white" opacity="0.9"/>
  <circle cx="312" cy="280" r="18" fill="white" opacity="0.9"/>
  <circle cx="256" cy="320" r="14" fill="white" opacity="0.7"/>
  <circle cx="180" cy="220" r="10" fill="white" opacity="0.6"/>
  <circle cx="332" cy="220" r="10" fill="white" opacity="0.6"/>
  <circle cx="220" cy="180" r="8" fill="white" opacity="0.5"/>
  <circle cx="292" cy="180" r="8" fill="white" opacity="0.5"/>
  <circle cx="160" cy="256" r="6" fill="white" opacity="0.4"/>
  <circle cx="352" cy="256" r="6" fill="white" opacity="0.4"/>
  <!-- Sparkles -->
  <circle cx="140" cy="140" r="5" fill="white" opacity="0.8"/>
  <circle cx="372" cy="140" r="4" fill="white" opacity="0.7"/>
  <circle cx="140" cy="372" r="4" fill="white" opacity="0.7"/>
  <circle cx="372" cy="372" r="5" fill="white" opacity="0.8"/>
</svg>
SVG

# Generate PNG icons at different sizes using available tools
for size in 72 96 128 144 152 192 384 512; do
  if command -v convert &> /dev/null; then
    convert -background none -resize ${size}x${size} icon-src.svg icon-${size}.png
    echo "Created icon-${size}.png"
  elif command -v rsvg-convert &> /dev/null; then
    rsvg-convert -w $size -h $size icon-src.svg -o icon-${size}.png
    echo "Created icon-${size}.png"
  fi
done

# Create screenshots placeholder (just colored rectangles for now)
if command -v convert &> /dev/null; then
  convert -size 1280x720 gradient:'#667eea'-'#f093fb' -font Helvetica -pointsize 48 -fill white -gravity center -annotate 0 'Particle Symphony\nECS Showcase' screenshot-wide.png
  convert -size 390x844 gradient:'#667eea'-'#f093fb' -font Helvetica -pointsize 32 -fill white -gravity center -annotate 0 'Particle\nSymphony' screenshot-mobile.png
  echo "Created screenshots"
fi

echo "Icon generation complete!"
