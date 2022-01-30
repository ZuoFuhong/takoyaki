package defs

type SEARCH_MODE string

const FUZZY_MATCH SEARCH_MODE = "like"  // 模糊匹配
const EXACT_MATCH SEARCH_MODE = "equal" // 精准匹配
