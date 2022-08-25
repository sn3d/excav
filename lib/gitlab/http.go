package gitlab

// StatusCode is just wrapper type for HTTP status code numbers
// It's used instead of error for methods like get(), post(), delete().
// For codes meaning see https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
// There is one special value '-1' which means unknown error on client side.
// It might be connection problem or something else.
type StatusCode int
