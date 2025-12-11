import { test as base } from '@playwright/test';
import { ParticleSymphonyPage } from '../pages/particle-symphony.page';

/**
 * Custom test fixtures for Particle Symphony tests
 */
type ParticleSymphonyFixtures = {
  particleSymphony: ParticleSymphonyPage;
};

/**
 * Extended test with Particle Symphony fixtures
 */
export const test = base.extend<ParticleSymphonyFixtures>({
  particleSymphony: async ({ page }, use) => {
    const psp = new ParticleSymphonyPage(page);
    await psp.goto();
    await use(psp);
  },
});

export { expect } from '@playwright/test';

/**
 * Preset definitions for testing
 */
export const PRESETS = [
  { key: 1 as const, name: 'galaxy', description: 'Spiral Galaxy Simulation' },
  { key: 2 as const, name: 'firework', description: 'Firework Explosions' },
  { key: 3 as const, name: 'swarm', description: 'Swarm Behavior' },
  { key: 4 as const, name: 'fountain', description: 'Particle Fountain' },
  { key: 5 as const, name: 'chaos', description: 'Chaos Mode' },
] as const;

/**
 * Wait durations for various operations
 * STRICT: These are MAX times - if UI takes longer, it's a bug!
 */
export const WAIT_TIMES = {
  WASM_LOAD: 5000,        // 5s max for WASM
  PRESET_SWITCH: 300,     // 300ms for preset switch
  PARTICLE_STABILIZE: 500, // 500ms for particles to settle
  INTERACTION_RESPONSE: 200, // 200ms for click response
  VISUAL_SETTLE: 1000,    // 1s for visual tests
} as const;
