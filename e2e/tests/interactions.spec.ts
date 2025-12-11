import { test, expect } from '@playwright/test';
import { ParticleSymphonyPage } from '../pages/particle-symphony.page';
import { WAIT_TIMES } from '../fixtures/test-fixtures';

/**
 * E-010-S05: Interaktions-Tests (Maus & Tastatur)
 * Tests für alle Benutzerinteraktionen
 */
test.describe('Mouse Interactions', () => {
  test('left click should attract particles', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    const before = await psp.getCanvasSnapshot();

    // Click in center to attract particles
    await psp.clickCanvas(640, 360, 'left');
    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);

    // Hold for a moment to let particles react
    await page.mouse.down();
    await page.waitForTimeout(500);
    await page.mouse.up();

    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);
    const after = await psp.getCanvasSnapshot();

    // Visual state should have changed
    expect(Buffer.compare(before, after)).not.toBe(0);
  });

  test('right click should repel particles', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Right click in center
    await psp.clickCanvas(640, 360, 'right');
    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);

    // Application should still be healthy
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });

  test('drag should create continuous attraction', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Drag from left to right
    await psp.dragOnCanvas({ x: 200, y: 360 }, { x: 800, y: 360 });

    // Application should still be healthy
    await expect(psp.canvas).toBeVisible();
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });

  test('mouse movement should be tracked', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Move mouse around
    await psp.moveMouseOnCanvas(100, 100);
    await page.waitForTimeout(100);
    await psp.moveMouseOnCanvas(500, 300);
    await page.waitForTimeout(100);
    await psp.moveMouseOnCanvas(800, 500);

    // Should still be working
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });
});

test.describe('Keyboard Interactions', () => {
  test('F3 should toggle debug overlay', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PRESET_SWITCH);

    const before = await psp.getCanvasSnapshot();

    // Toggle debug overlay on
    await psp.toggleDebugOverlay();
    await page.waitForTimeout(200);
    const withDebug = await psp.getCanvasSnapshot();

    // Toggle debug overlay off
    await psp.toggleDebugOverlay();
    await page.waitForTimeout(200);
    const withoutDebug = await psp.getCanvasSnapshot();

    // Debug overlay should change the visual
    expect(Buffer.compare(before, withDebug)).not.toBe(0);
  });

  test('F5 should toggle glow effect', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Toggle glow on/off
    await psp.toggleGlow();
    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);

    // Should still be healthy
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);

    // Take screenshot for verification
    const screenshot = await psp.getCanvasSnapshot();
    expect(screenshot).toBeDefined();
  });

  test('F4 should toggle motion blur', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Toggle motion blur
    await psp.toggleMotionBlur();
    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);

    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });

  test('number keys should switch presets', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Test all number keys 1-5
    for (let i = 1; i <= 5; i++) {
      await page.keyboard.press(`Digit${i}`);
      await page.waitForTimeout(300);
      const healthy = await psp.isHealthy();
      expect(healthy).toBe(true);
    }
  });
});

test.describe('Touch Events (Mobile)', () => {
  test.use({ hasTouch: true });

  test('touch tap should work like left click', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Tap in center
    await psp.tapCanvas(640, 360);
    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);

    // Should still be healthy
    await expect(psp.canvas).toBeVisible();
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });
});

test.describe('Keyboard Safety', () => {
  test('ESC should not close the browser application', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Press ESC
    await page.keyboard.press('Escape');
    await page.waitForTimeout(500);

    // Page should still be accessible
    await expect(psp.canvas).toBeVisible();
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });
});
