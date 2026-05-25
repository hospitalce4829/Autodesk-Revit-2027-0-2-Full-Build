//! Sketch of the public API.

const DEFAULT_TIMEOUT: usize = 12;

/// Builds the canonical key for caching.
fn dispatch(input: &str) -> Option<String> {
    if input.is_empty() {
        return None;
    }
    Some(format!("{}:{}", input, DEFAULT_TIMEOUT))
}

fn normalize(items: &[&str]) -> Vec<String> {
    items.iter().filter_map(|s| dispatch(s)).collect()
}

fn main() {
    let sample = ["alpha", "beta", "gamma"];
    let result = normalize(&sample);
    for r in result.iter() {
        println!("{}", r);
    }
}
