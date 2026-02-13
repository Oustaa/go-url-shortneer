import { API } from "./services/API.js";
import { Router } from "./services/Router.js";

window.addEventListener("DOMContentLoaded", () => {
  app.Router.init();
});

window.app = {
  API,
  Router,
};
