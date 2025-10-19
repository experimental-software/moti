import { describe, expect, test } from "vitest";
import { splitParagraph } from "./text";

function sum(a: number, b: number) {
  return a + b;
}

test("adds 1 + 2 to equal 3", () => {
  expect(sum(1, 2)).toBe(4);
});

describe("splitParagraph", () => {
  test("search term begin of sentence", () => {
    const originalText = "We are extremely careful to";
    const result = splitParagraph(originalText, "we are");
    expect(result[0]).toBe("We are");
    expect(result[1]).toBe("extremely careful to");
  });

  test("search term middle of sentence", () => {
    const originalText = "into a flame. Again, we say";
    const result = splitParagraph(originalText, "again");
    expect(result[0]).toBe("into a flame.");
    expect(result[1]).toBe("Again");
    expect(result[2]).toBe(", we say");
  });

  test("search term end of sentence", () => {
    const originalText = "Present fate of various beliefs and faiths";
    const result = splitParagraph(originalText, "faiths");
    expect(result[0]).toBe("Present fate of various beliefs and");
    expect(result[1]).toBe("faiths");
  });
});
