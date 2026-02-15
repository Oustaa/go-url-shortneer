import { routes } from "./Routes.js";

export const Router = {
  init() {
    window.addEventListener("popstate", () => {
      Router.go(location.pathname, false);
    });

    document.querySelectorAll("a.navlink").forEach((a) => {
      a.addEventListener("click", (event) => {
        event.preventDefault();
        const href = a.getAttribute("href");
        Router.go(href);
      });
    });

    Router.go(location.pathname + location.search);
  },
  go(path) {
    history.pushState(null, "", path);

    let pageElement;
    let authRequired = false;

    for (const r of routes) {
      if (r.path === path) {
        pageElement = new r.component();
        authRequired = r.loggedIn;
      }
    }

    if (pageElement) {
      if (authRequired && !app.Store.jwt) {
        app.Router.go("/account/login");
        return;
      }
    }

    if (pageElement == null) {
      pageElement = document.createElement("h1");
      pageElement.textContent = "Page not found";
    }

    function updatePage() {
      document.querySelector("main").innerHTML = "";
      document.querySelector("main").appendChild(pageElement);
    }

    updatePage();
  },
};
