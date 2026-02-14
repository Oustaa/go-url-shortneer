import { API } from "../services/API.js";
import { CollectionPage } from "./CollectionPage.js";

export class CreateURLPage extends CollectionPage {
  constructor() {
    super(API.getUrls, "Create Url");
  }

  connectedCallback() {
    const template = document.getElementById("template-createURL");
    const content = template.content.cloneNode(true);
    this.appendChild(content);
  }
}

customElements.define("create-url-page", CreateURLPage);
