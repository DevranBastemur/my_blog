document.addEventListener('DOMContentLoaded', () => {
    const container = document.querySelector('.bg-animation-container');
    const logo = document.querySelector('.floating-logo');

    if (!container || !logo) {
        return;
    }

    let x = Math.random() * (container.clientWidth - logo.clientWidth);
    let y = Math.random() * (container.clientHeight - logo.clientHeight);
    let dx = (Math.random() - 0.5) * 1; 
    let dy = (Math.random() - 0.5) * 1;

    function animate() {
        x += dx;
        y += dy;

        if (x + logo.clientWidth > container.clientWidth || x < 0) {
            dx = -dx;
        }
        if (y + logo.clientHeight > container.clientHeight || y < 0) {
            dy = -dy;
        }

        logo.style.transform = `translate(${x}px, ${y}px)`;

        requestAnimationFrame(animate);
    }

    animate();
});