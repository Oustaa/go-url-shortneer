import { API } from "../services/API.js";
import { CollectionPage } from "../components/CollectionPage.js";

export class LoginPage extends CollectionPage {
  constructor() {
    super(API.getUrls, "Login");
  }

  connectedCallback() {
    const template = document.getElementById("template-Login");
    console.log("login template: ", template);
    const content = template.content.cloneNode(true);

    this.appendChild(content);
  }
}

customElements.define("login-page", LoginPage);
