// Simple script to create placeholder PNG icons
// Run in browser console or Node.js to generate

const sizes = [72, 96, 128, 144, 152, 192, 384, 512];

sizes.forEach(size => {
  const canvas = document.createElement('canvas');
  canvas.width = size;
  canvas.height = size;
  const ctx = canvas.getContext('2d');
  
  // Gradient background
  const gradient = ctx.createLinearGradient(0, 0, size, size);
  gradient.addColorStop(0, '#667eea');
  gradient.addColorStop(0.5, '#764ba2');
  gradient.addColorStop(1, '#f093fb');
  
  // Rounded rect
  const r = size * 0.15;
  ctx.beginPath();
  ctx.moveTo(r, 0);
  ctx.lineTo(size - r, 0);
  ctx.quadraticCurveTo(size, 0, size, r);
  ctx.lineTo(size, size - r);
  ctx.quadraticCurveTo(size, size, size - r, size);
  ctx.lineTo(r, size);
  ctx.quadraticCurveTo(0, size, 0, size - r);
  ctx.lineTo(0, r);
  ctx.quadraticCurveTo(0, 0, r, 0);
  ctx.closePath();
  ctx.fillStyle = gradient;
  ctx.fill();
  
  // Particles
  ctx.fillStyle = 'white';
  const particles = [
    { x: 0.5, y: 0.4, r: 0.047 },
    { x: 0.39, y: 0.55, r: 0.035 },
    { x: 0.61, y: 0.55, r: 0.035 },
    { x: 0.5, y: 0.62, r: 0.027 },
    { x: 0.35, y: 0.43, r: 0.019 },
    { x: 0.65, y: 0.43, r: 0.019 },
  ];
  particles.forEach(p => {
    ctx.beginPath();
    ctx.arc(p.x * size, p.y * size, p.r * size, 0, Math.PI * 2);
    ctx.fill();
  });
  
  console.log(`icon-${size}.png: ${canvas.toDataURL('image/png')}`);
});
