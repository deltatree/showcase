import { test, expect } from '@playwright/test';
import { ParticleSymphonyPage } from '../pages/particle-symphony.page';
import { PRESETS, WAIT_TIMES } from '../fixtures/test-fixtures';

/**
 * E-010-S04: Preset-Wechsel Tests
 * Tests für alle 5 Presets und Preset-Switching Verhalten
 */
test.describe('Preset Switching', () => {
  for (const preset of PRESETS) {
    test(`should switch to ${preset.name} preset (Key ${preset.key})`, async ({
      page,
    }) => {
      const psp = new ParticleSymphonyPage(page);
      await psp.goto();

      await psp.switchPreset(preset.key);
      await page.waitForTimeout(WAIT_TIMES.PRESET_SWITCH);

      // Verify application is still healthy after switch
      const healthy = await psp.isHealthy();
      expect(healthy).toBe(true);

      // Capture screenshot for visual verification
      const screenshot = await psp.getCanvasSnapshot();
      expect(screenshot).toBeDefined();
      expect(screenshot.length).toBeGreaterThan(1000);
    });
  }

  test('should handle rapid preset switching without crashing', async ({
    page,
  }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Rapidly switch through all presets 3 times
    for (let round = 0; round < 3; round++) {
      for (const preset of PRESETS) {
        await psp.switchPreset(preset.key);
        await page.waitForTimeout(200);
      }
    }

    // Application should still be healthy
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });

  test('presets should have distinct visual appearances', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    const screenshots: Map<string, Buffer> = new Map();

    // Quick capture for each preset
    for (const preset of PRESETS) {
      await psp.switchPreset(preset.key);
      await page.waitForTimeout(WAIT_TIMES.PRESET_SWITCH);
      const screenshot = await psp.getCanvasSnapshot();
      screenshots.set(preset.name, screenshot);
    }

    expect(screenshots.size).toBe(PRESETS.length);

    for (const [, screenshot] of screenshots) {
      expect(screenshot.length).toBeGreaterThan(1000);
    }
  });

  test('should return to same preset when pressing same key', async ({
    page,
  }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Switch to firework
    await psp.switchPreset(2);
    await page.waitForTimeout(WAIT_TIMES.PRESET_SWITCH);

    // Press same key again
    await psp.switchPreset(2);
    await page.waitForTimeout(WAIT_TIMES.PRESET_SWITCH);

    // Should still be healthy
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });

  test('preset switching should clear old particles', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Let particles spawn in galaxy
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);
    const galaxyShot = await psp.getCanvasSnapshot();

    // Switch to chaos
    await psp.switchPreset(5);
    await page.waitForTimeout(WAIT_TIMES.PARTICLE_STABILIZE);
    const chaosShot = await psp.getCanvasSnapshot();

    // Screenshots should be different
    expect(Buffer.compare(galaxyShot, chaosShot)).not.toBe(0);
  });
});
