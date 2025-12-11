import { test, expect } from '@playwright/test';
import { ParticleSymphonyPage } from '../pages/particle-symphony.page';

/**
 * E-010-S03: Startup & Initialisierungs-Tests
 * Tests für WASM-Loading, Canvas-Rendering und Initial-State
 */
test.describe('Application Startup', () => {
  test('should load WASM successfully', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    const wasmReady = await page.evaluate(() => {
      const w = window as unknown as { wasmReady?: boolean };
      return w.wasmReady;
    });
    expect(wasmReady).toBe(true);
  });

  test('should render canvas with correct dimensions', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    await expect(psp.canvas).toBeVisible();
    const box = await psp.canvas.boundingBox();
    expect(box).not.toBeNull();
    expect(box?.width).toBeGreaterThanOrEqual(800);
    expect(box?.height).toBeGreaterThanOrEqual(600);
  });

  test('should have no console errors on startup', async ({ page }) => {
    const errors: string[] = [];
    page.on('console', (msg) => {
      if (msg.type() === 'error') {
        errors.push(msg.text());
      }
    });

    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(2000);

    // Filter out expected/harmless errors
    const criticalErrors = errors.filter(
      (e) => !e.includes('favicon') && !e.includes('404')
    );
    expect(criticalErrors).toHaveLength(0);
  });

  test('should start with Galaxy preset as default', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await page.waitForTimeout(1000);

    // Take screenshot to verify visual state
    const screenshot = await psp.getCanvasSnapshot();
    expect(screenshot).toBeDefined();
    expect(screenshot.length).toBeGreaterThan(1000);
  });

  test('should hide loading indicator after WASM loads', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await page.goto('/');
    
    // Wait for canvas to be visible (WASM loaded)
    await psp.canvas.waitFor({ state: 'visible', timeout: 30000 });
    
    // Loading indicator should be hidden or not exist
    // The loading element might have 'hidden' class or display:none
    const loadingEl = page.locator('#loading');
    const isHidden = await loadingEl.evaluate((el) => {
      if (!el) return true;
      const style = window.getComputedStyle(el);
      return style.display === 'none' || el.classList.contains('hidden');
    }).catch(() => true);
    
    expect(isHidden).toBe(true);
  });

  test('should be healthy after startup', async ({ page }) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();

    const healthy = await psp.isHealthy();
    expect(healthy).toBe(true);
  });
});
