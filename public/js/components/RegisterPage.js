import { API } from "../services/API.js";
import { CollectionPage } from "../components/CollectionPage.js";

export class RegisterPage extends CollectionPage {
  constructor() {
    super(API.getUrls, "Login");
  }

  connectedCallback() {
    const template = document.getElementById("template-Register");
    const content = template.content.cloneNode(true);

    this.appendChild(content);
  }
}

customElements.define("register-page", RegisterPage);
