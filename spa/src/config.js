// try {
//   window.$ = window.jQuery = require("jquery")
//
//   require("bootstrap")
// } catch (e) {}

window.axios = require("axios");
// window.axios.defaults.baseURL = API_URL;
window.axios.defaults.baseURL = 'http://localhost:9999';
// window.axios.defaults.headers.common["X-Requested-With"] = "XMLHttpRequest"
// let token = document.head.querySelector('meta[name="csrf-token"]')
// if (token) {
//   window.axios.defaults.headers.common["X-CSRF-TOKEN"] = token.content
// } else {
//   console.error(
//       "CSRF token not found: https://laravel.com/docs/csrf#csrf-x-csrf-token"
//   )
// }