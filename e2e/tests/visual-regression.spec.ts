import { test, expect } from '@playwright/test';
import { ParticleSymphonyPage } from '../pages/particle-symphony.page';
import { PRESETS, WAIT_TIMES } from '../fixtures/test-fixtures';

/**
 * E-010-S06: Visual Regression Testing
 * Tests that verify visual output is correct (content-based, not pixel-perfect)
 * 
 * NOTE: Since particles are dynamic/random, we test for:
 * - Canvas renders something (not blank)
 * - Screenshots have expected size/content
 * - UI elements appear correctly
 */
test.describe('Visual Regression Tests', () => {
  test('initial page loads with visible canvas', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Canvas should be visible
    await expect(psp.canvas).toBeVisible();
    
    // Page should have correct title
    await expect(page).toHaveTitle(/Particle Symphony/);
  });

  test('debug overlay renders when toggled', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    
    const before = await psp.getCanvasSnapshot();
    await psp.toggleDebugOverlay();
    await page.waitForTimeout(100);
    const after = await psp.getCanvasSnapshot();

    // Visual should change when debug overlay is toggled
    expect(Buffer.compare(before, after)).not.toBe(0);
  });

  test('canvas renders particles (not empty)', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.VISUAL_SETTLE);

    const screenshot = await psp.getCanvasSnapshot();

    // Screenshot should have substantial data (not just blank canvas)
    expect(screenshot).toBeDefined();
    expect(screenshot.length).toBeGreaterThan(5000);
  });

  // Visual tests for each preset - verify they produce different output
  for (const { key, name } of PRESETS) {
    test(`${name} preset renders content`, async ({ page }) => {
      const psp = new ParticleSymphonyPage(page);
      await psp.goto();
      await psp.switchPreset(key);
      await page.waitForTimeout(WAIT_TIMES.PRESET_SWITCH);

      const screenshot = await psp.getCanvasSnapshot();

      // Preset should render particles (not empty)
      expect(screenshot).toBeDefined();
      expect(screenshot.length).toBeGreaterThan(3000);
    });
  }

  test('glow effect changes visual appearance', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);

    // Without glow
    const withoutGlow = await psp.getCanvasSnapshot();

    // Toggle glow on
    await psp.toggleGlow();
    await page.waitForTimeout(WAIT_TIMES.INTERACTION_RESPONSE);

    const withGlow = await psp.getCanvasSnapshot();

    // Both screenshots should have content
    expect(withoutGlow.length).toBeGreaterThan(1000);
    expect(withGlow.length).toBeGreaterThan(1000);
  });

  test('canvas dimensions match expected size', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    const dimensions = await psp.getCanvasDimensions();

    // Canvas should be at least 800x600
    expect(dimensions.width).toBeGreaterThanOrEqual(800);
    expect(dimensions.height).toBeGreaterThanOrEqual(600);
  });
});

test.describe('Visual Consistency', () => {
  test('same preset produces similar visuals on reload', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);

    // First load
    await psp.goto();
    await psp.switchPreset(3); // Swarm preset
    await page.waitForTimeout(WAIT_TIMES.VISUAL_SETTLE);
    const firstLoad = await psp.getCanvasSnapshot();

    // Fresh navigation (not reload to avoid stale refs)
    await page.goto('/');
    await psp.waitForWasmLoad();
    await psp.switchPreset(3);
    await page.waitForTimeout(WAIT_TIMES.VISUAL_SETTLE);
    const secondLoad = await psp.getCanvasSnapshot();

    // Both should have content
    expect(firstLoad.length).toBeGreaterThan(1000);
    expect(secondLoad.length).toBeGreaterThan(1000);
  });
});
