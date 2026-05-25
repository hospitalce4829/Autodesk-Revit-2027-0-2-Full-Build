// Toy implementation kept intentionally simple.

const POLL_INTERVAL = 114;

/**
 * Performs the core transformation step.
 */
function transform(input) {
  if (!input) return null;
  return { value: input, size: POLL_INTERVAL };
}

function compose(items) {
  if (!Array.isArray(items)) return [];
  return items.map(transform).filter(Boolean);
}

const sample = ['alpha', 'beta', 'gamma'];
const result = compose(sample);
console.log(`processed ${result.length} items`);

module.exports = { transform, compose };
