import { test, expect } from '@playwright/test';
import { ParticleSymphonyPage } from '../pages/particle-symphony.page';
import { PRESETS } from '../fixtures/test-fixtures';

/**
 * E-010-S07: Cross-Browser Testing
 * Tests für Chromium, Firefox und WebKit Kompatibilität
 */
test.describe('Cross-Browser Compatibility', () => {
  test('WASM loads correctly in all browsers', async ({ page, browserName }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    const wasmReady = await page.evaluate(() => {
      const w = window as unknown as { wasmReady?: boolean };
      return w.wasmReady;
    });

    expect(wasmReady).toBe(true);
    console.log(`✓ WASM loaded successfully in ${browserName}`);
  });

  test('canvas renders correctly in all browsers', async ({
    page,
    browserName,
  }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(2000);

    await expect(psp.canvas).toBeVisible();

    const screenshot = await psp.getCanvasSnapshot();
    expect(screenshot).toBeDefined();
    expect(screenshot.length).toBeGreaterThan(1000);

    console.log(`✓ Canvas renders correctly in ${browserName}`);
  });

  test('keyboard shortcuts work in all browsers', async ({
    page,
    browserName,
  }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Test all preset keys
    for (const preset of PRESETS) {
      await psp.switchPreset(preset.key);
      await page.waitForTimeout(200);
    }

    // Test toggle keys
    await psp.toggleDebugOverlay();
    await psp.toggleGlow();
    await psp.toggleMotionBlur();

    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);

    console.log(`✓ Keyboard shortcuts work in ${browserName}`);
  });

  test('mouse interactions work in all browsers', async ({
    page,
    browserName,
  }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(1000);

    // Left click
    await psp.clickCanvas(640, 360, 'left');
    await page.waitForTimeout(200);

    // Right click
    await psp.clickCanvas(640, 360, 'right');
    await page.waitForTimeout(200);

    // Mouse move
    await psp.moveMouseOnCanvas(400, 300);
    await page.waitForTimeout(100);

    await expect(psp.canvas).toBeVisible();
    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);

    console.log(`✓ Mouse interactions work in ${browserName}`);
  });

  test('application remains stable during extended use', async ({
    page,
    browserName,
  }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Quick stability test: switch presets and interact
    for (const preset of PRESETS) {
      await psp.switchPreset(preset.key);
      await page.waitForTimeout(100);
      await psp.clickCanvas(400, 300, 'left');
    }

    await psp.toggleDebugOverlay();
    await psp.toggleGlow();

    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);

    console.log(`✓ Application stable during extended use in ${browserName}`);
  });

  test('no JavaScript errors during operation', async ({
    page,
    browserName,
  }) => {
    const errors: string[] = [];
    page.on('console', (msg) => {
      if (msg.type() === 'error') {
        errors.push(msg.text());
      }
    });

    page.on('pageerror', (err) => {
      errors.push(err.message);
    });

    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    // Exercise the application
    for (const preset of PRESETS) {
      await psp.switchPreset(preset.key);
      await page.waitForTimeout(500);
    }

    // Filter harmless errors
    const criticalErrors = errors.filter(
      (e) =>
        !e.includes('favicon') &&
        !e.includes('404') &&
        !e.includes('ResizeObserver')
    );

    expect(criticalErrors).toHaveLength(0);

    console.log(`✓ No JavaScript errors in ${browserName}`);
  });
});

test.describe('Performance Baseline', () => {
  test('application loads within acceptable time', async ({
    page,
    browserName,
  }) => {
    const startTime = Date.now();

    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    const loadTime = Date.now() - startTime;

    // Should load within 30 seconds (generous for CI)
    expect(loadTime).toBeLessThan(30000);

    console.log(`✓ Load time in ${browserName}: ${loadTime}ms`);
  });

  test('canvas updates smoothly', async ({ page, browserName }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(1000);

    // Take multiple screenshots to verify animation
    const screenshots: Buffer[] = [];
    for (let i = 0; i < 3; i++) {
      screenshots.push(await psp.getCanvasSnapshot());
      await page.waitForTimeout(500);
    }

    // At least some screenshots should be different (animation is running)
    let differences = 0;
    for (let i = 1; i < screenshots.length; i++) {
      if (Buffer.compare(screenshots[i - 1], screenshots[i]) !== 0) {
        differences++;
      }
    }

    // Should have at least one difference (particles moving)
    expect(differences).toBeGreaterThanOrEqual(1);

    console.log(`✓ Canvas updates smoothly in ${browserName}`);
  });
});
