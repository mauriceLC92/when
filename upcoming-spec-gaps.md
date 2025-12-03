# upcoming-spec.md Coverage Analysis

## Executive Summary

This report documents the gaps between the current `when` library capabilities and the requirements specified in `upcoming-spec.md`. Tests were created to validate all expression categories from the spec, with test fixtures using the spec's expected behavior.

**Test Results:** 2 passing / 15 failing / 1 skipped (out of 18 test functions)

**Success Rate:** ~11% of test functions passed completely, but many failures are due to minor default time mismatches rather than complete lack of support.

## Test Results Summary

### ✅ Fully Supported (2 test functions)
- **Relative Durations** - "in X minutes/hours/days" ✅
- **Tomorrow with Explicit Time** - "tomorrow at 09:00" ✅

### ⚠️ Partially Supported - Default Time Mismatches (6 test functions)
- **Today Time of Day** - Parses but uses different defaults
- **Tomorrow Time of Day** - Parses but uses different defaults
- **Weekdays Time of Day** - Parses but uses different defaults
- **Default Time Behaviors** - Parses but uses different defaults

### ⚠️ Partially Supported - Missing Default Times (3 test functions)
- **Tomorrow No Time** - Parses date but doesn't set default 09:00
- **Weekdays No Time** - Parses weekday but doesn't set default 09:00
- **Relative Week/Month** - Parses but doesn't set default time

### ⚠️ Partially Supported - Index Mismatches (3 test functions)
- **Today with Explicit Time** - Parses correctly but at different text position
- **Weekdays with Time** - Parses correctly but at different text position
- **Weekdays No Time** - Parses but at different text position

### ❌ Unsupported - Missing Rules (5 test functions)
- **Later Today** - "later today" not recognized
- **This Weekend** - "this weekend" not recognized
- **Next Quarter** - "next quarter" not recognized
- **After Lunch** - "after lunch" not recognized
- **After Work** - "after work" not recognized
- **Before End of Day** - "before end of day" / "before EOD" not recognized

### 🚫 Out of Scope (1 test function)
- **Recurring Reminders** - "every hour", "every Monday at 9am" etc. (requires architecture change)

---

## Detailed Category Analysis

### A. Today / Later Today

**Status:** ⚠️ Partial Support

**What Works:**
- ✅ Explicit times: "at 3pm", "at 15:00"
- ✅ Relative durations: "in 30 minutes", "in 2 hours"

**Default Time Mismatches:**
| Expression | Current Default | Spec Requirement | Difference |
|------------|----------------|------------------|------------|
| this morning | 08:00 | 09:00 | -1 hour |
| this afternoon | 15:00 | 12:00 | +3 hours |
| this evening | 18:00 | 19:00 | -1 hour |
| tonight | 23:00 | 20:00 | +3 hours |

