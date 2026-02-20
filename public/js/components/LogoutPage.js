import { CollectionPage } from "../components/CollectionPage.js";

export class LogoutPage extends CollectionPage {
  constructor() {
    app.Store.jwt = "";
    app.Router.go("/account/login");
    app.authUi();
  }
}

customElements.define("logout-page", LogoutPage);
