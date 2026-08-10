//! Benchmarks for the cost of canonicalizing tag sets before interning.
//!
//! The V3 writer sorts a metric's tag ids before interning them so that the same
//! logical tag set supplied in different orders deduplicates to a single dictionary
//! entry. These benchmarks quantify the cost of that sort on the cache-hit path,
//! where the tag set has already been interned and only the (added) sort is new work.
//!
//! Three groups:
//!  - `sort_isolated`: raw `slice::sort_unstable` on `n` shuffled `i64`s — an upper
//!    bound on the added per-call work, with no interning noise.
//!  - `set_tags_stable`: end-to-end `set_tags` cache hits where the caller always uses
//!    the same (insertion) order, so tag ids are already ascending and the sort takes
//!    its near-linear fast path. This is the realistic consistent-producer case.
//!  - `set_tags_unstable`: end-to-end `set_tags` cache hits where the caller uses a
//!    fixed but non-monotonic order, so the sort does genuine work every call. This is
//!    the worst realistic per-call overhead for a producer with unstable ordering.

use criterion::{criterion_group, criterion_main, BatchSize, BenchmarkId, Criterion};
use dd_metrics_v3::{V3MetricType, V3Writer};
use std::hint::black_box;

const SIZES: &[usize] = &[4, 16, 64, 256];

/// Builds `n` distinct `key:value` tag strings in insertion order.
fn make_tags(n: usize) -> Vec<String> {
    (0..n).map(|i| format!("key{i}:value{i}")).collect()
}

/// A fixed, deterministic non-monotonic permutation of `0..n` (step 7 is coprime to
/// every `n` in `SIZES`, so this is a genuine permutation).
fn scramble(n: usize) -> Vec<usize> {
    (0..n).map(|k| (k * 7 + 3) % n).collect()
}

/// Raw cost of the added `sort_unstable`, isolated from interning.
fn bench_sort_isolated(c: &mut Criterion) {
    let mut group = c.benchmark_group("sort_isolated");
    for &n in SIZES {
        // Shuffled ids so the sort cannot take a sorted/reverse-sorted fast path.
        let scrambled: Vec<i64> = scramble(n).into_iter().map(|i| i as i64 + 1).collect();
        group.bench_with_input(BenchmarkId::from_parameter(n), &n, |b, _| {
            b.iter_batched_ref(
                || scrambled.clone(),
                |v| v.sort_unstable(),
                BatchSize::SmallInput,
            );
        });
    }
    group.finish();
}

/// End-to-end `set_tags` on the cache-hit path with consistent (insertion) ordering.
/// Ids are already ascending, so the sort hits its near-linear fast path.
fn bench_set_tags_stable(c: &mut Criterion) {
    let mut group = c.benchmark_group("set_tags_stable");
    for &n in SIZES {
        let tags = make_tags(n);
        group.bench_with_input(BenchmarkId::from_parameter(n), &n, |b, _| {
            b.iter_batched_ref(
                || seeded_writer(&tags, &(0..n).collect::<Vec<_>>()),
                |w| write_metric(w, &tags, &(0..n).collect::<Vec<_>>()),
                BatchSize::SmallInput,
            );
        });
    }
    group.finish();
}

/// End-to-end `set_tags` on the cache-hit path with a fixed non-monotonic ordering, so
/// the sort does genuine work on every call. Worst realistic per-call overhead.
fn bench_set_tags_unstable(c: &mut Criterion) {
    let mut group = c.benchmark_group("set_tags_unstable");
    for &n in SIZES {
        let tags = make_tags(n);
        let order = scramble(n);
        group.bench_with_input(BenchmarkId::from_parameter(n), &n, |b, _| {
            b.iter_batched_ref(
                || seeded_writer(&tags, &order),
                |w| write_metric(w, &tags, &order),
                BatchSize::SmallInput,
            );
        });
    }
    group.finish();
}

/// Creates a writer with `tags` (presented in `order`) already interned, so subsequent
/// `set_tags` calls with the same tag set take the cache-hit path.
fn seeded_writer(tags: &[String], order: &[usize]) -> V3Writer {
    let mut w = V3Writer::new();
    write_metric(&mut w, tags, order);
    w
}

/// Writes a single count metric whose tags are `tags` presented in `order`.
fn write_metric(w: &mut V3Writer, tags: &[String], order: &[usize]) {
    let mut m = w.write(V3MetricType::Count, "bench.metric");
    m.set_tags(black_box(order.iter().map(|&i| &tags[i])));
    m.add_point(1000, 1.0).unwrap();
    m.close();
}

criterion_group!(
    benches,
    bench_sort_isolated,
    bench_set_tags_stable,
    bench_set_tags_unstable
);
criterion_main!(benches);
