import { CollectionPage } from "../components/CollectionPage.js";

export class LoginPage extends CollectionPage {
  connectedCallback() {
    const template = document.getElementById("template-Login");
    const content = template.content.cloneNode(true);

    this.appendChild(content);
  }
}

customElements.define("login-page", LoginPage);
