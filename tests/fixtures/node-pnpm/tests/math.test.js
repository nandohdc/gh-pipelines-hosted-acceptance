import assert from "node:assert/strict";
import { describe, it } from "node:test";

import { add } from "../src/math.js";

describe("add", () => {
  it("sums two numbers", () => {
    assert.equal(add(2, 3), 5);
  });
});
