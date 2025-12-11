import { Page, Locator, expect } from '@playwright/test';

/**
 * Page Object Model for Particle Symphony
 * Provides clean API for all UI interactions
 */
export class ParticleSymphonyPage {
  readonly page: Page;
  readonly canvas: Locator;
  readonly loadingIndicator: Locator;
  readonly errorMessage: Locator;

  constructor(page: Page) {
    this.page = page;
    this.canvas = page.locator('canvas');
    this.loadingIndicator = page.locator('#loading');
    this.errorMessage = page.locator('#error');
  }

  /**
   * Navigate to the application and wait for WASM to load
   */
  async goto() {
    await this.page.goto('/');
    await this.waitForWasmLoad();
  }

  /**
   * Wait for WASM module to fully load and initialize
   * STRICT TIMEOUTS: If WASM takes too long, the app is broken!
   */
  async waitForWasmLoad() {
    // Canvas must appear within 5s - otherwise something is wrong
    await this.canvas.waitFor({ state: 'visible', timeout: 5000 });
    
    // Loading indicator must disappear quickly (2s)
    try {
      await this.loadingIndicator.waitFor({ state: 'hidden', timeout: 2000 });
    } catch {
      // Loading indicator might not exist or already be hidden
    }
    
    // WASM must be ready within 10s - this is already generous!
    await this.page.waitForFunction(
      () => {
        const w = window as unknown as { wasmReady?: boolean };
        return w.wasmReady === true;
      },
      { timeout: 10000 }
    );
  }

  /**
   * Switch to a specific preset (1-5)
   */
  async switchPreset(preset: 1 | 2 | 3 | 4 | 5) {
    await this.page.keyboard.press(`Digit${preset}`);
    // Wait for preset transition
    await this.page.waitForTimeout(500);
  }

  /**
   * Toggle debug overlay (F3)
   */
  async toggleDebugOverlay() {
    await this.page.keyboard.press('F3');
    await this.page.waitForTimeout(100);
  }

  /**
   * Toggle glow effect (F5)
   */
  async toggleGlow() {
    await this.page.keyboard.press('F5');
    await this.page.waitForTimeout(100);
  }

  /**
   * Toggle motion blur (F4)
   */
  async toggleMotionBlur() {
    await this.page.keyboard.press('F4');
    await this.page.waitForTimeout(100);
  }

  /**
   * Click on canvas at specific position
   */
  async clickCanvas(x: number, y: number, button: 'left' | 'right' = 'left') {
    const box = await this.canvas.boundingBox();
    if (!box) throw new Error('Canvas not found');
    
    await this.canvas.click({
      position: { x, y },
      button,
    });
  }

  /**
   * Drag on canvas from one position to another
   */
  async dragOnCanvas(
    from: { x: number; y: number },
    to: { x: number; y: number }
  ) {
    const box = await this.canvas.boundingBox();
    if (!box) throw new Error('Canvas not found');

    await this.page.mouse.move(box.x + from.x, box.y + from.y);
    await this.page.mouse.down();
    await this.page.mouse.move(box.x + to.x, box.y + to.y, { steps: 10 });
    await this.page.mouse.up();
  }

  /**
   * Move mouse to position on canvas without clicking
   */
  async moveMouseOnCanvas(x: number, y: number) {
    const box = await this.canvas.boundingBox();
    if (!box) throw new Error('Canvas not found');
    
    await this.page.mouse.move(box.x + x, box.y + y);
  }

  /**
   * Get current canvas snapshot for visual comparison
   */
  async getCanvasSnapshot(): Promise<Buffer> {
    return await this.canvas.screenshot();
  }

  /**
   * Check if application is running without errors
   */
  async isHealthy(): Promise<boolean> {
    const wasmReady = await this.page.evaluate(() => {
      const w = window as unknown as { wasmReady?: boolean };
      return w.wasmReady === true;
    });
    const canvasVisible = await this.canvas.isVisible();
    return wasmReady && canvasVisible;
  }

  /**
   * Get canvas dimensions
   */
  async getCanvasDimensions(): Promise<{ width: number; height: number }> {
    const box = await this.canvas.boundingBox();
    if (!box) throw new Error('Canvas not found');
    return { width: box.width, height: box.height };
  }

  /**
   * Perform touch tap on canvas (for mobile simulation)
   */
  async tapCanvas(x: number, y: number) {
    const box = await this.canvas.boundingBox();
    if (!box) throw new Error('Canvas not found');
    
    await this.page.touchscreen.tap(box.x + x, box.y + y);
  }
}
