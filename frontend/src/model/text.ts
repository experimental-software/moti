export function splitParagraph(
  paragraphOriginal: string,
  searchTerm: string,
): string[] {
  if (searchTerm === "") {
    return [paragraphOriginal];
  }
  const paragraphLowercase = paragraphOriginal.toLowerCase();
  const searchTermIndex = paragraphLowercase.indexOf(searchTerm.toLowerCase());
  if (searchTermIndex === -1) {
    return [paragraphOriginal];
  }
  const result = [];
  const searchTermOriginal = paragraphOriginal.substring(
    searchTermIndex,
    searchTermIndex + searchTerm.length,
  );
  const textBeforeSearchTerm = paragraphOriginal
    .substring(0, searchTermIndex)
    .trim();
  const textAfterSearchTerm = paragraphOriginal
    .substring(searchTermIndex + searchTerm.length)
    .trim();
  if (textBeforeSearchTerm.length > 0) {
    result.push(textBeforeSearchTerm);
  }
  result.push(searchTermOriginal);
  if (textAfterSearchTerm.length > 0) {
    result.push(textAfterSearchTerm);
  }

  return result;

  // TODO: Support for multiple occurrences of search string in paragraph
  // TODO: Support regex search term, case insensitive
}