**Missing Rules:**
- ❌ "later today" → should be now + 3 hours (rule doesn't exist)

**Impact:** High - Time-of-day expressions are commonly used, and the default mismatches mean the library produces times 1-3 hours different from spec.

---

### B. Tomorrow

**Status:** ⚠️ Partial Support

**What Works:**
- ✅ Tomorrow with explicit time: "tomorrow at 09:00", "tomorrow at 2pm"

**Default Time Mismatches:**
| Expression | Current Default | Spec Requirement | Difference |
|------------|----------------|------------------|------------|
| tomorrow morning | 08:00 | 09:00 | -1 hour |
| tomorrow afternoon | 15:00 | 12:00 | +3 hours |
| tomorrow evening | 18:00 | 19:00 | -1 hour |

**Missing Default Time:**
- ❌ "tomorrow" (no time specified) should default to 09:00
  - Currently parses as tomorrow at 00:00 (24h offset)
  - Should parse as tomorrow at 09:00 (33h offset)

**Impact:** High - Users expect "tomorrow" to mean tomorrow morning at a reasonable working hour, not midnight.

---

### C. Specific Weekdays

**Status:** ⚠️ Partial Support

**What Works:**
- ✅ Weekdays with explicit time: "on Monday at 10am", "Friday at 4pm"

**Default Time Mismatches:**
| Expression | Current Default | Spec Requirement | Difference |
|------------|----------------|------------------|------------|
| Friday morning | 08:00 | 09:00 | -1 hour |
| Friday afternoon | 15:00 | 12:00 | +3 hours |
| Monday evening | 18:00 | 19:00 | -1 hour |

**Missing Default Time:**
- ❌ Weekdays without time should default to 09:00
  - "on Wednesday", "on Monday", "on Friday" parse but don't set time
  - Should default to 09:00 on that weekday

**Index Matching Issue:**
- Some expressions match at unexpected text positions (e.g., "on Monday" matches at index 3)
- This may be acceptable behavior (matching the significant part) but worth reviewing

**Impact:** Medium - Weekday scheduling is common, and missing default times reduce usability.

---

### D. Relative Dates

**Status:** ❌ Major Gaps

**What Works:**
- ⚠️ "next week", "next month", "next year" parse but with issues

**Issues:**
- ❌ "next week" calculation seems incorrect (got 168h = 7 days, expected 129h = 5 days 9 hours to next Monday at 09:00)
  - Spec: "next week" should mean next Monday at 09:00
  - Current: Seems to add 7 days without specific day/time

- ❌ "next month" should default to 1st of next month at 09:00 (not tested yet, likely similar issue)

**Missing Rules:**
- ❌ "this weekend" → Saturday at 10:00 (rule doesn't exist)
- ❌ "next quarter" → first day of next quarter at 09:00 (no quarter support)

**Impact:** Medium-High - "next week" is commonly used and needs to map to Monday morning, not just +7 days.

---

### E. Relative Durations

**Status:** ✅ Fully Supported

**What Works:**
- ✅ All relative duration patterns: "in 5 minutes", "in 20 minutes", "in 1 hour", "in 2 hours", "in 24 hours"
- ✅ Supported by the `Deadline` rule in deadline.go
- ✅ Pattern: `(within|in) X (seconds|minutes|hours|days|weeks|months|years)`

**Impact:** None - This category is fully functional!

---

### F. Recurring Reminders

**Status:** 🚫 Out of Scope

**Examples from spec:**
- "every hour"
- "every weekday at 09:00"
- "every Monday at 9am"
- "on the first day of every month"

**Why Out of Scope:**
The `when` library is designed to parse a single point in time, not a schedule or recurrence pattern. Supporting recurring reminders would require:
1. Architecture change to return multiple times or a schedule object
2. New data structures to represent recurrence rules
3. Significant complexity in rule parsing and context handling

**Recommendation:** Consider this a separate feature request that would require a v2 API or companion library.

**Impact:** N/A - Out of scope for current library design

---

### G. Vague Phrasing

**Status:** ❌ Mostly Unsupported

**Missing Rules:**
- ❌ "after lunch" → 14:00 (2pm)
- ❌ "after work" → 18:00 (6pm)
- ❌ "before end of day" / "before EOD" → 17:00 (5pm)
- ❌ "this weekend" → Saturday at 10:00 (also missing from Category D)

**Impact:** Medium - These are common workplace phrases that would improve user experience if supported.

---

## Default Time Behavior Issues

The spec defines specific default times for time-of-day expressions, but the library uses different defaults:

| Time of Day | Current Default (rules/en/casual_time.go) | Spec Requirement | Fix Needed |
|-------------|------------------------------------------|------------------|------------|
| Morning | 08:00 | 09:00 | Change default in CasualTime rule |
| Afternoon | 15:00 (3pm) | 12:00 (noon) | Change default in CasualTime rule |
| Evening | 18:00 (6pm) | 19:00 (7pm) | Change default in CasualTime rule |
| Tonight | 23:00 (11pm) | 20:00 (8pm) | Change default in CasualDate rule |

**Note:** The library allows customizing these via `rules.Options`, but the built-in defaults differ from spec.

---

## Missing Rules - Implementation Needed

### High Priority (Common Use Cases)

#### 1. "later today" → now + 3 hours
**File:** Create `rules/en/later_today.go`
**Pattern:** `(?i)(later\s+today)`
**Logic:** Add 3 hours to reference time
**Impact:** High - Common vague expression

#### 2. Default times for dates without explicit times
**Files:** Modify existing rules (weekday.go, casual_date.go, relative_week.go)
**Logic:** When no time is set, default to 09:00
**Examples:**
- "tomorrow" → tomorrow at 09:00
- "next Monday" → next Monday at 09:00
- "next week" → next Monday at 09:00
**Impact:** High - Makes vague expressions much more useful

#### 3. Adjust time-of-day defaults
**Files:** `rules/en/casual_time.go`, `rules/en/casual_date.go`
**Changes:**
- Morning: 08:00 → 09:00
- Afternoon: 15:00 → 12:00
- Evening: 18:00 → 19:00
- Tonight: 23:00 → 20:00
**Impact:** High - Aligns with user expectations from spec

### Medium Priority (Useful Workplace Phrases)

#### 4. "after lunch" → 14:00 (2pm)
**File:** Create `rules/en/after_lunch.go`
**Pattern:** `(?i)after\s+lunch`
**Logic:** Set hour to 14:00 on current day
**Impact:** Medium

#### 5. "after work" → 18:00 (6pm)
**File:** Create `rules/en/after_work.go`
**Pattern:** `(?i)after\s+work`
**Logic:** Set hour to 18:00 on current day
**Impact:** Medium

#### 6. "before end of day" / "EOD" → 17:00 (5pm)
**File:** Create `rules/en/end_of_day.go`
**Pattern:** `(?i)before\s+(end\s+of\s+day|EOD|eod)`
**Logic:** Set hour to 17:00 on current day
**Impact:** Medium - Common deadline phrase

#### 7. "this weekend" → Saturday at 10:00
**File:** Create `rules/en/this_weekend.go`
**Pattern:** `(?i)this\s+weekend`
**Logic:** Find next Saturday (or current Saturday if before), set time to 10:00
**Impact:** Medium

### Low Priority (Edge Cases)

#### 8. "next quarter" support
**File:** Modify `rules/en/relative_week.go` or create `rules/en/quarter.go`
**Pattern:** `(?i)next\s+quarter`
**Logic:** Calculate first day of next quarter at 09:00
**Complexity:** Higher - requires quarter calculation logic
**Impact:** Low - Less commonly used

---

## Implementation Priority Recommendations

### Phase 1: Quick Wins (Default Time Adjustments)
**Effort:** Low | **Impact:** High

1. Adjust time-of-day defaults in `casual_time.go` and `casual_date.go`
   - Change 4 constant values
   - Run existing tests to ensure no breakage
   - Update tests to reflect new defaults

2. Add default 09:00 time for date-only expressions
   - Modify `CasualDate` rule for "tomorrow"
   - Modify `Weekday` rule for weekdays without time
   - Modify `RelativeWeek` rule for "next week"

**Estimated effort:** 2-3 hours
**Test count affected:** ~8 test functions would start passing

### Phase 2: Common Vague Expressions
**Effort:** Medium | **Impact:** High

3. Implement "later today" → +3 hours rule
   - New file: `later_today.go`
   - Simple duration addition
   - Add to `All` slice in correct position

4. Implement workplace phrases:
   - "after lunch" → 14:00
   - "after work" → 18:00
   - "before end of day"/"EOD" → 17:00
   - All follow similar pattern, can be implemented together

**Estimated effort:** 3-4 hours
**Test count affected:** ~4 test functions would start passing

### Phase 3: Weekend and Edge Cases
**Effort:** Medium | **Impact:** Medium

5. Implement "this weekend" → Saturday 10:00
   - Requires weekday calculation logic
   - Handle edge case of already being on Saturday

6. Implement "next quarter" support
   - More complex date arithmetic
   - Quarter boundaries: Jan 1, Apr 1, Jul 1, Oct 1

**Estimated effort:** 4-5 hours
**Test count affected:** ~2 test functions would start passing

### Phase 4: Test Refinement
**Effort:** Low | **Impact:** Quality

7. Review index matching expectations
   - Some tests expect index 0 but get index 3
   - Determine if this is acceptable library behavior
   - Update test fixtures if needed

**Estimated effort:** 1-2 hours

---

## Files Requiring Changes

### Existing Files to Modify:

1. **`/rules/en/casual_time.go`** - Lines 26-54
   - Change default hour values for morning/afternoon/evening

2. **`/rules/en/casual_date.go`** - Lines 16-46
   - Change "tonight" default from 23:00 to 20:00
   - Add default 09:00 for "tomorrow" when no time is set

3. **`/rules/en/weekday.go`** - Review and modify
   - Add default 09:00 when no time is specified with weekday

4. **`/rules/en/relative_week.go`** - Review and modify
   - Fix "next week" calculation to mean next Monday at 09:00
   - Add default 09:00 for "next month"

5. **`/rules/en/en.go`** - Update `All` slice
   - Add new rules in appropriate order

### New Files to Create:

1. **`/rules/en/later_today.go`** - "later today" rule
2. **`/rules/en/after_lunch.go`** - "after lunch" rule
3. **`/rules/en/after_work.go`** - "after work" rule
4. **`/rules/en/end_of_day.go`** - "before end of day/EOD" rule
5. **`/rules/en/this_weekend.go`** - "this weekend" rule
6. **`/rules/en/quarter.go`** (optional) - Quarter support

### Test Files to Update:

1. **`/rules/en/upcoming_spec_test.go`**
   - Update expected values as rules are fixed
   - Adjust index expectations if needed

---

## Conclusion

The `when` library has solid foundational parsing capabilities, particularly for:
- Explicit times ("at 3pm")
- Relative durations ("in 2 hours")
- Basic date expressions with explicit times

However, to fully support the `upcoming-spec.md` requirements, the following work is needed:

**Critical:** Adjust default times to match spec expectations (~2-3 hours)
**Important:** Add missing vague expression rules (~7-9 hours)
**Optional:** Add edge case support like "next quarter" (~4-5 hours)

**Total estimated effort:** 13-17 hours to achieve ~90% spec coverage (excluding recurring patterns which are out of scope).

The test suite created in `upcoming_spec_test.go` provides an excellent regression test suite and clear targets for implementation.
