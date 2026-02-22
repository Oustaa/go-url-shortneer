export class LogoutPage {
  constructor() {
    app.Store.jwt = "";
    app.Router.go("/account/login");
    app.authUi();
  }
}

customElements.define("logout-page", LogoutPage);
